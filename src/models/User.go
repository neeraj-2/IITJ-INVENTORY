package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Issued  []Issued
	isAdmin bool `gorm: "DEFAULT:false"`
}
