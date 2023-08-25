package main

import (
	"github.com/lwzphper/go-mall/pkg/common/config"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	cfgHelper "github.com/lwzphper/go-mall/pkg/config"
	"github.com/lwzphper/go-mall/pkg/logger"
	"github.com/lwzphper/go-mall/pkg/server"
	authpb "github.com/lwzphper/go-mall/server/auth/api/gen/v1"
	"github.com/lwzphper/go-mall/server/auth/client/member"
	"github.com/lwzphper/go-mall/server/auth/service"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config 配置文件
type Config struct {
	App     *app.App        `toml:"app" yaml:"app" mapstructure:"app"`
	Logging *config.Logging `toml:"logging" yaml:"logging" mapstructure:"logging"`
}

func main() {
	var err error

	// 初始化配置
	cfg := &Config{
		App:     app.NewDefaultApp(),
		Logging: config.NewDefaultLogging(),
	}
	err = cfgHelper.LoadConfigFromToml("server/auth/etc/config.yaml", cfg)
	if err != nil {
		panic("load config error: " + err.Error())
	}

	// 初始化日志
	log := cfg.Logging.InitLogger(cfg.App.Env)

	// 会员服务
	memberConn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect member service", zap.Error(err))
	}

	log.L.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   cfg.App.Name,
		Addr:   cfg.App.Addr,
		Logger: log,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &service.AuthService{
				Logger: log,
				MemberManger: &member.Manager{
					MemberService: memberpb.NewMemberServiceClient(memberConn),
				},
			})
		},
	}))
}
