package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/handler"
)

// type routeConfig struct {
// 	g *gin.Engine
// }

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

	image := v1.Group("/file")
	image.Static("/image", "./storage/logos/")

	v1.POST("/login", h.userHandler.Login)

	register := v1.Group("register")
	register.POST("/user", h.userHandler.RegisterUserRoot)
	register.POST("/store", h.storeHandler.StoreRegister)

	// userV1 := v1.Group("/user")

	return c
}
