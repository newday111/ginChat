package userrouter

import (
	userhandler "go_jichu/internal/handlers/userHandler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", userhandler.RegisterHandler)
	}

}
