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
		db: db.Model(&entity.Member{}).Session(&gorm.Session{}),
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

// GetItem 查询用户信息
func (m *Member) GetItem(ctx context.Context, where *entity.Member) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.Where(where).First(record).Error
}
