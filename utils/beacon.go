package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func GetExitStatus(pubKey string) (bool) {
	reqUrl := "https://goerli.beaconcha.in/api/v1/validator/" + pubKey 
	resp, err := http.Get(reqUrl)
	if err == nil {
		var response schemas.BeaconResponse
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal([]byte(data), &response)
		if err != nil {
			fmt.Println("Error in marshalling")
			return false
		}

		if response.Data.Status == "exited" {
			return true
		}
	}
	return false
}