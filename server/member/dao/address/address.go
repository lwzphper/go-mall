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
	result := m.db.Select([]string{"id", "name", "phone", "is_default", "post_code", "province", "city", "region", "address"}).
		Where(&entity.Address{MemberId: mId.Uint64()}).
		Find(&address)
	return address, result.Error
}

// GetItemById 通过 id 获取数据
func (m *Address) GetItemById(ctx context.Context, id id.AddressID) (*entity.Address, error) {
	item := &entity.Address{}
	return item, m.db.Take(item, id.Uint64()).Error
}

// UpdateById 通过id更新
func (m *Address) UpdateById(ctx context.Context, id id.AddressID, data map[string]interface{}) error {
	return m.db.Model(&entity.Address{}).Where("id", id).Save(data).Error
}

// UpdateUserItem 更新用户数据
func (m *Address) UpdateUserItem(ctx context.Context, mId id.MemberID, aId id.AddressID, data entity.Address) error {
	return m.db.Model(&entity.Address{}).
		Select("name", "phone", "is_default", "post_code", "province", "city", "region", "address").
		Where("member_id", mId).
		Where("id", aId).
		Updates(data).
		Error
}

// DeleteById 删除数据
func (m *Address) DeleteById(ctx context.Context, aId id.AddressID) error {
	var address entity.Address
	address.Id = aId.Uint64()
	return m.db.Delete(&address).Error
}
