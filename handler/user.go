package handler

import (
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"go_echo/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserRequestBody struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success",
	})
}

func SignIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		util.BadRequestResponseWithLog(c, err)
	}

	plainPassword := []byte(user.Password)

	db := config.DBConnection()
	result := db.Find(&user, "username=?", user.Username)

	// len(result) == 0 으로도 비슷하게 에러 핸들링 가능할듯??
	if result.RowsAffected == 0 {
		return echo.ErrBadRequest
	}

	checkHashed := bcrypt.CompareHashAndPassword([]byte(user.Password), plainPassword)
	// TODO: 체크 했으면 그 다음에 토큰 생성으로
	// if 체크해시드 그리고 토큰 생성

	if checkHashed != nil {
		return echo.ErrUnauthorized
	}

	genToken, err := util.CreateJWT(user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal Server Error",
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = genToken
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"token":   genToken,
	})
}
