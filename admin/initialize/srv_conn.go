package initialize

func InitSrvConn() {
	//memberSrv()
}

/*func memberSrv() {
	//fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
	conn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		global.L.Fatal("[InitSrvConn] 连接 【会员服务失败】")
	}
	global.MemberSrvClient = memberpb.NewMemberServiceClient(conn)
}*/
