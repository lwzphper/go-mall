package server

import (
	"fmt"
	"github.com/lwzphper/go-mall/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

type GRPCConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(*grpc.Server)
	Logger       *logger.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption

	// 设置拦截器
	/*if c.AuthPublicKeyFile != "" {
		interceptor, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create auth Interceptor", zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}*/

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)

	// 监听中断信号，优雅关闭服务
	hook := NewHook()
	go hook.Close(func(sg os.Signal) {
		switch v := sg.(type) {
		// todo 根据不同的类型，做不同的处理
		default:
			c.Logger.Info(fmt.Sprintf("receive signal '%v', start graceful shutdown", v.String()))
			s.Stop()
			return
		}
	})

	c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
