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
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

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
	getValidatorRegisteredEvents(db, config)
	return nil

	return fetchValidatorKeys(config, db)
}

func getValidatorRegisteredEvents(db *sql.DB, cfg schemas.Config) error {

	var (
		stakingManagerAddr      = common.HexToAddress("0x25e821b7197B146F7713C3b89B6A4D83516B912d")
		auctionManagerAddr      = common.HexToAddress("0x00C452aFFee3a17d9Cecc1Bcd2B8d5C7635C4CB9")
		etherFiNodesManagerAddr = common.HexToAddress("0x8B71140AD2e5d1E7018d2a7f8a288BD3CD38916F")
	)
	rpcClient, err := ethclient.Dial(cfg.RPC_URL)
	if err != nil {
		panic(err)
	}
	stakingManager, _ := NewStakingManager(stakingManagerAddr, rpcClient)
	auctionManager, _ := NewAuctionManager(auctionManagerAddr, rpcClient)
	etherFiNodesManager, _ := NewEtherFiNodesManager(etherFiNodesManagerAddr, rpcClient)

	// operator / bnft / tnft
	operator := common.HexToAddress("0xcA2e28756379f660d15f1113dB5Cf34A73CbE89b")
	events, err := stakingManager.FilterValidatorRegistered(nil, []common.Address{operator}, nil, nil)
	if err != nil {
		panic(err)
	}
	count := 0
	for events.Next() {
		count += 1
		event := events.Event
		fmt.Printf("Register: %+v\n", event)

		// the original authors did not add a primay key or any sort of key to database :(
		// so we check if there are any number of entries tied to this validator id
		count, err := utils.GetIDCount(db, event.ValidatorId.String())
		if err != nil {
			return fmt.Errorf("failed to check existence of validator in db")
		}
		if count > 0 {
			// we have already downloaded keys for this validator
			continue
		}

		// fetch the bid data from auctionManager
		bid, err := auctionManager.Bids(&bind.CallOpts{}, event.ValidatorId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", bid)

		// grab the associated node
		nodeAddr, err := etherFiNodesManager.EtherfiNodeAddress(nil, event.ValidatorId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("nodeAddr: %+v\n", nodeAddr.Hex())

		fmt.Println(`Processing stake request for validator: ` + event.ValidatorId.String() + ` and ipfs path: ` + event.IpfsHashForEncryptedValidatorKey)

		IPFSResponse, err := utils.FetchFromIPFS(cfg.IPFS_GATEWAY, event.IpfsHashForEncryptedValidatorKey)
		if err != nil {
			fmt.Println("failed to fetch encrypted key from IPFS, Retrying Once: " + event.IpfsHashForEncryptedValidatorKey)
			IPFSResponse, err = utils.FetchFromIPFS(cfg.IPFS_GATEWAY, event.IpfsHashForEncryptedValidatorKey)
			if err != nil {
				return fmt.Errorf("FetchFromIPFS: %w", err)
			}
		}
		fmt.Println(IPFSResponse)

		/*
			var validatorKey schemas.DecryptedDataJSON
			if isUsingCBC {
				validatorKey, err = utils.DecryptPrivateKeysCBC(privateKey, config.PASSWORD)
			} else {
				validatorKey, err = utils.DecryptPrivateKeysGCM(privateKey, config.PASSWORD)
			}
			if err != nil {
				return fmt.Errorf("DecryptPrivateKeys: %w", err)
			}

		*/

		/*
			pubKeyArray := validatorKey.PublicKeys
			privKeyArray := validatorKey.PrivateKeys
			keypairForIndex, err := utils.GetKeyPairByPubKeyIndex(bid.PubKeyIndex, privKeyArray, pubKeyArray)
			if err != nil {
				return fmt.Errorf("GetKeyPairByPubKeyIndex: %w", err)
			}

				type ValidatorKeyInfo struct {
					ValidatorKeyFile     []byte `json:"validatorKeyFile"`
					ValidatorKeyPassword []byte `json:"validatorKeyPassword"`
					KeystoreName         []byte `json:"keystoreName"`
				}
		*/

		//data := utils.DecryptValidatorKeyInfo(IPFSResponse, keypairForIndex)

		data := schemas.ValidatorKeyInfo{
			ValidatorKeyFile:     []byte("A"),
			ValidatorKeyPassword: []byte("B"),
			KeystoreName:         []byte("C"),
		}

		if err := utils.SaveKeysToFS(cfg.OUTPUT_LOCATION, data, event.ValidatorId.String(), int64(bid.BidderPubKeyIndex), hexutil.Encode(event.ValidatorPubKey), nodeAddr.Hex(), db); err != nil {
			return fmt.Errorf("SaveKeysToFS: %w", err)
		}

	}
	fmt.Println("event count:", count)

	// for each registered validator
	// see if already in the DB
	// if not, fetch the bid data
	// then decrypt from IPFS
	// then save in DB

	//check liquidityPool validatorRegistered?

	return nil
}

func fetchValidatorKeys(config schemas.Config, db *sql.DB) error {
	fmt.Println("Fetching Validator Keys from IPFS...")
	privateKey, err := utils.ParseKeystoreFile(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("parsing keystore: %w", err)
	}
	spew.Dump(privateKey)

	// For compatibility, if the authTag is empty, we know it's CBC mode
	isUsingCBC := false
	if privateKey.AuthTag == "" {
		isUsingCBC = true
	}
	fmt.Println(isUsingCBC)

	pubkeyIndex, err := utils.GetLastPubkeyIndex(db)
	if err != nil {
		return fmt.Errorf("GetLastPubkeyIndex: %w", err)
	}
	// TODO
	// there is no way of sorting against latest won bids. beacuse sync-client did not store keys which is not status:won in data.db.
	// at the moment, iterating all keys is the clear way to get recent won bids.
	if config.USE_LAST_VALIDATOR_INDEX == true {
		pubkeyIndex = -1
	}

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

			/*
				if bid.Validator.Phase == "READY_FOR_DEPOSIT" || bid.Validator.Phase == "STAKE_DEPOSITED" {
					skipCount++
					continue
				}
			*/

			fmt.Println(`Processing stake request for validator: ` + bid.Id + ` and phase: ` + bid.Validator.Phase + ` and BNFT Holder: ` + bid.Validator.BNFTHolder + ` and ipfs path: ` + bid.Validator.IpfsHashForEncryptedValidatorKey)

			validator := bid.Validator
			ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey
			fmt.Printf("V-pubkey: %+v\n", validator.ValidatorPubKey)

			IPFSResponse, err := utils.FetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
			if err != nil {
				fmt.Println("Fetch timed out, Retrying Once: " + ipfsHashForEncryptedValidatorKey)
				IPFSResponse, err = utils.FetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
				if err != nil {
					return fmt.Errorf("FetchFromIPFS: %w", err)
				}
			}
			fmt.Println(IPFSResponse)

			/*
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
			*/
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

	fmt.Println("graph result")
	spew.Dump(result)

	uniquePhases := make(map[string]bool)
	for _, x := range result.Data.Bids {
		uniquePhases[x.Validator.Phase] = true
	}

	fmt.Println("unique phases")
	for k, _ := range uniquePhases {
		fmt.Println(k)
	}

	return result.Data.Bids, nil
}
