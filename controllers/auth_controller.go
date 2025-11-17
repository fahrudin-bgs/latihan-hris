package controllers

import (
	"fmt"
	"latihan-hris/config"
	"latihan-hris/dto"
	"latihan-hris/models"
	"latihan-hris/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	var user models.User
	if err := config.DB.First(&user, "email = ?", email).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	//  create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(), // 30 menit
	})
	accessTokenString, _ := accessToken.SignedString([]byte(os.Getenv("SECRET")))

	//  create refresh token
	exp := time.Now().Add(time.Hour * 24 * 7)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": exp.Unix(), // 7 hari
	})
	refreshTokenString, _ := refreshToken.SignedString([]byte(os.Getenv("SECRET")))

	// save refresh token in database
	refreshData := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshTokenString,
		IPAddress: &ipAddress,
		UserAgent: &userAgent,
		ExpiresAt: exp,
	}
	config.DB.Create(&refreshData)

	c.JSON(200, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}

func Register(c *gin.Context) {
	var req struct {
		Username   string `json:"username" form:"username" binding:"required"`
		Email      string `json:"email" form:"email" binding:"required,email"`
		Password   string `json:"password" form:"password" binding:"required"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}

	if err := c.ShouldBind(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var existing models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Email already registered")
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	defaultRole := uint64(2)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   &defaultRole,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Kirim email jika IsVerified = true
	if req.IsVerified {
		userId := user.ID
		token := utils.CreateAccessToken(uint(userId))
		println(token)

		appURL := os.Getenv("APP_URL")
		urlVerified := fmt.Sprintf("%s/verifikasi?token=%s", appURL, token)

		data := struct {
			Name string
			Link string
		}{
			Name: req.Username,
			Link: urlVerified,
		}

		err := utils.SendEmail(
			[]string{req.Email},
			"Verifikasi Akun Anda",
			"templates/verify_email.html",
			data,
		)
		if err != nil {
			fmt.Println("Gagal mengirim email:", err)
			// tidak return, agar registrasi tetap sukses walau email gagal
		}
	}

	res := dto.ToResUser(user)

	utils.SuccessResponse(c, http.StatusCreated, "Registration successful", res)
}

func Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token"`
	}

	if err := c.ShouldBind(&req); err != nil || req.RefreshToken == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Refresh token is required")
		return
	}

	var token models.RefreshToken
	if err := config.DB.Where("token = ?", req.RefreshToken).First(&token).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Refresh token not found or already deleted")
		return
	}

	if err := config.DB.Delete(&token).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete refresh token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func CurrentUser(c *gin.Context) {
	id, exists := c.Get("user_id")

	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User Id not found")
	}

	var user models.User

	if err := config.DB.Preload("Role").First(&user, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResUser(user)

	utils.SuccessResponse(c, http.StatusOK, "User Found", res)
}

func VerifiedUser(c *gin.Context) {
	accessToken := c.Query("token")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid claims"})
		c.Abort()
		return
	}

	userID := fmt.Sprintf("%v", claims["sub"])

	var user models.User
	if err := config.DB.Preload("Role").First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	res := dto.ToResUserDetail(user)

	utils.SuccessResponse(c, http.StatusOK, "User Verified", res)
}
