package models

import (
	"gorm.io/gorm"
	"time"
)

type Cliente struct {
	ID_Cliente        uint           `gorm:"primaryKey"`
	Nombre            string         `gorm:"not null"`
	CorreoElectronico string         `gorm:"unique;not null"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at"`
	UpdatedAt         time.Time      `gorm:"column:update_at"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
}
