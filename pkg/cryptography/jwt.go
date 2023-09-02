package cryptography

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenMetada struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	Exp  int    `json:"exp"`
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userId string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
