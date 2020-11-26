package main

import (
	"fmt"
	"net/http"
	"modules/pack"
	"github.com/labstack/echo"
)

func init() {
	fmt.Println("init in server.go")
}

func main() {
	pack.Test()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
