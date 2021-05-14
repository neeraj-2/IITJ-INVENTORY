package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Society struct {
	gorm.Model
	Name     string `gorm:"not null"` 
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Details  string `json:"details"`
	Items    []Item
	UUID     uuid.UUID `gorm:"primaryKey"`
}

func CheckSocietyExists(db *gorm.DB, soc Society) (Society, error) {
	var society Society
	err := db.Find(&society, "Email = ? AND Password = ?", soc.Email, soc.Password).Error
	return society, err
}

func (society *Society) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("UUID", uuid)
}
