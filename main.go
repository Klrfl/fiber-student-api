package main

import (
	"database/sql"
	"fiber-student-api/database"
	"fiber-student-api/router"

	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

func main() {
	database.Init()

	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":8080")
}
