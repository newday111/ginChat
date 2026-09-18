package main

import (
	"go_jichu/conf"
	inithandle "go_jichu/initHandle"
)

func main() {
	configFilePath := "./conf/devConfig.yaml"

	cfg, err := conf.LoadConfig(configFilePath)
	if err != nil {
		panic(err)

	}
	inithandle.InitGinModel(cfg)

	// // 可以区分线上或者是线下环境
	// gin.SetMode(gin.DebugMode)
	// server := gin.Default()

	// server.Run(":8080")
}
