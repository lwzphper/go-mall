package initialize

import (
	"fmt"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	"github.com/lwzphper/go-mall/server/member/global"
	"gorm.io/gorm/logger"
)

func InitDB() {
	cfg := global.Config
	envLogLevelMap := map[app.Env]logger.LogLevel{
		app.ENV_DEVELOPMENT: logger.Info,
		app.ENV_TEST:        logger.Info,
		app.ENV_PRODUCTION:  logger.Warn,
	}
	level, ok := envLogLevelMap[cfg.App.Env]
	if !ok {
		level = logger.Warn
	}

	db := cfg.Mysql
	db.LogLevel = level
	err := db.InitDB()
	if err != nil {
		panic(fmt.Sprintf("init mysql error:%v", err))
	}
	global.DB = db.GetDB()
}
