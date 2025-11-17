package models

import "time"

type Employee struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement"`
	UserID         uint64    `gorm:"not null"`
	Name           string    `gorm:"type:varchar(255);not null"`
	EmployeeNumber string    `gorm:"type:varchar(50);unique;not null"`
	EmployeeStatus string    `gorm:"type:enum('active','inactive','resigned');default:'active'"`
	JoinDate       time.Time `gorm:"not null"`
	EndDate        *time.Time
	DivisionID     *uint64   `gorm:""`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	// Relasi
	User           User            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Division       *Division       `gorm:"foreignKey:DivisionID;constraint:OnDelete:SET NULL;"`
	EmployeeDetail *EmployeeDetail `gorm:"foreignKey:EmployeeID"`
}
