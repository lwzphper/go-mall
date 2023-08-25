package service

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/logger"
	authpb "github.com/lwzphper/go-mall/server/auth/api/gen/v1"
)

type AuthService struct {
	authpb.UnimplementedAuthServiceServer
	MemberAuth   MemberAuth
	MemberManger MemberManger
	Logger       *logger.Logger
}

// MemberManger defines the ACL for member management. 会员管理 防止入侵层
type MemberManger interface {
	RegisterMemberInfo(c context.Context, username, password string) error
	LoginVerify(ctx context.Context, username, password string) error
}

type MemberAuth interface {
	// CheckMemberPassword 检查会员密码是否正确
	CheckMemberPassword(c context.Context, name, password string) (bool, error)
}

// Register 注册
func (a *AuthService) Register(context.Context, *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return nil, nil
}

// Login 登录
func (a *AuthService) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return nil, nil
}
