package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func parseDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func UpdatePositionHistory(c *gin.Context) {
	idParam := c.Param("id")

	var req dto.ReqPositionHistory
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	var history models.PositionHistory
	if err := config.DB.First(&history, idParam).Error; err != nil {
		utils.ErrorResponse(c, 404, "History not found")
		return
	}

	// Convert tanggal
	start, _ := parseDate(req.StartDate)
	end, _ := parseDate(req.EndDate)

	// Update field jika tidak nil
	if req.PositionID != nil {
		history.PositionID = req.PositionID
	}
	if req.Description != nil {
		history.Description = req.Description
	}
	if start != nil {
		history.StartDate = start
	}
	if end != nil {
		history.EndDate = end
	}

	// Save (akan memicu BeforeUpdate)
	if err := config.DB.Save(&history).Error; err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, "Updated", history)
}

func GetAllPositionHistories(c *gin.Context) {
	id := c.Param("employee_id")
	var history []models.PositionHistory

	if err := config.DB.Where("employee_id = ?", id).Find(&history).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var res []dto.ResPositionHistory
	for _, u := range history {
		res = append(res, dto.ToResPositionHistory(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Get All Position Histories", res)
}
