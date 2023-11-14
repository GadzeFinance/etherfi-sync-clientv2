package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

func SaveKeysToFS(output_location string, validatorInfo schemas.ValidatorKeyInfo, bidId string, validatorPublicKey string, nodeAddress string, db *sql.DB) error {

	// Step 1: Create directory and add data to the directory
	if err := createDir(output_location); err != nil {
		return err
	}

	bidPath := filepath.Join(output_location, bidId)
	if err := createDir(bidPath); err != nil {
		return err
	}

	// Passwords
	// Passwords are stored in a non-destructive manner in two places:
	// 			1. ./output/passwords/<bidId>.txt
	// 			2. ./output/<bidId>/pw.txt
	// The first is for user friendly completeness.
	// The second is for teku validator client to read from.
	if err := createFile(filepath.Join(bidPath, "pw.txt"), string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return err
	}
	if err := createFile(output_location + "/passwords/" + bidId + ".txt", string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return err
	}

	// Keystores
	// We also duplicate the keystores with a .json file for teku clieant to read from.
	if err := createFile(filepath.Join(bidPath, string(validatorInfo.KeystoreName)), string(validatorInfo.ValidatorKeyFile)); err != nil {
		return err
	}
	if err := createFile(output_location + "/keys/" + bidId + ".json", string(validatorInfo.ValidatorKeyFile)); err != nil { 
		return err
	}

	// Validator Public key
	if err := createFile(filepath.Join(bidPath, "pubkey.txt"), validatorPublicKey); err != nil {
		return err
	}

	// Withdrawal contract address
	if err := createFile(filepath.Join(bidPath, "node_address.txt"), nodeAddress); err != nil {
		return err
	}

	query := "REPLACE INTO winning_bids (id, pubkey, password, nodeAddress, keystore) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(bidId, validatorPublicKey, string(validatorInfo.ValidatorKeyPassword), nodeAddress, string(validatorInfo.ValidatorKeyFile))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func createDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return err
    }
	// Create keys and passwords directory only if we're creating our output directory
	if path == "output" {
		keysPath := filepath.Join(path, "keys")
		passwordPath := filepath.Join(path, "passwords")

		if err := os.MkdirAll(keysPath, os.ModePerm); err != nil {
			return err
		}

		if err := os.MkdirAll(passwordPath, os.ModePerm); err != nil {
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