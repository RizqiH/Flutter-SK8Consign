package database

import (
	"log"
	"sk8consign-backend/models"
	"sk8consign-backend/utils"

	"github.com/google/uuid"
)

// SeedData - seed initial data untuk development
func SeedData() {
	if shouldSkipSeeding() {
		log.Println("   ⏭️  Seeding skipped - data already exists")
		return
	}

	log.Println("🌱 Seeding initial data...")
	seedUsers()
	seedProducts()
	seedNotifications()
	log.Println("✅ Seeding completed")
}

// shouldSkipSeeding - cek apakah perlu seeding
func shouldSkipSeeding() bool {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	return count > 0
}

// seedUsers - seed default users
func seedUsers() {
	log.Println("   📝 Seeding users...")

	defaultUsers := []struct {
		Username string
		Email    string
		Password string
		FullName string
		Phone    string
		Role     string
	}{
		{
			Username: "admin",
			Email:    "admin@sk8consign.com",
			Password: "admin123",
			FullName: "Admin SK8 Consign",
			Phone:    "081234567890",
			Role:     "admin",
		},
		{
			Username: "user",
			Email:    "user@sk8consign.com",
			Password: "user123",
			FullName: "Regular User",
			Phone:    "081234567891",
			Role:     "user",
		},
	}

	for _, userData := range defaultUsers {
		// Hash password
		hashedPassword, err := utils.HashPassword(userData.Password)
		if err != nil {
			log.Printf("   ❌ Failed to hash password for %s: %v", userData.Username, err)
			continue
		}

		// Create user
		user := models.User{
			ID:       uuid.New().String(),
			Username: userData.Username,
			Email:    userData.Email,
			Password: hashedPassword,
			FullName: userData.FullName,
			Phone:    userData.Phone,
			Role:     userData.Role,
			IsActive: true,
		}

		if err := DB.Create(&user).Error; err != nil {
			log.Printf("   ❌ Failed to seed user %s: %v", userData.Username, err)
		} else {
			log.Printf("   ✅ User created: %s (password: %s, role: %s)", userData.Username, userData.Password, userData.Role)
		}
	}
}

// seedProducts - seed sample products
func seedProducts() {
	log.Println("   📦 Seeding products...")

	// Get seeded users
	var adminUser, regularUser models.User
	DB.Where("username = ?", "admin").First(&adminUser)
	DB.Where("username = ?", "user").First(&regularUser)

	sampleProducts := []models.Product{
		{
			ID:          uuid.New().String(),
			UserID:      adminUser.ID,
			Name:        "PlayStation 5 Console",
			Description: "Brand new PS5 console with controller, barely used. Perfect condition.",
			Price:       7500000,
			Category:    "gaming",
			Condition:   "like_new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1606813907291-d86efa9b94db?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      adminUser.ID,
			Name:        "MacBook Pro M3 14-inch",
			Description: "MacBook Pro 14-inch with M3 chip, 16GB RAM, 512GB SSD. Like new condition with box.",
			Price:       25000000,
			Category:    "laptop",
			Condition:   "like_new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      regularUser.ID,
			Name:        "iPhone 15 Pro Max 256GB",
			Description: "iPhone 15 Pro Max, Titanium Blue, 256GB. Mint condition, complete with box and accessories.",
			Price:       18500000,
			Category:    "phone",
			Condition:   "new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1695048133142-1a20484d2569?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      regularUser.ID,
			Name:        "Sony WH-1000XM5 Headphones",
			Description: "Premium noise-cancelling headphones. Excellent sound quality and battery life.",
			Price:       4500000,
			Category:    "audio",
			Condition:   "good",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1546435770-a3e426bf472b?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      adminUser.ID,
			Name:        "Canon EOS R6 Camera Body",
			Description: "Professional mirrorless camera, perfect for photography and videography. Low shutter count.",
			Price:       32000000,
			Category:    "camera",
			Condition:   "like_new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1516035069371-29a1b244cc32?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      regularUser.ID,
			Name:        "Apple Watch Series 9",
			Description: "Apple Watch Series 9, 45mm, GPS + Cellular. Midnight Aluminum Case with Sport Band.",
			Price:       6500000,
			Category:    "watch",
			Condition:   "new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1434493789847-2f02dc6ca35d?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      adminUser.ID,
			Name:        "iPad Pro 12.9 M2",
			Description: "iPad Pro 12.9-inch with M2 chip, 256GB, WiFi + Cellular. Includes Magic Keyboard.",
			Price:       16500000,
			Category:    "tablet",
			Condition:   "like_new",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=500",
			IsActive:    true,
		},
		{
			ID:          uuid.New().String(),
			UserID:      regularUser.ID,
			Name:        "Razer BlackWidow V4 Pro",
			Description: "Mechanical gaming keyboard with RGB lighting, green switches. Perfect for gaming.",
			Price:       2500000,
			Category:    "accessories",
			Condition:   "good",
			Status:      "available",
			ImageURL:    "https://images.unsplash.com/photo-1587829741301-dc798b83add3?w=500",
			IsActive:    true,
		},
	}

	for _, product := range sampleProducts {
		if err := DB.Create(&product).Error; err != nil {
			log.Printf("   ❌ Failed to seed product %s: %v", product.Name, err)
		} else {
			log.Printf("   ✅ Product created: %s (Rp %.0f)", product.Name, product.Price)
		}
	}
}

// ClearData - clear all data (untuk testing/development)
func seedNotifications() {
	log.Println("   🔔 Seeding notifications...")

	var adminUser, regularUser models.User
	DB.Where("username = ?", "admin").First(&adminUser)
	DB.Where("username = ?", "user").First(&regularUser)

	sampleNotifications := []models.Notification{
		{
			ID:      uuid.New().String(),
			UserID:  regularUser.ID,
			Title:   "Welcome to SK8 Consign!",
			Message: "Start buying and selling gaming gear and electronics today!",
			Type:    "promo",
			IsRead:  false,
		},
		{
			ID:      uuid.New().String(),
			UserID:  regularUser.ID,
			Title:   "New Product Alert",
			Message: "Check out the latest PlayStation 5 listings!",
			Type:    "product",
			IsRead:  false,
		},
		{
			ID:      uuid.New().String(),
			UserID:  adminUser.ID,
			Title:   "System Update",
			Message: "SK8 Consign has been updated with new features!",
			Type:    "promo",
			IsRead:  true,
		},
	}

	for _, notification := range sampleNotifications {
		if err := DB.Create(&notification).Error; err != nil {
			log.Printf("   ❌ Failed to seed notification: %v", err)
		}
	}

	log.Println("   ✅ Notifications seeded")
}

func ClearData() {
	log.Println("⚠️  Clearing all data...")

	DB.Unscoped().Where("1 = 1").Delete(&models.Notification{})
	DB.Unscoped().Where("1 = 1").Delete(&models.OrderItem{})
	DB.Unscoped().Where("1 = 1").Delete(&models.Order{})
	DB.Unscoped().Where("1 = 1").Delete(&models.Cart{})
	DB.Unscoped().Where("1 = 1").Delete(&models.Product{})
	DB.Unscoped().Where("1 = 1").Delete(&models.User{})

	log.Println("✅ All data cleared")
}

// ResetData - clear dan seed ulang
func ResetData() {
	ClearData()
	SeedData()
}
