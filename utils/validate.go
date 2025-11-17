package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ValidateFile(file *multipart.FileHeader, allowedExts []string) bool {
	if file == nil {
		return false
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		return false
	}

	for _, allowed := range allowedExts {
		if ext == strings.ToLower(allowed) {
			return true
		}
	}
	return false
}
