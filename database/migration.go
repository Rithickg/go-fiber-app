package database

import (
	"log"
	"my-fiber-app/models"
)

// AutoMigrateModels automatically migrates all the database models
func AutoMigrateModels() {
	if DB == nil {
		log.Fatal("Database connection is not initialized. Cannot run migrations.")
	}

	log.Println("Starting database migrations...")

	// Auto-migrate the User model (add other models here as necessary)
	err := DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	log.Println("Migration for User model completed.")
}
