package api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"

	"myURL.com/inventory/controllers"
	"myURL.com/inventory/database"
	"myURL.com/inventory/helpers"

	"github.com/gorilla/mux"
)

func StartApi() {
	err := godotenv.Load(".env")
	helpers.CheckError(err)
	database.InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/getHome", controllers.GetHome).Methods("GET")
	router.HandleFunc("/addUser", controllers.AddUser).Methods("POST")
	fmt.Println("App is working on port :4040")
	log.Fatal(http.ListenAndServe(":4040", router))
}
