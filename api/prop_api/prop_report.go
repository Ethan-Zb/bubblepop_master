package prop_api

import (
	"bubblepop_master/global"
	"bubblepop_master/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (PropApi) PropReportView(c *gin.Context) {

	GETUsers(c)

}

func GETUsers(c *gin.Context) {

	userID := c.Query("user_id")

	var id uint
	// 转换 userID 为 uint
	if parsedID, err := strconv.ParseUint(userID, 10, 32); err == nil {
		id = uint(parsedID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	//var db *gorm.DB
	result, err := models.GetBaseInfo(global.DB, id)
	if err != nil {
		c.JSON(101, gin.H{"error": "not found data"})
	}

	c.JSON(200, gin.H{"user": result})
}
