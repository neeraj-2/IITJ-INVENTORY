package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"primaryKey;not null"`
	Password string `gorm:"not null"`
}

type Email struct {
	Email string `json:"email"`
}

func EmailExists(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Find(&user, "Email = ?", email).Error
	return user, err
}
