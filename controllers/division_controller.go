package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDivision(c *gin.Context) {
	var req dto.ReqDivision

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.FindOrError(config.DB, &models.Department{}, req.DepartmentID, c, "Department ID Not Found"); err != nil {
		return
	}

	division := dto.ToModelDivision(req)

	if err := config.DB.Create(&division).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Department").First(&division, division.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResDivision(division)

	utils.SuccessResponse(c, http.StatusCreated, "Division created successfully", res)
}

func GetDivisions(c *gin.Context) {
	var divisions []models.Division

	if err := config.DB.Preload("Department").Find(&divisions).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResDivision
	for _, u := range divisions {
		res = append(res, dto.ToResDivision(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Get All Divisions", res)
}

func DeleteDivision(c *gin.Context) {
	var division models.Division
	id := c.Param("id")

	if err := config.DB.First(&division, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&division).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Division Deleted Successfully", nil)
}
