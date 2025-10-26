package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            string         `gorm:"type:char(36);primaryKey" json:"id"`
	UserID        string         `gorm:"type:char(36);not null;index" json:"user_id"`
	TotalAmount   float64        `gorm:"type:decimal(12,2);not null" json:"total_amount"`
	Status        string         `gorm:"type:varchar(20);default:'pending';index" json:"status"`
	PaymentMethod string         `gorm:"type:varchar(50)" json:"payment_method"`
	PaymentStatus string         `gorm:"type:varchar(20);default:'pending'" json:"payment_status"`
	ShippingAddr  string         `gorm:"type:text" json:"shipping_address"`
	Notes         string         `gorm:"type:text" json:"notes"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID        string         `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID   string         `gorm:"type:char(36);not null;index" json:"order_id"`
	ProductID string         `gorm:"type:char(36);not null;index" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"type:decimal(12,2);not null" json:"price"`
	Subtotal  float64        `gorm:"type:decimal(12,2);not null" json:"subtotal"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderResponse struct {
	ID            string              `json:"id"`
	UserID        string              `json:"user_id"`
	TotalAmount   float64             `json:"total_amount"`
	Status        string              `json:"status"`
	PaymentMethod string              `json:"payment_method"`
	PaymentStatus string              `json:"payment_status"`
	ShippingAddr  string              `json:"shipping_address"`
	Notes         string              `json:"notes"`
	OrderItems    []OrderItemResponse `json:"order_items"`
	CreatedAt     time.Time           `json:"created_at"`
}

type OrderItemResponse struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Price     float64         `json:"price"`
	Subtotal  float64         `json:"subtotal"`
	Product   ProductResponse `json:"product"`
}

func (o *Order) ToResponse() OrderResponse {
	items := make([]OrderItemResponse, len(o.OrderItems))
	for i, item := range o.OrderItems {
		items[i] = OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
			Product:   item.Product.ToResponse(),
		}
	}

	return OrderResponse{
		ID:            o.ID,
		UserID:        o.UserID,
		TotalAmount:   o.TotalAmount,
		Status:        o.Status,
		PaymentMethod: o.PaymentMethod,
		PaymentStatus: o.PaymentStatus,
		ShippingAddr:  o.ShippingAddr,
		Notes:         o.Notes,
		OrderItems:    items,
		CreatedAt:     o.CreatedAt,
	}
}



