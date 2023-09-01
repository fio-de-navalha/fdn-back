package cryptography

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
