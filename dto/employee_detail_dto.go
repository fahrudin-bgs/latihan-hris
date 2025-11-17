package dto

import (
	"latihan-hris/models"
	"time"
)

type ReqEmployeeDetail struct {
	Gender      string `json:"gender" form:"gender"`
	BirthDate   string `json:"birth_date" form:"birth_date"  binding:"datetime=2006-01-02"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

func ToUpdateEmployeeDetail(empdetail *models.EmployeeDetail, req ReqEmployeeDetail) {
	empdetail.Gender = req.Gender
	if req.BirthDate != "" {
		parsedJoinDate, _ := time.Parse("2006-01-02", req.BirthDate)
		empdetail.BirthDate = &parsedJoinDate
	}
	empdetail.PhoneNumber = req.PhoneNumber
	empdetail.Address = req.Address
}

type ResEmployeeDetail struct {
	ID          uint64  `json:"id"`
	EmployeeID  uint64  `json:"employee_id"`
	Gender      string  `json:"gender"`
	BirthDate   *string `json:"birth_date"`
	PhoneNumber string  `json:"phone_number"`
	Address     string  `json:"address"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ToResEmployeeDetail(model models.EmployeeDetail) ResEmployeeDetail {
	var birthDateStr *string
	if model.BirthDate != nil {
		formatted := model.BirthDate.Format("2006-01-02")
		birthDateStr = &formatted
	}

	res := ResEmployeeDetail{
		ID:          model.ID,
		EmployeeID:  model.EmployeeID,
		Gender:      model.Gender,
		BirthDate:   birthDateStr,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
		CreatedAt:   model.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return res
}
