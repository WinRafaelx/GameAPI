package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second, // Slow SQL thresulthold
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // Disable color
	},
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
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
