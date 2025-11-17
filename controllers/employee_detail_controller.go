package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateEmployeeDetail(c *gin.Context) {
	id := c.Param("id")
	var employeeDetail models.EmployeeDetail

	if err := config.DB.First(&employeeDetail, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Employee detail not found")
		return
	}

	var req dto.ReqEmployeeDetail
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dto.ToUpdateEmployeeDetail(&employeeDetail, req)

	if err := config.DB.Save(&employeeDetail).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.First(&employeeDetail, employeeDetail.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResEmployeeDetail(employeeDetail)
	utils.SuccessResponse(c, http.StatusOK, "Success Update Employee", res)
}
