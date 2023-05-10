package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
		"log"
		"crypto/sha256"

		"golang.org/x/crypto/pbkdf2"
)


// ********************************
// ******* GRAPH TYPES ************
// ********************************

//TODO: Split everything into different files for readibility
type IPFSResponseType struct {
	EncryptedKeystoreName string `json:"encryptedKeystoreName"`
	EncryptedValidatorKey string `json:"encryptedValidatorKey"`
	EncryptedPassword string `json:"encryptedPassword"`
	StakerPublicKey string `json:"stakerPublicKey"`
	NodeOperatorPublicKey string `json:"nodeOperatorPublicKey"`
	EtherfiDesktopAppVersion string `json:"etherfiDesktopAppVersion"`
}

type GQLResponseType struct {
	Data struct {
		Bids []BidType `json:"bids"`
	} `json:"data"`
}

type BidType struct {
	Id string `json:"id"`
	BidderAddress string `json:"bidderAddress"`
  	PubKeyIndex string `json:"pubKeyIndex"`
  	Validator ValidatorType `json:"validator"`
}

type ValidatorType struct {
	Id string `json:"id"`
	Phase string `json:"phase"`
  	IpfsHashForEncryptedValidatorKey string `json:"ipfsHashForEncryptedValidatorKey"`
  	ValidatorPubKey string `json:"validatorPubKey"`        	
}


// ******************************
// ********** TYPES *************
// ******************************


type Config struct {
	GRAPH_URL string `json:"GRAPH_URL"`
	BIDDER string `json:"BIDDER"`
	PRIVATE_KEYS_FILE_LOCATION string `json:"PRIVATE_KEYS_FILE_LOCATION"`
	OUTPUT_LOCATION string `json:"OUTPUT_LOCATION"`
	PASSWORD string `json:"PASSWORD"`
	IPFS_GATEWAY string `json:"IPFS_GATEWAY"`
}

type KeyStoreFile struct {
	Iv string `json:"iv"`
	Salt string `json:"salt"`
	Data string `json:"data"`
	EtherfiDesktopAppVersion string `json:"etherfiDesktopAppVersion"`
}

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := getConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}
	fmt.Println(PrettyPrint(config))

	// STEP 2: extract private keys from file
	privateKeys, err := extractPrivateKeysFromFS(config.PRIVATE_KEYS_FILE_LOCATION)
	if err != nil {
		fmt.Println("Error: ", err);
		return
	}

	// fmt.Println(PrettyPrint(privateKeys))

	// STEP 3: fetch bids from subgraph
	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		fmt.Println("Error: ", err);
		return
	}

	// fmt.Println(PrettyPrint(bids))

	// TODO: STEP 4: a loop to process each bid 

	for _, bid := range bids {
		fmt.Println("> start processing bid with id: " + string(bid.Id))

		validator := bid.Validator
		pubKeyIndex := bid.PubKeyIndex
		ipfsHashForEncryptedValidatorKey := validator.IpfsHashForEncryptedValidatorKey
		validatorPubKey := validator.ValidatorPubKey
		
		IPFSResponse := fetchFromIPFS(ipfsHashForEncryptedValidatorKey)

		// TODO: still need to implement following
		// const validatorKey = decryptKeyPairJSON(privateKeys, PASSWORD)
    // const { pubKeyArray, privKeyArray } = validatorKey
    // const keypairForIndex = getKeyPairByPubKeyIndex(pubKeyIndex, privKeyArray, pubKeyArray)
    // const data = decryptValidatorKeyInfo(file, keypairForIndex)
    // console.log(`creating ${data.keystoreName} for bid:${bid.id}`)
    // createFSBidOutput(OUTPUT_LOCATION, data, bid.id, validatorPubKey)
    // console.log(`< end processing bid with id:${bid.id}`)

	}

}

func decryptKeyPairJSON(privateKeysJson KeyStoreFile, password string) () {

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



func extractPrivateKeysFromFS(location string) (KeyStoreFile, error) {
	content, err := ioutil.ReadFile(location)
  if err != nil {
    log.Fatal("Error when opening file: ", err)
		return KeyStoreFile{}, err
  }
  
	var payload KeyStoreFile
  err = json.Unmarshal(content, &payload)
  if err != nil {
    log.Fatal("keystore file has invalid form: ", err)
		return KeyStoreFile{}, err
	}

	return payload, nil
 
}


func getConfig () (Config, error) {

	// will read from config.json file which exists in the same directory

	// TODO: Check if the file exists

	// read the file
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return Config{}, err
	}

	// parse the config data from the json
	var data Config
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("config.json has invalid form", err)
		return Config{}, err
	}

	return data, nil

}


// This function fetch bids from the Graph
func retrieveBidsFromSubgraph (GRAPH_URL string, BIDDER string) ([]BidType, error) {

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

	var result GQLResponseType
	if err := json.Unmarshal(data, &result); err != nil {   // Parse []byte to go struct pointer
    	fmt.Println("Can not unmarshal JSON")
		return nil, err
	}
	
	return result.Data.Bids, nil
}

func fetchFromIPFS (gatewayURL string, cid string) (*IPFSResponseType, error) {

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

	var ipfsResponse IPFSResponseType
	if err := json.Unmarshal(body, &ipfsResponse); err != nil {   // Parse []byte to go struct pointer
    	fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	fmt.Println(PrettyPrint(ipfsResponse))
	return &ipfsResponse, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}