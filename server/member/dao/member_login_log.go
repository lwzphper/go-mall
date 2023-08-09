package dao

import "github.com/lwzphper/go-mall/pkg/gorm"

type LoginType int8

const (
	LoginTypePc LoginType = iota + 1
	LoginTypeAndroid
	LoginTypeIos
	LoginTypeMinProgram
)

var loginTypeNameMap = map[LoginType]string{
	LoginTypePc:         "pc",
	LoginTypeAndroid:    "android",
	LoginTypeIos:        "ios",
	LoginTypeMinProgram: "小程序",
}

type MemberLoginLog struct {
	gorm.BigIdField
	LoginType string `json:"login_type" gorm:"column:login_type"`
	Ip        string `json:"ip" gorm:"column:ip"`
	Province  string `json:"province" gorm:"column:province"`
	City      string `json:"city" gorm:"column:city"`
	MemberId  uint64 `json:"member_id" gorm:"column:member_id"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (l MemberLoginLog) TableName() string {
	return "member_login_log"
}

func (l LoginType) String() string {
	name, ok := loginTypeNameMap[l]
	if !ok {
		return "未知"
	}
	return name
}
