package main

import (
	"github.com/lwzphper/go-mall/bff/initialize"
)

func main() {
	initialize.InitConfig()                       // 配置
	initialize.InitLogger()                       // 日志
	initialize.InitSrvConn()                      // 服务
	initialize.InitValidator(initialize.ZhLocale) // 初始验证器
	initialize.InitGin()                          // gin
}
