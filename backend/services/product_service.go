package services

import (
	"errors"
	"sk8consign-backend/database"
	"sk8consign-backend/models"
	"strings"

	"github.com/google/uuid"
)

// SearchProducts - search products dengan filters
func SearchProducts(query string, category string, minPrice float64, maxPrice float64, status string, limit int, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	// Build query
	db := database.DB.Model(&models.Product{}).Preload("User")

	// Filter by search query (nama atau deskripsi)
	if query != "" {
		searchTerm := "%" + strings.ToLower(query) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", searchTerm, searchTerm)
	}

	// Filter by category
	if category != "" && category != "all" {
		db = db.Where("category = ?", category)
	}

	// Filter by price range
	if minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}

	// Filter by status
	if status != "" && status != "all" {
		db = db.Where("status = ?", status)
	}

	// Only active products
	db = db.Where("is_active = ?", true)

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get products dengan pagination
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// Order by created_at desc
	db = db.Order("created_at DESC")

	if err := db.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductByID - get product by ID
func GetProductByID(productID string) (*models.Product, error) {
	var product models.Product

	if err := database.DB.Preload("User").Where("id = ? AND is_active = ?", productID, true).First(&product).Error; err != nil {
		return nil, errors.New("product not found")
	}

	// Increment view count
	database.DB.Model(&product).Update("view_count", product.ViewCount+1)

	return &product, nil
}

// GetUserProducts - get products by user ID
func GetUserProducts(userID string, status string, limit int, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	db := database.DB.Model(&models.Product{}).Where("user_id = ?", userID)

	if status != "" && status != "all" {
		db = db.Where("status = ?", status)
	}

	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get products
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("created_at DESC")

	if err := db.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// CreateProduct - create new product
func CreateProduct(userID string, name string, description string, price float64, category string, condition string, imageURL string) (*models.Product, error) {
	product := models.Product{
		ID:          uuid.New().String(),
		UserID:      userID,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Condition:   condition,
		Status:      "available",
		ImageURL:    imageURL,
		IsActive:    true,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	// Load user relation
	database.DB.Preload("User").First(&product, "id = ?", product.ID)

	return &product, nil
}

// UpdateProduct - update product
func UpdateProduct(productID string, userID string, name string, description string, price float64, category string, condition string, status string, imageURL string) (*models.Product, error) {
	var product models.Product

	// Check if product exists dan milik user
	if err := database.DB.Where("id = ? AND user_id = ?", productID, userID).First(&product).Error; err != nil {
		return nil, errors.New("product not found or unauthorized")
	}

	// Update fields
	updates := map[string]interface{}{
		"name":        name,
		"description": description,
		"price":       price,
		"category":    category,
		"condition":   condition,
		"status":      status,
	}

	if imageURL != "" {
		updates["image_url"] = imageURL
	}

	if err := database.DB.Model(&product).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Reload product with user
	database.DB.Preload("User").First(&product, "id = ?", product.ID)

	return &product, nil
}

// DeleteProduct - soft delete product
func DeleteProduct(productID string, userID string) error {
	var product models.Product

	// Check if product exists dan milik user
	if err := database.DB.Where("id = ? AND user_id = ?", productID, userID).First(&product).Error; err != nil {
		return errors.New("product not found or unauthorized")
	}

	// Soft delete
	if err := database.DB.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

// GetCategories - get available categories
func GetCategories() []string {
	return []string{
		"gaming",
		"laptop",
		"phone",
		"audio",
		"camera",
		"watch",
		"tablet",
		"accessories",
	}
}
