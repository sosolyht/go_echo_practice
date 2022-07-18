package handler

import (
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"net/http"
)

type Media struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Size       string `json:"size"`
	Codec      string `json:"codec"`
	Duration   string `json:"duration"`
	Resolution string `json:"resolution"`
}

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

func Post(c echo.Context) error {
	var binder Media
	err := c.Bind(&binder)
	if err != nil {
		log.Debug(err)
	}

	newMedia := model.Media{
		Title:      binder.Title,
		Size:       binder.Size,
		Codec:      binder.Codec,
		Duration:   binder.Duration,
		Resolution: binder.Resolution,
	}

	err = db.Create(&newMedia).Error
	if err != nil {
		log.Debug(err)
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success",
	})
}

func Update(c echo.Context) error {
	var binder Media
	err := c.Bind(&binder)
	if err != nil {
		log.Debug(err)
	}

	newMedia := model.Media{
		Id:    binder.Id,
		Title: binder.Title,
	}

	db.Select("id", "title").Updates(&newMedia)
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success",
	})
}

func Delete(c echo.Context) error {
	var binder Media
	err := c.Bind(&binder)
	if err != nil {
		log.Debug(err)
	}
	newMedia := model.Media{
		Id: binder.Id,
	}
	db.Delete(&newMedia)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "removed",
	})
}
