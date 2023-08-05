package gorm

import (
	"gorm.io/plugin/soft_delete"
)

type CreatedAtField struct {
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
}

type UpdatedAtFiled struct {
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}

type SoftDeleteField struct {
	IsDelete soft_delete.DeletedAt `json:"is_delete" gorm:"softDelete:flag"`
}
