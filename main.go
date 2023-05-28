package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/GadzeFinance/etherfi-sync-clientv2/utils"
	"github.com/robfig/cron"
)

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := getConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}

	fmt.Println("Starting Sync Client!")
	fmt.Println("Configuration values: ")
	fmt.Println(PrettyPrint(config))

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

		// Create the table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS winning_bids (
			id STRING PRIMARY KEY,
			pubkey TEXT,
			password TEXT
		);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Specify 'listen' or 'query' argument")
		return
	}

	programType := os.Args[1]
	if programType == "listen" {
		c := cron.New()
		c.AddFunc("*/1 * * * *", func() {
	
			if err := cronjob(config, db); err != nil {
				fmt.Printf("Error executing function: %s\n", err)
				os.Exit(1)
			}
		})
	
		c.Start()
	
		for {
			time.Sleep(time.Second)
		}
	} else if programType == "query" {

		fmt.Println("Querying!")
		data := []schemas.WinningBids{}
		query := "SELECT * FROM winning_bids"
		row, err := db.Query(query)
		if err != nil {
			fmt.Println("Error querying database")
			return
		} 

		fmt.Println(row)

		defer row.Close()
		for row.Next() { // Iterate and fetch the records from result cursor

			var id string
			var pubkey string
			var password string
			row.Scan(&id, &pubkey, &password)
			fmt.Println(password)

			data = append(data, schemas.WinningBids{
				Id: id,
				Pubkey: pubkey,
				Password: password,
			})
			fmt.Println(data)
		}
		dataInJson, err := json.MarshalIndent(row, "", "  ")
		if err != nil {
			fmt.Println("Error formatting data: ", err)
		}

		fmt.Println(dataInJson)



		fmt.Println("Getting file")
	} else {
		fmt.Println("Specify 'listen' or 'query' argument")
	}
}

func cronjob(config schemas.Config, db *sql.DB) error {

	privateKey, err := utils.ExtractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return err
	}

	isUsingCBC := false
	// For compatibility, if the authTag is empty, we know it's CBC mode
	if privateKey.AuthTag == "" {
		isUsingCBC = true
	}

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	
	fmt.Println(bids)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	for i, bid := range bids {
		_ = i

		query := "SELECT COUNT(*) FROM winning_bids WHERE id = ?"
		var count int
		err = db.QueryRow(query, bid.Id).Scan(&count)
		if err != nil {
			fmt.Println("Error querying database")
			return err
		} 

		if count > 0 {
			continue
		}

		fmt.Println(`> start processing bid with id:` + bid.Id)
		
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

		keypairForIndex, err := getKeyPairByPubKeyIndex(bid.PubKeyIndex, privKeyArray, pubKeyArray)

		if err != nil {
			return err
		}

		data := utils.DecryptValidatorKeyInfo(IPFSResponse, keypairForIndex)
		
		if err := utils.SaveKeysToFS(config.OUTPUT_LOCATION, config.CONSENSUS_FOLDER_LOCATION, config.ETHERFI_SC_CLIENT_LOCATION, data, bid.Id, validator.ValidatorPubKey, db); err != nil {
			return err
		}

	}

	return nil
}

func getKeyPairByPubKeyIndex(pubkeyIndexString string, privateKeys []string, publicKeys []string) (schemas.KeyPair, error) {
	index, err := strconv.ParseInt(pubkeyIndexString, 10, 0)
	if err != nil {
		return schemas.KeyPair{}, err
	}
	return schemas.KeyPair{
		PrivateKey: privateKeys[index],
		PublicKey:  publicKeys[index],
	}, nil
}

func getConfig() (schemas.Config, error) {

	err := utils.FileExists("./config.json")
	if err != nil {
		return schemas.Config{}, err
	}
	// file exists, do something with it

	// read the file
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return schemas.Config{}, err
	}

	// parse the config data from the json
	var data schemas.Config
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("config.json has invalid form", err)
		return schemas.Config{}, err
	}

	dataValue := reflect.ValueOf(&data).Elem()
	typeOfData := dataValue.Type()

	for i := 0; i < dataValue.NumField(); i++ {
		fieldValue := dataValue.Field(i).Interface()
		fieldName := typeOfData.Field(i).Name

		if fieldValue == "" {
			field := dataValue.Field(i)
			if field.Kind() == reflect.String {
				fmt.Printf("Value for %s is missing, enter value: ", fieldName)
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				value := scanner.Text()
				field.SetString(value)
			}
		}
	}

	return data, nil
}

// This function fetch bids from the Graph
func retrieveBidsFromSubgraph(GRAPH_URL string, BIDDER string) ([]schemas.BidType, error) {

	// the query to fetch bids
	queryJsonData := map[string]string{
		"query": `
		  {
      	bids(where: { bidderAddress: "` + BIDDER + `", status: "WON", validator_not: null, validator_: { phase: VALIDATOR_REGISTERED} }) {
        	id
        	bidderAddress
        	pubKeyIndex
        	validator {
            id
            phase
            ipfsHashForEncryptedValidatorKey
            validatorPubKey
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
