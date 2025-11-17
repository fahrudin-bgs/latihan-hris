package dto

import "latihan-hris/models"

type ReqPosition struct {
	Name        string  `json:"name" form:"name" binding:"required"`
	Description *string `json:"description" form:"description"`
}

func ToModelPosition(req ReqPosition) models.Position {
	return models.Position{
		Name:        req.Name,
		Description: req.Description,
	}
}

type ResPosition struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ToResPosition(p models.Position) ResPosition {
	return ResPosition{
		ID:          uint64(p.ID),
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUpdatePosition(p *models.Position, req ReqPosition) {
	p.Name = req.Name
	p.Description = req.Description
}
