package emialhandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterCodeHandler(c *gin.Context) {
	fmt.Println("发送邮件")
}
