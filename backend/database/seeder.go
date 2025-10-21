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

// ClearData - clear all data (untuk testing/development)
func ClearData() {
	log.Println("⚠️  Clearing all data...")
	
	// Delete all users (hard delete)
	DB.Unscoped().Where("1 = 1").Delete(&models.User{})
	
	log.Println("✅ All data cleared")
}

// ResetData - clear dan seed ulang
func ResetData() {
	ClearData()
	SeedData()
}
