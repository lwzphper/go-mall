package member

import (
	"context"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
)

// Manager defines a member manager.
type Manager struct {
	memberService memberpb.MemberServiceClient
}

// LoginVerify 登录验证
func (m *Manager) LoginVerify(c context.Context, username, password string) error {
	return nil
}
