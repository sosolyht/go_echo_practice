package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetBoardList(c echo.Context) error {
	return c.String(http.StatusOK, "Get Board")
}
