package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

// func FormatValidationError(err error) []string {
// 	var errors []string

// 	// Periksa apakah error adalah tipe validator.ValidationErrors
// 	if validationErrors, ok := err.(validator.ValidationErrors); ok {
// 		for _, e := range validationErrors {
// 			errors = append(errors, e.Error())
// 			fmt.Println(e.Field())
// 			fmt.Println(e.Tag())
// 		}
// 	} else {
// 		// Jika error bukan tipe validator.ValidationErrors, tambahkan pesan default
// 		errors = append(errors, err.Error())
// 	}

// 	return errors
// }

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field() // Nama field
			tag := e.Tag()     // Tag validasi yang gagal
			errors[field] = fmt.Sprintf("Validation failed on '%s' tag", tag)
		}
	} else {
		errors["general"] = err.Error() // Error umum jika bukan Validation Errors
	}

	return errors
}
