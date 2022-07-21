package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignUpCognito(c echo.Context) error {
	
	return c.JSON(http.StatusOK, "hello")
}
