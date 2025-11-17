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

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var res []dto.ResUser
	for _, u := range users {
		res = append(res, dto.ToResUser(u))
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Get All User", res)
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.Preload("Role").First(&user, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResUserDetail(user)

	utils.SuccessResponse(c, http.StatusOK, "User Found", res)
}

func CreateUser(c *gin.Context) {
	var (
		req      dto.ReqUser
		role     models.Role
		existing models.User
	)

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.First(&role, req.RoleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusBadRequest, "Role Id Not Found")
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Email already registered")
		return
	}

	data := dto.ToModelUser(req)
	if result := config.DB.Create(&data).Error; result != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error())
		return
	}

	if err := config.DB.Preload("Role").First(&data, data.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.ToResUserDetail(data)
	utils.SuccessResponse(c, http.StatusCreated, "Success Create New User", res)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Cek apakah user dengan ID tersebut ada
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	// Bind form-data
	var req dto.ReqUser
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Cek apakah email sudah digunakan oleh user lain
	var existing models.User
	if err := config.DB.
		Where("email = ? AND id <> ?", req.Email, id).
		First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Email already registered by another user")
		return
	}

	// Update field user
	dto.ToUpdateUser(&user, req)

	// Simpan perubahan
	if err := config.DB.Save(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Preload role
	if err := config.DB.Preload("Role").First(&user, user.ID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Response
	res := dto.ToResUserDetail(user)
	utils.SuccessResponse(c, http.StatusOK, "Success Update User", res)
}


func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User Deleted Successfully", nil)
}
