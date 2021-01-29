package app

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/http"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/db"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/rest"
	oauth_access_token "github.com/aipetto/go-aipetto-oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(oauth_access_token.NewService(db.NewRepository(), rest.NewRestUserRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8082")
}