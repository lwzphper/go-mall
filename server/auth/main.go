package main

import (
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	"github.com/lwzphper/go-mall/pkg/common/config/db"
	config2 "github.com/lwzphper/go-mall/pkg/config"
	"github.com/lwzphper/go-mall/pkg/logger"
	authpb "github.com/lwzphper/go-mall/server/auth/api/gen/v1"
	"github.com/lwzphper/go-mall/server/auth/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var (
	logFileDir    = "/log/auth.log"
	configFileDir = "./etc/config.toml"
)

// Config 配置文件
type Config struct {
	App     *app.App    `toml:"app" yaml:"app" mapstructure:"app"`
	Mysql   *db.Mysql   `toml:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Mongodb *db.Mongodb `toml:"Mongodb" yaml:"mongodb" mapstructure:"mongodb"`
}

func main() {
	var err error

	// 初始化配置
	cfg := newConfig()
	err = config2.LoadConfigFromToml(configFileDir, cfg)
	if err != nil {
		panic("load config error: " + err.Error())
	}

	// 初始化日志
	logCfg := logger.SizeRotateLogConfig{
		Level:      logger.DebugLevel,
		FileName:   logFileDir,
		MaxSize:    100,
		MaxAge:     7,
		MaxBackups: 10,
	}
	rotateLog := logger.NewWithSizeRotate(logCfg)
	rotateLog.Name(cfg.App.Name)

	lis, err := net.Listen("tcp", cfg.App.Addr)
	if err != nil {
		rotateLog.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &service.AuthService{
		Logger:     rotateLog,
		MemberAuth: nil, // 这里设置 member
	})
	err = s.Serve(lis)
	if err != nil {
		rotateLog.Fatalf("grpc server error: %v", err)
	}
}

func newConfig() *Config {
	return &Config{
		Mysql:   db.NewDefaultMysql(),
		Mongodb: db.NewDefaultMongoDB(),
	}
}
