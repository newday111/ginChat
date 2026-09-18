package inithandle

import (
	"go_jichu/conf"

	"github.com/gin-gonic/gin"
)

func InitGinModel(cfg *conf.ConfigStruct) {
	switch cfg.App.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	}
}
