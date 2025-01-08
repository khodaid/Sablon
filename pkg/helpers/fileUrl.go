package helpers

import (
	"fmt"
	"net/url"
)

func UrlLogo(app_url string, file_name string) string {
	link := fmt.Sprintf("%s/file/image/%s", app_url, url.PathEscape(file_name))
	return link
}
