package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"

	"github.com/gin-gonic/gin"
)

//GetHome ..
func (s *Server) GetHome(ctx *gin.Context) {
	db := s.DB

	var users []models.User

	db.Find(&users)
	fmt.Println("{}", users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})

}

//AddUser ...
func (s *Server) AddUser(ctx *gin.Context) {
	db := s.DB

	var user models.User

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	helpers.CheckError(err)

	fmt.Println(user)

	db.Create(&user)

	fmt.Println(ctx.PostForm("user"))

	ctx.String(http.StatusOK, "User Added")
	defer db.Close()
}
