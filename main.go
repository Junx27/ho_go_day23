package main

import (
	"log"
	"os"

	"github.com/Junx27/ho_go_day23/config"
	"github.com/Junx27/ho_go_day23/service/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		panic(err)
	}
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.SetTrustedProxies([]string{"127.0.0.1"})
	config.LoadEnv()
	config.InitializeDatabase()
	router.ServerRoutes()

	r.Run(":" + port)
}
