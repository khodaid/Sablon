package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/service"
)

type UserHandler interface {
	Login(c *gin.Context)
}

type userHandler struct {
	userService service.Service
}

func NewUserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *gin.Context) {
	log.Println("masuk user handler user")
}
