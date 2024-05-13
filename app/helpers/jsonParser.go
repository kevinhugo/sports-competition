package helpers

import (
	"encoding/json"
	"fmt"
	"sports-competition/app/logger"
)

func JsonToStruct(jsonData interface{}, structObject any) {
	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		logger.Warning(err)
	}
	json.Unmarshal(jsonString, &structObject)
}

func JsonStringToStruct(jsonString string, structObject any) {
	json.Unmarshal([]byte(jsonString), &structObject)
}

func StructToMap(structData interface{}) map[string]interface{} {
	b, _ := json.Marshal(&structData)
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}

func ToJson(data any) string {
	var jsonStr []byte
	var err error
	switch data.(type) {
	case map[string]interface{}:
		if data.(map[string]interface{}) == nil {
			return ""
		}
	}
	jsonStr, err = json.Marshal(data)
	if err != nil {
		fmt.Println("Error while trying to parse data to JSON.")
		fmt.Println(fmt.Sprintf("Raw Data : %s", data))
		fmt.Println(err)
	}
	return string(jsonStr)
}

func MinifyJSON(jsonData []byte) string {
	var structConverter map[string]interface{}
	var err error = json.Unmarshal(jsonData, &structConverter)
	if err != nil {
		fmt.Println("Error while trying to unmarshal jsonData.")
		fmt.Println(err)
	}
	var minifiedJSON []byte
	minifiedJSON, err = json.Marshal(structConverter)
	if err != nil {
		fmt.Println("Error while trying to marshal minified json data.")
		fmt.Println(err)
	}
	return string(minifiedJSON)
}
