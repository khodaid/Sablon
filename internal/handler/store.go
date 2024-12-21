package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/service"
	"github.com/khodaid/Sablon/internal/validation"
	"github.com/khodaid/Sablon/pkg/helpers"
	"gorm.io/gorm"
)

type StoreHandler interface {
	StoreRegister(*gin.Context)
}

type storeHandler struct {
	db      *gorm.DB
	service service.StoreService
}

func NewStoreHandler(db *gorm.DB, storeService service.StoreService) *storeHandler {
	return &storeHandler{db: db, service: storeService}
}

func (h *storeHandler) StoreRegister(c *gin.Context) {
	var storeInput validation.RegisterStoreInput
	// var userStoreInput validation.RegisterUserStoreAdminInput

	// checking file upload
	file, err := c.FormFile("logo")
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Logo is required", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// checking size limit file upload
	fileLimit := helpers.MaxFileSizeMB(2, int(file.Size))
	if fileLimit != nil {
		errorMessage := gin.H{"errors": fileLimit}
		response := helpers.APIResponse("File more than limit", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// checking ext logo upload
	fileExt := filepath.Ext(file.Filename)
	extNotAllowed := helpers.ValidationLogoExtensions(fileExt)
	if extNotAllowed != nil {
		errorMessage := gin.H{"errors": extNotAllowed}
		response := helpers.APIResponse("Invalid logo file type", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	trx := h.db.Begin()

	store, err := h.service.StoreRegister(storeInput)
	if err != nil {
		trx.Rollback()
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"erorrs": errors}

		response := helpers.APIResponse("Failed create new store", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save file
	dst := fmt.Sprintf("./logos/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		trx.Rollback()
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"erorrs": errors}

		response := helpers.APIResponse("Failed to save logo", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	trx.Commit()
	response := helpers.APIResponse("Success create new store", http.StatusCreated, "success", store)
	c.JSON(http.StatusOK, response)
}