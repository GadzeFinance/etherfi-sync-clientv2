package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func GetExitStatus(pubKey string, beacon_url string) (bool) {
	reqUrl := fmt.Sprintf("%s%s", beacon_url, pubKey)
	resp, err := http.Get(reqUrl)
	if err == nil {
		defer resp.Body.Close()
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
	} else {
		fmt.Println(err)
	}
	return false
}