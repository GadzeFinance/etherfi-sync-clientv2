package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

	pubkeyIndex, err := utils.GetLastPubkeyIndex(db)
	if err != nil {
		return fmt.Errorf("GetLastPubkeyIndex: %w", err)
	}
	// TODO
	// there is no way of sorting against latest won bids. beacuse sync-client did not store keys which is not status:won in data.db.
	// at the moment, iterating all keys is the clear way to get recent won bids.
	pubkeyIndex = 0

	for {
		fmt.Printf("Begin pubkeyIndex : %d\n", pubkeyIndex)
		bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER, pubkeyIndex)
		if err != nil {
			return fmt.Errorf("retrieveBidsFromSubgraph: %w", err)
		}

		if len(bids) == 0 {
			fmt.Printf("complete: fetched all keys\n")
			return nil
		}
		fmt.Println("Found ", len(bids), " stake requests.")
		skipCount := 0
		for _, bid := range bids {
			pi, err := strconv.ParseInt(bid.PubKeyIndex, 10, 64)
			if err != nil {
				return fmt.Errorf("ParseInt: %w", err)
			}
			if pubkeyIndex < pi {
				pubkeyIndex = pi
			}

			count, err := utils.GetIDCount(db, bid.Id)
			if err != nil {
				return fmt.Errorf("GetIDCount: %w", err)
			}

			if count > 0 {
				skipCount++
				// fmt.Printf("Skipping stake request for validator: %s because it has already been processed.\n", bid)
				continue
			}

			if bid.Validator.Phase == "READY_FOR_DEPOSIT" || bid.Validator.Phase == "STAKE_DEPOSITED" {
				skipCount++
				continue
			}

			fmt.Println(`Processing stake request for validator: ` + bid.Id + ` and phase: ` + bid.Validator.Phase + ` and BNFT Holder: ` + bid.Validator.BNFTHolder + ` and ipfs path: ` + bid.Validator.IpfsHashForEncryptedValidatorKey)

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

			if err := utils.SaveKeysToFS(config.OUTPUT_LOCATION, data, bid.Id, pi, validator.ValidatorPubKey, bid.Validator.EtherfiNode, db); err != nil {
				return fmt.Errorf("SaveKeysToFS: %w", err)
			}
		}
		fmt.Printf("Skipping %d stake requests because these have already been processed.\n", skipCount)
	}
}

// This function fetch bids from the Graph
func retrieveBidsFromSubgraph(GRAPH_URL string, BIDDER string, pubkeyIndex int64) ([]schemas.BidType, error) {
	// the query to fetch bids

	limit := "500"

	queryJsonData := map[string]string{
		"query": `
		{
			bids(
				where: { 
					pubKeyIndex_gt: ` + fmt.Sprintf("%d", pubkeyIndex) + `
					bidderAddress: "` + BIDDER + `"
					status: "WON"
					validator_not: null 
				} 
				first: ` + limit + `
				orderBy: pubKeyIndex
				orderDirection: asc
			) {
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
