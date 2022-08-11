package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go_echo/util"
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

	AdminUserDeleteRequest struct {
		Username string `json:"username"`
	}
)

func SignUp(c echo.Context) {

	var binder SignUpRequest
	c.Bind(&binder) // TODO 에러처리

	awsReq := &idp.SignUpInput{
		ClientId: aws.String(os.Getenv("APP_CLIENT_ID")),
		Username: aws.String(binder.Email),
		Password: aws.String(binder.Password),
	}

	_, err := util.CognitoConnect().SignUp(context.TODO(), awsReq)
	if err != nil {
		log.Print(err)
	}
}
