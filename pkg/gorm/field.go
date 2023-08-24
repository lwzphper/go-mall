package gorm

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Timestamp time.Time

type BigIdField struct {
	Id uint64 `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
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
