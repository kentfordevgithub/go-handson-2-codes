package main

import (
	"net/http"

	"github.com/kentfordevgithub/go-handson-2-codes/1_go_commands/pack"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!"+pack.String())
	})
	e.Start(":80")
}
