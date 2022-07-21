package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c echo.Context) error {
	var binder SignUpRequest
	//var cognitoClient util.CognitoClient

	godotenv.Load(".env")

	// TODO 에러처리
	// c.Logger().Error(err)
	c.Bind(&binder) // TODO 에러처리

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		panic(err)
	}

	client := idp.NewFromConfig(cfg)

	awsReq := &idp.SignUpInput{
		ClientId: aws.String(os.Getenv("APP_CLIENT_ID")),
		Username: aws.String(binder.Email),
		Password: aws.String(binder.Password),
	}

	response, err := client.SignUp(context.Background(), awsReq)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "에러 처리 해야함",
			// TODO err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"result":  response,
	})
}
