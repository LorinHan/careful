package model

import (
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/database"
	"go.uber.org/zap"
)

func Init() {
	var (
		option = database.NewOption()
	)

	//option.SlowQueryTime = time.Duration(viper.GetInt64("DB.SlowQueryTime")) * time.Millisecond  // 设置慢查询时间
	//logFile, _ := os.OpenFile(viper.GetString("DB.SlowQueryFile") , os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	//option.SlowQueryLog = logFile  // 设置慢查询日志

	if err := database.Init(viper.GetString("DB.Mode"), viper.GetString("DB.Cfg"), option); err != nil {
		zap.L().Fatal(err.Error())
	}

	db := database.GetDB()
	if err := db.AutoMigrate(&Server{}); err != nil {
		zap.L().Error(err.Error())
	}
	if err := db.AutoMigrate(&Folder{}); err != nil {
		zap.L().Error(err.Error())
	}
	if err := db.AutoMigrate(&Node{}); err != nil {
		zap.L().Error(err.Error())
	}
	if err := db.AutoMigrate(&ServerPoint{}); err != nil {
		zap.L().Error(err.Error())
	}
}
