package dto

import "time"

type ReqEmployeeDocument struct {
	EmployeeID  uint64  `form:"employee_id" binding:"required"`
	FileType    string  `form:"file_type" binding:"required"`
	Description *string `form:"description"`
}

type ResEmployeeDocument struct {
	ID          uint64    `json:"id"`
	EmployeeID  uint64    `json:"employee_id"`
	FileType    string    `json:"file_type"`
	Description *string   `json:"description,omitempty"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
