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
	return record, m.db.Where(where).Take(record).Error
}

// GetItemById GetItemById 查询用户信息
func (m *Member) GetItemById(ctx context.Context, id id.MemberID) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.Take(record, id.Uint64()).Error
}

// UpdateById 通过id更新
func (m *Member) UpdateById(ctx context.Context, id id.MemberID, data map[string]interface{}) error {
	return m.db.Model(&entity.Member{}).Where("id", id).Save(data).Error
}

// UpdateByEntity 通过实体更新（id 不存在会新增）
func (m *Member) UpdateByEntity(ctx context.Context, member *entity.Member) error {
	return m.db.Save(member).Error
}
