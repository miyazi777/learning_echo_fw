package main

import (
	"crud_sample/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	db.InitDb()

	e := echo.New()
	Routing(e)
	e.Logger.Fatal(e.Start(":2222"))
}

func Routing(e *echo.Echo) {
	e.GET("/items", getItems)
	e.GET("/item/:id", getItem)
	e.POST("/item", createItem)
	e.PUT("/item/:id", updateItem)
	e.DELETE("/item/:id", deleteItem)
}

func getItems(c echo.Context) error {
	repo := db.ItemRepositoryImpl{}

	items := repo.GetList()
	for _, item := range *items {
		fmt.Println(item)
	}
	return c.JSON(http.StatusOK, items)
}

func getItem(c echo.Context) error {
	repo := db.ItemRepositoryImpl{}

	id := c.Param("id")
	item := repo.FindById(id)
	return c.JSON(http.StatusOK, item)
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
