package entity

import "github.com/lwzphper/go-mall/pkg/gorm"

type Address struct {
	gorm.BigIdField
	Name      string `json:"name" gorm:"column:name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	IsDefault uint32 `json:"is_default" gorm:"column:is_default"`
	PostCode  string `json:"post_code" gorm:"column:post_code"`
	Province  string `json:"province" gorm:"column:province"`
	City      string `json:"city" gorm:"column:city"`
	Region    string `json:"region" gorm:"column:region"`
	Detail    string `json:"detail" gorm:"column:detail"`
	MemberId  uint64 `json:"member_id" gorm:"column:member_id"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (a Address) TableName() string {
	return "address"
}
