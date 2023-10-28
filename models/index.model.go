package models

import (
	"time"

	"gorm.io/gorm"
)

// Custom JSON marshaling for the User struct

type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	GormModel
	UserSchema
	Photo Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Photo struct {
	GormModel
	PhotoSchema
	UserID uint `json:"user_id"`
}
