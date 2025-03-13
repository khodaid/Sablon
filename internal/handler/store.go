package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/khodaid/Sablon/internal/dto"
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

	if err := c.ShouldBind(&storeInput); err != nil {
		fmt.Println("Error binding data:", err)
		return
	}

	fmt.Println(storeInput)

	// // coba goroutine
	// wg := sync.WaitGroup{}
	// var file *multipart.FileHeader
	// go func() {
	// 	wg.Add(1)
	// 	var err error
	// 	file, err = c.FormFile("logo")
	// 	if err != nil {
	// 		errors := helpers.FormatValidationError(err)
	// 		errorMessage := gin.H{"errors": errors}

	// 		response := helpers.APIResponse("Logo is required", http.StatusUnprocessableEntity, "error", errorMessage)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	wg.Add(2)
	// 	fileLimit := helpers.MaxFileSizeMB(2, int(file.Size))
	// 	if fileLimit != nil {
	// 		errorMessage := gin.H{"errors": fileLimit}
	// 		response := helpers.APIResponse("File more than limit", http.StatusUnprocessableEntity, "error", errorMessage)
	// 		c.JSON(http.StatusUnprocessableEntity, response)
	// 		return
	// 	}
	// 	wg.Done()
	// }()

	// var fileExt string
	// go func() {
	// 	wg.Add(3)
	// 	fileExt = filepath.Ext(file.Filename)
	// 	extNotAllowed := helpers.ValidationLogoExtensions(fileExt)
	// 	if extNotAllowed != nil {
	// 		errorMessage := gin.H{"errors": extNotAllowed}
	// 		response := helpers.APIResponse("Invalid logo file type", http.StatusUnprocessableEntity, "error", errorMessage)
	// 		c.JSON(http.StatusUnprocessableEntity, response)
	// 		return
	// 	}
	// 	wg.Done()
	// }()

	// wg.Wait()

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

	newFileName := fmt.Sprintf("%s%s", helpers.GenerateRandomString(16), fileExt)

	store, err := h.service.StoreRegister(storeInput, newFileName)
	if err != nil {
		fmt.Println("masuk if save")
		// trx.Rollback()
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"erorrs": errors}

		response := helpers.APIResponse("Failed create new store", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save file
	dst := fmt.Sprintf("./storage/logos/%s", newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		trx.Rollback()
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"erorrs": errors}

		response := helpers.APIResponse("Failed to save logo", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	trx.Commit()
	// response := helpers.APIResponse("Success create new store", http.StatusCreated, "success", dto.FormatStoreRegister(c.Request.Host+"/api/v1", store))
	response := helpers.APIResponse("Success create new store", http.StatusCreated, "success", dto.FormatStoreRegister(helpers.GetFullBaseURL(c, "/api/v1"), store))
	c.JSON(http.StatusOK, response)
}

func (h *storeHandler) GetAllStoreWithOutSoftDelete(c *gin.Context) {
	stores, err := h.service.GetAllStoreWithOutSoftDelete()
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("error get all store", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, respone)
		return
	}

	respone := helpers.APIResponse("success get data all store", http.StatusOK, "success", stores)
	c.JSON(http.StatusOK, respone)
}

func (h *storeHandler) UpdateStore(c *gin.Context) {
	var input validation.UpdateStoreInput

	err := c.ShouldBind(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("failed binding input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, respone)
		return
	}
	storeId := c.Param("id")

	result, err := h.service.UpdateStore(storeId, input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}
		respone := helpers.APIResponse("failed update store", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, respone)
		return
	}

	respone := helpers.APIResponse("success updated store", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, respone)
}
