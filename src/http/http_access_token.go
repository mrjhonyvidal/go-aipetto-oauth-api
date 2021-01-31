package http

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	access_token2 "github.com/aipetto/go-aipetto-oauth-api/src/services/access_token"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewAccessTokenHandler(service access_token2.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context){
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessTokenRequest

	// validate if our data struct is a correct valid json format
	if err := c.ShouldBindJSON(&at); err != nil{
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	token, err := handler.service.Create(at);
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}