package handler

import (
	"net/http"

	"github.com/Bapi-Reddy/htmx-test/templates/common"
	"github.com/labstack/echo"
)

func ComponentHandler(c echo.Context) error {
	t := c.QueryParam("type")
	switch t {
	case "add-todo":
		component := common.Input("add-todo")
		return render(c, component)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Provide Valid Type")
	}
}
