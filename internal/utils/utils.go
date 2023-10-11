package utils

import (
	"encoding/json"
	"time"
)

func StructStringfy(data interface{}) string {
	dataJSON, _ := json.Marshal(data)
	return string(dataJSON)
}

func ConvertToGMTMinus3 (t time.Time) time.Time {
	gmtMinus3 := time.FixedZone("GMT-3", -3*60*60)
	return t.In(gmtMinus3)
}