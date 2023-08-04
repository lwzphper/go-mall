package gorm

import (
	"gorm.io/plugin/soft_delete"
)

type CreatedAtField struct {
	CreatedAt int64                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64                 `json:"updated_at" gorm:"autoUpdateTime"`
	IsDelete  soft_delete.DeletedAt `json:"is_delete" gorm:"softDelete:flag"`
}
