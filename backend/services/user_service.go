package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"
	"sk8consign-backend/utils"
)

// GetUserProfile - get user profile by ID
func GetUserProfile(userID string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("id = ? AND is_active = ?", userID, true).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// UpdateUserProfile - update user profile
func UpdateUserProfile(userID string, fullName string, phone string, email string) (*models.User, error) {
	var user models.User

	// Check if user exists
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Check if email already taken by another user
	if email != user.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", email, userID).First(&existingUser).Error; err == nil {
			return nil, errors.New("email already taken")
		}
	}

	// Update fields
	updates := map[string]interface{}{
		"full_name": fullName,
		"phone":     phone,
		"email":     email,
	}

	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Reload user
	database.DB.First(&user, "id = ?", userID)

	return &user, nil
}

// ChangePassword - change user password
func ChangePassword(userID string, oldPassword string, newPassword string) error {
	var user models.User

	// Get user
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	// Verify old password
	if !utils.CheckPassword(user.Password, oldPassword) {
		return errors.New("invalid old password")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Update password
	if err := database.DB.Model(&user).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}

// DeactivateAccount - deactivate user account (soft delete)
func DeactivateAccount(userID string) error {
	var user models.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	// Set is_active to false dan soft delete
	if err := database.DB.Model(&user).Update("is_active", false).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
