package database

import (
	"log"
	"os"

	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto-migrate the User model to create the users table
	if err := DB.AutoMigrate(&customer.Customer{}); err != nil {
		return err
	}

	return nil
}
