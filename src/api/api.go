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

func StartApi() {
	err := godotenv.Load(".env")
	helpers.CheckError(err)
	database.InitialMigration()
	router := gin.Default()
	router.GET("/getHome", controllers.GetHome)
	router.POST("/addUser", controllers.AddUser)
	fmt.Println("App is working on port :8080")
	log.Fatal(router.Run(":8080"))
}
