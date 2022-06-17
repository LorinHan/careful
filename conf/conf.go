package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() {
	// 初始化
	viper.SetConfigFile("conf/conf.ini")
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
