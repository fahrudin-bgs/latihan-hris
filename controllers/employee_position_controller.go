package controllers

import (
	"errors"
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getLastHistoryPosition(employeeID uint64) (*models.PositionHistory, error) {
	var history models.PositionHistory

	if err := config.DB.
		Where("employee_id = ?", employeeID).
		Where("end_date IS NULL").
		Last(&history).Error; err != nil {
		return nil, err
	}

	return &history, nil
}

func CreateEmployeePosition(c *gin.Context) {
	var req dto.ReqEmployeePosition

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.FindOrError(config.DB, &models.Employee{}, req.EmployeeID, c, "Employee ID Not Found"); err != nil {
		return
	}

	if err := utils.FindOrError(config.DB, &models.Position{}, req.PositionID, c, "Position ID Not Found"); err != nil {
		return
	}

	employeePosition := dto.ToModelEmployeePosition(req)

	tx := config.DB.Begin()

	if err := tx.Create(&employeePosition).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	history := models.PositionHistory{
		EmployeeID:  employeePosition.EmployeeID,
		PositionID:  employeePosition.PositionID,
		StartDate:   employeePosition.AssignedAt,
		Description: employeePosition.Description,
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit().Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResEmployeePosition(employeePosition)

	utils.SuccessResponse(c, http.StatusOK, "Success Create Employee Position", res)
}

func UpdateEmployeePosition(c *gin.Context) {
	id := c.Param("id")
	var req dto.ReqEmployeePosition

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var employeePosition models.EmployeePosition
	if err := config.DB.First(&employeePosition, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Employee Position not found")
		return
	}

	tx := config.DB.Begin()

	parsedTime, err := time.Parse("2006-01-02", req.AssignedAt)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid assigned_at format, use YYYY-MM-DD")
		return
	}

	employeePosition.PositionID = &req.PositionID
	employeePosition.Description = &req.Description
	employeePosition.AssignedAt = &parsedTime

	if err := tx.Save(&employeePosition).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lastHistory, err := getLastHistoryPosition(req.EmployeeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if lastHistory != nil {
		lastHistory.PositionID = &req.PositionID
		lastHistory.Description = &req.Description
		lastHistory.StartDate = &parsedTime

		if err := config.DB.Save(lastHistory).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		newHistory := models.PositionHistory{
			EmployeeID:  req.EmployeeID,
			PositionID:  &req.PositionID,
			Description: &req.Description,
			StartDate:   &parsedTime,
		}

		if err := config.DB.Create(&newHistory).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	tx.Commit()

	utils.SuccessResponse(c, http.StatusOK, "Employee position updated successfully", nil)
}

func EndEmployeePosition(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		EmployeeID uint64 `json:"employee_id" form:"employee_id" binding:"required"`
		EndDate    string `json:"end_date" form:"end_date" binding:"required,datetime=2006-01-02"`
	}

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var employeePosition models.EmployeePosition
	if err := config.DB.First(&employeePosition, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Employee Position not found")
		return
	}

	parsedTime, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid assigned_at format, use YYYY-MM-DD")
		return
	}

	tx := config.DB.Begin()

	employeePosition.PositionID = nil
	employeePosition.Description = nil
	employeePosition.AssignedAt = nil

	if err := tx.Save(&employeePosition).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lastHistory, err := getLastHistoryPosition(req.EmployeeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Jika tidak ada history → return error + rollback
	if lastHistory == nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusNotFound, "History posisi tidak ditemukan")
		return
	}

	// Jika ada history → update
	lastHistory.EndDate = &parsedTime

	if err := tx.Save(lastHistory).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()

	utils.SuccessResponse(c, http.StatusOK, "Employee position updated successfully", nil)
}
