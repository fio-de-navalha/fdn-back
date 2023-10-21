package utils

import (
	"time"

	"github.com/bytedance/sonic"
)

func StructStringfy(data interface{}) string {
	dataJSON, _ := sonic.Marshal(data)
	return string(dataJSON)
}

func ConvertToGMTMinus3(t time.Time) time.Time {
	gmtMinus3 := time.FixedZone("GMT-3", -3*60*60)
	return t.In(gmtMinus3)
}
