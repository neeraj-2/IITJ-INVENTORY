package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User Model
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string
	Issued   []Issued
	IsAdmin  bool      `gorm:"default:false"`
	UUID     uuid.UUID `gorm:"primaryKey"`
}

//BeforeCreate for user
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("UUID", uuid)
}

//CheckUserExists ...
func CheckUserExists(db *gorm.DB, us User) (User, error) {
	var user User
	err := db.Find(&user, "Email = ? AND Password = ?", us.Email, us.Password).Error
	return user, err
}

//CheckUserExistsFromEmail ...
func CheckUserExistsFromEmail(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Find(&user, "Email = ?", email).Error
	return user, err
}
