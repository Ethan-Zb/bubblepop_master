package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// response 结构体定义
type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Time string      `json:"time"`
}

// 获取当前时间戳
func getCurrentTimestamp() string {
	return time.Now().Format(time.RFC3339) // 格式化时间戳
}

// RetResponse 封装响应格式
func RetResponse(c *gin.Context, code int, data interface{}, msg string) {
	response := response{
		Code: code,
		Data: data,
		Msg:  msg,
		Time: getCurrentTimestamp(),
	}

	c.JSON(http.StatusOK, response) // 返回 JSON 响应
}
