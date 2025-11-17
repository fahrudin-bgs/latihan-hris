package controllers

import (
	"fmt"
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadPhoto(c *gin.Context) {
	var req dto.ReqUploadPhoto
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.FindOrError(config.DB, &models.Employee{}, req.EmployeeID, c, "Employee ID Not Found"); err != nil {
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Photo file is required")
		return
	}

	if !utils.ValidateFile(file, []string{".jpg", ".jpeg", ".png"}) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file extension")
		return
	}

	folderPath := filepath.Join(config.UploadPath, "employees/images")
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create upload directory")
		return
	}

	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d_%d%s", req.EmployeeID, time.Now().Unix(), ext)
	dstPath := filepath.Join(folderPath, newFileName)

	if err := c.SaveUploadedFile(file, dstPath); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to save file")
		return
	}

	photo := models.EmployeePhoto{
		EmployeeID: req.EmployeeID,
		FilePath:   dstPath,
		IsProfile:  req.IsProfile,
	}

	if err := config.DB.Create(&photo).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Photo upload successfully", nil)
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	var photo models.EmployeePhoto

	if err := config.DB.First(&photo, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	filePath := photo.FilePath

	if err := config.DB.Delete(&photo).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if photo.FilePath != "" {
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete file from disk")
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Photo Deleted Successfully", nil)
}

func UpdateIsProfile(c *gin.Context) {
	id := c.Param("id")

	var photo models.EmployeePhoto

	if err := config.DB.First(&photo, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		employeeId := photo.EmployeeID

		if err := tx.Model(&models.EmployeePhoto{}).
			Where("employee_id = ? AND id <> ?", employeeId, photo.ID).
			Update("is_profile", false).Error; err != nil {
			return err
		}

		return tx.Model(&photo).Update("is_profile", true).Error
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Status updated successfully", nil)
}

func GetPhotos(c *gin.Context) {
	employeeId := c.Param("employee_id")

	var photos []models.EmployeePhoto

	if err := config.DB.Where("employee_id = ?", employeeId).
		Find(&photos).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResEmployeePhoto
	for _, u := range photos {
		res = append(res, dto.ToResEmployeePhoto(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Get Photo Employee", res)
}

func GetPhotoProfile(c *gin.Context) {
	employeeId := c.Param("employee_id")

	var photo models.EmployeePhoto

	if err := config.DB.Where("employee_id = ?", employeeId).
		Where("is_profile = ?", true).
		First(&photo).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResEmployeePhoto(photo)

	utils.SuccessResponse(c, http.StatusOK, "Get Photo Profile", res)
}
