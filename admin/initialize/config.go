package initialize

import (
	"flag"
	"fmt"
	"github.com/lwzphper/go-mall/admin/config"
	"github.com/lwzphper/go-mall/admin/global"
	cfgHelper "github.com/lwzphper/go-mall/pkg/config"
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
	global.C = cfg
}
