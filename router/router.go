package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	// [WARNING] Running in "debug" mode. Switch to "release" mode in production
	//gin.SetMode(global.Config.System.Env)
	//gin.SetMode("release")

	router := gin.Default()

	FrontendRouter(router)
	ManageRouter(router)

	return router
}
