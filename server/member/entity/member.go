package entity

import (
	"github.com/lwzphper/go-mall/pkg/gorm"
	"time"
)

type MemberStatus int8
type Gender int8
type SourceType int8

type Member struct {
	gorm.BigIdField
	MemberLevelId         uint64       `json:"member_level_id" gorm:"column:member_level_id"`
	Username              string       `json:"username" gorm:"column:username;uniqueIndex:uk_username"`
	Password              string       `json:"password" gorm:"column:password"`
	Nickname              string       `json:"nickname" gorm:"column:nickname"`
	Phone                 string       `json:"phone" gorm:"column:phone"`
	Status                MemberStatus `json:"status" gorm:"column:status"`
	Icon                  string       `json:"icon" gorm:"column:icon"`
	Gender                Gender       `json:"gender" gorm:"column:gender"`
	Birthday              *time.Time   `json:"birthday" gorm:"column:birthday;default:null"`
	City                  string       `json:"city" gorm:"column:city"`
	PersonalizedSignature string       `json:"personalized_signature" gorm:"column:personalized_signature"`
	SourceType            SourceType   `json:"source_type" gorm:"column:source_type"`
	Integration           int          `json:"integration" gorm:"column:integration"`
	Growth                int          `json:"growth" gorm:"column:growth"`
	LuckyCount            int          `json:"lucky_count" gorm:"column:lucky_count"`
	HistoryIntegration    int          `json:"history_integration" gorm:"column:history_integration"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (m Member) TableName() string {
	return "member"
}

const (
	StatusDisable MemberStatus = iota
	StatusEnable
)

var memberStatusNameMap = map[MemberStatus]string{
	StatusDisable: "禁用",
	StatusEnable:  "启用",
}

func (m MemberStatus) String() string {
	name, ok := memberStatusNameMap[m]
	if !ok {
		return "禁用"
	}
	return name
}

const (
	GenderUnknown Gender = iota
	GenderMan
	GenderWoman
)

var genderNameMap = map[Gender]string{
	GenderUnknown: "未知",
	GenderMan:     "男",
	GenderWoman:   "女",
}

func (g Gender) String() string {
	name, ok := genderNameMap[g]
	if !ok {
		return "未知"
	}
	return name
}

func (s SourceType) String() string {
	return ""
}
