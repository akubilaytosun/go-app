package main

import (
	"TodoApp/database"
	"TodoApp/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=goTodo port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	database.DBConn.AutoMigrate(&models.Todo{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateToDo)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	setupRoutes(app)
	app.Listen(":8000")
}
