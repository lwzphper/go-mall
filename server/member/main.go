package main

import (
	"github.com/lwzphper/go-mall/pkg/server"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/lwzphper/go-mall/server/member/initialize"
	"github.com/lwzphper/go-mall/server/member/service"
	"google.golang.org/grpc"
)

func main() {
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()

	// 启动 grpc
	global.Logger.L.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   global.Config.App.Name,
		Addr:   global.Config.App.Addr,
		Logger: global.Logger,
		RegisterFunc: func(s *grpc.Server) {
			memberpb.RegisterMemberServiceServer(s, &service.MemberService{
				Logger:    global.Logger,
				MemberDao: dao.NewMember(),
			})
		},
	}))
}
