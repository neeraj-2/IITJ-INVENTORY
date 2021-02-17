package database

import (
	"fmt"

	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 3000
	user     = "postgres"
	password = "Tarun@2001"
	dbname   = "test"
)

func InitialMigration() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}

func OpenConnectionToDb() *gorm.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlconn)
	helpers.CheckError(err)

	return db

}
