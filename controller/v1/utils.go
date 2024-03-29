package v1

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
