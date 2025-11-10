package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "secretString"

func GenerateTokens(email, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		id:    id,
		email: email,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func VarifyTokens(token string) {

}
