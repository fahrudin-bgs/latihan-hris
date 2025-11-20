package dto

import (
	"latihan-hris/models"
)

type ResPositionHistory struct {
	ID           uint64  `json:"id"`
	EmployeeID   uint64  `json:"employee_id"`
	PositionID   uint64  `json:"position_id"`
	PositionName string  `json:"position_name"`
	StartDate    string  `json:"start_date"`
	EndDate      *string `json:"end_date"`
	Description  *string `json:"description"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

func ToResPositionHistory(history models.PositionHistory) ResPositionHistory {
	var enddate *string = nil
	if history.EndDate != nil {
		str := history.EndDate.Format("2006-01-02")
		enddate = &str
	}
	return ResPositionHistory{
		ID:           history.ID,
		EmployeeID:   history.EmployeeID,
		PositionID:   *history.PositionID,
		PositionName: history.PositionName,
		StartDate:    history.StartDate.Format("2006-01-02"),
		EndDate:      enddate,
		Description:  history.Description,
		CreatedAt:    history.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    history.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type ReqPositionHistory struct {
	PositionID  *uint64   `json:"position_id"`  
    StartDate   *string   `json:"start_date"`    
    EndDate     *string   `json:"end_date"`      
    Description *string   `json:"description"`   
}