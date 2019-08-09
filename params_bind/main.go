package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Item struct {
	Name  string `json:"name" form:"name" query:"name"`
	Price int    `json:"price" form:"price" query:"price"`
}

func main() {
	e := echo.New()
	e.GET("/items", getItems)
	e.POST("/item", createItem)
	e.Logger.Fatal(e.Start(":1111"))
}

func getItems(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return nil
	}
	return c.String(http.StatusOK, "name = "+item.Name+" price = "+strconv.Itoa(item.Price))
}

func createItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return nil
	}
	return c.String(http.StatusOK, "name = "+item.Name+" price = "+strconv.Itoa(item.Price))
}
