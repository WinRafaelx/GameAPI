package main

import (
	"GameAPI/controller"
	"GameAPI/model"

	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	// Try to connect 5 times because the DB might be slow to start
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("DB is still napping... retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
	app := fiber.New()

	controller := controller.NewGameController(db)
	app.Post("/games", controller.CreateGame)
	app.Get("/games/:id", controller.GetGameByID)
	app.Get("/games", controller.GetAllGames)

	db.AutoMigrate(&model.Game{})

	app.Listen(":8000")
}
