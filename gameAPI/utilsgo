package gameAPI

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateGame(db *gorm.DB, c *fiber.Ctx) error {
	input := new(Input)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Check if game already exists
	game := new(Game)
	res := db.Preload("Studio").Preload("Platforms").First(&game, "name = ?", input.Name)
	if res.Error == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Game already exists")
	}

	studio := new(Studio)
	res = db.First(&studio, "name = ?", input.StudioName)
	if res.Error != nil {
		// Create Studio
		studio.Name = input.StudioName
		result := db.Create(studio)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}
	}
	db.First(&studio, "name = ?", input.StudioName)
	id := studio.ID

	// Create Platforms
	var platforms []Platform
	for _, platformName := range input.PlatformName {
		platform := new(Platform)
		platform.Name = platformName
		res = db.First(&platform, "name = ?", platformName)
		if res.Error != nil {
			// Create Platform
			result := db.Create(platform)
			if result.Error != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
			}
		}
		db.First(&platform, "name = ?", platformName)
		platforms = append(platforms, *platform)
	}

	// Create Game
	game = new(Game)
	game.Name = input.Name
	game.StudioID = int(id)
	game.Platforms = platforms
	game.Price = input.Price
	res = db.First(&game, "name = ?", input.Name)
	if res.Error != nil {
		result := db.Create(game)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}
	} else {
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Game successfully created!")
}

func GetGame(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	game := new(Game)
	res := db.Preload("Studio").Preload("Platforms").First(&game, id)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}
	return c.JSON(game)
}

func GetGames(db *gorm.DB, c *fiber.Ctx) error {
	var games []Game
	res := db.Preload("Studio").Preload("Platforms").Find(&games)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}
	return c.JSON(games)
}

func UpdateGame(db *gorm.DB, c *fiber.Ctx) error {
	// Get the game ID from the request parameters
	id := c.Params("id")

	// Retrieve the existing game from the database
	game := new(Game)
	res := db.Preload("Studio").Preload("Platforms").First(&game, id)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	// Parse the request body into an Input struct
	input := new(Input)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Update the game fields based on the request body
	game.Name = input.Name
	game.Price = input.Price

	// Update Studio
	studio := new(Studio)
	res = db.First(&studio, "name = ?", input.StudioName)
	if res.Error != nil {
		// Create Studio if not found
		studio.Name = input.StudioName
		result := db.Create(studio)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
		}
	}
	db.First(&studio, "name = ?", input.StudioName)
	game.StudioID = int(studio.ID)

	// Update or create Platforms
	var platforms []Platform
	for _, platformName := range input.PlatformName {
		platform := new(Platform)
		platform.Name = platformName
		res = db.First(&platform, "name = ?", platformName)
		if res.Error != nil {
			// Create Platform if not found
			result := db.Create(platform)
			if result.Error != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
			}
		}
		db.First(&platform, "name = ?", platformName)
		platforms = append(platforms, *platform)
	}
	game.Platforms = platforms

	// Save the changes to the database
	result := db.Save(&game)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Game successfully updated!")
}

func DeleteGame(db *gorm.DB, c *fiber.Ctx) error {
	// Get the game ID from the request parameters
	id := c.Params("id")

	// Retrieve the existing game from the database
	game := new(Game)
	res := db.Preload("Studio").Preload("Platforms").First(&game, id)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	// Delete the game from the database
	result := db.Delete(&game)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Game successfully deleted!")
}
