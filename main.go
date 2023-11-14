package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	config, err := utils.GetAndCheckConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	fmt.Println("Starting EtherFi Sync Client:")
	fmt.Println("Operator Address: ", config.BIDDER)
	fmt.Println("Output directory: ", config.OUTPUT_LOCATION)

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS winning_bids (
			id STRING PRIMARY KEY,
			pubkey TEXT,
			password TEXT,
			nodeAddress TEXT,
			executed BOOLEAN DEFAULT false
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	fetchValidatorKeys(config, db)
}

func fetchValidatorKeys(config schemas.Config, db *sql.DB) error {

	fmt.Println("Fetching Validator Keys from IPFS...")
	privateKey, err := utils.ExtractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return err
	}

	// For compatibility, if the authTag is empty, we know it's CBC mode
	isUsingCBC := false
	if privateKey.AuthTag == "" {
		isUsingCBC = true
	}

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		return err
	}

	fmt.Println("Found ", len(bids), " new stake requests.")
	for i, bid := range bids {
		_ = i

		count, err := utils.GetIDCount(db, bid.Id)
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		fmt.Println(`Processing stake request for validator: ` + bid.Id + ` and BNFT Holder: ` + bid.Validator.BNFTHolder)

		validator := bid.Validator
		ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey

		IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
		if err != nil {
			return err
		}

		var validatorKey schemas.DecryptedDataJSON
		if isUsingCBC {
			validatorKey, err = utils.DecryptPrivateKeysCBC(privateKey, config.PASSWORD)
		} else {
			validatorKey, err = utils.DecryptPrivateKeysGCM(privateKey, config.PASSWORD)
		}
		if err != nil {
			return err
		}

		pubKeyArray := validatorKey.PublicKeys
		privKeyArray := validatorKey.PrivateKeys

		keypairForIndex, err := utils.GetKeyPairByPubKeyIndex(bid.PubKeyIndex, privKeyArray, pubKeyArray)

		if err != nil {
			return err
		}

		data := utils.DecryptValidatorKeyInfo(IPFSResponse, keypairForIndex)

		if err := utils.SaveKeysToFS(config.OUTPUT_LOCATION, data, bid.Id, validator.ValidatorPubKey, bid.Validator.EtherfiNode, db); err != nil {
			return err
		}
	}

	return nil
}

// This function fetch bids from the Graph
func retrieveBidsFromSubgraph(GRAPH_URL string, BIDDER string) ([]schemas.BidType, error) {
	// the query to fetch bids
	// TODO we have first 1000 here that's going to have to be fixed in the near future
	queryJsonData := map[string]string{
		"query": `
		  {
      	bids(where: { bidderAddress: "` + BIDDER + `", status: "WON", validator_not: null }, first: 1000) {
        	id
        	bidderAddress
        	pubKeyIndex
        	validator {
				id
				phase
				ipfsHashForEncryptedValidatorKey
				validatorPubKey
				etherfiNode
				BNFTHolder
        	}
      	}
    	}`,
	}
	jsonValue, _ := json.Marshal(queryJsonData)

	request, err := http.NewRequest("POST", GRAPH_URL, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		// TODO: return []
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		// TODO: return []
		return nil, err
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	var result schemas.GQLResponseType
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return result.Data.Bids, nil
}

