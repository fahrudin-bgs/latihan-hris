package dto

import "latihan-hris/models"

type ReqUploadPhoto struct {
	EmployeeID uint64 `form:"employee_id" binding:"required"`
	IsProfile  bool   `form:"is_profile"`
}

type ResEmployeePhoto struct {
	ID         uint64 `json:"id"`
	EmployeeID uint64 `json:"employee_id"`
	FilePath   string `json:"file_path"`
	IsProfile  bool   `json:"is_profile"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func ToResEmployeePhoto(photo models.EmployeePhoto) ResEmployeePhoto {
	return ResEmployeePhoto{
		ID:         photo.ID,
		EmployeeID: photo.EmployeeID,
		FilePath:   photo.FilePath,
		IsProfile:  photo.IsProfile,
		CreatedAt:  photo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  photo.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
