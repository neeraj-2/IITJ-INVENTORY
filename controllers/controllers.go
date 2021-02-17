package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"myURL.com/inventory/database"
	"myURL.com/inventory/helpers"
	"myURL.com/inventory/models"

	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	db := database.OpenConnectionToDb()

	var users []models.User
	db.Find(&users)
	fmt.Println("{}", users)

	// var data string

	// json.NewEncoder(data).Encode(users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})

	defer db.Close()
}

func AddUser(ctx *gin.Context) {
	db := database.OpenConnectionToDb()

	var user models.User

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	helpers.CheckError(err)

	fmt.Println(user)

	db.Create(&user)

	fmt.Println(ctx.PostForm("user"))

	ctx.String(http.StatusOK, "User Added")
	defer db.Close()
}
