package controllers

import (
	"latihan-hris/config"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePositionHistory(c *gin.Context) {
	id := c.Param("id")
	var history models.PositionHistory

	if err := config.DB.First(&history, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
}
