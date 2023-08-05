package member

import "github.com/lwzphper/go-mall/pkg/gorm"

type Member struct {
	Id                    uint64       `json:"id" gorm:"primaryKey"`
	MemberLevelId         uint64       `json:"member_level_id"`
	Username              string       `json:"username"`
	Password              string       `json:"password"`
	Nickname              string       `json:"nickname"`
	Phone                 string       `json:"phone"`
	Status                MemberStatus `json:"status"`
	Icon                  string       `json:"icon"`
	Gender                Gender       `json:"gender"`
	Birthday              string       `json:"birthday"`
	City                  string       `json:"city"`
	PersonalizedSignature string       `json:"personalized_signature"`
	SourceType            SourceType   `json:"source_type"`
	Integration           int          `json:"integration"`
	Growth                int          `json:"growth"`
	LuckCount             int          `json:"luck_count"`
	HistoryIntegration    int          `json:"history_integration"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (Member) TableName() string {
	return "user"
}
