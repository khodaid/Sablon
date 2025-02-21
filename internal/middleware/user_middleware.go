package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
}

type jwtMiddleware struct {
	jwt jwt.JwtService
}

func NewAuthMiddleware(jwt jwt.JwtService) *jwtMiddleware {
	return &jwtMiddleware{jwt}
}

func (j *jwtMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			respone := helpers.APIResponse("Autorization not found", http.StatusNotFound, "error", "")
			ctx.AbortWithStatusJSON(http.StatusNotFound, respone)
			return
		}

		token := strings.Split(authHeader, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			respone := helpers.APIResponse("token invalid", http.StatusMethodNotAllowed, "error", "")
			ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, respone)
			return
		}

		jwtToken, err := j.jwt.ValidateToken(token[1])
		if err != nil || !jwtToken.Valid {
			fmt.Println(err)
			errors := helpers.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			respone := helpers.APIResponse("token expired", http.StatusUnauthorized, "error", errorMessage)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
			return
		}

		ctx.Next()
	}
}
