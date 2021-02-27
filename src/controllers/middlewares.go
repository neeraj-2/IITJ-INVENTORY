package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myurl.com/inventory/helpers"
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
