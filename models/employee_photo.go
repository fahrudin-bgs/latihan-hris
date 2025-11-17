package models

import (
	"time"
)

type EmployeePhoto struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	EmployeeID uint64    `gorm:"not null"`
	FilePath   string    `gorm:"type:varchar(255);not null"`
	IsProfile  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// Relasi ke model Employee
	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE;"`
}
