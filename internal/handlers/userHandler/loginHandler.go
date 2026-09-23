package userhandler

import (
	"go_jichu/internal/model"
	"go_jichu/internal/utils"
	"go_jichu/internal/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginReqParams struct {
	UserName string `json:"username" binding:"required,email"`
	PassWord string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var loginReqParams LoginReqParams
	err := c.ShouldBindJSON(&loginReqParams)
	if err != nil {
		utils.ErrorLog.Error("login function request params error",
			zap.Error(err))
		response.Error(c, response.ParamErrorCode, response.ResponseMessage[response.ParamErrorCode])
		return
	}

	userInfo, err := model.QueryUserInfoToEmail(&model.LoginUser{
		Email:    loginReqParams.UserName,
		Password: loginReqParams.PassWord,
	})

	if err != nil {
		utils.ErrorLog.Error("query login user infomatio failed",
			zap.Error(err))
		response.Error(c, response.QueryLoginUserInfoFailedCode, response.ResponseMessage[response.QueryLoginUserInfoFailedCode])
		return
	}

	encryptionPassword, err := utils.HashPassword(loginReqParams.PassWord)
	if err != nil {
		utils.ErrorLog.Error("login check encryption password failed",
			zap.Error(err))
		response.Error(c, response.ServerErrorCode, response.ResponseMessage[response.ServerErrorCode])
		return
	}

	if userInfo.Password != encryptionPassword {
		utils.AccessLog.Info("login function check password not same",
			zap.String("email", loginReqParams.UserName))
		response.Error(c, response.LoginPassWordInconsistentCode, response.ResponseMessage[response.LoginPassWordInconsistentCode])
		return
	}

}
