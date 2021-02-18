package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Server intialsation ...
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}
