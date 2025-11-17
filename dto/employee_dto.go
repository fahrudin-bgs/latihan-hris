package dto

import (
	"latihan-hris/models"
	"time"
)

type ReqEmployee struct {
	UserID         uint64  `json:"user_id" form:"user_id"`
	Name           string  `json:"name" form:"name"`
	EmployeeNumber string  `json:"employee_number" form:"employee_number"`
	EmployeeStatus string  `json:"employee_status" form:"employee_status" binding:"omitempty,oneof=active inactive resigned"`
	JoinDate       string  `json:"join_date" form:"join_date" binding:"required,datetime=2006-01-02"`
	EndDate        *string `json:"end_date" form:"end_date" binding:"omitempty,datetime=2006-01-02"`
	DivisionID     *uint64 `json:"division_id" form:"division_id"`
}

func ToModelEmployee(req ReqEmployee) models.Employee {
	var joinDate time.Time
	var endDate *time.Time
	var divisionID *uint64

	// Parsing join_date (gunakan default zero time jika gagal)
	joinDate, _ = time.Parse("2006-01-02", req.JoinDate)

	// Parsing end_date (jika ada)
	if req.EndDate != nil && *req.EndDate != "" {
		if t, err := time.Parse("2006-01-02", *req.EndDate); err == nil {
			endDate = &t
		}
	}

	if req.DivisionID != nil && *req.DivisionID != 0 {
		divisionID = req.DivisionID
	} else {
		divisionID = nil
	}

	return models.Employee{
		UserID:         req.UserID,
		Name:           req.Name,
		EmployeeNumber: req.EmployeeNumber,
		EmployeeStatus: req.EmployeeStatus,
		JoinDate:       joinDate,
		EndDate:        endDate,
		DivisionID:     divisionID,
	}
}

type ResEmployee struct {
	ID             uint64             `json:"id"`
	UserID         uint64             `json:"user_id"`
	User           *ResUser           `json:"user,omitempty"`
	Name           string             `json:"name"`
	EmployeeNumber string             `json:"employee_number"`
	EmployeeStatus string             `json:"employee_status"`
	JoinDate       string             `json:"join_date"`
	EndDate        *string            `json:"end_date"`
	Division       *string            `json:"division"`
	EmployeeDetail *ResEmployeeDetail `json:"employee_detail,omitempty"`
	CreatedAt      string             `json:"created_at"`
	UpdatedAt      string             `json:"updated_at"`
}

func ToResEmployee(employee models.Employee) ResEmployee {
	var endDate *string
	if employee.EndDate != nil {
		formatted := employee.EndDate.Format("2006-01-02")
		endDate = &formatted
	}

	var division *string
	if employee.Division != nil {
		division = &employee.Division.Name
	}

	var user *ResUser
	if employee.User.ID != 0 {
		user = &ResUser{
			ID:        employee.User.ID,
			Username:  employee.User.Username,
			Email:     employee.User.Email,
			Role:      employee.User.Role.Name,
			CreatedAt: employee.User.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.User.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	var employeeDetail *ResEmployeeDetail
	if employee.EmployeeDetail.ID != 0 {
		var birthDateStr *string
		if employee.EmployeeDetail.BirthDate != nil {
			formatted := employee.EmployeeDetail.BirthDate.Format("2006-01-02")
			birthDateStr = &formatted
		}

		employeeDetail = &ResEmployeeDetail{
			ID:          employee.EmployeeDetail.ID,
			EmployeeID:  employee.EmployeeDetail.EmployeeID,
			Gender:      employee.EmployeeDetail.Gender,
			BirthDate:   birthDateStr,
			PhoneNumber: employee.EmployeeDetail.PhoneNumber,
			Address:     employee.EmployeeDetail.Address,
			CreatedAt:   employee.EmployeeDetail.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   employee.EmployeeDetail.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return ResEmployee{
		ID:             employee.ID,
		UserID:         employee.UserID,
		User:           user,
		Name:           employee.Name,
		EmployeeNumber: employee.EmployeeNumber,
		EmployeeStatus: employee.EmployeeStatus,
		JoinDate:       employee.JoinDate.Format("2006-01-02"),
		EndDate:        endDate,
		Division:       division,
		EmployeeDetail: employeeDetail,
		CreatedAt:      employee.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      employee.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type ResEmployeeSimple struct {
	ID             uint64  `json:"id"`
	UserID         uint64  `json:"user_id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	EmployeeNumber string  `json:"employee_number"`
	EmployeeStatus string  `json:"employee_status"`
	Division       *string `json:"division"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func ToResEmployeeSimple(employee models.Employee) ResEmployeeSimple {
	var email string

	if employee.User.ID != 0 {
		email = employee.User.Email
	}

	var division *string

	if employee.Division != nil {
		division = &employee.Division.Name
	} else {
		division = nil
	}
	return ResEmployeeSimple{
		ID:             employee.ID,
		UserID:         employee.UserID,
		Name:           employee.Name,
		Email:          email,
		EmployeeNumber: employee.EmployeeNumber,
		EmployeeStatus: employee.EmployeeStatus,
		Division:       division,
		CreatedAt:      employee.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      employee.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUpdateEmployee(employee *models.Employee, req ReqEmployee) {
	employee.Name = req.Name
	employee.EmployeeNumber = req.EmployeeNumber
	employee.EmployeeStatus = req.EmployeeStatus
	if req.JoinDate != "" {
		parsedJoinDate, _ := time.Parse("2006-01-02", req.JoinDate)
		employee.JoinDate = parsedJoinDate
	}

	if req.EndDate != nil && *req.EndDate != "" {
		parsedEndDate, _ := time.Parse("2006-01-02", *req.EndDate)
		employee.EndDate = &parsedEndDate
	}

	if req.DivisionID != nil && *req.DivisionID != 0 {
		employee.DivisionID = req.DivisionID
	} else {
		employee.DivisionID = nil
	}
}
