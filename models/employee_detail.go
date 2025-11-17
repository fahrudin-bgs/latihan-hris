package models

import "time"

type EmployeeDetail struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement"`
	EmployeeID  uint64     `gorm:"not null"`
	Gender      string     `gorm:"type:enum('Male','Female');default:null"`
	BirthDate   *time.Time `gorm:"type:date;default:null"`
	PhoneNumber string     `gorm:"type:varchar(20);default:null"`
	Address     string     `gorm:"type:text;default:null"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`

	// Relasi
	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE;"`
}
