package db

import (
	"log"
	"user-registration/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// All the GORM DB related method can be checked from here
// https://pkg.go.dev/gorm.io/gorm@v1.24.7-0.20230306060331-85eaf9eeda11
var DB *gorm.DB

func InitDB() {
	cfg := configs.LoadConfig()
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
