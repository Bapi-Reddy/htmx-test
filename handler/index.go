package handler

import (
	"github.com/Bapi-Reddy/htmx-test/templates"
	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	wordComp := templates.Index()
	return render(c, wordComp)
}
