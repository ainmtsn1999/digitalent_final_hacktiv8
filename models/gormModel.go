package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
