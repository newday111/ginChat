package routers

import (
	emailrouter "go_jichu/internal/routers/emailRouter"
	userrouter "go_jichu/internal/routers/userRouter"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouter(r *gin.Engine) {
	userrouter.RegisterUserRouter(r)
	emailrouter.SendEmailRouter(r)
}
