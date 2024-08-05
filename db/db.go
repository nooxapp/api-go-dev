package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Ping() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		log.Fatalf("Missing DATABSE_URL environment variable")
	}
	fmt.Println("Successfully connected!")
}
