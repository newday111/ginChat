package userrouter

import "github.com/gin-gonic/gin"

func registerUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register")
	}

}
