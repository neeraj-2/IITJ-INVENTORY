package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"myURL.com/inventory/database"
	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnectionToDb()

	var users []models.User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)

	defer db.Close()
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnectionToDb()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	helpers.CheckError(err)

	fmt.Println(user)

	db.Create(&user)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
