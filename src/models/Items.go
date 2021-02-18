package models

import (
	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	ItemKey  string `gorm:"primaryKey;not null"` // identifier or ItemKey
	Society  string `gorm:"not null"`            // society or society key
	Details  string
	Quantity int 	`gorm:"not null"`
	Available int 	`gorm:"not null"`
}