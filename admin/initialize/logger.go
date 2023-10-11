package initialize

import (
	"github.com/lwzphper/go-mall/admin/global"
)

func InitLogger() {
	env := global.C.App.Env
	global.L = global.C.Logging.InitLogger(env)
}
