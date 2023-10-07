package address

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/server/member/entity"
	"github.com/lwzphper/go-mall/server/member/global"
	"gorm.io/gorm"
)

type Address struct {
	db *gorm.DB
}

func NewAddress() *Address {
	return &Address{
		db: global.DB,
	}
}

// Create 创建
func (m *Address) Create(ctx context.Context, address *entity.Address) error {
	return m.db.Create(address).Error
}

// GetList 获取列表
func (m *Address) GetList(ctx context.Context, mId id.MemberID) ([]entity.Address, error) {
	var address []entity.Address
	result := m.db.Select([]string{"name", "phone", "is_default", "post_code", "province", "city", "region", "address"}).
		Where(&entity.Address{MemberId: mId.Uint64()}).
		Find(&address)
	return address, result.Error
}

// UpdateById 通过id更新
func (m *Address) UpdateById(ctx context.Context, id id.AddressID, data map[string]interface{}) error {
	return m.db.Model(&entity.Address{}).Where("id", id).Save(data).Error
}

// UpdateByWhere 通过条件更新
func (m *Address) UpdateByWhere(ctx context.Context, where map[string]interface{}, data map[string]interface{}) error {
	// todo 更新操作
	return nil
}
