package emailrouter

import (
	emialhandler "go_jichu/internal/handlers/emialHandler"

	"github.com/gin-gonic/gin"
)

func SendEmailRouter(c *gin.Engine) {
	email := c.Group("/email")
	{
		email.POST("/register/code", emialhandler.RegisterCodeHandler)
	}
}
