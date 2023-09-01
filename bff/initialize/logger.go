package initialize

import "github.com/lwzphper/go-mall/bff/global"

func InitLogger() {
	env := global.Config.App.Env
	global.Logger = global.Config.Logging.InitLogger(env)
}
