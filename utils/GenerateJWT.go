package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("hellogolang")

func GenerateJWT(email string, role string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,

		jwt.MapClaims{
			"email": email,
			"role":  role,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(SecretKey)
	fmt.Println("token string: ", tokenString)

	return tokenString, err
}
