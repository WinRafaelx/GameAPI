package gameAPI

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name      string
	StudioID  int
	Studio    Studio
	Platforms []Platform `gorm:"many2many:game_platforms;"`
	Price     int
}

type Studio struct {
	ID   uint
	Name string
}

type Platform struct {
	ID    uint
	Name  string
	Games []Game `gorm:"many2many:game_platforms;"`
}

type Input struct {
	Name         string   `json:"name"`
	StudioName   string   `json:"studio_name"`
	PlatformName []string `json:"platform_name"`
	Price        int      `json:"price"`
}
