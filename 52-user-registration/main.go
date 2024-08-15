package main

import (
	"user-registration/database"
	"user-registration/models"
	"user-registration/routes"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.User{})

	router := routes.SetupRouter()
	router.Run(":8080")
}
