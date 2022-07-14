package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/config"
	"go_echo/handler"
)

func main() {
	e := echo.New()
	config.DBConnection()
	e.Use(middleware.Logger())

	// Route
	e.GET("/", handler.GetBoardList)
	e.GET("/board/:title", handler.GetBoardPathParameter)
	e.POST("/boards", handler.CreateBoard)

	// 유저 생성
	e.POST("/sign-up", handler.SignUp)

	// 로그인
	e.POST("/sign-in", handler.SignIn)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
