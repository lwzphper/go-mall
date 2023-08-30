package dao

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	memberpb "github.com/lwzphper/go-mall/server/member/api/gen/v1"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/lwzphper/go-mall/server/member/global"
	"gorm.io/gorm"
)

type Member struct {
	db *gorm.DB
}

func NewMember() *Member {
	return &Member{
		db: global.DB,
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

// GetItemByWhere 查询用户信息
func (m *Member) GetItemByWhere(ctx context.Context, where *entity.Member) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.Where(where).First(record).Error
}

// GetItemById GetItemById 查询用户信息
func (m *Member) GetItemById(ctx context.Context, id id.MemberID) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.First(record, id.Uint64()).Error
}

// Update 更新
func (m *Member) Update(ctx context.Context, member *entity.Member) error {
	return m.db.Save(member).Error
}
