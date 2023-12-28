package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL thresulthold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Book{})

	app := fiber.New()

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		return getBook(db, c)
	})
	app.Get("/books", func(c *fiber.Ctx) error {
		return getBooks(db, c)
	})
	app.Post("/books", func(c *fiber.Ctx) error {
		return createBook(db, c)
	})
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		return updateBook(db, c)
	})
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		return deleteBook(db, c)
	})

	log.Fatal(app.Listen(":8000"))
}
