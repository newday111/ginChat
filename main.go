package main

import (
	"go_jichu/conf"
	inithandle "go_jichu/initHandle"

	mysqldb "go_jichu/internal/db/mysqlDB"
	redisdb "go_jichu/internal/db/redisDB"
	"go_jichu/internal/middleware"
	"go_jichu/internal/routers"
	utils "go_jichu/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	configFilePath := "./conf/devConfig.yaml"

	//	加载配置文件
	cfg, err := conf.LoadConfig(configFilePath)
	if err != nil {
		panic(err)

	}

	//	程序结束防止没有落盘日志丢失
	defer utils.Sync()

	if err := utils.InitLogger(cfg); err != nil {
		panic(err)
	}

	utils.InitWorkerPool(10, 1000)
	utils.InitEmailConfig(cfg)

	// 初始化redis链接
	err = redisdb.InitRedisConn(cfg)
	if err != nil {
		utils.ErrorLog.Error("init redis connection failed",
			zap.String("redis connection", "failed"),
			zap.Error(err))
		return
	}
	utils.AccessLog.Info("redis init connection", zap.String("status", "success"))

	//	初始化mysql连接
	err = mysqldb.InitMySQL(cfg)
	if err != nil {
		utils.ErrorLog.Error("init mysql connection failed",
			zap.String("mysql connection", "failed"),
			zap.Error(err))
		return
	}
	utils.AccessLog.Info("mysql init connection", zap.String("status", "success"))

	// 初始化gin环境
	inithandle.InitGinModel(cfg)

	// 可以区分线上或者是线下环境
	server := gin.New()
	server.SetTrustedProxies([]string{"127.0.0.1"})

	server.Use(
		middleware.LoggerMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	routers.RegisterAllRouter(server)
	// fmt.Sprintf(":%s", cfg.App.Port)
	server.Run(":8081")
}
