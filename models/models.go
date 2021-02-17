package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
