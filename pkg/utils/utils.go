package utils

import (
	"crypto/rand"
	"log"
	"math/big"
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

func GenerateSixDigitCode() int {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	res := int(n.Int64())
	return res
}
