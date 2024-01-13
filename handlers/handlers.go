package handlers

import (
	"fiber-student-api/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Student struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Major string    `json:"major"`
	Grade int       `json:"grade"`
}

func GetAllStudents(c *fiber.Ctx) error {
	var students []Student

	rows, err := database.DB.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}

	var student Student

	for rows.Next() {
		rows.Scan(&student.Id, &student.Name, &student.Major, &student.Grade)
		students = append(students, student)
	}

	return c.JSON(fiber.Map{
		"err":  false,
		"data": students,
	})
}
