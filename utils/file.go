package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"path/filepath"
	"os"
	"log"
	"bufio"
	"reflect"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func FetchFromIPFS(gatewayURL string, cid string) (schemas.IPFSResponseType, error) {

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

func SaveKeysToFS(output_location string, consensus_location string, client_location string, validatorInfo schemas.ValidatorKeyInfo, bidId string, validatorPublicKey string) error {

	// Step 1: Create directory and add data to the directory
	if err := createDir(output_location); err != nil {
		return err
	}

	bidPath := filepath.Join(output_location, bidId)
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

func ExtractPrivateKeysFromFS(location string) (schemas.KeyStoreFile, error) {
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

func getConfig() (schemas.Config, error) {

	err := FileExists("./config.json")
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

func FileExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}
	return err
}
