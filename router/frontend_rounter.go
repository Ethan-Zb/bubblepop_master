package router

import (
	"bubblepop_master/api"
	"github.com/gin-gonic/gin"
)

func FrontendRouter(router *gin.Engine) {

	propApi := api.GroupApp.PropApi
	userApi := api.GroupApp.UserApi

	frontendGroup := router.Group("/frontend")
	{
		frontendGroup.GET("/", propApi.PropReportView)
		frontendGroup.GET("/login", userApi.DeviceIdLoginView)
	}

}
