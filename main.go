package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/config"
	"go_echo/handler"
	"go_echo/handler/cognito"
	mw "go_echo/middleware"
)

func main() {
	e := echo.New()
	config.DBConnection()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Board Route
	e.GET("/", handler.GetBoardList)
	e.GET("/board/:title", handler.GetBoardPathParameter)
	e.POST("/boards", handler.CreateBoard)

	// CRUD Route
	e.GET("/crud", handler.Get)
	e.POST("/crud", handler.Post)
	e.PUT("/crud", handler.Update)
	e.DELETE("/crud", handler.Delete)

	// 유저 생성
	e.POST("/sign-up", handler.SignUp)

	// 로그인
	e.POST("/sign-in", handler.SignIn)

	// JWT Test
	e.GET("/jwt", handler.CheckJWT, mw.IsLoggedIn)

	// S3 test
	e.Static("/s3", "templates")
	e.POST("/s3", handler.Upload)

	// Cognito test
	e.POST("/cognito/signup", cognito.SignUp)

	// Start Server
	e.Logger.Fatal(e.Start(":8000"))
}
