package main

import (
	auth "gorm_prac/authentication"
	game "gorm_prac/gameAPI"

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
	app.Post("/login", func(c *fiber.Ctx) error {
		return auth.Login(db, c)
	})

	app.Use("/games", auth.AuthRequired)

	app.Post("/games", func(c *fiber.Ctx) error {
		return game.CreateGame(db, c)
	})
	app.Get("/games", func(c *fiber.Ctx) error {
		return game.GetGames(db, c)
	})
	app.Get("/games/:id", func(c *fiber.Ctx) error {
		return game.GetGame(db, c)
	})
	app.Put("/games/:id", func(c *fiber.Ctx) error {
		return game.UpdateGame(db, c)
	})
	app.Delete("/games/:id", func(c *fiber.Ctx) error {
		return game.DeleteGame(db, c)
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&game.Game{}, &game.Studio{}, &game.Platform{}, &auth.User{})

	app.Listen(":8000")
}
