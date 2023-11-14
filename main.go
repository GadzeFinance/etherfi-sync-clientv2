package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	config, err := utils.GetAndCheckConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return err
	}

	fmt.Println("Starting EtherFi Sync Client:")
	fmt.Println("Operator Address: ", config.BIDDER)
	fmt.Println("Output directory: ", config.OUTPUT_LOCATION)

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	err = utils.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fetchValidatorKeys(config, db)
	return nil
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

	fmt.Println("Found ", len(bids), " stake requests.")
	for i, bid := range bids {
		_ = i

		count, err := utils.GetIDCount(db, bid.Id)
		if err != nil {
			return err
		}

		if count > 0 {
			fmt.Println(`Skipping stake request for validator: ` + bid.Id + ` because it has already been processed.`)
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

func retrieveBidsFromSubgraph(GRAPH_URL string, BIDDER string) ([]schemas.BidType, error) {
	var allBids []schemas.BidType
	itemsPerPage := 100 // Number of items to fetch per page

	for skip := 0; ; skip += itemsPerPage {
		fmt.Printf("Fetching bids %d to %d\n", skip+1, skip+itemsPerPage)

		query := fmt.Sprintf(`
		  {
			bids(where: { bidderAddress: "%s", status: "WON", validator_not: null }, first: %d, skip: %d) {
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
		  }`, BIDDER, itemsPerPage, skip)

		queryJsonData := map[string]string{"query": query}
		jsonValue, _ := json.Marshal(queryJsonData)

		request, err := http.NewRequest("POST", GRAPH_URL, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Second * 10}
		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return nil, err
		}
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)

		var result schemas.GQLResponseType
		if err := json.Unmarshal(data, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
			return nil, err
		}

		if len(result.Data.Bids) == 0 {
			break // Break the loop if no more bids are found
		}

		allBids = append(allBids, result.Data.Bids...)
	}

	return allBids, nil
}
