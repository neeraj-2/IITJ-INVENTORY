package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Details     string `json:"details"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	Available   int    `gorm:"not null" json:"available"`
	DefectiveId uint32
	SocietyId   uint32
	IssuedId    uint32
	InboundId   uint32
}
