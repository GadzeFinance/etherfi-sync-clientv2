package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {

	// TODO: fetch GRAPH_URL from env
	GRAPH_URL := "https://api.studio.thegraph.com/query/41778/goerli-dressrehearsal-1/0.0.6"

	// TODO: fetch BIDDER from env
	BIDDER := "0xF88866238ecE28A41e050b04360423a5d1181d49"

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
	}

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}