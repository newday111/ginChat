package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode                         = 0
	RegisterUserSuccessCode             = 10000
	ParamErrorCode                      = 10001
	RegisterVerificationCode            = 10002
	RegisterVerificationCodeNotSame     = 10003
	RegisterUserFailedCode              = 10004
	RegisterVerificationCodeEffective   = 10005
	RegisterVerificationCodeSendSuccess = 10006
	ServerErrorCode                     = 50000
)

var ResponseMessage = map[int]string{
	ParamErrorCode:                      "参数错误",
	RegisterUserSuccessCode:             "用户注册成功,请登录",
	RegisterVerificationCode:            "验证码已失效",
	RegisterVerificationCodeNotSame:     "验证码不一致",
	RegisterUserFailedCode:              "用户注册失败,请稍后重试",
	RegisterVerificationCodeEffective:   "验证码生效中",
	RegisterVerificationCodeSendSuccess: "验证码已发送至邮箱,请注意查收",
	ServerErrorCode:                     "服务器错误请稍后重试",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}
