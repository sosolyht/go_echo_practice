package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserRequestBody struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func SignUp(c echo.Context) error {
	var binder UserRequestBody
	err := c.Bind(&binder)
	if err != nil {
		c.Logger().Debug(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}
	db := config.DBConnection()

	password := []byte(binder.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	post := model.User{
		Id:       binder.Id,
		Username: binder.Username,
		Password: string(hashedPassword),
	}

	db.Create(&post)
	return c.JSON(http.StatusCreated, post)
}
