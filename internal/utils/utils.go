package utils

import "encoding/json"

func StructPrettify(data interface{}) string {
	dataJSON, _ := json.MarshalIndent(data, "", "  ")
	return string(dataJSON)
}
