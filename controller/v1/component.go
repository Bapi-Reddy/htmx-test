package v1

import (
	"net/http"

	"github.com/Bapi-Reddy/htmx-test/templates"
	"github.com/labstack/echo"
)

func (h *Handler) MakeComponent(todoService ToDoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.QueryParam("type")
		id := c.QueryParam("id")
		switch t {
		case "add-todo":
			component := templates.AddTodoInput()
			return render(c, component)
		case "add-todo-btn":
			component := templates.AddTodoButton()
			return render(c, component)
		case "edit-todo-input":
			todo := todoService.GetTodo(id)
			component := templates.EditTodoInput(todo)
			return render(c, component)
		case "edit-todo-btn":
			todo := todoService.GetTodo(id)
			component := templates.TodoCard(*todo)
			return render(c, component)
		default:
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid element")
		}
	}
}
