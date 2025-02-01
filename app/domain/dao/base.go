package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}
