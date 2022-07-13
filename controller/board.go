package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"net/http"
)

func GetBoardList() echo.HandlerFunc {
	return func(c echo.Context) error {

		db := config.DBConnection()

		result := []model.Board{}
		//err := c.Bind(result)
		//if err != nil {
		//	return err
		//}
		db.Find(&result)
		return c.JSON(http.StatusOK, result)
	}
}

func BoardTitlePathParameter() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.Param("title")
		db := config.DBConnection()

		boards := model.Board{}
		err := c.Bind(boards)
		if err != nil {
			return err
		}
		db.Find(&boards, "title = ?", title)
		fmt.Println(boards)
		return c.JSON(http.StatusOK, boards)
	}
}

func BoardQueryParameter() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("title")
		db := config.DBConnection()

		boards := model.Board{}
		err := c.Bind(boards)
		if err != nil {
			return err
		}
		db.Find(&boards, "title = ?", title)
		return c.JSON(http.StatusOK, boards)
	}
}

//func GetBoard() ([]model.Board, error) {
//	db := config.GetDB()
//	boards := []model.Board{}
//
//	if err := db.Find(&boards).Error; err != nil {
//		return nil, err
//	}
//
//	return boards, nil
//}
