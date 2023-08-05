package dao

import "github.com/lwzphper/go-mall/pkg/gorm"

type MemberLevel struct {
	gorm.BigIdField
	Name                 string `json:"name" gorm:"column:username,uniqueIndex:uk_username"`
	GrowthPoint          int    `json:"growth_point" gorm:"column:growth_point"`
	IsDefault            int8   `json:"is_default" gorm:"column:is_default"`
	CommentGrowthPoint   int8   `json:"comment_growth_point" gorm:"column:comment_growth_point"`
	PrivilegeFreeFreight int8   `json:"privilege_free_freight" gorm:"column:privilege_free_freight"`
	PrivilegeSignIn      int8   `json:"privilege_sign_in" gorm:"column:privilege_sign_in"`
	PrivilegeComment     int8   `json:"privilege_comment" gorm:"column:privilege_comment"`
	PrivilegePromotion   int8   `json:"privilege_promotion" gorm:"column:privilege_promotion"`
	PrivilegeMemberPrice int8   `json:"privilege_member_price" gorm:"column:privilege_member_price"`
	PrivilegeBirthday    int8   `json:"privilege_birthday" gorm:"column:privilege_birthday"`
	gorm.CreatedAtField
	gorm.UpdatedAtFiled
	gorm.SoftDeleteField
}

func (l MemberLevel) TableName() string {
	return "member_level"
}
