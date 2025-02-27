package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHeaderToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return ""
	}

	token := strings.Split(authHeader, " ")
	if len(token) != 2 || token[0] != "Bearer" {
		return ""
	}

	return token[1]
}
