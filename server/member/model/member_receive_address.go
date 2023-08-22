package model

import "github.com/lwzphper/go-mall/pkg/gorm"

type MemberReceiveAddress struct {
	gorm.BigIdField
	Name          string `json:"name" gorm:"column:name"`
	Phone         string `json:"phone_number" gorm:"column:phone_number"`
	IsDefault     int8   `json:"is_default" gorm:"column:is_default"`
	PostCode      string `json:"post_code" gorm:"column:post_code"`
	Province      string `json:"province" gorm:"column:province"`
	City          string `json:"city" gorm:"column:city"`
	Region        string `json:"region" gorm:"column:region"`
	DetailAddress string `json:"detail_address" gorm:"column:detail_address"`
	MemberId      uint64 `json:"member_id" gorm:"column:member_id"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (a MemberReceiveAddress) TableName() string {
	return "member_receive_address"
}
