package router

import (
	"fiber-student-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	api.Use(logger.New())
	api.Get("/", handlers.GetStudents)
	api.Post("/", handlers.CreateNewStudent)

	api.Get("/:id", handlers.GetStudentByID)

	// basic auth middleware
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin",
			"jose":  "password",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			c.Set("WWW-authenticate", `Basic realm=localhost`)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"err":     true,
				"message": "Please provide the right credentials to continue your operation",
			})
		},
	}))

	api.Patch("/:id", handlers.UpdateStudentData)
	api.Delete("/:id", handlers.DeleteStudent)
}
