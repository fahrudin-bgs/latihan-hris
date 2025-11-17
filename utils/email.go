package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendEmail(to []string, subject string, templateFile string, data interface{}) error {
	email := os.Getenv("EMAIL_ADDRESS")
	appPassword := os.Getenv("EMAIL_APP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if email == "" || appPassword == "" {
		return fmt.Errorf("email atau app password tidak ditemukan di environment variable")
	}

	auth := smtp.PlainAuth(
		"",
		email,
		appPassword,
		smtpHost,
	)

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("gagal parse template: %v", err)
	}

	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: %s\r\n", subject)))
	body.Write([]byte("MIME-Version: 1.0\r\n"))
	body.Write([]byte("Content-Type: text/html; charset=\"UTF-8\"\r\n"))
	body.Write([]byte("\r\n"))

	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("gagal eksekusi template: %v", err)
	}

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		email,
		to,
		body.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("gagal kirim email: %v", err)
	}

	return nil
}
