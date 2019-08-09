package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", headerTest)
	e.Logger.Fatal(e.Start(":1111"))
}

func headerTest(c echo.Context) error {
	header := c.Request().Header
	ua := header.Get("User-Agent")
	contentType := header.Get("ContentType")
	return c.String(http.StatusOK, "user agent = "+ua+" content type = "+contentType)
}
