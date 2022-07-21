package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func CognitoConnect() *idp.Client {
	godotenv.Load(".env")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Debug(err)
	}

	client := idp.NewFromConfig(cfg)
	return client
}

func AdminCognitoConnect() *idp.Client {
	godotenv.Load(".env")
	cfg := aws.Config{
		Region: os.Getenv("AWS_REGION"),
		Credentials: credentials.NewStaticCredentialsProvider(
			os.Getenv("ACCESS_KEY"),
			os.Getenv("SECRET_KEY"),
			"",
		),
	}
	client := idp.NewFromConfig(cfg)
	return client
}

//	staticProvider := credentials.NewStaticCredentialsProvider(
//		os.Getenv("ACCESS_KEY"),
//		os.Getenv("SECRET_KEY"),
//		"")
//	cfg, err := config.LoadDefaultConfig(context.TODO(),
//		config.WithCredentialsProvider(staticProvider),
//	)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	client := idp.NewFromConfig(cfg)
//	return client
//}
