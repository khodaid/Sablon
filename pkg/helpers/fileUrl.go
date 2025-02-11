package helpers

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func UrlLogo(app_url string, file_name string) string {
	link := fmt.Sprintf("%s/file/image/%s", app_url, url.PathEscape(file_name))
	return link
}

func GetFullBaseURL(c *gin.Context, path any) string {
	scheme := "http"

	if c.Request.TLS != nil { // TLS aktif berarti HTTPS
		scheme = "https"
	}

	if path == nil {
		return scheme + "://" + c.Request.Host
	}
	return scheme + "://" + c.Request.Host + path.(string)
}
