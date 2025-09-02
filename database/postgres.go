package database

import (
	"fmt"
	"log"

	"gin-auth-project/config"
	"gin-auth-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitPostgres() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// 自动迁移数据库表
	err = DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")

	// 创建默认管理员用户
	createDefaultAdmin()
}

func createDefaultAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count)

	if count == 0 {
		adminUser := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
			Role:     models.RoleAdmin,
			IsActive: true,
		}

		if err := DB.Create(&adminUser).Error; err != nil {
			log.Printf("Failed to create default admin user: %v", err)
		} else {
			log.Println("Default admin user created successfully")
		}
	}
}
