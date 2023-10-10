package initialize

import (
	"github.com/lwzphper/go-mall/bff/global"
	addresspb "github.com/lwzphper/go-mall/server/member/api/gen/v1/address"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1/member"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSrvConn() {
	memberSrv()
	addressSrv()
}

func memberSrv() {
	//fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
	conn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		global.L.Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	global.MemberSrvClient = memberpb.NewMemberServiceClient(conn)
}

func addressSrv() {
	//fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
	conn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		global.L.Fatal("[InitSrvConn] 连接 【收货地址服务失败】")
	}
	global.AddressSrvClient = addresspb.NewAddressServiceClient(conn)
}
