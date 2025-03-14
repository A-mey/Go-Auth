package sql

import (
	"fmt"
	"log"

	"github.com/A-mey/Auth_db/config/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeSQL() {
	// dsn := os.Getenv("DB_DSN")
	cfg := config.LoadConfig()

	log.Println(cfg)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection established")
}
