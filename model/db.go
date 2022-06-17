package model

import (
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/database"
	"go.uber.org/zap"
	"os"
	"time"
)

func Init() {
	var (
		option = database.NewOption()
	)

	option.SlowQueryTime = 300 * time.Millisecond  // 设置慢查询时间
	logFile, _ := os.OpenFile(viper.GetString("DB.SlowQueryFile") , os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	option.SlowQueryLog = logFile  // 设置慢查询日志

	if err := database.Init(viper.GetString("DB.Mode"), viper.GetString("DB.Cfg"), option); err != nil {
		zap.L().Fatal(err.Error())
	}
}
