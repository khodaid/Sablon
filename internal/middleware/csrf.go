package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type CSRFMiddleware interface {
	CsrfMiddleware() gin.HandlerFunc
}

type csrfService struct {
	csrf jwt.CsrfService
}

func NewCSRFMiddleware(csrf jwt.CsrfService) *csrfService {
	return &csrfService{csrf: csrf}
}

func (c *csrfService) CsrfMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			csrfToken := ctx.PostForm("_csrf")
			if csrfToken == "" {
				respone := helpers.APIResponse("token csrf not found", http.StatusBadRequest, "error", csrfToken)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
				return
			}
			_, err := c.csrf.ValidateToken(csrfToken)

			if err != nil {
				errors := helpers.FormatValidationError(err)
				errorMessage := gin.H{"errors": errors}
				respone := helpers.APIResponse("token csrf expired", http.StatusUnauthorized, "error", errorMessage)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
				return
			}
		}

		ctx.Next()
	}
}
