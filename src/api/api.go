package api

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"myURL.com/inventory/controllers"
	"myURL.com/inventory/database"
	"myURL.com/inventory/helpers"

	"github.com/gin-gonic/gin"
)

var server = controllers.Server{}

//StartAPI ...
func StartAPI() {

	// Loadind env variables
	err := godotenv.Load(".env")
	helpers.CheckError(err)

	// Intialising the Server
	server.DB = database.InitialMigration()
	server.Router = gin.Default()

	// Intialising the Routes
	server.InitializeRoutes()

	//Starting the Server
	fmt.Println("App is working on port :8080")
	log.Fatal(server.Router.Run(":8080"))
}
