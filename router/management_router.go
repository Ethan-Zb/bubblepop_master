package router

import "github.com/gin-gonic/gin"

func ManageRouter(router *gin.Engine) {
	ManageGroup := router.Group("/manage")
	{
		ManageGroup.POST("/", Login)
	}
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World!"})
}
