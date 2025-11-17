package models

import "time"

type PositionHistory struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	EmployeeID   uint64 `gorm:"not null"`
	PositionID   *uint64
	PositionName string     `gorm:"type:varchar(255);not null"`
	StartDate    *time.Time `gorm:"type:date"`
	EndDate      *time.Time `gorm:"type:date"`
	Description  *string    `gorm:"type:text"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}
