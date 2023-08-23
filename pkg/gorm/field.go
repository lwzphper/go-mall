package gorm

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BigIdField struct {
	Id time.Time `json:"id" gorm:"column:id;primaryKey"`
}

type CreatedAtField struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

type UpdatedAtFiled struct {
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type SoftDeleteField struct {
	IsDelete soft_delete.DeletedAt `json:"is_delete" gorm:"column:is_delete;softDelete:flag"`
}
