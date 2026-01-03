package controller

import (
	"GameAPI/model"
	"GameAPI/repository"
	"GameAPI/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IGameController interface {
	CreateGame(c *fiber.Ctx) error
	GetGameByID(c *fiber.Ctx) error
	GetAllGames(c *fiber.Ctx) error
}

type GameController struct {
	svc service.IGameService
}

func NewGameController(db *gorm.DB) *GameController {
	return &GameController{
		svc: service.NewGameService(repository.NewGameRepository(db)),
	}
}

func (gc *GameController) CreateGame(c *fiber.Ctx) error {
	input := new(model.GameInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	game := model.Game{
		Name:  input.Name,
		Price: input.Price,
	}

	if err := gc.svc.CreateGame(&game); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(game)
}

func (gc *GameController) GetGameByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	game, err := gc.svc.GetGameByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(game)
}

func (gc *GameController) GetAllGames(c *fiber.Ctx) error {
	games, err := gc.svc.GetAllGames()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(games)
}
