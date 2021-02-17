package api

import (
	"fmt"
	"log"

	"myURL.com/inventory/controllers"
	"myURL.com/inventory/database"

	"github.com/gin-gonic/gin"
)

func StartApi() {
	database.InitialMigration()
	router := gin.Default()
	router.GET("/getHome", controllers.GetHome)
	router.POST("/addUser", controllers.AddUser)
	fmt.Println("App is working on port :4040")
	log.Fatal(router.Run(":8080"))
}
