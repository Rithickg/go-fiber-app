// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func LoadConfig() {
//     // Load .env file
//     err := godotenv.Load()
//     if err != nil {
//         log.Fatalf("Error loading .env file")
//     }

//     // Fetching environment variables
//     dbHost := os.Getenv("DB_HOST")
//     dbUser := os.Getenv("DB_USER")
//     dbPassword := os.Getenv("DB_PASSWORD")
//     dbName := os.Getenv("DB_NAME")

//     // Display values (for debugging purposes)
//     fmt.Printf("Database Config: %s, %s, %s, %s\n", dbHost, dbUser, dbPassword, dbName)
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
}
func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg := &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBPort:     os.Getenv("DB_PORT"),
    }

    log.Printf("Loaded configuration: %+v\n", cfg) // Add this line to debug

    return cfg
}
