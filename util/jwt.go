package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySecretKey = []byte("dWkrCnyQr2w8e7lmnLkapQtMy0LKDMM4")
var expTime = time.Now().Unix()

func CreateJWT(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = name
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", fmt.Errorf("token signed error")
	} else {
		return tk, nil
	}
}
