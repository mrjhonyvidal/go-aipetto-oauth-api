package app

import (
	"fmt"
	"github.com/aipetto/go-aipetto-oauth-api/src/clients/cassandra"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-oauth-api/src/http"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Check connection with cassandra up front before any operation and close session
	retryCount := 10
	for {
		if session, dbErr := cassandra.GetSession(); dbErr != nil {
			if retryCount == 0 {
				log.Fatalf("It was not able to establish connection to db.")
			}
			log.Printf(fmt.Sprintf("Could not connect to database. Wait 5 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(5 * time.Second)
		}else{
			session.Close()
			break;
		}
	}

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8082")
}