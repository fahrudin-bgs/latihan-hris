package controllers

import (
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePosition(c *gin.Context) {
	var req dto.ReqPosition

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	position := dto.ToModelPosition(req)

	if err := config.DB.Create(&position).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResPosition(position)

	utils.SuccessResponse(c, http.StatusOK, "Success Create New Position", res)
}

func GetPositions(c *gin.Context) {
	var positions []models.Position

	if err := config.DB.Find(&positions).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResPosition
	for _, u := range positions {
		res = append(res, dto.ToResPosition(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Get All Position", res)
}

func GetPositonById(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	if err := config.DB.First(&position, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResPosition(position)

	utils.SuccessResponse(c, http.StatusOK, "Position Found", res)
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	if err := config.DB.First(&position, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&position).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Position Deleted Successfully", nil)
}

func UpdatePosition(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	if err := config.DB.First(&position, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	var req dto.ReqPosition
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dto.ToUpdatePosition(&position, req)

	if err := config.DB.Save(&position).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.First(&position, position.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResPosition(position)
	utils.SuccessResponse(c, http.StatusOK, "Success Update User", res)
}
