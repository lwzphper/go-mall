package service

import (
	"github.com/lwzphper/go-mall/pkg/logger"
	authpb "github.com/lwzphper/go-mall/server/auth/api/gen/v1"
)

type AuthService struct {
	authpb.UnimplementedAuthServiceServer
	Logger *logger.Logger
}
