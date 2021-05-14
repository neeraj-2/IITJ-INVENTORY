package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

// Get Society
func (s *Server) GetSociety(ctx *gin.Context) {
	db:= s.DB

	var societies[] models.Society

	db.Find(&societies).Where("")

	ctx.JSON(http.StatusOK, gin.H{"socities": societies})
}

// Request Item
func (s *Server) RequestItem(ctx *gin.Context){
	db:= s.DB

	var user models.User

	studentId, err := helpers.ExtractTokenID(ctx.Request)
	helpers.CheckError(err)

	db.Find(&user).Where("uuid = ?", studentId)

	keys := ctx.Request.URL.Query()
	name := keys.Get("item-name")
	details := keys.Get("item-details")
	societyName := keys.Get("item-society-name")

	// dt, err := time.Parse(time.RFC3339, keys.Get("due-date"))
	helpers.CheckError(err)
	issuedItem := models.Issued{
		IssueDate:time.Now(),
		DueDate: time.Now().AddDate(0, 0, 5),
		Approved:"false",
		Denied:"false",
		Purpose:"None",
		UserId:studentId,
		Name: name,
		Details: details,
		SocietyName: societyName,
	}

	db.Create(&issuedItem)

	ctx.JSON(http.StatusOK, gin.H{"issuedItem": issuedItem})

}


type ItemWithSocietyInfo struct {
	Item models.Item
	Society models.Society
}

// Search ItemsPage
func (s *Server) SearchItems(ctx *gin.Context){
	db:= s.DB

	var allItems[] models.Item

	db.Find(&allItems).Where("")

	fmt.Println("here we are", allItems[0])

	keys := ctx.Request.URL.Query()
	name := keys.Get("item-name")

	var items[] ItemWithSocietyInfo
	var society models.Society

	for i:=0; i<len(allItems); i++{
		if strings.Contains(strings.ToLower(allItems[i].Name), strings.ToLower(name)){
			db.Find(&society).Where("UUID = ?", allItems[i].SocietyId)
			temp := ItemWithSocietyInfo{Item: allItems[i], Society: society}
			items = append(items, temp)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"items": items})

}


