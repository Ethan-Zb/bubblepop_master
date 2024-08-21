package router

import (
	"bubblepop_master/api"
	"github.com/gin-gonic/gin"
)

func FrontendRouter(router *gin.Engine) {

	propApi := api.ApiGroupApp.PropApi

	frontendGroup := router.Group("/frontend")
	{
		frontendGroup.GET("/", propApi.PropReportView)
	}
}
