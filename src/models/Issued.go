package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Issued struct {
	gorm.Model
	IssueDate time.Time `gorm:"not null"`
	DueDate   time.Time `gorm:"not null"`
	Approved  string    `gorm:"not null"`
	Denied    string    `gorm:"not null"`
	Purpose   string    `gorm:"not null"`
	UserId uint32
	Item Item
}