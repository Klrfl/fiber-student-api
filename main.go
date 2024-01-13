package main

import (
	"database/sql"
	"fiber-student-api/database"

	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

func main() {
	database.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("halo")
	})

	app.Listen(":8080")
}
