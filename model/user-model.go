package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type Role struct {
	ID   int    `gorm:"column:id; varchar(255); primary_key; not null" json:"id"`
	Role string `gorm:"column:role" json:"role"`
	BaseModel
}

type User struct {
	ID       int    `gorm:"column:id; varchar(255); unique; not null" json:"id"`
	Name     string `gorm:"column:name; unique" json:"name"`
	Email    string `gorm:"column:email; unique" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Status   int    `gorm:"column:status" json:"status"`
	RoleID   int    `gorm:"column:role_id;not null; serial" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
