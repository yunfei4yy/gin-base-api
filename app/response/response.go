package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义一个统一的返回值结构体
type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

// Success 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

// Error 失败返回
func Error(c *gin.Context, errorCode int, message string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		message,
	})
}

func ValidateFailed(c *gin.Context, message string) {
	//c.JSON(http.StatusOK, Response{
	//	1,
	//	nil,
	//	message,
	//})
	Error(c, ValidateError, message)
}

func PermissionDenied(c *gin.Context) {
	Error(c, HaveNoPermission, ErrorMap[HaveNoPermission])
}

func BusinessFail(c *gin.Context, message string) {
	Error(c, BusinessError, message)
}
