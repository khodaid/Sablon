package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type CsrfHandler interface {
	GenerateCSRFToken(*gin.Context)
}

type csrfHandler struct {
	csrf jwt.CsrfService
}

func NewCsrfHandler(csrfService jwt.CsrfService) *csrfHandler {
	return &csrfHandler{csrf: csrfService}
}

func (csrf *csrfHandler) GenerateCSRFToken(c *gin.Context) {
	token, err := csrf.csrf.GenerateToken()
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("error generated token csrf", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, respone)
		return
	}
	c.Header("_csrf", token)
	respone := helpers.APIResponse("success get csrf token", http.StatusOK, "success", "")
	c.JSON(http.StatusOK, respone)
}
