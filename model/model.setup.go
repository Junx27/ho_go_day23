package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(user, password, host, port, dbname string) *gorm.DB {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = database
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Auto-migration gagal: %v", err)
	}
	if err := DB.AutoMigrate(&Product{}); err != nil {
		log.Fatalf("Auto-migration gagal: %v", err)
	}

	return DB
}
