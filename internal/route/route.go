package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/handler"
)

type routeConfig struct {
	g *gin.Engine
}

type handlers struct {
	userHandler handler.UserHandler
}

func NewRoute(user handler.UserHandler) *handlers {
	return &handlers{user}
}

func (h *handlers) InitRoute() *gin.Engine {
	c := gin.Default()

	api := c.Group("/api")

	v1 := api.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "space persons",
		})
	})

	userV1 := v1.Group("/user")
	userV1.POST("/login", h.userHandler.Login)

	return c
}
