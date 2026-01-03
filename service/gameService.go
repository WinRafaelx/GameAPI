package service

import (
	"GameAPI/model"
	"GameAPI/repository"
)

type IGameService interface {
	CreateGame(game *model.Game) error
	GetGameByID(id uint) (*model.Game, error)
	GetAllGames() ([]model.Game, error)
	UpdateGame(id uint, game *model.Game) error
	DeleteGame(id uint) error
}

type GameService struct {
	repo repository.IGameRepository
}

func NewGameService(repo repository.IGameRepository) *GameService {
	return &GameService{repo: repo}
}

func (gs *GameService) CreateGame(game *model.Game) error {
	// Check if the game already exists
	existingGame, err := gs.repo.FindByID(game.ID)
	if err != nil {
		return err
	} else if existingGame != nil {
		return model.ErrGameExists
	}

	// Save the game
	return gs.repo.Save(game)
}

func (gs *GameService) GetGameByID(id uint) (*model.Game, error) {
	return gs.repo.FindByID(id)
}

func (gs *GameService) GetAllGames() ([]model.Game, error) {
	return gs.repo.FindAll()
}

func (gs *GameService) UpdateGame(id uint, game *model.Game) error {
	return gs.repo.Update(id, game)
}

func (gs *GameService) DeleteGame(id uint) error {
	return gs.repo.Delete(id)
}
