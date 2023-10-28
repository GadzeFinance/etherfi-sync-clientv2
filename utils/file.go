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
	"strings"
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

func SaveKeysToFS(
	output_location string,
	validatorInfo schemas.ValidatorKeyInfo,
	bidId string,
	validatorPublicKey string,
	nodeAddress string,
	db *sql.DB,
) error {

	// Step 1: Create directory and add data to the directory
	if err := createDir(output_location); err != nil {
		return err
	}

	bidPath := filepath.Join(output_location, bidId)
	if err := createDir(bidPath); err != nil {
		return err
	} else {
		fmt.Println("Created directory: ", bidPath)
	}

	if err := createFile(filepath.Join(bidPath, "pw.txt"), string(validatorInfo.ValidatorKeyPassword)); err != nil {
		return err
	}

	if err := createFile(filepath.Join(bidPath, "pubkey.txt"), validatorPublicKey); err != nil {
		return err
	}

	if err := createFile(filepath.Join(bidPath, "node_address.txt"), nodeAddress); err != nil {
		return err
	}

	if err := createFile(filepath.Join(bidPath, string(validatorInfo.KeystoreName)), string(validatorInfo.ValidatorKeyFile)); err != nil {
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

func AddToTeku(validatorPath string, bidId string, password string, validatorKeyFile string) error {

	passwordFilename := fmt.Sprintf("keystore-%s.txt", bidId)
	if err := createFile(filepath.Join(validatorPath, "passwords", passwordFilename), string(password)); err != nil {
		return err
	}

	keystoreFileName := fmt.Sprintf("keystore-%s.json", bidId)
	if err := createFile(filepath.Join(validatorPath, "keys", keystoreFileName), string(validatorKeyFile)); err != nil {
		return err
	}

	return nil
}

func DeleteFromTeku(validatorPath string, bidId string) error {
	passwordFilename := fmt.Sprintf("keystore-%s.txt", bidId)
	if err := deleteFile(filepath.Join(validatorPath, "passwords", passwordFilename)); err != nil {
		return err
	}

	keystoreFileName := fmt.Sprintf("keystore-%s.json", bidId)
	if err := deleteFile(filepath.Join(validatorPath, "keys", keystoreFileName)); err != nil {
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

func deleteFile(filePath string) error {
	if exists := FileExists(filePath); !exists {
		return nil
	}

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

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

func SaveTekuProposerConfig(validatorPath, pubKey, feeRecipient string) error {

	tekuProposerConfigFile := filepath.Join(validatorPath, "teku_proposer_config.json")
	if exists := FileExists(tekuProposerConfigFile); !exists {
		fmt.Println("teku_proposer_config.json does not exist")
		return nil
	}

	fileContent, err := ioutil.ReadFile(tekuProposerConfigFile)
	if err != nil {
		return err
	}

	var config schemas.Configuration
	err = json.Unmarshal(fileContent, &config)
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return err
	}

	if config.ProposerConfig == nil {
		config.ProposerConfig = make(map[string]schemas.ProposerEntry)
	}

	pubKey = strings.ToLower(strings.TrimSpace(pubKey))
	feeRecipient = strings.ToLower(strings.TrimSpace(feeRecipient))

	config.ProposerConfig[pubKey] = schemas.ProposerEntry{
		FeeRecipient: feeRecipient,
	}

	updatedJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(tekuProposerConfigFile, updatedJSON, 0644); err != nil {
		return err
	}

	return nil
}

func RemoveTekuProposerConfig(validatorPath, pubKey string) error {
	tekuProposerConfigFile := filepath.Join(validatorPath, "teku_proposer_config.json")

	if exists := FileExists(tekuProposerConfigFile); !exists {
		fmt.Println("teku_proposer_config.json does not exist")
		return nil
	}

	fileContent, err := ioutil.ReadFile(tekuProposerConfigFile)
	if err != nil {
		return err
	}

	var config schemas.Configuration
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return err
	}

	pubKey = strings.ToLower(strings.TrimSpace(pubKey))

	if config.ProposerConfig != nil {
		delete(config.ProposerConfig, pubKey)
	}

	updatedJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(tekuProposerConfigFile, updatedJSON, 0644); err != nil {
		return err
	}

	return nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
