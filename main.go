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

	app := fiber.New()

	app.Post("/games", func(c *fiber.Ctx) error {
		return createGame(db, c)
	})
	app.Get("/games", func(c *fiber.Ctx) error {
		return getGames(db, c)
	})
	app.Get("/games/:id", func(c *fiber.Ctx) error {
		return getGame(db, c)
	})
	app.Put("/games/:id", func(c *fiber.Ctx) error {
		return updateGame(db, c)
	})
	app.Delete("/games/:id", func(c *fiber.Ctx) error {
		return deleteGame(db, c)
	})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Game{}, &Studio{}, &Platform{})

	app.Listen(":8000")
}
