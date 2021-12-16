package models

import (
	"TodoApp/database"
	"net/http"
	"strings"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(t *testing.T) {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=goTodo port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func TestCreateToDo(t *testing.T) {
	const todo = `{
        Title:"Test",
	}`

	err, _ := http.Post("http://localhost.com/todos", "application/json", strings.NewReader(todo))

	if err != nil {
		t.Error("Couldnt Create To Do")
	}
}

func TestGetTodos(t *testing.T) {
	db := database.DBConn
	var todos []Todo
	response := db.Find(&todos)
	if response == nil {
		t.Error("No Data")
	}
}
