package server

import (
	"github.com/lwzphper/go-mall/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
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

	c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
