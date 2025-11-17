package dto

import "latihan-hris/models"

type ReqDivision struct {
	DepartmentID uint64 `json:"department_id" form:"department_id" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	Description  string `json:"description" form:"description"`
}

func ToModelDivision(req ReqDivision) models.Division {
	return models.Division{
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		Description:  req.Description,
	}
}

type ResDivision struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DepartmenID uint64 `json:"department_id"`
	Department  string `json:"department"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToResDivision(division models.Division) ResDivision {
	return ResDivision{
		ID:          division.ID,
		Name:        division.Name,
		Description: division.Description,
		DepartmenID: division.DepartmentID,
		Department:  division.Department.Name,
		CreatedAt:   division.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   division.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
