package database

import (
	"fmt"
	"os"

	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)






func InitialMigration() {
	
	host     := os.Getenv("HOST")
	port     := 5432
	user     := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname   := os.Getenv("DBNAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
	
	host     := os.Getenv("HOST")
	port     := 5432
	user     := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname   := os.Getenv("DBNAME")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)

	return db

}
