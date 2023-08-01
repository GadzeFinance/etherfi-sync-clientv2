package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func GetConfig(pathToFile string) (schemas.Config, error) {

	if exists := FileExists(pathToFile); !exists {
		return schemas.Config{}, fmt.Errorf("Config file does not exist")
	}

	// read the file
	content, err := ioutil.ReadFile(pathToFile)
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

	dataValue := reflect.ValueOf(&data).Elem()
	typeOfData := dataValue.Type()

	for i := 0; i < dataValue.NumField(); i++ {
		fieldValue := dataValue.Field(i).Interface()
		fieldName := typeOfData.Field(i).Name

		// STAKER is optional
		if fieldName == "STAKER" {
			continue
		}

		if fieldValue == "" && fieldName != "PATH_TO_VALIDATOR" {
			field := dataValue.Field(i)
			if field.Kind() == reflect.String {
				fmt.Printf("Value for %s is missing, enter value: ", fieldName)
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				value := scanner.Text()
				field.SetString(value)
			}
		}
	}

	return data, nil
}
