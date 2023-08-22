package dao

import (
	"context"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"gorm.io/gorm"
)

type Member struct {
	db *gorm.DB
}

func NewMember(db *gorm.DB) *Member {
	return &Member{
		db: db,
	}
}

// MemberRecord 定义 member 记录
type MemberRecord struct {
	Member memberpb.BasicInfo
}

// GetMemberByUsername 通过用户名获取用户信息
func (m *Member) GetMemberByUsername(ctx context.Context, username string) (*MemberRecord, error) {
	return nil, nil
}
