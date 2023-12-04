package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/GadzeFinance/etherfi-sync-clientv2/schemas"
)

func GetAndCheckConfig(path string) (schemas.Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return schemas.Config{}, fmt.Errorf("reading config: %w", err)
	}

	var data schemas.Config
	err = json.Unmarshal(content, &data)
	if err != nil {
		return schemas.Config{}, fmt.Errorf("invalid config format: %w", err)
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
