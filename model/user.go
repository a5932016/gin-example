package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        uint                  `json:"-" gorm:"primaryKey"`
	Name      string                `json:"name" gorm:"not null"`
	Email     string                `json:"email" gorm:"not null;uniqueIndex:idx_email"`
	Password  string                `json:"password" gorm:"not null;size:255"`
	CreatedAt time.Time             `json:"createdAt" example:"2022-11-03"`
	UpdatedAt time.Time             `json:"updatedAt" example:"2022-11-03"`
	IsDel     soft_delete.DeletedAt `json:"isDel" example:"0" gorm:"not null;uniqueIndex:idx_email;softDelete:,DeletedAtField:DeletedAt;default:0"`
}
