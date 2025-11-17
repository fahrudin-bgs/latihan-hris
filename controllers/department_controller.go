package controllers

import (
	"errors"
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDepartments(c *gin.Context) {
	var departments []models.Department

	if err := config.DB.Find(&departments).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResDepartment
	for _, u := range departments {
		res = append(res, dto.ToResDepartment(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Get All Departments", res)
}

func GetDepartmentById(c *gin.Context) {
	var department models.Department
	id := c.Param("id")

	if err := config.DB.First(&department, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	res := dto.ToResDepartment(department)
	utils.SuccessResponse(c, http.StatusOK, "Department Found", res)
}

func CreateDepartment(c *gin.Context) {
	var req dto.ReqDepartment

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := dto.ToModelDepartment(req)

	if err := config.DB.Create(&data).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResDepartment(data)
	utils.SuccessResponse(c, http.StatusOK, "Success Create New Department", res)
}

func DeleteDepartment(c *gin.Context) {
	var department models.Department
	id := c.Param("id")

	if err := config.DB.First(&department, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&department).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Department Deleted Successfully", nil)
}
