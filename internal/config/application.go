package config

import (
	"github.com/gin-gonic/gin"
)

type appEnv struct {
	APP_NAME string
	APP_ENV  string
	APP_HOST string
	APP_PORT string
}

// type appConfig struct {
// 	g *gin.Engine
// }

func (env *appEnv) InitApp(r *gin.Engine) {
	r.Run(env.APP_HOST + ":" + env.APP_PORT)
	// r.Run(":" + env.APP_PORT)
}
