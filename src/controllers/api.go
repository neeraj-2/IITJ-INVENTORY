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

//Check If email Exists ...
func (s *Server) CheckIfEmailExists(ctx *gin.Context) {
	db := s.DB
	var email models.Email
	err := json.NewDecoder(ctx.Request.Body).Decode(&email)
	helpers.CheckError(err)

	user, er := models.EmailExists(db, email.Email)
	if er != nil {
		ctx.JSON(http.StatusOK, gin.H{"user": nil})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	}
}
