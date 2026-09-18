package middleware

import (
	"fmt"
	"runtime/debug"

	utils "go_jichu/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {

			if err := recover(); err != nil {

				// 记录 panic
				utils.ErrorLog.Error(
					"panic recovered",
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				c.AbortWithStatusJSON(500, gin.H{
					"code":    500,
					"message": "internal server error",
					"error":   fmt.Sprintf("%v", err),
				})
			}

		}()

		c.Next()
	}
}
