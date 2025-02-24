package helpers

import (
	"errors"
	"fmt"
)

// Maksimal Ukuran File Upload
func MaxFileSizeMB(mb int, fileSize int) error {
	size := mb * 1024 * 1024
	if size < fileSize {
		errMessage := fmt.Errorf("logo size exceeds %dMB", mb)
		return errors.New(errMessage.Error())
	}
	return nil
}

// Ekstensi yang diperbolehkan untuk logo
func ValidationLogoExtensions(ext string) error {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := allowedExtensions[ext]; !ok {
		return errors.New("invalid logo file type")
	}
	return nil
}
