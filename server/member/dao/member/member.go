package member

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/entity"
	"gorm.io/gorm"
)

type Member struct {
	db *gorm.DB
}

func NewMember(ctx context.Context) *Member {
	return &Member{
		db: dao.GetDB(ctx),
	}
}

// CreateMember 创建会员
func (m *Member) CreateMember(ctx context.Context, member *entity.Member) error {
	return m.db.WithContext(ctx).Create(member).Error
}

// GetItemByWhere 查询用户信息
func (m *Member) GetItemByWhere(ctx context.Context, where *entity.Member) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.WithContext(ctx).Where(where).Take(record).Error
}

// GetItemById GetItemById 查询用户信息
func (m *Member) GetItemById(ctx context.Context, id id.MemberID) (*entity.Member, error) {
	record := &entity.Member{}
	return record, m.db.WithContext(ctx).Take(record, id.Uint64()).Error
}

// UpdateById 通过id更新
func (m *Member) UpdateById(ctx context.Context, id id.MemberID, data map[string]interface{}) error {
	return m.db.WithContext(ctx).Model(&entity.Member{}).Where("id", id).Save(data).Error
}

// UpdateByEntity 通过实体更新（id 不存在会新增）
func (m *Member) UpdateByEntity(ctx context.Context, member *entity.Member) error {
	return m.db.WithContext(ctx).Save(member).Error
}
