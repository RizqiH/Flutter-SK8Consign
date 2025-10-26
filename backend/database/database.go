package database

import (
	"fmt"
	"log"
	"time"

	"sk8consign-backend/config"
	"sk8consign-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect - establish connection to MySQL database
func Connect() {
	cfg := config.AppConfig

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// GORM configuration
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		// Disable foreign key constraint untuk flexibility
		DisableForeignKeyConstraintWhenMigrating: false,
	}

	// Set logger based on environment
	if cfg.Env == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	// Connect to database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("❌ Failed to configure database connection pool:", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)           // Max idle connections
	sqlDB.SetMaxOpenConns(100)          // Max open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Max connection lifetime

	log.Println("✅ Database connected successfully")
}

// AutoMigrate - automatically migrate database schema
func AutoMigrate() {
	log.Println("🔄 Running database migrations...")

	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
		&models.Notification{},
	}

	if err := DB.AutoMigrate(modelsToMigrate...); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Database migration completed")
}

// Close - close database connection
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("⚠️  Failed to get database instance:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("⚠️  Failed to close database connection:", err)
		return
	}

	log.Println("✅ Database connection closed")
}

// Ping - check database connection health
func Ping() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
