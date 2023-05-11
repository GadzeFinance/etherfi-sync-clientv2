package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	//"github.com/robfig/cron"
	"golang.org/x/crypto/pbkdf2"
	"path/filepath"
)

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := getConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}
	//fmt.Println(PrettyPrint(config))

	// c := cron.New()
	// c.AddFunc("*/1 * * * *", func() {

	// 	if err := cronjob(config); err != nil {
	// 		fmt.Printf("Error executing function: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// })

	// c.Start()

	// for {
	// 	time.Sleep(time.Second)
	// }

	cronjob(config)
}

func cronjob(config schemas.Config) error {

	privateKey, err := extractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		return err
	}

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	for i, bid := range bids {

		if i >= 1 {
			break
		}

		fmt.Println(`> start processing bid with id:` + bid.Id)

		validator := bid.Validator
		ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey
		
		IPFSResponse, err := fetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
		if err != nil {
			return err
		}

		fmt.Println(PrettyPrint(IPFSResponse))

		validatorKey, err := decryptPrivateKeys(privateKey, config.PASSWORD)
		if err != nil {
			return err
		}
		pubKeyArray := validatorKey.PublicKeys
		privKeyArray := validatorKey.PrivateKeys

		keypairForIndex, err := getKeyPairByPubKeyIndex(bid.PubKeyIndex, pubKeyArray, privKeyArray)

		if err != nil {
			return err
		}

		var data schemas.ValidatorKeyInfo
		
		fmt.Println(PrettyPrint(keypairForIndex))
		if err := saveKeysToFS(config.OUTPUT_LOCATION, config.CONSENSUS_FOLDER_LOCATION, config.ETHERFI_SC_CLIENT_LOCATION, data, bid.Id, validator.ValidatorPubKey); err != nil {
			return err
		}

	}

	return nil
}


func saveKeysToFS(ouput_location string, consensus_location string, client_location string , validatorInfo schemas.ValidatorKeyInfo, bidId string, validatorPublicKey string) error {

	// Step 1: Create directory and add data to the directory
	if err := createDir(ouput_location); err != nil {
		return err
	}

	bidPath := filepath.Join(ouput_location, bidId)
	if err := createDir(bidPath); err != nil {
		return err
	}

	if err := createFile(filepath.Join(bidPath, "pw.txt"), string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return err
	}	

	if err := createFile(filepath.Join(bidPath, "pubkey.txt"), validatorPublicKey); err != nil {
		return err
	}	

	if err := createFile(filepath.Join(bidPath, string(validatorInfo.KeystoreName)), string(validatorInfo.ValidatorKeyFile)); err != nil {
		return err
	}	

	// Step 2: Create an easy to run script (not sure if we need to do this)
	bashHeader := "#!/bin/bash -xe \n"
	echoLine := fmt.Sprintf("echo \"Adding keystore to prysm for validator with pubkey:%s ...\" \n", validatorPublicKey[:10])
	changeDirLine := fmt.Sprintf("cd %s \n", consensus_location)
	keysDir := filepath.Join(client_location, "storage", "output", bidId, string(validatorInfo.KeystoreName))

	prysmCommand := fmt.Sprintf("sudo ./prysm.sh validator accounts import --goerli --wallet-dir=%s --keys-dir=%s", consensus_location, keysDir)
	scriptContent := fmt.Sprintf("%s %s %s %s", bashHeader, echoLine, changeDirLine, prysmCommand)
	if err := createFile(filepath.Join(bidPath, "add.sh"), scriptContent); err != nil {
		return err
	}	

	return nil
}

func createDir(location string) error {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.Mkdir(location, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func createFile(location string, content string) error {
    if _, err := os.Stat(location); !os.IsNotExist(err) {
        return err
    }
    ioutil.WriteFile(location, []byte(content), 0644)
	return nil
}

func getKeyPairByPubKeyIndex(pubkeyIndexString string, privateKeys []string, publicKeys []string) (schemas.KeyPair, error) {
	//fmt.Println("index:", pubkeyIndexString)
	index, err := strconv.ParseInt(pubkeyIndexString, 10, 0)
	if err != nil {
		panic(err)
		return schemas.KeyPair{}, err
	}
	return schemas.KeyPair {
		PrivateKey: privateKeys[index],
		PublicKey: publicKeys[index],
	}, nil
}


func extractPrivateKeysFromFS(location string) (schemas.KeyStoreFile, error) {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return schemas.KeyStoreFile{}, err
	}

	var payload schemas.KeyStoreFile
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("keystore file has invalid form: ", err)
		return schemas.KeyStoreFile{}, err
	}

	return payload, nil
}


func decryptPrivateKeys(privateKeys schemas.KeyStoreFile, privKeyPassword string) (schemas.DecryptedDataJSON, error) {
	iv, err := hex.DecodeString(privateKeys.Iv)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}
	salt, err := hex.DecodeString(privateKeys.Salt)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}
	ciphertext, err := hex.DecodeString(privateKeys.Data)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	// TODO: we need to figure out how the last big of the cryptography works 
	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// TODO: didn't handle error from PKCS5UnPadding for now, maybe use same function from some packages
	decryptedData := PKCS5UnPadding(ciphertext)

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(decryptedData, &decryptedDataJSON)
	if err != nil {
		panic(err)
		return schemas.DecryptedDataJSON{}, err
	}

	// fmt.Println(PrettyPrint(decryptedDataJSON))
	
	return decryptedDataJSON, nil

}

func getConfig() (schemas.Config, error) {

	err := fileExists("config.json")
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

func fetchFromIPFS(gatewayURL string, cid string) (schemas.IPFSResponseType, error) {

	reqURL := gatewayURL + "/" + cid
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Printf("Unable to create IPFS request")
		return schemas.IPFSResponseType{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		// TODO: return []
		return schemas.IPFSResponseType{}, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		return schemas.IPFSResponseType{}, err
	}

	var ipfsResponse schemas.IPFSResponseType
	if err := json.Unmarshal(body, &ipfsResponse); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return schemas.IPFSResponseType{}, err
	}

	return ipfsResponse, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func fileExists(filename string) error {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return fmt.Errorf("file %s does not exist", filename)
    }
    return err
}
