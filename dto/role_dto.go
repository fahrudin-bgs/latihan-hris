package dto

import "latihan-hris/models"

type ReqRole struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

func ToModelRole(req ReqRole) models.Role {
	return models.Role{
		Name:        req.Name,
		Description: req.Description,
	}
}

type ResRole struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResRole(role models.Role) ResRole {
	return ResRole{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
