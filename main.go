package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

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


func main() {

	// TODO: fetch GRAPH_URL from env
	GRAPH_URL := "https://api.studio.thegraph.com/query/41778/goerli-dressrehearsal-1/0.0.6"

	// TODO: fetch BIDDER from env
	BIDDER := "0xF88866238ecE28A41e050b04360423a5d1181d49"

	bids := retrieveBidsFromSubgraph(GRAPH_URL, BIDDER)

	fmt.Println(PrettyPrint(bids))

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