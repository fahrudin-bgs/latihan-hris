package models

import "time"

type User struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	Username   string    `gorm:"type:varchar(100);not null"`
	Email      string    `gorm:"type:varchar(100);not null;unique"`
	Password   string    `gorm:"type:varchar(255);not null"`
	RoleID     *uint64   `gorm:""`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// Relasi
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL;"`
}
