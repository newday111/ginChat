package main

import (
	"go_jichu/conf"
	inithandle "go_jichu/initHandle"

	"github.com/gin-gonic/gin"
)

func main() {
	configFilePath := "./conf/devConfig.yaml"

	//	加载配置文件
	cfg, err := conf.LoadConfig(configFilePath)
	if err != nil {
		panic(err)

	}

	// 初始化gin环境
	inithandle.InitGinModel(cfg)

	// 可以区分线上或者是线下环境
	server := gin.Default()

	server.Run(":8080")
}
