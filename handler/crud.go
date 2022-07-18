package handler

import (
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"net/http"
)

var db = config.DBConnection()

func Get(c echo.Context) error {
	result := []model.Media{}
	err := c.Bind(result)
	if err != nil {
		log.Debug(err)
	}
	db.Find(&result)
	return c.JSON(http.StatusOK, result)
}
