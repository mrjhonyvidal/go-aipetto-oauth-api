package http

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context){
	c.JSON(http.StatusNotImplemented, "implement me!")
}