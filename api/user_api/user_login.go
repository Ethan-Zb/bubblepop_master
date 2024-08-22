package user_api

import (
	"bubblepop_master/helper"
	"github.com/gin-gonic/gin"
)

func (UserApi) DeviceIdLoginView(c *gin.Context) {

	helper.RetResponse(
		c,
		200,
		map[string]interface{}{"user_id": 1},
		"success",
	)

}
