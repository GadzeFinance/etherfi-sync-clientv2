package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
	"os"
	"github.com/robfig/cron"
)

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

	c := cron.New()
	c.AddFunc("0 * * * *", func() { 

		if err := cronjob(config) ; err != nil {
			fmt.Printf("Error executing function: %s\n", err)
			os.Exit(1)
		}
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func cronjob(config Config) error {

	bids, err := retrieveBidsFromSubgraph(config.GRAPH_URL, config.BIDDER)
	if err != nil {
		fmt.Println("Error: ", err);
		return err
	}
	fmt.Println("Hello")
	for _, bid := range bids {
		fmt.Println(bid)
		// validator, pubKeyIndex := bid.Validator, bid.PubKeyIndex
		response, err := fetchFromIPFS(config.IPFS_GATEWAY, bid.Validator.IpfsHashForEncryptedValidatorKey)
		if err != nil {
			fmt.Println("ERROR")
			return err
		}
		fmt.Println(PrettyPrint(*response))

	}	

	return nil
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

	return &ipfsResponse, nil
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}