package v1

import (
	"net/http"

	"github.com/Bapi-Reddy/htmx-test/templates"
	"github.com/labstack/echo"
)

func (h *Handler) MakePostTodos(todoService ToDoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		text := c.FormValue("add-todo-input")
		if text == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid text")
		}
		todoService.CreateTodo(text)
		todos := todoService.GetTodos()
		component := templates.TodoCardsWithBtn(todos)
		return render(c, component)
	}
}

func (h *Handler) MakePutTodos(todoService ToDoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}

		oldTodo := todoService.GetTodo(id)

		text := c.FormValue("edit-todo-input")
		if text == "" {
			text = oldTodo.Text
		}

		checkedString := c.FormValue("checked")
		var checked bool
		if checkedString == "on" {
			checked = true
		} else {
			checked = false
		}

		todo := todoService.UpdateTodo(id, text, checked)

		component := templates.TodoCard(*todo)
		return render(c, component)
	}
}

func (h *Handler) MakeDeleteTodo(todoService ToDoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}
		todoService.DeleteTodo(id)
		todos := todoService.GetTodos()
		component := templates.TodoCards(todos)
		return render(c, component)

	}
}
