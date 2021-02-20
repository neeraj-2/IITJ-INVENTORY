package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// defective items
type Defective struct {
	gorm.Model
	DateOfDest time.Time `gorm:"not null"`
	Quantity   string    `gorm:"not null"`
	Item Item
}