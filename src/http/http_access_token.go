package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domainAT "github.com/mohammadshabab/bookstore_oauth-api/src/domain/access_token"
	"github.com/mohammadshabab/bookstore_oauth-api/src/service/access_token"
	"github.com/mohammadshabab/bookstore_utils-go/rest_errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	access_token, err := h.service.GetById(c.Param("access_token_id"))
	//accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, access_token)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request domainAT.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
