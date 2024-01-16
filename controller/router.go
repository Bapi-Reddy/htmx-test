package controller

import (
	"log"

	v1 "github.com/Bapi-Reddy/htmx-test/controller/v1"
	"github.com/labstack/echo"
)

func SetupRouter(
	todoService v1.ToDoService,
	logger *log.Logger,
	router *echo.Echo,
) {
	handler := v1.NewHandler(logger)
	router.GET("/", handler.MakeIndexPage(todoService))
	router.POST("/todos", handler.MakePostTodos(todoService))
	router.PUT("/todos/:id", handler.MakePutTodos(todoService))
	router.DELETE("/todos/:id", handler.MakeDeleteTodo(todoService))
	router.GET("/components", handler.MakeComponent(todoService))
	router.Static("/css", "./css")
}
