package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routeConfig struct {
	g *gin.Engine
}

func InitRoute() *gin.Engine {
	// c := routeConfig{}
	c := gin.Default()

	// g := c.g.Group("/api/account")
	g := c.Group("/api/account")

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "space persons",
		})
	})

	return c
}
