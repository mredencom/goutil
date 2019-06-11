package http

import "github.com/gin-gonic/gin"

//设置接口响应的值
func SetResponse(code int, success bool, msg string, data interface{}) map[string]interface{} {
	response := gin.H{
		"code":    code,
		"success": success,
		"msg":     msg,
		"data":    data,
	}
	return response
}