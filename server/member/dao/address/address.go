package address

import (
	"context"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/server/member/dao"
	"github.com/lwzphper/go-mall/server/member/entity"
	"gorm.io/gorm"
)

type Address struct {
	db *gorm.DB
}

func NewAddress(ctx context.Context) *Address {
	return &Address{
		db: dao.GetDB(ctx),
	}
}

var allowUpdateColumn = []string{"name", "phone", "is_default", "post_code", "province", "city", "region", "detail"}

// Create 创建
func (m *Address) Create(ctx context.Context, address *entity.Address) error {
	return m.db.Create(address).Error
}

// GetList 获取列表
func (m *Address) GetList(ctx context.Context, mId id.MemberID) ([]entity.Address, error) {
	var address []entity.Address
	result := m.db.Select([]string{"id", "name", "phone", "is_default", "post_code", "province", "city", "region", "detail"}).
		Where(&entity.Address{MemberId: mId.Uint64()}).
		Order("is_default desc,id desc").
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

// UpdateByMemberId 通过会员id更新数据
func (m *Address) UpdateByMemberId(ctx context.Context, mId id.MemberID, data map[string]interface{}) error {
	return m.db.Model(&entity.Address{}).Select(allowUpdateColumn).
		Where("member_id", mId).
		Updates(data).Error
}

// UpdateUserItem 更新用户数据
func (m *Address) UpdateUserItem(ctx context.Context, data entity.Address) error {
	return m.db.Model(&entity.Address{}).
		Select(allowUpdateColumn).
		Where("member_id", data.MemberId).
		Where("id", data.Id).
		Updates(data).
		Error
}

// DeleteById 删除数据
func (m *Address) DeleteById(ctx context.Context, aId id.AddressID) error {
	var address entity.Address
	address.Id = aId.Uint64()
	return m.db.Delete(&address).Error
}
