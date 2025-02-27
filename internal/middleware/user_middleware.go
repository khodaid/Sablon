package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/internal/service"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
	UserRootStoreMiddleware() gin.HandlerFunc
}

type jwtMiddleware struct {
	jwt         jwt.JwtService
	userService service.UserService
}

func NewAuthMiddleware(jwt jwt.JwtService, userService service.UserService) *jwtMiddleware {
	return &jwtMiddleware{jwt: jwt, userService: userService}
}

func (j *jwtMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			respone := helpers.APIResponse("Authorization not found", http.StatusNotFound, "error", "")
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

func (j *jwtMiddleware) UserRootStoreMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := helpers.GetHeaderToken(ctx)

		if token == "" {
			respone := helpers.APIResponse("token is invalid", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
			return
		}

		payload, err := jwt.DecodeJWT(token)
		if err != nil {
			respone := helpers.APIResponse("failed get payload token", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
			return
		}

		user, err := j.userService.GetUserById(payload["user_id"].(string))

		if err != nil {
			errors := helpers.FormatValidationError(err)
			errorMessages := gin.H{"message": errors}
			respone := helpers.APIResponse("failed get data user by token", http.StatusUnauthorized, "error", errorMessages)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
			return
		}

		if user.UserRoleAdmin.IsBackoffice || (user.UserRoleAdmin.Role.Value != "user root" && user.UserRoleAdmin.Role.ForLogin != "store") {
			respone := helpers.APIResponse("user not authorization", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respone)
			return
		}

		ctx.Next()
	}

}
