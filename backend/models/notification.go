package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        string         `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    string         `gorm:"type:char(36);not null;index" json:"user_id"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	Type      string         `gorm:"type:varchar(50);index" json:"type"`
	IsRead    bool           `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}

type NotificationResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Notification) ToResponse() NotificationResponse {
	return NotificationResponse{
		ID:        n.ID,
		Title:     n.Title,
		Message:   n.Message,
		Type:      n.Type,
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	}
}



