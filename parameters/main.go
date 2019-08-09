package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/item/:id", getItem) // path parameter
	e.GET("/items", getItems)   // query parameter
	e.POST("/item", createItem) // form parameter
	e.Logger.Fatal(e.Start(":1111"))
}

func getItem(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id = "+id)
}

func getItems(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "name = "+name)
}

func createItem(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, "create name = "+name)
}
