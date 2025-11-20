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
	ID          uint64       `json:"id"`
	EmployeeID  uint64       `json:"employee_id"`
	PositionID  *uint64      `json:"position_id"`
	Description *string      `json:"description"`
	AssignedAt  *string      `json:"assigned_at"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Position    *ResPosition `json:"position,omitempty"`
}

func ToResEmployeePosition(p models.EmployeePosition) ResEmployeePosition {
	var assignedAt *string = nil
	if p.AssignedAt != nil {
		str := p.AssignedAt.Format("2006-01-02")
		assignedAt = &str
	}

	var position *ResPosition = nil
	if p.PositionID != nil && *p.PositionID != 0 {
		position = &ResPosition{
			ID:          uint64(p.Position.ID),
			Name:        p.Position.Name,
			Description: p.Position.Description,
			CreatedAt:   p.Position.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   p.Position.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return ResEmployeePosition{
		ID:          p.ID,
		EmployeeID:  p.EmployeeID,
		PositionID:  p.PositionID,
		Description: p.Description,
		AssignedAt:  assignedAt,
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
		Position:    position,
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
