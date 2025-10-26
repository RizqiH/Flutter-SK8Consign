package handlers

import (
	"encoding/json"
	"net/http"
	"sk8consign-backend/services"
	"strconv"
)

// SearchProductsRequest - request structure
type SearchProductsRequest struct {
	Query    string  `json:"query"`
	Category string  `json:"category"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Status   string  `json:"status"`
	Page     int     `json:"page"`
	Limit    int     `json:"limit"`
}

// CreateProductRequest - request structure
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Condition   string  `json:"condition"`
	ImageURL    string  `json:"image_url"`
}

// UpdateProductRequest - request structure
type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Condition   string  `json:"condition"`
	Status      string  `json:"status"`
	ImageURL    string  `json:"image_url"`
}

// SearchProducts handler
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// Parse request
	var req SearchProductsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	// Default values
	if req.Limit == 0 {
		req.Limit = 20
	}
	if req.Page == 0 {
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit

	// Search products
	products, total, err := services.SearchProducts(
		req.Query,
		req.Category,
		req.MinPrice,
		req.MaxPrice,
		req.Status,
		req.Limit,
		offset,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to search products",
			"error":   err.Error(),
		})
		return
	}

	// Convert to response
	var productResponses []interface{}
	for _, product := range products {
		productResponses = append(productResponses, product.ToResponse())
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Products retrieved successfully",
		"data": map[string]interface{}{
			"products": productResponses,
			"total":    total,
			"page":     req.Page,
			"limit":    req.Limit,
		},
	})
}

// GetProductDetail handler
func GetProductDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// Get product ID from query
	productID := r.URL.Query().Get("id")
	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Product ID is required",
		})
		return
	}

	// Get product
	product, err := services.GetProductByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product.ToResponse(),
	})
}

// GetUserProducts handler
func GetUserProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// Get user ID from query
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "User ID is required",
		})
		return
	}

	// Get filters
	status := r.URL.Query().Get("status")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit == 0 {
		limit = 20
	}
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * limit

	// Get products
	products, total, err := services.GetUserProducts(userID, status, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to get products",
			"error":   err.Error(),
		})
		return
	}

	// Convert to response
	var productResponses []interface{}
	for _, product := range products {
		productResponses = append(productResponses, product.ToResponse())
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Products retrieved successfully",
		"data": map[string]interface{}{
			"products": productResponses,
			"total":    total,
			"page":     page,
			"limit":    limit,
		},
	})
}

// CreateProduct handler
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// Get user ID from context (will be set by auth middleware)
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	// Parse request
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	// Validate
	if req.Name == "" || req.Price <= 0 || req.Category == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Name, price, and category are required",
		})
		return
	}

	// Create product
	product, err := services.CreateProduct(
		userID,
		req.Name,
		req.Description,
		req.Price,
		req.Category,
		req.Condition,
		req.ImageURL,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to create product",
			"error":   err.Error(),
		})
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product created successfully",
		"data":    product.ToResponse(),
	})
}

// GetCategories handler
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	categories := services.GetCategories()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Categories retrieved successfully",
		"data":    categories,
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	productID := r.URL.Query().Get("id")
	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Product ID is required",
		})
		return
	}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.Price <= 0 || req.Category == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Name, price, and category are required",
		})
		return
	}

	product, err := services.UpdateProduct(
		productID,
		userID,
		req.Name,
		req.Description,
		req.Price,
		req.Category,
		req.Condition,
		req.Status,
		req.ImageURL,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product updated successfully",
		"data":    product.ToResponse(),
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	productID := r.URL.Query().Get("id")
	if productID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Product ID is required",
		})
		return
	}

	err := services.DeleteProduct(productID, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Product deleted successfully",
	})
}
