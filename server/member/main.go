package main

import (
	"fmt"
	"github.com/lwzphper/go-mall/pkg/common/config"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	configDB "github.com/lwzphper/go-mall/pkg/common/config/db"
	configHelper "github.com/lwzphper/go-mall/pkg/config"
	"github.com/lwzphper/go-mall/pkg/server"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	App     *app.App        `toml:"app" yaml:"app" mapstructure:"app"`
	Mysql   *configDB.Mysql `toml:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Logging *config.Logging `toml:"logging" yaml:"logging" mapstructure:"logging"`
}

// NewConfig 创建配置文件
func NewConfig() *Config {
	return &Config{
		App:     app.NewDefaultApp(),
		Mysql:   configDB.NewDefaultMysql(),
		Logging: config.NewDefaultLogging(),
	}
}

var (
	cfg    *Config
	gormDB *gorm.DB
)

func main() {
	// 初始化配置文件
	cfg = NewConfig()
	err := configHelper.LoadConfigFromYml("server/member/etc/config.yaml", cfg)
	if err != nil {
		panic(fmt.Sprintf("load config from env error:%v", err))
	}

	// 初始化日志
	logger := cfg.Logging.InitLogger(cfg.App.Env)

	// 初始化数据库
	db := cfg.Mysql
	err = db.InitDB()
	if err != nil {
		panic(fmt.Sprintf("init mysql error:%v", err))
	}
	gormDB = db.GetDB()

	// 启动 grpc
	logger.L.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   cfg.App.Name,
		Addr:   cfg.App.Addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			memberpb.RegisterMemberServiceServer(s, &service.MemberService{
				Logger:    logger,
				MemberDao: dao.NewMember(gormDB),
			})
		},
	}))
}
