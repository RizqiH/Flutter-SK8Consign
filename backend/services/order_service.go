package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrder(userID, paymentMethod, shippingAddr, notes string) (*models.Order, error) {
	var carts []models.Cart
	err := database.DB.Where("user_id = ?", userID).Preload("Product").Find(&carts).Error
	if err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	var totalAmount float64
	var orderItems []models.OrderItem

	for _, cart := range carts {
		if !cart.Product.IsActive || cart.Product.Status != "available" {
			return nil, errors.New("some products are not available")
		}

		subtotal := cart.Product.Price * float64(cart.Quantity)
		totalAmount += subtotal

		orderItem := models.OrderItem{
			ID:        uuid.New().String(),
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
			Subtotal:  subtotal,
		}
		orderItems = append(orderItems, orderItem)
	}

	order := &models.Order{
		ID:            uuid.New().String(),
		UserID:        userID,
		TotalAmount:   totalAmount,
		Status:        "pending",
		PaymentMethod: paymentMethod,
		PaymentStatus: "pending",
		ShippingAddr:  shippingAddr,
		Notes:         notes,
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for i := range orderItems {
			orderItems[i].OrderID = order.ID
			if err := tx.Create(&orderItems[i]).Error; err != nil {
				return err
			}

			if err := tx.Model(&models.Product{}).
				Where("id = ?", orderItems[i].ProductID).
				Update("status", "reserved").Error; err != nil {
				return err
			}
		}

		if err := tx.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	database.DB.Preload("OrderItems.Product.User").First(order, order.ID)
	return order, nil
}

func GetUserOrders(userID string, status string, limit, offset int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := database.DB.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Model(&models.Order{}).Count(&total)

	err := query.Preload("OrderItems.Product.User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&orders).Error

	return orders, total, err
}

func GetOrderByID(orderID, userID string) (*models.Order, error) {
	var order models.Order
	err := database.DB.Where("id = ? AND user_id = ?", orderID, userID).
		Preload("OrderItems.Product.User").
		First(&order).Error

	if err != nil {
		return nil, errors.New("order not found")
	}

	return &order, nil
}

func UpdateOrderStatus(orderID, userID, status string) error {
	validStatuses := []string{"pending", "confirmed", "shipped", "delivered", "cancelled"}
	valid := false
	for _, s := range validStatuses {
		if s == status {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New("invalid status")
	}

	result := database.DB.Model(&models.Order{}).
		Where("id = ? AND user_id = ?", orderID, userID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func UpdatePaymentStatus(orderID, userID, paymentStatus string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var order models.Order
		if err := tx.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
			return errors.New("order not found")
		}

		if err := tx.Model(&order).Update("payment_status", paymentStatus).Error; err != nil {
			return err
		}

		if paymentStatus == "paid" {
			if err := tx.Model(&order).Update("status", "confirmed").Error; err != nil {
				return err
			}

			var orderItems []models.OrderItem
			if err := tx.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
				return err
			}

			for _, item := range orderItems {
				if err := tx.Model(&models.Product{}).
					Where("id = ?", item.ProductID).
					Update("status", "sold").Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}



