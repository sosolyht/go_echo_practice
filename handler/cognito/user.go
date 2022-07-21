package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	ConfirmSignUpRequest struct {
		ConfirmCode string `json:"confirm_code"`
		Email       string `json:"email"`
	}
)

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

func ConfirmSignUp(c echo.Context) error {
	var binder ConfirmSignUpRequest
	c.Bind(&binder)

	// LoadConfig 부분 따로 빼놓아야함

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Debug(err)
	}

	client := idp.NewFromConfig(cfg)

	awsReq := &idp.ConfirmSignUpInput{
		ClientId:         aws.String(os.Getenv("APP_CLIENT_ID")),
		ConfirmationCode: aws.String(binder.ConfirmCode),
		Username:         aws.String(binder.Email),
	}

	response, err := client.ConfirmSignUp(context.Background(), awsReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "나도 모르겠다 이거는")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"result":  response,
	})

}
