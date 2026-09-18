package routers

import (
	userrouter "go_jichu/internal/routers/userRouter"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouter(r *gin.Engine) {
	userrouter.RegisterUserRouter(r)
}
