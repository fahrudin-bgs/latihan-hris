package models

import (
	"time"
)

type EmployeeDocument struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	EmployeeID  uint64    `gorm:"not null"`
	FileType    string    `gorm:"type:varchar(255);not null"`
	Description *string   `gorm:"type:varchar(255)"`
	FilePath    string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE"`
}
