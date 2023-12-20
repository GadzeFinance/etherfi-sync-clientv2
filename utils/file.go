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

func FetchFromIPFS(gatewayURL string, cid string) (*schemas.IPFSResponseType, error) {

	reqURL := gatewayURL + "/" + cid
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating IPFS request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 30}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from IPFS: %w", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Reading IPFS response: %w", err)
	}

	var ipfsResponse schemas.IPFSResponseType
	if err := json.Unmarshal(body, &ipfsResponse); err != nil { // Parse []byte to go struct pointer
		return nil, fmt.Errorf("Unmarshalling IPFS response: %w", err)
	}

	return &ipfsResponse, nil
}

func SaveKeysToFS(output_location string, validatorInfo schemas.ValidatorKeyInfo, bidId string, validatorPublicKey string, nodeAddress string, db *sql.DB) error {

	// Step 1: Create directory and add data to the directory
	if err := createDir(output_location); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	bidPath := filepath.Join(output_location, bidId)
	if err := createDir(bidPath); err != nil {
		return fmt.Errorf("creating bid directory: %w", err)
	}

	// Passwords
	// Passwords are stored in a non-destructive manner in two places:
	// 			1. ./output/passwords/<bidId>.txt
	// 			2. ./output/<bidId>/pw.txt
	// The first is for user friendly completeness.
	// The second is for teku validator client to read from.
	if err := createFile(filepath.Join(bidPath, "pw.txt"), string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return fmt.Errorf("creating pw.txt: %w", err)
	}
	if err := createFile(filepath.Join(output_location, "passwords", fmt.Sprintf("%s.txt", bidId)), string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return fmt.Errorf("creating bid.txt: %w", err)
	}

	// Keystores
	// We also duplicate the keystores with a .json file for teku clieant to read from.
	if err := createFile(filepath.Join(bidPath, string(validatorInfo.KeystoreName)), string(validatorInfo.ValidatorKeyFile)); err != nil {
		return fmt.Errorf("creating keystore file: %w", err)
	}
	if err := createFile(filepath.Join(output_location, "keys", fmt.Sprintf("%s.json", bidId)), string(validatorInfo.ValidatorKeyFile)); err != nil {
		return fmt.Errorf("creating %s.json: %w", bidId, err)
	}

	// Validator Public key
	if err := createFile(filepath.Join(bidPath, "pubkey.txt"), validatorPublicKey); err != nil {
		return fmt.Errorf("creating pubkey.txt: %w", err)
	}

	// Withdrawal contract address
	if err := createFile(filepath.Join(bidPath, "node_address.txt"), nodeAddress); err != nil {
		return fmt.Errorf("creating node_address.txt: %w", err)
	}

	query := "REPLACE INTO winning_bids (id, pubkey, password, nodeAddress, keystore) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("creating bid update query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(bidId, validatorPublicKey, string(validatorInfo.ValidatorKeyPassword), nodeAddress, string(validatorInfo.ValidatorKeyFile))
	if err != nil {
		log.Fatalf("storing bids in DB: %v", err)
	}

	return nil
}

func createDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("creating directory '%s': %w", path, err)
	}
	// Create keys and passwords directory only if we're creating our output directory
	if path == "output" {
		keysPath := filepath.Join(path, "keys")
		passwordPath := filepath.Join(path, "passwords")

		if err := os.MkdirAll(keysPath, os.ModePerm); err != nil {
			return fmt.Errorf("creating keys directory: %w", err)
		}

		if err := os.MkdirAll(passwordPath, os.ModePerm); err != nil {
			return fmt.Errorf("creating passwords directory: %w", err)
		}
	}

	return nil
}

func createFile(location string, content string) error {
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		return err
	}
	return ioutil.WriteFile(location, []byte(content), 0644)
}

func ParseKeystoreFile(location string) (*schemas.KeyStoreFile, error) {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatalf("Error when opening file: %v", err)
		return nil, err
	}

	var payload schemas.KeyStoreFile
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("keystore file has invalid form: %v", err)
		return nil, err
	}

	return &payload, nil
}
