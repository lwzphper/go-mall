package initialize

import "github.com/lwzphper/go-mall/server/member/global"

func InitLogger() {
	env := global.Config.App.Env
	global.Logger = global.Config.Logging.InitLogger(env)
}
