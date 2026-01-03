package repository

import (
	"GameAPI/model"

	"gorm.io/gorm"
)

// IGameRepository defines the "contract" for what a Game librarian does.
type IGameRepository interface {
	Save(game *model.Game) error
	FindByID(id uint) (*model.Game, error)
	FindAll() ([]model.Game, error)
	Update(id uint, data interface{}) error // interface{} allows maps or structs
	Delete(id uint) error
}

type GameRepository struct {
	db *gorm.DB
}

// NewGameRepository is a "Constructor" - it creates the librarian.
func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (gr *GameRepository) Save(game *model.Game) error {
	// Create the record in the DB
	err := gr.db.Create(game).Error
	return err
}

func (gr *GameRepository) FindByID(id uint) (*model.Game, error) {
	var game model.Game
	err := gr.db.First(&game, id).Error
	return &game, err
}

func (gr *GameRepository) FindAll() ([]model.Game, error) {
	var games []model.Game
	err := gr.db.Find(&games).Error
	return games, err
}

func (gr *GameRepository) Update(id uint, data interface{}) error {
	// Using model(&model.Game{}) tells GORM which table to hit.
	return gr.db.Model(&model.Game{}).Where("id = ?", id).Updates(data).Error
}

func (gr *GameRepository) Delete(id uint) error {
	return gr.db.Delete(&model.Game{}, id).Error
}
