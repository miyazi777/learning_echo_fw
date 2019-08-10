package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name string `json:"name" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/", test)
	e.Logger.Fatal(e.Start(":1113"))
}

func test(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, &Error{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
