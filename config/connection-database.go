package config

import (
	"fmt"
	"log"
	"os"

	"github.com/IBHMM/jwtauth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")
	timezone := os.Getenv("Zone")

	if port == "" || user == "" || password == "" || name == "" || host == "" || timezone == "" {
		return fmt.Errorf("one or more required environment variables are not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s timezone=%s", host, user, password, name, port, timezone)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := connection.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	log.Println("Database Connection is succesfull")
	DB = connection
	return nil
}

func Disconnect(connection *gorm.DB) error {
	db, err := connection.DB()

	if err != nil {
		return fmt.Errorf("failed to get underlying database connection: %v", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %v", err)
	}

	return nil
}
