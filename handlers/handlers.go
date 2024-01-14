package handlers

import (
	"database/sql"
	"fiber-student-api/database"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Student struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Major string    `json:"major"`
	Grade int       `json:"grade"`
}

func GetStudents(c *fiber.Ctx) error {
	students := []Student{}

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

func GetStudentByID(c *fiber.Ctx) error {
	studentId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "message": "error processing ID"})
	}

	var student Student
	row := database.DB.QueryRow("SELECT * FROM students WHERE id=$1", studentId)

	switch err = row.Scan(&student.Id, &student.Name, &student.Major, &student.Grade); err {
	case sql.ErrNoRows:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"err": false, "message": fmt.Sprintf("Student with ID %s not found", studentId)})
	case nil:
		return c.JSON(fiber.Map{"err": false, "data": student})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": true, "message": "sorry we didn't know what happened"})
	}
}

func CreateNewStudent(c *fiber.Ctx) error {
	c.Accepts("application/json")
	student := new(Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "message": "Something wrong with your payload."})
	}

	query := "INSERT INTO students(id, name, major, grade) VALUES($1, $2, $3, $4)"

	_, err := database.DB.Exec(query, uuid.New(), student.Name, student.Major, student.Grade)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":     true,
			"message": "Error when inserting new student data into database",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"err": false, "message": "student successfully created"})
}

func UpdateStudentData(c *fiber.Ctx) error {
	c.Accepts("application/json")
	studentId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": true, "message": "error processing id"})
	}
	var student Student
	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "messsage": "something wrong with your payload."})
	}

	// check request body params one by one
	// surely there's a better way to do this
	// source: https://stackoverflow.com/questions/38206479/golang-rest-patch-and-building-an-update-query
	query := `UPDATE students SET `
	queryParts := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)

	if student.Name != "" {
		queryParts = append(queryParts, `name = '%s'`)
		args = append(args, student.Name)
	}
	if student.Major != "" {
		queryParts = append(queryParts, `major = '%s'`)
		args = append(args, student.Major)
	}
	if student.Grade != 0 {
		queryParts = append(queryParts, `grade = %d`)
		args = append(args, student.Grade)
	}

	query += strings.Join(queryParts, ", ") + ` WHERE id = '%s'`
	args = append(args, studentId)

	_, err = database.DB.Exec(fmt.Sprintf(query, args...))

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":     true,
			"message": "Error when updating student data",
		})
	}

	return c.JSON(fiber.Map{"err": false, "message": "Student successfully updated"})
}

func DeleteStudent(c *fiber.Ctx) error {
	studentId, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": false, "message": "cannot process student ID"})
	}

	_, err = database.DB.Exec("DELETE FROM students WHERE id=$1", studentId)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": true, "message": "Error when deleting student from database"})
	}

	return c.JSON(fiber.Map{"err": false, "message": "Student successfully deleted"})
}
