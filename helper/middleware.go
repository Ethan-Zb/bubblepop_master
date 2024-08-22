package helper

import (
	"bubblepop_master/global"
	"bubblepop_master/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 错误码
const (
	Error999InvalidBody      = 999
	Error997InvalidToken     = 997
	Error996ExpiredToken     = 996
	Error995InvalidRequest   = 995
	Error994InvalidSig       = 994
	Error1001RequestFrequent = 1001
	Error1002UserAccountBan  = 1002
	// 其他错误代码...
)

var appConfig = struct {
	IsMaintenance bool
}{}

// Helper functions (dummy implementations)
func helperRet(code string, msg string) gin.H {
	return gin.H{"code": code, "message": msg}
}

func verifyAuthorization(token string) (int, error) {
	// 模拟token验证，返回token包含的用户ID
	// 这里可以替换为实际的验证逻辑
	if token == "valid-token" {
		return 1, nil
	}
	return 0, fmt.Errorf("invalid token")
}

func verifyParamSignature(postData map[string]interface{}) bool {
	// 模拟参数签名校验
	// 这里可以替换为实际的签名校验逻辑
	return true
}

func setUserLang(userID int, lang string) {
	// 设置用户语言的逻辑，这里可以实现与数据库交互等
}

func AuthMiddleware(checkAuthorization bool, checkSignature bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 解析请求体
		var postData map[string]interface{}
		if err := json.NewDecoder(c.Request.Body).Decode(&postData); err != nil {
			RetResponse(c, Error999InvalidBody, nil, "Invalid request body")
			c.Abort()
			return
		}

		public := postData["public"].(map[string]interface{})
		params := postData["params"].(map[string]interface{})
		c.Set("public", public)
		c.Set("params", params)

		authorization := c.GetHeader("Authorization")
		global.Log.Infof(
			"path:%s, authorization:%s, post_data:%v\n", c.Request.URL.Path, authorization, postData,
		)

		// 服务器是否维护
		//appConfig.IsMaintenance = true
		if appConfig.IsMaintenance {
			RetResponse(c, 4999, nil, "server is under maintenance")
			c.Abort()
			return
		}

		// 请求频率
		if utils.IsRequestTooFrequent(c.Request, "") {
			RetResponse(c, 4999, nil, "server is under maintenance")
			c.Abort()
			return
		}

		// 权限检查
		var tokenUID int
		var err error
		if checkAuthorization {
			tokenUID, err = verifyAuthorization(strings.TrimSpace(authorization))
			if err != nil || tokenUID <= 0 {
				RetResponse(c, Error997InvalidToken, nil, "checkAuthorization error")
				c.Abort()
				return
			}
		}

		// 参数签名校验
		if checkSignature && !verifyParamSignature(postData) {
			RetResponse(c, Error994InvalidSig, nil, "checkSignature error")
			c.Abort()
			return
		}

		// 继续请求处理
		c.Next()
	}
}
