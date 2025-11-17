package dto

import (
	"latihan-hris/models"

	"golang.org/x/crypto/bcrypt"
)

type ReqUser struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	RoleID     uint64 `json:"role_id" form:"role_id" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
}

func ToModelUser(req ReqUser) models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	return models.User{
		Username: req.Username,
		Email:    req.Email,
		RoleID:   &req.RoleID,
		Password: string(hashedPassword),
	}
}

type ResUser struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResUser(user models.User) ResUser {
	roleName := user.Role.Name
	return ResUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      roleName,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type ResUserDetail struct {
	ID        uint64   `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Role      *ResRole `json:"role,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func ToResUserDetail(user models.User) ResUserDetail {
	var role *ResRole
	if user.Role.ID != 0 {
		role = &ResRole{
			ID:        user.Role.ID,
			Name:      user.Role.Name,
			CreatedAt: user.Role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.Role.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

	}
	return ResUserDetail{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUpdateUser(user *models.User, req ReqUser) {
	user.Username = req.Username
	user.Email = req.Email
	user.RoleID = &req.RoleID

	if req.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}
}
