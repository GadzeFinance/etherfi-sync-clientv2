package main

import (
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
	"time"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
	"github.com/robfig/cron"
	"golang.org/x/crypto/pbkdf2"
)

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := getConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}
	fmt.Println(PrettyPrint(config))

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {

		if err := cronjob(config); err != nil {
			fmt.Printf("Error executing function: %s\n", err)
			os.Exit(1)
		}
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func cronjob(config schemas.Config) error {

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	for _, bid := range bids {

		validator := bid.Validator
		ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey
		IPFSResponse, err := fetchFromIPFS(config.IPFS_GATEWAY, ipfsHashForEncryptedValidatorKey)
		if err != nil {
			return err
		}

		privateKey, err := extractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
		if err != nil {
			return err
		}

		fmt.Println(PrettyPrint(IPFSResponse))
		fmt.Println(bid.Id)
		validatorKey, err := decryptPrivateKeys(privateKey, config.PASSWORD)
		fmt.Println(PrettyPrint(validatorKey))

	}

	return nil
}

func decryptKeyPairJSON(privateKeysJson schemas.KeyStoreFile, password string) () {

	// from hex string to byte array
	iv, err := hex.DecodeString(privateKeysJson.iv)
	if err != nil {
		log.Fatal("cannot decode iv: ", err)
		return err
	}
	salt, err := hex.DecodeString(privateKeysJson.Salt)
	if err != nil {
		log.Fatal("cannot decode salt: ", err)
		return err
	}

	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)


	// TODO: need to implement following
  // const encryptedData = Buffer.from(privateKeysJSON.data, "hex");

  // const decipher = crypto.createDecipheriv("aes-256-cbc", key, iv);
  // const decryptedData = Buffer.concat([
  //   decipher.update(encryptedData),
  //   decipher.final(),
  // ]);
  // let decryptedDataJSON = JSON.parse(decryptedData.toString("utf8"));
  // return decryptedDataJSON;
}


// This function comes from https://gist.github.com/brettscott/2ac58ab7cb1c66e2b4a32d6c1c3908a7#file-aes-256-cbc-go-L64
// Still trying to understand aes-256-cbc decipher
func Decrypt(encrypted string) (string, error) {
	key := []byte(CIPHER_KEY)
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%s", cipherText), nil
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

func decryptPrivateKeys(privateKeys schemas.KeyStoreFile, privKeyPassword string) (*schemas.DecryptedDataJSON, error) {
	iv, err := hex.DecodeString(privateKeys.Iv)
	if err != nil {
		return nil, err
	}
	salt, err := hex.DecodeString(privateKeys.Salt)
	if err != nil {
		return nil, err
	}
	encryptedData, err := hex.DecodeString(privateKeys.Data)
	if err != nil {
		return nil, err
	}

	// TODO: we need to figure out how the last big of the cryptography works 
	key := pbkdf2.Key([]byte(privKeyPassword), salt, 100000, 32, sha256.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptedData := make([]byte, len(encryptedData))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(decryptedData, encryptedData)

	var decryptedDataJSON schemas.DecryptedDataJSON
	err = json.Unmarshal(decryptedData, &decryptedDataJSON)
	if err != nil {
		fmt.Println(err)
		fmt.Println(decryptedDataJSON)
		return nil, err
	}

	fmt.Println(decryptedDataJSON)
	return &decryptedDataJSON, nil

}

func getConfig() (schemas.Config, error) {

	// will read from config.json file which exists in the same directory

	// TODO: Check if the file exists

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

func fetchFromIPFS(gatewayURL string, cid string) (*schemas.IPFSResponseType, error) {

	reqURL := gatewayURL + "/" + cid
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Printf("Unable to create IPFS request")
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

	body, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var ipfsResponse schemas.IPFSResponseType
	if err := json.Unmarshal(body, &ipfsResponse); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return &ipfsResponse, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
