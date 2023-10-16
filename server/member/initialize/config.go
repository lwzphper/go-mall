package initialize

import (
	"flag"
	"fmt"
	cfgHelper "github.com/lwzphper/go-mall/pkg/config"
	"github.com/lwzphper/go-mall/server/member/config"
	"github.com/lwzphper/go-mall/server/member/global"
)

// InitConfig 初始化配置文件
func InitConfig() {
	conf := flag.String("c", "admin/etc/config.yaml", "配置文件")
	flag.Parse()

	cfg := config.NewDefaultConfig()
	err := cfgHelper.LoadConfigFromYml(*conf, cfg)
	if err != nil {
		panic(fmt.Sprintf("load config from env error:%v", err))
	}
	global.Config = cfg
}
