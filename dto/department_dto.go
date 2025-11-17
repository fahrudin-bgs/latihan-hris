package dto

import "latihan-hris/models"

type ReqDepartment struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

func ToModelDepartment(req ReqDepartment) models.Department {
	return models.Department{
		Name:        req.Name,
		Description: req.Description,
	}
}

type ResDepartment struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToResDepartment(department models.Department) ResDepartment {
	return ResDepartment{
		ID:          uint64(department.ID),
		Name:        department.Name,
		Description: department.Description,
		CreatedAt:   department.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   department.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
