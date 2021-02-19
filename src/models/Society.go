package models

import (
	"github.com/jinzhu/gorm"
)

type Society struct {
	gorm.Model
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Details  string	`json:"details"`
	Items []Item
}