package models

import (
	"time"

	"gorm.io/gorm"
)

// Product model - tabel products
type Product struct {
	ID          string         `gorm:"type:char(36);primaryKey" json:"id"`
	UserID      string         `gorm:"type:char(36);not null;index" json:"user_id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(12,2);not null" json:"price"`
	Category    string         `gorm:"type:varchar(50);index" json:"category"`
	Condition   string         `gorm:"type:varchar(20)" json:"condition"` // new, like_new, good, fair
	Status      string         `gorm:"type:varchar(20);default:'available';index" json:"status"` // available, sold, reserved
	ImageURL    string         `gorm:"type:varchar(500)" json:"image_url"`
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relation
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName override nama tabel
func (Product) TableName() string {
	return "products"
}

// ProductResponse - response dengan seller info
type ProductResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Condition   string    `json:"condition"`
	Status      string    `json:"status"`
	ImageURL    string    `json:"image_url"`
	ViewCount   int       `json:"view_count"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	// Seller info
	SellerName     string `json:"seller_name,omitempty"`
	SellerUsername string `json:"seller_username,omitempty"`
}

// ToResponse convert Product ke ProductResponse
func (p *Product) ToResponse() ProductResponse {
	return ProductResponse{
		ID:             p.ID,
		UserID:         p.UserID,
		Name:           p.Name,
		Description:    p.Description,
		Price:          p.Price,
		Category:       p.Category,
		Condition:      p.Condition,
		Status:         p.Status,
		ImageURL:       p.ImageURL,
		ViewCount:      p.ViewCount,
		IsActive:       p.IsActive,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		SellerName:     p.User.FullName,
		SellerUsername: p.User.Username,
	}
}
