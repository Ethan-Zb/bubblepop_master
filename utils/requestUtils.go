package utils

import "net/http"

func IsRequestTooFrequent(r *http.Request, funcName string) bool {
	// 这里模拟请求频率限制逻辑，实际应根据业务逻辑实现
	return false
}
