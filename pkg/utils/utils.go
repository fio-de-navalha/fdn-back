package utils

import (
	"crypto/rand"
	"encoding/base64"
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

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
