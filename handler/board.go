package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
		log.Debug(err)
		return util.BadRequestResponseWithLog(c, err)
	}

	newBoard := model.Board{
		Title:   binder.Title,
		Content: binder.Content,
		UserId:  binder.UserId,
	}
	err = db.Create(&newBoard).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}
	log.Debug(err)
	return c.JSON(http.StatusCreated, newBoard)
}

func GetBoardList(c echo.Context) error {

	result := []model.Board{}
	err := c.Bind(result)
	if err != nil {
		log.Debug(err)
		return err
	}
	db.Find(&result)
	return c.JSON(http.StatusOK, result)
}

func GetBoardPathParameter(c echo.Context) error {
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
