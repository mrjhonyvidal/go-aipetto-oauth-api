package app

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/clients/cassandra"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-oauth-api/src/http"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	cassandra.CheckConnectionOnStartApplication()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8082")
}