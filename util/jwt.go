package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySecretKey = []byte("dWkrCnyQr2w8e7lmnLkapQtMy0LKDMM4")
var expTime = time.Now().Unix()

func CreateJWT(userid int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", fmt.Errorf("token signed error")
	} else {
		return tk, nil
	}
}

//func ParseToken(myToken string) {
//	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
//		return []byte(mySecretKey), nil
//	})
//	if err != nil {
//		fmt.Println("not good")
//	} else if token.Valid {
//		fmt.Println("good")
//	}
//}
