package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/handler"
	"github.com/khodaid/Sablon/internal/middleware"
)

type newRoute struct {
	handler    *handlers
	middleware *middlewares
}

type handlers struct {
	userHandler  handler.UserHandler
	storeHandler handler.StoreHandler
}

type middlewares struct {
	auth middleware.Middleware
}

func NewRouteHandler(user handler.UserHandler, store handler.StoreHandler) *handlers {
	return &handlers{userHandler: user, storeHandler: store}
}

func NewRouteMiddleware(auth middleware.Middleware) *middlewares {
	return &middlewares{auth: auth}
}

func NewRoute(handler *handlers, middleware *middlewares) *newRoute {
	return &newRoute{handler: handler, middleware: middleware}
}

func (r *newRoute) InitRoute() *gin.Engine {
	c := gin.Default()

	api := c.Group("/api")

	v1 := api.Group("/v1")

	image := v1.Group("/file")
	image.Static("/image", "./storage/logos/")

	v1.POST("/login", r.handler.userHandler.Login)

	register := v1.Group("/register")
	register.POST("/user", r.handler.userHandler.RegisterUserRoot)
	register.POST("/store", r.handler.storeHandler.StoreRegister)

	// authProtected := v1.Use(r.middleware.auth.AuthMiddleware())

	store := v1.Group("/store").Use(r.middleware.auth.AuthMiddleware())
	store.GET("/users", r.handler.userHandler.GetUsersStore)
	store.GET("users/:id", r.handler.userHandler.GetUserById)

	// userV1 := v1.Group("/user")

	return c
}
