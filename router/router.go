package router

import (
	"fiber-student-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	api.Get("/", handlers.GetAllStudents)
	api.Post("/", handlers.CreateNewStudent)

	api.Get("/:id", handlers.GetStudentByID)
	api.Patch("/:id", handlers.UpdateStudentData)
	api.Delete("/:id", handlers.DeleteStudent)
}
