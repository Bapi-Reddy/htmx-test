package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Bapi-Reddy/htmx-test/controller"
	"github.com/Bapi-Reddy/htmx-test/services"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	logger := log.New(os.Stdout, "MyApp ", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists todo (id text not null primary key, text text, checked bool);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		logger.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	todoService := &services.TodoService{
		DB: db,
	}

	app := echo.New()
	controller.SetupRouter(
		todoService,
		logger,
		app,
	)
	app.Start(getPort())
}
