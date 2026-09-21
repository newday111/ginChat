package emialhandler

import (
	"fmt"
	redisdb "go_jichu/internal/db/redisDB"
	"go_jichu/internal/utils"
	"go_jichu/internal/utils/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type registerCodeEmail struct {
	Email string `json:"email" binding:"required,email"`
}

func RegisterCodeHandler(c *gin.Context) {
	var registerCodeRep registerCodeEmail
	err := c.ShouldBindJSON(&registerCodeRep)
	if err != nil {
		fmt.Println(err)
		utils.AccessLog.Warn("register code function params error",
			zap.String("path", c.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(c, response.ParamErrorCode, "参数错误")
		return
	}

	_, err = redisdb.GlobalRdb.Get(c, registerCodeRep.Email).Result()
	if err == nil {
		utils.AccessLog.Info("register verification code already exists",
			zap.String("email", registerCodeRep.Email),
		)
		resisterExistsCode := map[string]interface{}{
			"msg": "验证码还在生效期",
		}
		response.Success(c, resisterExistsCode)
		return
	}

	if err != nil {
		utils.ErrorLog.Error("redis get register verification code failed",
			zap.String("email", registerCodeRep.Email),
			zap.Error(err),
		)
		resisterExistsCode := map[string]interface{}{
			"msg": "服务器错误,请稍后重试",
		}
		response.Success(c, resisterExistsCode)
		return
	}

	if err != redis.Nil {
		utils.ErrorLog.Error("register code get redis key error",
			zap.String("failed", registerCodeRep.Email),
			zap.Error(err),
		)
	}

	registerCode, err := utils.GenerateRegisterCode()
	if err != nil {
		utils.AccessLog.Warn("create register code failed error",
			zap.String("path", c.Request.URL.Path),
			zap.Error(err),
		)
	}

	//	进行go协程的验证码发送

	utils.EmailPool.Submit(func() {
		err = utils.EmailSender.SendEmail(registerCodeRep.Email, registerCode)
		if err != nil {
			utils.AccessLog.Error("send register email result message: ",
				zap.String("path", c.Request.URL.Path),
				zap.String("register email", registerCodeRep.Email),
				zap.Error(err),
			)
		} else {
			utils.AccessLog.Info("send register code success",
				zap.String("path", c.Request.URL.Path),
				zap.String("register code email result", fmt.Sprintf("%s email send code success", registerCodeRep.Email)),
			)
			err = redisdb.GlobalRdb.Set(c, registerCodeRep.Email, registerCode, 1*time.Minute).Err()
			if err != nil {
				utils.ErrorLog.Error("redis storage register code failed",
					zap.String("email", registerCodeRep.Email),
					zap.Error(err),
				)
			}
		}

	})

	sendEmailData := map[string]interface{}{
		"msg": "注册验证码稍后发送至邮箱,请注意查收",
	}

	response.Success(c, sendEmailData)

}
