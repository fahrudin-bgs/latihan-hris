package controllers

import (
	"fmt"
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadDocument(c *gin.Context) {
	var req dto.ReqEmployeeDocument
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.FindOrError(config.DB, &models.Employee{}, req.EmployeeID, c, "Employee ID Not Found"); err != nil {
		return
	}

	file, err := c.FormFile("document")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Document file is required")
		return
	}

	if !utils.ValidateFile(file, []string{".pdf"}) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file extension")
		return
	}

	folderPath := filepath.Join(config.UploadPath, "employees/documents")
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create upload directory")
		return
	}

	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d_%d%s", req.EmployeeID, time.Now().Unix(), ext)
	dstPath := filepath.Join(folderPath, newFileName)

	if err := c.SaveUploadedFile(file, dstPath); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to save file")
		return
	}

	doc := models.EmployeeDocument{
		EmployeeID:  req.EmployeeID,
		FileType:    req.FileType,
		Description: req.Description,
		FilePath:    dstPath,
	}

	if err := config.DB.Create(&doc).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Document upload successfully", nil)
}

func DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	var doc models.EmployeeDocument

	if err := config.DB.First(&doc, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	filePath := doc.FilePath

	if err := config.DB.Delete(&doc).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if doc.FilePath != "" {
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete file from disk")
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Document Deleted Successfully", nil)
}

func UpdateDocument(c *gin.Context) {
	id := c.Param("id")

	var req dto.ReqEmployeeDocument
	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var existingDoc models.EmployeeDocument
	if err := config.DB.First(&existingDoc, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Document not found")
		return
	}

	// Update hanya field yang dikirim
	if req.FileType != "" {
		existingDoc.FileType = req.FileType
	}
	if req.Description != nil {
		existingDoc.Description = req.Description
	}
	if req.EmployeeID != 0 {
		existingDoc.EmployeeID = req.EmployeeID
	}

	// Cek apakah ada file baru
	file, err := c.FormFile("document")
	if err == nil {
		if !utils.ValidateFile(file, []string{".pdf"}) {
			utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file extension")
			return
		}

		folderPath := filepath.Join(config.UploadPath, "employees/documents")
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create upload directory")
			return
		}

		ext := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%d_%d%s", existingDoc.EmployeeID, time.Now().Unix(), ext)
		dstPath := filepath.Join(folderPath, newFileName)

		// Hapus file lama (kalau ada)
		if existingDoc.FilePath != "" {
			_ = os.Remove(existingDoc.FilePath)
		}

		// Simpan file baru
		if err := c.SaveUploadedFile(file, dstPath); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to save new file")
			return
		}

		existingDoc.FilePath = dstPath
	}

	// Simpan perubahan
	if err := config.DB.Save(&existingDoc).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Document updated successfully", existingDoc)
}

