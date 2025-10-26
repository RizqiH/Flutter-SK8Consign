package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"

	"github.com/google/uuid"
)

func AddToCart(userID, productID string, quantity int) (*models.Cart, error) {
	var product models.Product
	if err := database.DB.Where("id = ? AND is_active = ?", productID, true).First(&product).Error; err != nil {
		return nil, errors.New("product not found or inactive")
	}

	if product.Status != "available" {
		return nil, errors.New("product not available")
	}

	var existingCart models.Cart
	err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingCart).Error

	if err == nil {
		existingCart.Quantity += quantity
		if err := database.DB.Save(&existingCart).Error; err != nil {
			return nil, err
		}
		database.DB.Preload("Product.User").First(&existingCart, existingCart.ID)
		return &existingCart, nil
	}

	cart := &models.Cart{
		ID:        uuid.New().String(),
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := database.DB.Create(cart).Error; err != nil {
		return nil, err
	}

	database.DB.Preload("Product.User").First(cart, cart.ID)
	return cart, nil
}

func GetUserCart(userID string) ([]models.Cart, error) {
	var carts []models.Cart
	err := database.DB.Where("user_id = ?", userID).
		Preload("Product.User").
		Order("created_at DESC").
		Find(&carts).Error

	return carts, err
}

func UpdateCartQuantity(cartID, userID string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	result := database.DB.Model(&models.Cart{}).
		Where("id = ? AND user_id = ?", cartID, userID).
		Update("quantity", quantity)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("cart item not found")
	}

	return nil
}

func RemoveFromCart(cartID, userID string) error {
	result := database.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&models.Cart{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("cart item not found")
	}

	return nil
}

func ClearCart(userID string) error {
	return database.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}



