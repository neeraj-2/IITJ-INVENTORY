package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myurl.com/inventory/helpers"
	"myurl.com/inventory/models"
)

//SetMiddlewareJSON sets the header content type
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

//SetMiddlewareAuthentication checks for token in request
func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helpers.TokenValid(ctx.Request)
		if err != nil {
			respondWithError(ctx, 401, "Invalid API token")
			return
		}
		ctx.Next()
	}
}

//SetMiddlewareAuthenticationAdmin checks for token in request
func SetMiddlewareAuthenticationAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserUUID, err := helpers.ExtractTokenID(ctx.Request)
		if err != nil {
			respondWithError(ctx, 401, "Invalid API token")
			return
		}
		var user models.User
		result := db.Find(&user).Where("UUID = ? AND IsAdmin = ?", UserUUID, true)
		if result.Error != nil {
			respondWithError(ctx, 401, "Invalid API token")
			return
		}
		ctx.Next()
	}
}
