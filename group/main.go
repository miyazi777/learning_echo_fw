package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", hello)
	e.Use(filter)
	e.Logger.Fatal(e.Start(":1111"))
}

func hello(c echo.Context) error {
	fmt.Println("hello")
	return c.String(http.StatusOK, "Hello")
}

func filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("before")
		err := next(c)
		fmt.Println("action")
		return err
	}
}
