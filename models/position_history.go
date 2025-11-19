package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

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

func (p *PositionHistory) BeforeCreate(tx *gorm.DB) (err error) {
	return p.syncPositionName(tx)
}

func (p *PositionHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	return p.syncPositionName(tx)
}

// Helper untuk mengambil PositionName dari tabel positions
func (p *PositionHistory) syncPositionName(tx *gorm.DB) error {
	if p.PositionID == nil {
		return errors.New("position_id tidak boleh kosong")
	}

	var pos Position
	if err := tx.First(&pos, *p.PositionID).Error; err != nil {
		return err
	}

	p.PositionName = pos.Name
	return nil
}
