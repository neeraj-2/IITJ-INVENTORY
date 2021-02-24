package models

import (
	"github.com/jinzhu/gorm"
)

type Society struct {
	gorm.Model
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Details  string `json:"details"`
	Items    []Item
}

func CheckSocietyExists(db *gorm.DB, soc Society) (Society, error) {
	var society Society
	err := db.Find(&society, "Email = ? AND Password = ?", soc.Email, soc.Password).Error
	return society, err
}
