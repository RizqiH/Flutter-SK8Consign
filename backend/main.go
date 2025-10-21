package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"sk8consign-backend/config"
	"sk8consign-backend/database"
	"sk8consign-backend/handlers"

	"github.com/rs/cors"
)

func main() {
	// Banner
	printBanner()

	// Load configuration
	config.LoadConfig()

	// Connect to database
	database.Connect()

	// Auto migrate tables
	database.AutoMigrate()

	// Seed data (hanya jalan jika database kosong)
	// Untuk skip seeding, set environment: SKIP_SEED=true
	if os.Getenv("SKIP_SEED") != "true" {
		database.SeedData()
	}

	// Setup routes
	mux := setupRoutes()

	// CORS middleware
	handler := setupCORS(mux)

	// Start server
	startServer(handler)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()

	// API Routes
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/register", authHandler.Register)
	mux.HandleFunc("/api/health", handlers.HealthCheck)

	return mux
}

func setupCORS(mux *http.ServeMux) http.Handler {
	// Get allowed origins from env or use default
	allowedOrigins := []string{"*"}
	if config.AppConfig.Env == "production" {
		// Di production, specify origins
		allowedOrigins = []string{
			"https://sk8consign.com",
			"https://www.sk8consign.com",
		}
	}

	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	}).Handler(mux)
}

func startServer(handler http.Handler) {
	port := config.AppConfig.ServerPort

	log.Printf("🚀 Server running on http://localhost:%s\n", port)
	log.Printf("🌍 Environment: %s\n", config.AppConfig.Env)
	log.Println()
	log.Println("📱 Available Endpoints:")
	log.Println("   POST   /api/register     - Create new user account")
	log.Println("   POST   /api/login        - User authentication")
	log.Println("   GET    /api/health       - Health check")
	log.Println()
	log.Println("💾 Database:")
	log.Printf("   Host: %s:%s\n", config.AppConfig.DBHost, config.AppConfig.DBPort)
	log.Printf("   Name: %s\n", config.AppConfig.DBName)
	log.Println()
	log.Println("⏳ Server ready - waiting for requests...")

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

func printBanner() {
	banner := `
╔════════════════════════════════════════════════╗
║                                                ║
║         SK8 CONSIGN API SERVER v2.0            ║
║      Production-Ready Backend Service          ║
║                                                ║
║  Stack: Go + MySQL + GORM + JWT + Bcrypt       ║
║                                                ║
╚════════════════════════════════════════════════╝
`
	fmt.Println(banner)
}
