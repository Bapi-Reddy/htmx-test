package v1

import (
	"github.com/Bapi-Reddy/htmx-test/model"
)

type ToDoService interface {
	GetTodos() []*model.TodoCard
	GetTodo(id string) *model.TodoCard
	CreateTodo(text string) *model.TodoCard
	UpdateTodo(id string, text string, checked bool) *model.TodoCard
	DeleteTodo(id string) error
}
