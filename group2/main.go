package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// 全体に適用
	e := echo.New()
	e.Use(commonFilter)
	e.GET("/hello3", hello3)

	// /helloと/hello2に処理を適用
	g := e.Group("", filter, filter2)
	g.GET("/hello", hello)
	g.GET("/hello2", hello2)

	// /testパス以下に処理を適用
	g2 := e.Group("/test")
	g2.Use(filter2)
	g2.GET("/hello", hello)
	g2.GET("/hello2", hello2)

	e.Logger.Fatal(e.Start(":1111"))
}

func hello(c echo.Context) error {
	fmt.Println("hello")
	return c.String(http.StatusOK, "Hello")
}

func hello2(c echo.Context) error {
	fmt.Println("hello2")
	return c.String(http.StatusOK, "Hello2")
}

func hello3(c echo.Context) error {
	fmt.Println("hello3")
	return c.String(http.StatusOK, "Hello3")
}

func commonFilter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("before common")
		err := next(c)
		fmt.Println("after common")
		return err
	}
}

func filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("before")
		err := next(c)
		fmt.Println("after")
		return err
	}
}

func filter2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("before2")
		err := next(c)
		fmt.Println("after2")
		return err
	}
}
