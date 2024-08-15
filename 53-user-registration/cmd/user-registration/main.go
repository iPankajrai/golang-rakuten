package main

import (
	"log"
	// "user-registration/internal/app/repositories"
	"user-registration/internal/app/routes"
	"user-registration/internal/db"
	"user-registration/internal/worker"
	"user-registration/pkg/models"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.User{})

	// Start the worker pool with a specified number of workers
	worker.StartReadWorkerPool(10)

	log.Println("Starting server on port 8080")
	r := routes.SetupRouter()
	r.Run(":8080")
}
