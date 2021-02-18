package database

import (
	"fmt"
	"os"
	"strconv"

	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"

	"github.com/jinzhu/gorm"

	// adding driver for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitialMigration ...
func InitialMigration() *gorm.DB {

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Items{})
	db.AutoMigrate(&models.Societies{})
	db.AutoMigrate(&models.Issued{})
	db.AutoMigrate(&models.Inbound{})
	db.AutoMigrate(&models.Defective{})

	return db
}

//OpenConnectionToDb ...
func OpenConnectionToDb() *gorm.DB {

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)

	return db

}
