package emialhandler

import (
	"fmt"
	"go_jichu/internal/utils"
	"go_jichu/internal/utils/response"

	"github.com/gin-gonic/gin"
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

	registerCode, err := utils.GenerateRegisterCode()
	if err != nil {
		utils.AccessLog.Warn("create register code failed error",
			zap.String("path", c.Request.URL.Path),
			zap.Error(err),
		)
	}

	//	进行go协程的验证码发送

	utils.EmailPool.Submit(func() {
		utils.EmailSender.SendEmail(registerCodeRep.Email, registerCode)
	})

	// response.Success(c)

}
