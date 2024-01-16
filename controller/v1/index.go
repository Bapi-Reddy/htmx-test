package v1

import (
	"github.com/Bapi-Reddy/htmx-test/templates"
	"github.com/labstack/echo"
)

func (h *Handler) MakeIndexPage(todoService ToDoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos := todoService.GetTodos()
		component := templates.Index(todos)
		return render(c, component)
	}
}
