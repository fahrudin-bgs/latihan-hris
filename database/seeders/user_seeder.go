package seeders

import (
	"latihan-hris/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeds(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	users := []models.User{
		{
			Username: "superadmin",
			Email:    "superadmin@example.com",
			Password: string(hashedPassword),
			RoleID:   uint64Ptr(1),
		},
		{
			Username: "user",
			Email:    "user@example.com",
			Password: string(hashedPassword),
			RoleID:   uint64Ptr(2),
		},
	}

	for _, user := range users {
		var existing models.User
		if err := db.Where("email = ?", user.Email).First(&existing).Error; err == gorm.ErrRecordNotFound {
			db.Create(&user)
		}
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}
