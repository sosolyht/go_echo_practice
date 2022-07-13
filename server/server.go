package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/config"
)

func Init() {
	e := echo.New()
	config.DBConnection()
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8000"))

	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))
}
