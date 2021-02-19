package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Inbound struct {
	gorm.Model
	DateOfInbound  time.Time `gorm:"not null"`
	DateOfPurchase time.Time `gorm:"not null"`
	Bill           string    `gorm:"not null"` 
	Society        string    `gorm:"not null"`
	Price          int       `gorm:"not null"`
	Quantity       int       `gorm:"not null"`
	Item Item
}