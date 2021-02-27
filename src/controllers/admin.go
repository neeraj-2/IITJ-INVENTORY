package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"
)

//CreateSociety ...
func (s *Server) CreateSociety(ctx *gin.Context) {
	db := s.DB

	var society models.Society

	err := json.NewDecoder(ctx.Request.Body).Decode(&society)
	helpers.CheckError(err)

	db.Create(&society)

	ctx.String(http.StatusOK, "Society Added")
}
