package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/labstack/gommon/log"
	"os"
)

func CognitoConnect() *idp.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Debug(err)
	}

	client := idp.NewFromConfig(cfg)
	return client
}
