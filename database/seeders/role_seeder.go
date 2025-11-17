package seeders

import (
	"latihan-hris/models"

	"gorm.io/gorm"
)

func RoleSeeds(db *gorm.DB) {
	roles := []models.Role{
		{Name: "superadmin", Description: "-"},
		{Name: "user", Description: "-"},
	}

	for _, role := range roles {
		var existing models.Role
		if err := db.Where("name = ?", role.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			db.Create(&role)
		}
	}
}
