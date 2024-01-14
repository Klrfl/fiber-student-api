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

func CreateNewStudent(c *fiber.Ctx) error {
	c.Accepts("application/json")
	student := new(Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"err": true, "message": "Something wrong with your payload."})
	}

	query := "INSERT INTO students(id, name, major, grade) VALUES($1, $2, $3, $4)"

	_, err := database.DB.Exec(query, uuid.New(), student.Name, student.Major, student.Grade)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":     true,
			"message": "Error when inserting new student data into database",
		})
	}

	return c.JSON(fiber.Map{"err": false, "message": "student successfully created"})
}
