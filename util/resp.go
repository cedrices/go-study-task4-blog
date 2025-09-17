package util

import "github.com/gin-gonic/gin"

// response.go
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, code int, message string, data ...interface{}) {
	resp := Response{Code: code, Message: message}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	c.JSON(code, resp)
}

// 快捷函数
func Success(c *gin.Context, data interface{}) {
	JSON(c, 200, "success", data)
}

func Fail(c *gin.Context, code int, msg string) {
	JSON(c, code, msg)
}
