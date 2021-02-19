package controllers

import (
	"fmt"
	"io/ioutil"
	"myurl.com/inventory/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRETS"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	randomState = "random"
)

func (s *Server) Login(ctx *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *Server) Callback(ctx *gin.Context) {
	if ctx.Query("state") != randomState {
		panic("state is not valid")
		ctx.Redirect(http.StatusTemporaryRedirect, "auth/error")
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, ctx.Query("code"))
	helpers.CheckError(err)

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	helpers.CheckError(err)

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	helpers.CheckError(err)

	w := ctx.Writer

	fmt.Fprintf(w, "Response: %s", content)
}

func (s *Server) Error(ctx *gin.Context) {
	ctx.String(http.StatusOK, "<h1>Error in Login</h1>")
}
