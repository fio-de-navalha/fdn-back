package utils

import "encoding/json"

func StructStringfy(data interface{}) string {
	dataJSON, _ := json.Marshal(data)
	return string(dataJSON)
}
