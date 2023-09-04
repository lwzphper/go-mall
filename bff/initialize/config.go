package initialize

import (
	"fmt"
	"github.com/lwzphper/go-mall/bff/config"
	"github.com/lwzphper/go-mall/bff/global"
	cfgHelper "github.com/lwzphper/go-mall/pkg/config"
)

// InitConfig 初始化配置文件
func InitConfig() {
	cfg := config.NewDefaultConfig()
	err := cfgHelper.LoadConfigFromYml("bff/etc/config.yaml", cfg)
	if err != nil {
		panic(fmt.Sprintf("load config from env error:%v", err))
	}
	global.C = cfg
}
