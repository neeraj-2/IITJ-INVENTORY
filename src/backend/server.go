package backend

import (
	"fmt"
	"log"
	// "time"

	"github.com/joho/godotenv"

	"myurl.com/inventory/controllers"
	"myurl.com/inventory/database"
	"myurl.com/inventory/helpers"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var server = controllers.Server{}

//StartAPI ...
func StartAPI() {

	// Loading env variables
	err := godotenv.Load(".env")
	helpers.CheckError(err)

	// Initialising the Server
	server.DB = database.InitialMigration()
	server.Router = gin.Default()

	// server.Router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://127.0.0.1:3000"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://127.0.0.1:3000"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	// Initialising the Routes
	server.InitializeRoutes()

	//Starting the Server
	fmt.Println("App is working on port :8080")
	log.Fatal(server.Router.Run(":8080"))
}
