package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
	_ "github.com/glebarez/go-sqlite"
	"github.com/robfig/cron"
)

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := utils.GetConfig("./config.json")
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}

	fmt.Println("Starting Sync Client!")
	fmt.Println("Configuration values: ")
	fmt.Println(PrettyPrint(config))

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

	if len(os.Args) < 2 {
		fmt.Println("Specify 'listen' or 'add' argument")
		return
	}

	programType := os.Args[1]
	if programType == "listen" {
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
	} else if programType == "add" {

		if len(os.Args) < 3 {
			fmt.Println("Specify the bid id argument")
			return
		}

		bidId := os.Args[2]

		count, err := utils.GetIDCount(db, bidId)
		if err != nil {
			fmt.Println("Error querying database")
			return
		}

		if count == 0 {
			fmt.Println("No bids with the ID specified")
			return
		}

		stmt, err := db.Prepare("SELECT executed FROM winning_bids WHERE id = ?")
		if err != nil {
			fmt.Println("Error preparing statement:", err)
			return
		}
		defer stmt.Close()

		var executed bool
		err = stmt.QueryRow(bidId).Scan(&executed)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No row found with the specified ID")
			} else {
				fmt.Println("Error retrieving data:", err)
			}
			return
		}

		if executed {
			fmt.Println("These keys have already been added. Ending program")
			return
		}

		updateStmt, err := db.Prepare("UPDATE winning_bids SET executed = true WHERE id = ?")
		if err != nil {
			fmt.Println("Error preparing UPDATE statement:", err)
			return
		}
		defer updateStmt.Close()

		out, err := exec.Command(
			"sudo",
			config.PATH_TO_PRYSYM_SH,
			"validator",
			"accounts",
			"import",
			"--goerli",
			"--wallet-dir=",
			config.CONSENSUS_FOLDER_LOCATION,
			"--keys-dir=",
			config.ETHERFI_SC_CLIENT_LOCATION).Output()

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = updateStmt.Exec(bidId)
		if err != nil {
			fmt.Println("Error updating data:", err)
			return
		}

		fmt.Println(out)

	} else {
		fmt.Println("Specify 'listen' or 'add' argument")
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

		if err := utils.SaveKeysToFS(config.OUTPUT_LOCATION, config.CONSENSUS_FOLDER_LOCATION, config.ETHERFI_SC_CLIENT_LOCATION, data, bid.Id, validator.ValidatorPubKey, bid.Validator.EtherfiNode, db); err != nil {
			return err
		}

	}

	return nil
}

// This function fetch bids from the Graph
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

	fmt.Println("Bids:", len(result.Data.Bids))

	return result.Data.Bids, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
