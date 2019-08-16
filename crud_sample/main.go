package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/items", getItems)
	e.POST("/item", createItem)
	e.PUT("/item/:id", updateItem)
	e.DELETE("/item/:id", deleteItem)
	e.Logger.Fatal(e.Start(":1111"))
}

func getItems(c echo.Context) error {
	fmt.Println("items")
	return c.String(http.StatusOK, "items")
}

func createItem(c echo.Context) error {
	fmt.Println("create item")
	return c.String(http.StatusOK, "item")
}

func updateItem(c echo.Context) error {
	fmt.Println("update item")
	return c.String(http.StatusOK, "item")
}

func deleteItem(c echo.Context) error {
	fmt.Println("delete item")
	return c.String(http.StatusOK, "item")
}
