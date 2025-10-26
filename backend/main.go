package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"sk8consign-backend/config"
	"sk8consign-backend/database"
	"sk8consign-backend/handlers"
	"sk8consign-backend/middleware"

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

	authHandler := handlers.NewAuthHandler()

	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/register", authHandler.Register)

	mux.HandleFunc("/api/profile", middleware.AuthMiddleware(handlers.GetProfile))
	mux.HandleFunc("/api/profile/update", middleware.AuthMiddleware(handlers.UpdateProfile))
	mux.HandleFunc("/api/profile/change-password", middleware.AuthMiddleware(handlers.ChangePassword))

	mux.HandleFunc("/api/products/search", handlers.SearchProducts)
	mux.HandleFunc("/api/products/detail", handlers.GetProductDetail)
	mux.HandleFunc("/api/products/user", handlers.GetUserProducts)
	mux.HandleFunc("/api/products/create", middleware.RequireAdmin(handlers.CreateProduct))
	mux.HandleFunc("/api/products/update", middleware.RequireAuth(handlers.UpdateProduct))
	mux.HandleFunc("/api/products/delete", middleware.RequireAuth(handlers.DeleteProduct))
	mux.HandleFunc("/api/products/categories", handlers.GetCategories)

	mux.HandleFunc("/api/cart", middleware.AuthMiddleware(handlers.GetCart))
	mux.HandleFunc("/api/cart/add", middleware.AuthMiddleware(handlers.AddToCart))
	mux.HandleFunc("/api/cart/update", middleware.AuthMiddleware(handlers.UpdateCart))
	mux.HandleFunc("/api/cart/remove", middleware.AuthMiddleware(handlers.RemoveFromCart))
	mux.HandleFunc("/api/cart/clear", middleware.AuthMiddleware(handlers.ClearCart))

	mux.HandleFunc("/api/orders", middleware.AuthMiddleware(handlers.GetUserOrders))
	mux.HandleFunc("/api/orders/create", middleware.AuthMiddleware(handlers.CreateOrder))
	mux.HandleFunc("/api/orders/detail", middleware.AuthMiddleware(handlers.GetOrderDetail))
	mux.HandleFunc("/api/orders/update-status", middleware.AuthMiddleware(handlers.UpdateOrderStatus))
	mux.HandleFunc("/api/orders/update-payment", middleware.AuthMiddleware(handlers.UpdatePaymentStatus))

	mux.HandleFunc("/api/notifications", middleware.AuthMiddleware(handlers.GetNotifications))
	mux.HandleFunc("/api/notifications/read", middleware.AuthMiddleware(handlers.MarkNotificationRead))
	mux.HandleFunc("/api/notifications/read-all", middleware.AuthMiddleware(handlers.MarkAllNotificationsRead))
	mux.HandleFunc("/api/notifications/unread-count", middleware.AuthMiddleware(handlers.GetUnreadCount))

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
	log.Println("   [Auth]")
	log.Println("   POST   /api/register")
	log.Println("   POST   /api/login")
	log.Println()
	log.Println("   [Profile]")
	log.Println("   POST   /api/profile")
	log.Println("   PUT    /api/profile/update")
	log.Println("   PUT    /api/profile/change-password")
	log.Println()
	log.Println("   [Products]")
	log.Println("   POST   /api/products/search")
	log.Println("   GET    /api/products/detail")
	log.Println("   GET    /api/products/user")
	log.Println("   POST   /api/products/create")
	log.Println("   PUT    /api/products/update")
	log.Println("   DELETE /api/products/delete")
	log.Println("   GET    /api/products/categories")
	log.Println()
	log.Println("   [Cart]")
	log.Println("   GET    /api/cart")
	log.Println("   POST   /api/cart/add")
	log.Println("   PUT    /api/cart/update")
	log.Println("   DELETE /api/cart/remove")
	log.Println("   DELETE /api/cart/clear")
	log.Println()
	log.Println("   [Orders]")
	log.Println("   GET    /api/orders")
	log.Println("   POST   /api/orders/create")
	log.Println("   GET    /api/orders/detail")
	log.Println("   PUT    /api/orders/update-status")
	log.Println("   PUT    /api/orders/update-payment")
	log.Println()
	log.Println("   [Notifications]")
	log.Println("   GET    /api/notifications")
	log.Println("   PUT    /api/notifications/read")
	log.Println("   PUT    /api/notifications/read-all")
	log.Println("   GET    /api/notifications/unread-count")
	log.Println()
	log.Println("   [System]")
	log.Println("   GET    /api/health")
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
