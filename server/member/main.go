package main

import (
	"github.com/lwzphper/go-mall/pkg/server"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
	address2 "github.com/lwzphper/go-mall/server/member/dao/address"
	"github.com/lwzphper/go-mall/server/member/dao/member"
	"github.com/lwzphper/go-mall/server/member/global"
	"github.com/lwzphper/go-mall/server/member/initialize"
	"github.com/lwzphper/go-mall/server/member/service/address"
	member2 "github.com/lwzphper/go-mall/server/member/service/member"
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
			// 会员服务
			memberpb.RegisterMemberServiceServer(s, &member2.MemberService{
				Logger:    global.Logger,
				MemberDao: member.NewMember(),
			})
			// 地址服务
			addresspb.RegisterAddressServiceServer(s, &address.Service{
				Logger:     global.Logger,
				AddressDao: address2.NewAddress(),
			})
		},
	}))
}
