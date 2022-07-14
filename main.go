package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/config"
	"go_echo/controller"
)

func main() {
	e := echo.New()
	config.DBConnection()
	e.Use(middleware.Logger())

	// Route
	e.GET("/", controller.GetBoardList)
	e.GET("/board/:title", controller.GetBoardPathParameter)
	e.POST("/boards", controller.CreateBoard)

	// 유저 생성
	e.POST("/sign-up", controller.SignUp)

	// Start Server
	e.Logger.Fatal(e.Start(":8000"))
}
