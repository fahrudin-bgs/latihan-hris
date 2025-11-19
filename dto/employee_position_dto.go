package dto

import (
	"latihan-hris/models"
	"time"
)

type ReqEmployeePosition struct {
	EmployeeID  uint64 `json:"employee_id" form:"employee_id" binding:"required"`
	PositionID  uint64 `json:"position_id" form:"position_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	AssignedAt  string `json:"assigned_at" form:"assigned_at" binding:"required,datetime=2006-01-02"`
}

func ToModelEmployeePosition(req ReqEmployeePosition) models.EmployeePosition {
	var assigned_at time.Time
	assigned_at, _ = time.Parse("2006-01-02", req.AssignedAt)
	return models.EmployeePosition{
		EmployeeID:  req.EmployeeID,
		PositionID:  &req.PositionID,
		Description: &req.Description,
		AssignedAt:  &assigned_at,
	}
}

type ResEmployeePosition struct {
	ID          uint64 `json:"id"`
	EmployeeID  uint64 `json:"employee_id"`
	PositionID  uint64 `json:"position_id"`
	Description string `json:"description"`
	AssignedAt  string `json:"assigned_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToResEmployeePosition(p models.EmployeePosition) ResEmployeePosition {
	return ResEmployeePosition{
		ID:          p.ID,
		EmployeeID:  p.EmployeeID,
		PositionID:  *p.PositionID,
		Description: *p.Description,
		AssignedAt:  p.AssignedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUpdateEmployeePosition(p *models.EmployeePosition, req ReqEmployeePosition) {
	p.PositionID = &req.PositionID
	p.Description = &req.Description
	
	assignedAt, err := time.Parse("2006-01-02", req.AssignedAt)
	if err == nil {
		p.AssignedAt = &assignedAt
	}
}
