package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/config/jwt"
	"github.com/khodaid/Sablon/internal/dto"
	"github.com/khodaid/Sablon/internal/service"
	"github.com/khodaid/Sablon/internal/validation"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type UserHandler interface {
	Login(c *gin.Context)
	RegisterUserRoot(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserStore(c *gin.Context)
	GetAllWithOutSoftDelete(c *gin.Context)
	GetUsersStore(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  jwt.JwtService
}

func NewUserHandler(userService service.UserService, jwtService jwt.JwtService) *userHandler {
	return &userHandler{userService, jwtService}
}

func (h *userHandler) RegisterUserRoot(c *gin.Context) {
	var userInput validation.RegisterUserStoreAdminInput

	if err := c.ShouldBind(&userInput); err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := helpers.APIResponse("Error bindding data", http.StatusExpectationFailed, "error", errorMessage)
		c.JSON(http.StatusExpectationFailed, respone)
		return
	}

	user, err := h.userService.Register(userInput)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := helpers.APIResponse("Failed create new user store root", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	respone := helpers.APIResponse("Success create new user store root", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) Login(c *gin.Context) {
	// validasi login input
	var input validation.LoginUserInput

	// binding input
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := helpers.APIResponse("Error bindding data", http.StatusExpectationFailed, "error", errorMessage)
		c.JSON(http.StatusExpectationFailed, respone)
		return
	}

	user, err := h.userService.Login(input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := helpers.APIResponse("Failed find user", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respone := helpers.APIResponse("Failed find user", http.StatusExpectationFailed, "error", errorMessage)
		c.JSON(http.StatusExpectationFailed, respone)
		return
	}
	c.Header("X-AUTH", token)

	userData := dto.FormatDetailUserLogin(user)

	respone := helpers.APIResponse("Success findding user store", http.StatusOK, "success", map[string]interface{}{
		"token": token,
		"user":  userData,
	})
	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) GetUserById(c *gin.Context) {
	userId := c.Param("id")
	user, err := h.userService.GetUserById(userId)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorsMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("failed get user", http.StatusBadRequest, "error", errorsMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, respone)
		return
	}

	userData := dto.FormatAllUsersStore(user)
	respone := helpers.APIResponse("success get user", http.StatusOK, "success", userData)
	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) UpdateUserStore(c *gin.Context) {
	var input validation.UpdateUserStore

	err := c.ShouldBind(input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("failed binding input", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, respone)
		return
	}

	userId := c.Param("id")

	result, err := h.userService.UpdateUserById(userId, input)

	if err != nil {
		error := helpers.FormatValidationError(err)
		errorsMessage := gin.H{"message": error}
		respone := helpers.APIResponse("failed updated user", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, respone)
		return
	}
	respone := helpers.APIResponse("success updated user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) GetAllWithOutSoftDelete(c *gin.Context) {
	users, err := h.userService.GetAllWithOutSoftDelete()

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}
		respone := helpers.APIResponse("Failed fetch all user without soft delete", http.StatusBadRequest, "error", errorsMessage)
		c.AbortWithStatusJSON(http.StatusBadGateway, respone)
		return
	}

	respone := helpers.APIResponse("Success fetch all user without soft delete", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, respone)
}

func (h *userHandler) GetUsersStore(c *gin.Context) {
	tokenAuth := c.GetHeader("Authorization")
	if tokenAuth == "" {
		respone := helpers.APIResponse("Failed get users store", http.StatusUnauthorized, "error", "")
		c.AbortWithStatusJSON(http.StatusUnauthorized, respone)
	}

	token := strings.Split(tokenAuth, " ")
	decodeToken, err := jwt.DecodeJWT(token[1])

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}
		respone := helpers.APIResponse("Failed", http.StatusBadRequest, "error", errorsMessage)
		c.AbortWithStatusJSON(http.StatusBadGateway, respone)
		return
	}

	users, err := h.userService.GetAllUserByStore(decodeToken["user_id"].(string))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorsMessage := gin.H{"error": errors}
		respone := helpers.APIResponse("Failed", http.StatusBadRequest, "error", errorsMessage)
		c.AbortWithStatusJSON(http.StatusBadGateway, respone)
		return
	}

	respone := helpers.APIResponse("Success", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, respone)
}
