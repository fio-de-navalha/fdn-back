package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// // Auto-migrate the User model to create the users table
	// err = DB.AutoMigrate(
	// 	&customer.Customer{},
	// 	&barber.Barber{},
	// 	&service.Service{},
	// )
	// if err != nil {
	// 	return err
	// }

	return nil
}
