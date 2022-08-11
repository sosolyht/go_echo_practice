package cognito

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go_echo/util"
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

	AdminUserDeleteRequest struct {
		Username string `json:"username"`
	}
)

func SignUp(c echo.Context) error {
	var binder SignUpRequest

	godotenv.Load(".env")

	// TODO 에러처리
	// c.Logger().Error(err)
	c.Bind(&binder) // TODO 에러처리

	awsReq := &idp.SignUpInput{
		ClientId: aws.String(os.Getenv("APP_CLIENT_ID")),
		Username: aws.String(binder.Email),
		Password: aws.String(binder.Password),
	}

	_, err := util.CognitoConnect().SignUp(context.TODO(), awsReq)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "에러 처리 해야함",
			// TODO err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
	})
}

func ConfirmSignUp(c echo.Context) error {
	var binder ConfirmSignUpRequest
	c.Bind(&binder)

	awsReq := &idp.ConfirmSignUpInput{
		ClientId:         aws.String(os.Getenv("APP_CLIENT_ID")),
		ConfirmationCode: aws.String(binder.ConfirmCode),
		Username:         aws.String(binder.Email),
	}

	response, err := util.CognitoConnect().ConfirmSignUp(context.TODO(), awsReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "나도 모르겠다 이거는")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"result":  response,
	})

}

func ResendEmail(c echo.Context) error {
	godotenv.Load(".env")
	type Test struct {
		Username string `json:"username"`
	}
	var binder Test
	c.Bind(&binder)

	awsReq := &idp.AdminResetUserPasswordInput{
		UserPoolId:     aws.String("ap-northeast-2_J5YvnthYk"),
		Username:       aws.String(binder.Username),
		ClientMetadata: nil,
	}

	resp, err := util.AdminCognitoConnect().AdminResetUserPassword(context.TODO(), awsReq)
	fmt.Println(resp)
	fmt.Println()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "나도 모르겠다 이거는")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})

}

func AdminUserDelete(c echo.Context) error {
	godotenv.Load(".env")
	var binder AdminUserDeleteRequest
	c.Bind(&binder)
	awsReq := &idp.AdminDeleteUserInput{
		UserPoolId: aws.String(os.Getenv("POOL_ID")),
		Username:   aws.String(binder.Username),
	}

	res, err := util.AdminCognitoConnect().AdminDeleteUser(context.TODO(), awsReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.ResultMetadata)

	return c.JSON(http.StatusOK, echo.Map{
		"message": res,
	})
}
