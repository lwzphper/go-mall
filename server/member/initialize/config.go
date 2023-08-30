package initialize

import (
	"fmt"
	cfgHelper "github.com/lwzphper/go-mall/pkg/config"
	"github.com/lwzphper/go-mall/server/member/config"
	"github.com/lwzphper/go-mall/server/member/global"
)

// InitConfig 初始化配置文件
func InitConfig() {
	cfg := config.NewDefaultConfig()
	err := cfgHelper.LoadConfigFromYml("server/member/etc/config.yaml", cfg)
	if err != nil {
		panic(fmt.Sprintf("load config from env error:%v", err))
	}
	global.Config = cfg
}
