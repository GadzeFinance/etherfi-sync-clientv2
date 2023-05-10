package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
		"log"
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

func main() {

	// STEP 1: fetch env variables from json/.env file
	// NOTE: I'm using json now, but easy to switch
	config, err := getConfig()
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}
	fmt.Println(PrettyPrint(config))

	// TODO: STEP 2: extract private keys from file


	// STEP 3: fetch bids from subgraph
	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		fmt.Println("Error: ", err);
		return
	}

	// fmt.Println(PrettyPrint(bids))

	// TODO: STEP 4: a loop to process each bid 

	}	

	}


  for (const bid of bids) {
    console.log(`> start processing bid with id:${bid.id}`)
    const { validator, pubKeyIndex } = bid
    const { ipfsHashForEncryptedValidatorKey, validatorPubKey } = validator
    const file = await fetchFromIpfs(ipfsHashForEncryptedValidatorKey)
    const validatorKey = decryptKeyPairJSON(privateKeys, PASSWORD)
    const { pubKeyArray, privKeyArray } = validatorKey
    const keypairForIndex = getKeyPairByPubKeyIndex(pubKeyIndex, privKeyArray, pubKeyArray)
    const data = decryptValidatorKeyInfo(file, keypairForIndex)
    console.log(`creating ${data.keystoreName} for bid:${bid.id}`)
    createFSBidOutput(OUTPUT_LOCATION, data, bid.id, validatorPubKey)
    console.log(`< end processing bid with id:${bid.id}`)
  }

}

func fetchFromIpfs (cid string, IPFS_GATEWAY string) {
	url := IPFS_GATEWAY + "/" + cid
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
}

export const fetchFromIpfs = async (cid) => {
	const config = getConfig()
	const url = `${config.IPFS_GATEWAY}/${cid}`
	try {
			const resp = await fetch(url)
			const respJSON = await resp.json()
			return respJSON
	} catch (error) {
			console.error(error)
			return undefined
	}
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