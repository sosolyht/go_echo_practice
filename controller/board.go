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
type BoardRequestBody struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserId  uuid.UUID `json:"user_id"`
}

func CreateBoard(c echo.Context) error {
	var binder BoardRequestBody
	err := c.Bind(&binder)
	if err != nil {
		return badRequestResponseWithLog(c, err)
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

func badRequestResponseWithLog(c echo.Context, cause error) error {
	c.Logger().Debug(cause)
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": cause.Error(),
	})
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
