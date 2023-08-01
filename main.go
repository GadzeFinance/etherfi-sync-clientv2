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
	"bufio"
	"strings"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/glebarez/go-sqlite"
	"github.com/robfig/cron"
)

func main() {

	config, err := utils.GetConfig("./config.json")
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = utils.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Specify 'listen' or 'changes' argument")
		return
	}

	programType := os.Args[1]
	if programType == "listen" {
		fmt.Println("Starting Sync Client!")
		fmt.Println("Configuration values: ")
		fmt.Println(PrettyPrint(config))
		c := cron.New()
		c.AddFunc("1 * * * *", func() {
			if err := cronjob(config, db); err != nil {
				fmt.Printf("Error executing function: %s\n", err)
				os.Exit(1)
			}
		})
	
		c.Start()
		for {
			time.Sleep(time.Second)
		}
	} else if programType == "changes" {
	
		// Query database and find all PENDING
		pendingBids, err := utils.GetRowsByStatus(db, "PENDING")
		if err != nil {
			panic(err)
		}
		// Query all and find those that are EXITED
		exitedBids, err := utils.GetRowsByStatus(db, "EXITED")
		if err != nil {
			panic(err)
		}
		// Print them out each
		fmt.Println("The following validators with these bid IDs will be modified")
		t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Bid ID", "Public Key", "Change"})
		for _, bid := range pendingBids {
			t.AppendRow([]interface{}{bid.Id, bid.Pubkey, "ADD"})
		}

		for _, bid := range exitedBids {
			t.AppendRow([]interface{}{bid.Id, bid.Pubkey, "REMOVE"})
		}

		t.Render()


		// If yes is pressed, copy file contents of each bid and paste them into respective location and update validators in TEKU
		fmt.Println(`Type "CONFIRM" to apply these changes`)
		fmt.Print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)

		// Refresh Teku
		if input == "CONFIRM" {
			for _, bid := range exitedBids {
				if err := utils.DeleteFromTeku(config.PATH_TO_VALIDATOR, bid.Id); err != nil {
					panic(err)
				}
				if err := utils.UpdateRowStatus(db, bid.Id, "REMOVED"); err != nil {
					panic(err)
				}
				if err:= utils.RemoveTekuProposerConfig(config.PATH_TO_VALIDATOR, bid.Pubkey); err != nil {
					panic(err)
				}
			}
			for _, bidItem := range pendingBids {

				bid, err := utils.GetBid(db, bidItem.Id)
				if err != nil {
					panic(err)
				}
				if config.PATH_TO_VALIDATOR != "" {
					if err := utils.SaveTekuProposerConfig(config.PATH_TO_VALIDATOR, bid.Pubkey, bid.NodeAddress); err != nil {
						panic(err)
					}
					if err := utils.AddToTeku(config.PATH_TO_VALIDATOR, bid.Id, bid.Password, bid.Keystore); err != nil {
						panic(err)
					}
				}

				if err := utils.UpdateRowStatus(db, bidItem.Id, "ADDED"); err != nil {
					panic(err)
				}
			}
			utils.RefreshTeku()
			fmt.Println("Changes added")
		}
	}
}

func cronjob(config schemas.Config, db *sql.DB) error {

	fmt.Println("Searching for new stake requests 👀")

	privateKey, err := utils.ExtractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return err
	}

	isUsingCBC := false
	// For compatibility, if the authTag is empty, we know it's CBC mode
	if privateKey.AuthTag == "" {
		isUsingCBC = true
	}

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER, config.STAKER)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	for i, bid := range bids {
		_ = i

		count, err := utils.GetIDCount(db, bid.Id)
		if err != nil {
			fmt.Println("Error querying database")
			return err
		}

		if count > 0 {
			continue
		}

		fmt.Println(`> start processing stake request from: ` + bid.Validator.BNFTHolder)

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

		// Add query that checks beacon nodes for all active validators
		if utils.GetExitStatus(validator.ValidatorPubKey) {
			if err != utils.UpdateRowStatus(db, bid.Id, "EXITED") {
				panic(err)
			}
		}
	}

	return nil
}

// This function fetch bids from the Graph
// TODO: Paginate this
func retrieveBidsFromSubgraph(GRAPH_URL string, BIDDER string, STAKER string) ([]schemas.BidType, error) {

	validatorFilter := `{ phase: VALIDATOR_REGISTERED }`
	if STAKER != "" {
		validatorFilter = `{ phase: VALIDATOR_REGISTERED, BNFTHolder: "` + STAKER + `"}`
	}

	// the query to fetch bids
	queryJsonData := map[string]string{
		"query": `
		  {
      	bids(where: { bidderAddress: "` + BIDDER + `", status: "WON", validator_not: null, validator_: ` + validatorFilter + ` }) {
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

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
