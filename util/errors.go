package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BadRequestResponseWithLog(c echo.Context, cause error) error {
	c.Logger().Debug(cause)
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": cause.Error(),
	})
}
