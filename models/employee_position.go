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
}
