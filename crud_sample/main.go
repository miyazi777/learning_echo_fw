package main

import (
	"crud_sample/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ItemForm struct {
	Name string `json:"name" form:"name"`
}

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
	itemForm := new(ItemForm)
	if err := c.Bind(itemForm); err != nil {
		return nil
	}

	item := db.Item{}
	item.Name = itemForm.Name

	repo := db.ItemRepositoryImpl{}
	repo.Insert(&item)

	return c.JSON(http.StatusOK, item)
}

func updateItem(c echo.Context) error {
	repo := db.ItemRepositoryImpl{}

	id := c.Param("id")
	item := repo.FindById(id)
	if item == nil {
		return nil
	}

	newItem := new(db.Item)
	if err := c.Bind(newItem); err != nil {
		return nil
	}

	item.Name = newItem.Name
	repo.Update(item)
	return c.JSON(http.StatusOK, item)
}

func deleteItem(c echo.Context) error {
	id := c.Param("id")

	repo := db.ItemRepositoryImpl{}
	repo.Delete(id)

	return c.JSON(http.StatusOK, id)
}
