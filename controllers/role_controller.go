package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Get All User", roles)
}

func CreateRole(c *gin.Context) {
	var req dto.ReqRole

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	data := dto.ToModelRole(req)

	if result := config.DB.Create(&data).Error; result != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Create New Role", data)
}

func DeleteRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	if err := config.DB.First(&role, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&role).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Role Deleted Successfully", nil)
}
