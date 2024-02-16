package router

import (
	"fiber-student-api/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	api := app.Group("api/students")

	api.Get("/", handlers.GetStudents)
	api.Get("/:id", handlers.GetStudentByID)

	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			adminUsername: adminPassword,
		},
		Unauthorized: func(c *fiber.Ctx) error {
			c.Set("WWW-authenticate", `Basic realm=localhost`)

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"err":     true,
				"message": "Please provide the right credentials to continue your operation",
			})
		},
	}))

	api.Post("/", handlers.CreateNewStudent)
	api.Patch("/:id", handlers.UpdateStudentData)
	api.Delete("/:id", handlers.DeleteStudent)
}
