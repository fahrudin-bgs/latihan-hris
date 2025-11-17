package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var req dto.ReqEmployee

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.FindOrError(config.DB, &models.User{}, req.UserID, c, "User ID Not Found"); err != nil {
		return
	}

	if req.DivisionID != nil {
		if err := utils.FindOrError(config.DB, &models.Division{}, *req.DivisionID, c, "Division ID Not Found"); err != nil {
			return
		}
	}

	data := dto.ToModelEmployee(req)

	// DB Begin
	tx := config.DB.Begin()

	// create employee
	if err := tx.Where("user_id=?", req.UserID).FirstOrCreate(&data).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// create employe_detail
	if err := tx.FirstOrCreate(&models.EmployeeDetail{EmployeeID: data.ID}).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// DB Commit
	if err := tx.Commit().Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("User.Role").Preload("Division").Preload("EmployeeDetail").First(&data, data.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res := dto.ToResEmployee(data)
	utils.SuccessResponse(c, http.StatusCreated, "Employee created successfully", res)
}

func GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	if err := config.DB.Preload("User.Role").Preload("Division").Preload("EmployeeDetail").First(&employee, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResEmployee(employee)
	utils.SuccessResponse(c, http.StatusOK, "Employee Found", res)
}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	if err := config.DB.Preload("User.Role").Preload("Division").Find(&employees).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResEmployeeSimple
	for _, u := range employees {
		res = append(res, dto.ToResEmployeeSimple(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Get All Employees", res)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	// cek data employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Employee not found")
		return
	}

	var req dto.ReqEmployee
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dto.ToUpdateEmployee(&employee, req)

	if err := config.DB.Save(&employee).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("User.Role").Preload("Division").First(&employee, employee.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResEmployee(employee)
	utils.SuccessResponse(c, http.StatusOK, "Success Update Employee", res)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&employee).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employee Deleted Successfully", nil)
}
