package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Item struct {
	gorm.Model
	Details     string `json:"details"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	Available   int    `gorm:"not null" json:"available"`
	DefectiveId uint32
	SocietyId   uuid.UUID
	IssuedId    uint32
	InboundId   uint32
}
