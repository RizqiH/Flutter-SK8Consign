package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"
	"sk8consign-backend/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct{}

// Login - authenticate user
func (s *AuthService) Login(username, password string) (*models.User, string, error) {
	var user models.User

	// Cari user by username
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("username atau password salah")
		}
		return nil, "", err
	}

	// Cek apakah user aktif
	if !user.IsActive {
		return nil, "", errors.New("akun tidak aktif")
	}

	// Verify password
	if !utils.CheckPassword(user.Password, password) {
		return nil, "", errors.New("username atau password salah")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, "", errors.New("gagal generate token")
	}

	return &user, token, nil
}

// Register - create new user
func (s *AuthService) Register(username, email, password, fullName, phone string) error {
	// Validasi input
	if username == "" || email == "" || password == "" {
		return errors.New("username, email, dan password harus diisi")
	}

	// Cek username sudah ada
	var existingUser models.User
	if err := database.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errors.New("username sudah digunakan")
	}

	// Cek email sudah ada
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.New("email sudah digunakan")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("gagal hash password")
	}

	// Buat user baru
	user := models.User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
		Phone:    phone,
		Role:     "user", // default role
		IsActive: true,
	}

	// Save ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return errors.New("gagal membuat user")
	}

	return nil
}

// GetUserByID - get user by ID
func (s *AuthService) GetUserByID(userID string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser - update user data
func (s *AuthService) UpdateUser(userID string, updates map[string]interface{}) error {
	// Jangan update password langsung, harus lewat ChangePassword
	delete(updates, "password")
	delete(updates, "id")
	delete(updates, "role") // role tidak bisa diubah sendiri

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return errors.New("gagal update user")
	}

	return nil
}

// ChangePassword - change user password
func (s *AuthService) ChangePassword(userID, oldPassword, newPassword string) error {
	var user models.User

	// Get user
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user tidak ditemukan")
	}

	// Verify old password
	if !utils.CheckPassword(user.Password, oldPassword) {
		return errors.New("password lama salah")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("gagal hash password")
	}

	// Update password
	if err := database.DB.Model(&user).Update("password", hashedPassword).Error; err != nil {
		return errors.New("gagal update password")
	}

	return nil
}
