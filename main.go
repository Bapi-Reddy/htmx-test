package main

import (
	"github.com/Bapi-Reddy/htmx-test/handler"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	app.GET("/", handler.IndexHandler)
	app.Static("/css", "./css")
	app.Start(":3000")
}
