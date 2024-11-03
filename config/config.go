package config

import (
	"log"
	"os"

	"github.com/Junx27/ho_go_day23/model"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func InitializeDatabase() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	log.Printf("Connecting to mySQL database at %s:%s...", host, port)
	model.ConnectDatabase(user, password, host, port, dbname)
	log.Printf("mySQL database successfully connected!")
}
