package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo/config"
	"go_echo/model"
	"go_echo/util"
	"net/http"
)

//Board에 대한 구조체 생성
type BoardRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

func CreateBoard(c echo.Context) error {
	var binder BoardRequestBody
	err := c.Bind(&binder)
	if err != nil {
		return util.BadRequestResponseWithLog(c, err)
	}
	db := config.DBConnection()

	newBoard := model.Board{
		Title:   binder.Title,
		Content: binder.Content,
		UserId:  binder.UserId,
	}
	db.Create(&newBoard)
	return c.JSON(http.StatusCreated, newBoard)
}

func GetBoardList(c echo.Context) error {
	db := config.DBConnection()

	result := []model.Board{}
	err := c.Bind(result)
	if err != nil {
		return err
	}
	db.Find(&result)
	return c.JSON(http.StatusOK, result)
}

func GetBoardPathParameter(c echo.Context) error {
	db := config.DBConnection()
	boards := []model.Board{}
	title := c.Param("title")
	db.Select([]string{
		"id",
		"title",
		"content",
	}).Find(&boards, "Title = ?", title)
	fmt.Println(len(boards))
	if len(boards) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "content not found",
		})
	}
	return c.JSON(http.StatusOK, boards)
}
