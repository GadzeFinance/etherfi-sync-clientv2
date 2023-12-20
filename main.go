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
	config, err := utils.GetAndCheckConfig("config.json")
	if err != nil {
		return fmt.Errorf("Failed to load config: %w\n", err)
	}

	fmt.Println("Starting EtherFi Sync Client:")
	fmt.Println("Operator Address: ", config.BIDDER)
	fmt.Println("Output directory: ", config.OUTPUT_LOCATION)

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		return fmt.Errorf("failed to open sqlite db: %w", err)
	}
	defer db.Close()

	if err := utils.CreateTable(db); err != nil {
		return fmt.Errorf("crating bids table: %w", err)
	}

	return fetchValidatorKeys(config, db)
}

func fetchValidatorKeys(config schemas.Config, db *sql.DB) error {

	fmt.Println("Fetching Validator Keys from IPFS...")
	privateKey, err := utils.ParseKeystoreFile(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("parsing keystore: %w", err)
	}

	// For compatibility, if the authTag is empty, we know it's CBC mode
	isUsingCBC := false
	if privateKey.AuthTag == "" {
		isUsingCBC = true
	}

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		return fmt.Errorf("retrieveBidsFromSubgraph: %w", err)
	}

	fmt.Println("Found ", len(bids), " stake requests.")
	for _, bid := range bids {

		count, err := utils.GetIDCount(db, bid.Id)
		if err != nil {
			return fmt.Errorf("GetIDCount: %w", err)
		}

		if count > 0 {
			fmt.Printf("Skipping stake request for validator: %s because it has already been processed.", bid)
			continue
		}

		fmt.Println(`Processing stake request for validator: ` + bid.Id + ` and phase: ` + bid.Validator.Phase + ` and BNFT Holder: ` + bid.Validator.BNFTHolder + ` and ipfs path: ` + bid.Validator.IpfsHashForEncryptedValidatorKey)

		if bid.Validator.Phase == "READY_FOR_DEPOSIT" || bid.Validator.Phase == "STAKE_DEPOSITED" {
			continue
		}

		validator := bid.Validator
		ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey

		IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
		if err != nil {
			return fmt.Errorf("FetchFromIPFS: %w", err)
		}

		var validatorKey schemas.DecryptedDataJSON
		if isUsingCBC {
			validatorKey, err = utils.DecryptPrivateKeysCBC(privateKey, config.PASSWORD)
		} else {
			validatorKey, err = utils.DecryptPrivateKeysGCM(privateKey, config.PASSWORD)
		}
		if err != nil {
			return fmt.Errorf("DecryptPrivateKeys: %w", err)
		}

		pubKeyArray := validatorKey.PublicKeys
		privKeyArray := validatorKey.PrivateKeys
		keypairForIndex, err := utils.GetKeyPairByPubKeyIndex(bid.PubKeyIndex, privKeyArray, pubKeyArray)
		if err != nil {
			return fmt.Errorf("GetKeyPairByPubKeyIndex: %w", err)
		}

		data := utils.DecryptValidatorKeyInfo(IPFSResponse, keypairForIndex)

		if err := utils.SaveKeysToFS(config.OUTPUT_LOCATION, data, bid.Id, validator.ValidatorPubKey, bid.Validator.EtherfiNode, db); err != nil {
			return fmt.Errorf("SaveKeysToFS: %w", err)
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
		return nil, fmt.Errorf("creating request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Bid request: %w", err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading bid data: %w", err)
	}

	var result schemas.GQLResponseType
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("marshalling gql bid response: %w", err)
	}

	return result.Data.Bids, nil
}
