package userhandler

import (
	"go_jichu/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type registerUserStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var registerReq registerUserStruct
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		response.Error(c, response.ParamErrorCode, "参数错误")
		return
	}
}
