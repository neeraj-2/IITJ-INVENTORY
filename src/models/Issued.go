package models

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

type Issued struct {
	gorm.Model
	IssueDate time.Time `gorm:"not null"`
	DueDate   time.Time `gorm:"not null"`
	Approved  string    `gorm:"not null"`
	Denied    string    `gorm:"not null"`
	Purpose   string    `gorm:"not null"`
	UserId uuid.UUID
	Name 		string `gorm:"not null"`
	Details     string `json:"details"`
	SocietyName     string    `gorm:"not null"`
}