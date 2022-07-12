package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo/database"
)

func main() {
	database.Connect()
	e := echo.New()
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
