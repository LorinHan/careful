package main

import (
	"careful/conf"
	"careful/controller"
	"careful/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/logger"
	"github.com/zenfire-cn/webkit/web"
	"go.uber.org/zap"
)

func main() {
	conf.Init()

	webConf := &logger.Option{
		Path:    viper.GetString("Log.Path"),
		Level:   viper.GetString("Log.port"),
		MaxSize: viper.GetInt("Log.MaxSize"),
		Json:    viper.GetBool("Log.Json"),
		Std:     viper.GetBool("Log.Std"),
	}
	r := web.Init(gin.ReleaseMode, webConf) // 初始化：gin结合zap日志与lumberjack日志归档

	if viper.GetString("App.Mode") == "main" {
		model.Init()
	}

	controller.InitRouter(r)

	port := viper.GetString("App.port")
	zap.L().Info("server running on port: " + port)

	if err := r.Run(":" + port); err != nil {
		zap.L().Fatal(err.Error())
	}
}
