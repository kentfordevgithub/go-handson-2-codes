/*
go get github.com/labstack/echo
*/

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":80"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
