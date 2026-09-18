package userhandler

import (
	"go_jichu/internal/utils/response"

	utils "go_jichu/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type registerUserStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func RegisterHandler(c *gin.Context) {
	var registerReq registerUserStruct
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		utils.AccessLog.Warn("register function params error",
			zap.String("path", c.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(c, response.ParamErrorCode, "参数错误")
		return
	}

}
