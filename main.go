package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)


// ********************************
// ******* GRAPH TYPES ************
// ********************************


type ResponseType struct {
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
	GRAPH_URL string `json"GRAPH_URL"`
	BIDDER string `json"BIDDER"`
  PRIVATE_KEYS_FILE_LOCATION string `json"PRIVATE_KEYS_FILE_LOCATION"`
  OUTPUT_LOCATION string `json"OUTPUT_LOCATION"`
  PASSWORD string `json"PASSWORD"`
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
	bids := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	fmt.Println(PrettyPrint(bids))

	// TODO: STEP 4: a loop to process each bid 


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
func retrieveBidsFromSubgraph (GRAPH_URL string, BIDDER string) []BidType {

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
	request.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		// TODO: return []
	}

	data, _ := ioutil.ReadAll(response.Body)

	var result ResponseType
	if err := json.Unmarshal(data, &result); err != nil {   // Parse []byte to go struct pointer
    fmt.Println("Can not unmarshal JSON")
		// TODO: return []
	}
	
	return result.Data.Bids

}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}