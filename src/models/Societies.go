package models

import (
	"github.com/jinzhu/gorm"
)

type Societies struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Details  string
}