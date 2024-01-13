package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env file")
	}

	app := fiber.New()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to database")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("halo")
	})

	app.Listen(":8080")
}
