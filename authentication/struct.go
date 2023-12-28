package authentication

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Gmail    string `gorm:"unique"`
	Password string
}
