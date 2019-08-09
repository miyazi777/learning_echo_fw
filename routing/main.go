package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users", getUsers)
	e.POST("/user", createUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1111"))
}

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "GET Method")
}

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "POST Method")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "PUT Method")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE Method")
}
