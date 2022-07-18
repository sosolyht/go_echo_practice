package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/config"
	"go_echo/handler"
	middleware2 "go_echo/middleware"
)

func main() {
	e := echo.New()
	config.DBConnection()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Route
	e.GET("/", handler.GetBoardList)
	e.GET("/board/:title", handler.GetBoardPathParameter)
	e.POST("/boards", handler.CreateBoard)

	// 유저 생성
	e.POST("/sign-up", handler.SignUp)

	// 로그인
	e.POST("/sign-in", handler.SignIn)

	// Test
	e.GET("/jwt", handler.CheckJWT, middleware2.IsLoggedIn)

	// s3 test
	e.Static("/s3", "templates")
	e.POST("/s3", handler.Upload)

	// Start Server
	e.Logger.Fatal(e.Start(":8000"))
}
