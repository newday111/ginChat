package userhandler

import (
	redisdb "go_jichu/internal/db/redisDB"
	"go_jichu/internal/model"
	"go_jichu/internal/utils/response"

	utils "go_jichu/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RegisterUserStruct struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	RegisterCode string `json:"registerCode" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var registerReq RegisterUserStruct
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		utils.AccessLog.Warn("register function params error",
			zap.String("path", c.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(c, response.ParamErrorCode, "参数错误")
		return
	}

	registerCode, err := redisdb.GlobalRdb.Get(c, registerReq.Email).Result()

	if err == redis.Nil {
		utils.ErrorLog.Warn("register verification code invalid",
			zap.String("email", registerReq.Email),
			zap.Error(err))
		response.Error(c, response.RegisterVerificationCode, "验证码已失效")
		return
	}

	if err != nil {
		utils.ErrorLog.Warn("register function server error",
			zap.String("email", registerReq.Email),
			zap.Error(err))
		response.Error(c, response.ServerErrorCode, "服务器错误请稍后重试")
		return
	}

	if registerCode != registerReq.RegisterCode {
		utils.ErrorLog.Error("register verification code error",
			zap.String("对比注册验证码", "注册验证码不一致"))
		response.Error(c, response.RegisterVerificationCodeNotSame, "验证码不正确")
		return
	}

	registerReq.Password, err = utils.HashPassword(registerReq.Password)

	_, err = model.CreateRegisterUser(&model.User{
		Email:    registerReq.Email,
		Password: registerReq.Password,
	})
	if err != nil {
		utils.ErrorLog.Error("register user failed",
			zap.String("email", registerReq.Email),
			zap.Error(err),
		)
		response.Error(c, response.RegisterUserFailedCode, "注册用户失败,请稍后重试")
		return
	}

	utils.AccessLog.Info("register user success",
		zap.String("email", registerReq.Email))

	regsterSuccessMes := make(map[string]interface{})
	regsterSuccessMes["message"] = "register success"
	response.Success(c, regsterSuccessMes)
}
