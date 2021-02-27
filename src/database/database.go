package database

import (
	"fmt"
	"os"
	"strconv"

	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"

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

	debug := os.Getenv("DEBUG")
	if debug == "true" {
		// Drop Tables if exists
		db.DropTableIfExists(&models.User{})
		db.DropTableIfExists(&models.Item{})
		db.DropTableIfExists(&models.Society{})
		db.DropTableIfExists(&models.Issued{})
		db.DropTableIfExists(&models.Inbound{})
		db.DropTableIfExists(&models.Defective{})

		// Migrate the schema
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Item{})
		db.AutoMigrate(&models.Society{})
		db.AutoMigrate(&models.Issued{})
		db.AutoMigrate(&models.Inbound{})
		db.AutoMigrate(&models.Defective{})

		helpers.InsertDummyData(db)

	} else {
		// Migrate the schema
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Item{})
		db.AutoMigrate(&models.Society{})
		db.AutoMigrate(&models.Issued{})
		db.AutoMigrate(&models.Inbound{})
		db.AutoMigrate(&models.Defective{})
	}

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
