package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Bapi-Reddy/htmx-test/controller"
	"github.com/Bapi-Reddy/htmx-test/services"

	"github.com/labstack/echo"
)

func main() {
	logger := log.New(os.Stdout, "MyApp ", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./db/todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists todo (id text not null primary key, text text, checked bool);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
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
	app.Start(":3000")
}
