package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Item struct {
	gorm.Model
	SocietyName string `gorm:"not null"`
	Name 		string `gorm:"not null"`
	Details     string `json:"details"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	Available   int    `gorm:"not null" json:"available"`
	DefectiveId uint
	SocietyId   uuid.UUID
	IssuedId    uint
	InboundId   uint 
}
