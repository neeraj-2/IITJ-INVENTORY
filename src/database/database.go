package database

import (
	"fmt"
	"os"
	"strconv"

	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitialMigration() {

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Items{})
	db.AutoMigrate(&models.Societies{})
	db.AutoMigrate(&models.Issued{})
	db.AutoMigrate(&models.Inbound{})
	db.AutoMigrate(&models.Defective{})
}

func OpenConnectionToDb() *gorm.DB {

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)

	return db

}
