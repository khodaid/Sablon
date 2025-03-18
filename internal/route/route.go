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
	csrfHandler  handler.CsrfHandler
	userHandler  handler.UserHandler
	storeHandler handler.StoreHandler
}

type middlewares struct {
	csrf middleware.CSRFMiddleware
	auth middleware.Middleware
	cors middleware.CorsService
}

func NewRouteHandler(csrf handler.CsrfHandler, user handler.UserHandler, store handler.StoreHandler) *handlers {
	return &handlers{csrfHandler: csrf, userHandler: user, storeHandler: store}
}

func NewRouteMiddleware(auth middleware.Middleware, csrf middleware.CSRFMiddleware, cors middleware.CorsService) *middlewares {
	return &middlewares{auth: auth, csrf: csrf, cors: cors}
}

func NewRoute(handler *handlers, middleware *middlewares) *newRoute {
	return &newRoute{handler: handler, middleware: middleware}
}

func (r *newRoute) InitRoute() *gin.Engine {
	c := gin.Default()

	api := c.Group("/api")
	api.Use(r.middleware.cors.CorsMiddleware(), r.middleware.csrf.CsrfMiddleware())

	v1 := api.Group("/v1")
	v1.GET("/csrf-token", r.handler.csrfHandler.GenerateCSRFToken)

	image := v1.Group("/file")
	image.Static("/image", "./storage/logos/")

	v1.POST("/login", r.handler.userHandler.Login)

	register := v1.Group("/register")
	{
		register.POST("/user", r.handler.userHandler.RegisterUserRoot)
		register.POST("/store", r.handler.storeHandler.StoreRegister)
	}

	store := v1.Group("/store")
	store.Use(r.middleware.auth.AuthMiddleware())
	{
		store.GET("/")
	}

	storeUsers := store.Group("/users")
	{
		storeUsers.GET("/", r.handler.userHandler.GetUsersStore)
		storeUsers.GET("/:id", r.handler.userHandler.GetUserById)
		storeUsers.POST("create-employee", r.handler.userHandler.CreateUserEmployeeStore)
		storeUsers.PUT("/update/:id", r.handler.userHandler.UpdateUserStore)
		storeUsers.DELETE("/soft-delete/:id", r.handler.userHandler.SoftDeleteUser)

	}

	return c
}
