package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode                     = 0
	ParamErrorCode                  = 10001
	RegisterVerificationCode        = 10002
	RegisterVerificationCodeNotSame = 10003
	ServerErrorCode                 = 50000
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}
