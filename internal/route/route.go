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
	userHandler  handler.UserHandler
	storeHandler handler.StoreHandler
}

func NewRoute(user handler.UserHandler, store handler.StoreHandler) *handlers {
	return &handlers{user, store}
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

	v1.POST("/login", h.userHandler.Login)
	v1.POST("/register", h.storeHandler.StoreRegister)

	// userV1 := v1.Group("/user")

	return c
}
