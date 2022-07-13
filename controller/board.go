package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"net/http"
)

//Board에 대한 구조체 생성
type RequestBody struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserId  uuid.UUID `json:"user_id"`
}

func CreateBoard(c echo.Context) error {
	a := new(RequestBody)
	db := config.DBConnection()

	post := &model.Board{
		Title:   a.Title,
		Content: a.Content,
		UserId:  a.UserId,
	}

	err := c.Bind(post)
	if err != nil {
		return err
	}

	db.Create(post)
	return c.JSON(http.StatusCreated, post)
}

func GetBoardList(c echo.Context) error {

	db := config.DBConnection()

	result := []model.Board{}
	//err := c.Bind(result)
	//if err != nil {
	//	return err
	//}
	db.Find(&result)
	return c.JSON(http.StatusOK, result)
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
