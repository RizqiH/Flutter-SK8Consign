package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"

	"github.com/google/uuid"
)

func CreateNotification(userID, title, message, notifType string) (*models.Notification, error) {
	notification := &models.Notification{
		ID:      uuid.New().String(),
		UserID:  userID,
		Title:   title,
		Message: message,
		Type:    notifType,
		IsRead:  false,
	}

	if err := database.DB.Create(notification).Error; err != nil {
		return nil, err
	}

	return notification, nil
}

func GetUserNotifications(userID string, limit, offset int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	database.DB.Model(&models.Notification{}).Where("user_id = ?", userID).Count(&total)

	err := database.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error

	return notifications, total, err
}

func MarkAsRead(notificationID, userID string) error {
	result := database.DB.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("notification not found")
	}

	return nil
}

func MarkAllAsRead(userID string) error {
	return database.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true).Error
}

func GetUnreadCount(userID string) (int64, error) {
	var count int64
	err := database.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error

	return count, err
}



