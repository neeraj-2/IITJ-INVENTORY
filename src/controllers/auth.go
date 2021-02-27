package controllers

import (
	"encoding/json"
	"net/http"

	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"

	"github.com/gin-gonic/gin"
)

//StudentLogin ...
func (s *Server) StudentLogin(ctx *gin.Context) {

	db := s.DB

	keys := ctx.Request.URL.Query()
	token := keys.Get("token")

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if resp.StatusCode != 200 {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	defer resp.Body.Close()
	type response struct {
		Name  string
		Email string
	}

	var us response
	err = json.NewDecoder(resp.Body).Decode(&us)

	user, err := models.CheckUserExistsFromEmail(db, us.Email)
	if err != nil {
		db.Create(models.User{Name: us.Name, Email: us.Email})
	}

	token, err = helpers.CreateToken(user.UUID)
	helpers.CheckError(err)
	ctx.JSON(http.StatusOK, gin.H{"jwt": token})
}

//SocietyLogin ...
func (s *Server) SocietyLogin(ctx *gin.Context) {
	db := s.DB
	var soc models.Society
	err := json.NewDecoder(ctx.Request.Body).Decode(&soc)
	helpers.CheckError(err)
	society, err := models.CheckSocietyExists(db, soc)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	} else {
		token, err := helpers.CreateToken(society.UUID)
		helpers.CheckError(err)
		ctx.JSON(http.StatusOK, gin.H{"jwt": token})
	}
}

//AdminLogin ...
func (s *Server) AdminLogin(ctx *gin.Context) {
	db := s.DB
	var us models.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&us)
	helpers.CheckError(err)

	user, err := models.CheckUserExists(db, us)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	} else {
		token, err := helpers.CreateToken(user.UUID)
		helpers.CheckError(err)
		ctx.JSON(http.StatusOK, gin.H{"jwt": token})
	}

}
