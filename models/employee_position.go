package models

import "time"

type EmployeePosition struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	EmployeeID  uint64 `gorm:"not null"`
	PositionID  *uint64
	Description *string    `gorm:"type:text"`
	AssignedAt  *time.Time `gorm:"type:date"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`

	// Relasi ke model Employee
	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE;"`
	Position Position `gorm:"foreignKey:PositionID;constraint:OnDelete:SET NULL;"`
}
