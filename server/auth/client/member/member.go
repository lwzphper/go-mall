package member

import (
	"context"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
)

// Manager defines a member manager.
type Manager struct {
	MemberService memberpb.MemberServiceClient
}

// RegisterMemberInfo 注册用户信息
func (m *Manager) RegisterMemberInfo(c context.Context, username, password string) error {
	return nil
}

// LoginVerify 登录验证
func (m *Manager) LoginVerify(c context.Context, username, password string) error {
	return nil
}
