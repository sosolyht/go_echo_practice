package util

//
//import (
//	"context"
//	"github.com/aws/aws-sdk-go-v2/config"
//	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
//)
//
//type CognitoClient struct {
//	AppClientId string
//	Client      *idp.Client
//}
//
//func CognitoConfig() *CognitoClient {
//
//	// 기존의 session 방식에서 변경된 v2 방식 LoadDefaultConfig 방식
//	cfg, err := config.LoadDefaultConfig(context.Background(),
//		config.WithRegion("****"),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	return &CognitoClient{
//		AppClientId: "***",
//		Client:      idp.NewFromConfig(cfg),
//	}
//}
