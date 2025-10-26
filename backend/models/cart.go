package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        string         `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    string         `gorm:"type:char(36);not null;index" json:"user_id"`
	ProductID string         `gorm:"type:char(36);not null;index" json:"product_id"`
	Quantity  int            `gorm:"default:1" json:"quantity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartResponse struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	ProductID string          `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Product   ProductResponse `json:"product"`
	CreatedAt time.Time       `json:"created_at"`
}

func (c *Cart) ToResponse() CartResponse {
	return CartResponse{
		ID:        c.ID,
		UserID:    c.UserID,
		ProductID: c.ProductID,
		Quantity:  c.Quantity,
		Product:   c.Product.ToResponse(),
		CreatedAt: c.CreatedAt,
	}
}



