package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" query:"name"`
	Email string `json:"email" query:"email"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getFamily(c echo.Context) error {
	name := c.QueryParam("name")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "name:"+name+",member:"+member)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})
	e.GET("/users/:id", getUser)

	// /family?name=choi&member=family
	e.GET("/family", getFamily)
	e.Logger.Fatal(e.Start(":8000"))
}
