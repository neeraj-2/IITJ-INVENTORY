package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User Model
type User struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Issued  []Issued
	isAdmin bool      `gorm: "DEFAULT:false"`
	UUID    uuid.UUID `gorm:"primaryKey"`
}

//BeforeCreate for user
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("UUID", uuid)
}
