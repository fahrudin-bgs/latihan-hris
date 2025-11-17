package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var UploadPath string
var MaxUploadMB int64

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error .env file not found")
	}

	// Ambil variabel dari .env
	UploadPath = os.Getenv("UPLOAD_PATH")

	maxUpload := os.Getenv("MAX_UPLOAD_MB")
	if maxUpload == "" {
		MaxUploadMB = 8 // default 8 MB
	} else {
		size, err := strconv.ParseInt(maxUpload, 10, 64)
		if err != nil {
			log.Fatal("Invalid MAX_UPLOAD_MB value in .env")
		}
		MaxUploadMB = size
	}
}
