package models

import "time"

type RefreshToken struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"not null"`
	Token     string    `gorm:"type:text;not null"`
	IPAddress *string   `gorm:"type:varchar(45)"`
	UserAgent *string   `gorm:"type:text"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
