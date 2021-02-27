package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//GetUsers ...
func (s *Server) GetUsers(ctx *gin.Context) {
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
}

//AddItem ...
func (s *Server) AddItem(ctx *gin.Context) {
	db := s.DB

	var item models.Item

	err := json.NewDecoder(ctx.Request.Body).Decode(&item)
	helpers.CheckError(err)

	item.SocietyId, err = helpers.ExtractTokenID(ctx.Request)
	helpers.CheckError(err)

	db.Create(&item)

	fmt.Println(ctx.PostForm("item"))

	ctx.String(http.StatusOK, "Item Added")
}

//GetItems ...
func (s *Server) GetItems(ctx *gin.Context) {
	db := s.DB

	var items []models.Item

	keys := ctx.Request.URL.Query()
	societyID, err := uuid.FromString(keys.Get("society"))
	if err != nil {
		newSocietyID, err := helpers.ExtractTokenID(ctx.Request)
		helpers.CheckError(err)
		societyID = newSocietyID
	}
	db.Find(&items).Where("society_id = ?", societyID)

	ctx.JSON(http.StatusOK, gin.H{"items": items})
}
