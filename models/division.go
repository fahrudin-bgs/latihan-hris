package models

import "time"

type Division struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement"`
	DepartmentID uint64     `gorm:"not null"`
	Name         string     `gorm:"type:varchar(100);not null"`
	Description  string     `gorm:"type:varchar(100)"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`

	// Relasi
	Department Department `gorm:"foreignKey:DepartmentID;constraint:OnDelete:CASCADE;"`
}
