package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func GetAndCheckConfig() (schemas.Config, error) {
    content, err := ioutil.ReadFile("./config.json")
    if err != nil {
        fmt.Println("Error when opening file: ", err)
        return schemas.Config{}, err
    }

    var data schemas.Config
    err = json.Unmarshal(content, &data)
    if err != nil {
        fmt.Println("config.json has invalid form", err)
        return schemas.Config{}, err
    }

    // Use reflection to check for empty fields
    dataValue := reflect.ValueOf(&data).Elem()
    typeOfData := dataValue.Type()

    for i := 0; i < dataValue.NumField(); i++ {
        field := dataValue.Field(i)
        if field.Kind() == reflect.String && field.String() == "" {
            fieldName := typeOfData.Field(i).Name
            return schemas.Config{}, fmt.Errorf("missing value for required field: %s", fieldName)
        }
    }

    return data, nil
}

