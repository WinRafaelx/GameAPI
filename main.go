package main

import (
	auth "gorm_prac/authentication"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		return auth.Register(db, c)
	})

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

	db.AutoMigrate(&Game{}, &Studio{}, &Platform{}, &auth.User{})

	app.Listen(":8000")
}
