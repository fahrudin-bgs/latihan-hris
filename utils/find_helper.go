package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindOrError(db *gorm.DB, model interface{}, id interface{}, c *gin.Context, notFoundMsg string) error {
	if err := db.First(model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(c, http.StatusBadRequest, notFoundMsg)
			return err
		}

		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return nil
}
