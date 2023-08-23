package dao

import (
	"context"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/entity"
	"gorm.io/gorm"
)

type Member struct {
	db *gorm.DB
}

func NewMember(db *gorm.DB) *Member {
	return &Member{
		db: db.Model(&entity.Member{}),
	}
}

// MemberRecord 定义 member 记录
type MemberRecord struct {
	memberpb.BasicInfo
}

// CreateMember 创建会员
func (m *Member) CreateMember(ctx context.Context, member *entity.Member) error {
	return m.db.Create(member).Error
}

// GetMemberByUsername 通过用户名获取用户信息
func (m *Member) GetMemberByUsername(ctx context.Context, username string) (*MemberRecord, error) {
	record := &MemberRecord{}
	return record, m.db.Where("username", username).First(record).Error
}
