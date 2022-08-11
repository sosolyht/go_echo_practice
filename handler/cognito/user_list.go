package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	idp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go_echo/util"
	"log"
	"net/http"
	"os"
)

func UserList(c echo.Context) error {
	godotenv.Load(".env")
	var pToken *string
	awsReq := &idp.ListUsersInput{
		UserPoolId:      aws.String(os.Getenv("POOL_ID")),
		Filter:          aws.String("email = \"a29405a439aa@drmai.in\""),
		Limit:           nil,
		PaginationToken: pToken,
	}
	a, err := util.AdminCognitoConnect().ListUsers(context.Background(), awsReq)
	if err != nil {
		log.Fatal(err)
	}

	if len(a.Users) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "no users",
		})
	}

	pToken = a.PaginationToken

	return c.JSON(http.StatusOK, a)
}

func UserListGroup(c echo.Context) error {
	godotenv.Load(".env")
	awsReq := &idp.ListUsersInGroupInput{
		GroupName:  aws.String("GOOD"),
		UserPoolId: aws.String(os.Getenv("POOL_ID")),
		Limit:      aws.Int32(1),
		NextToken:  nil,
	}

	a, err := util.AdminCognitoConnect().ListUsersInGroup(context.Background(), awsReq)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, a)
}
