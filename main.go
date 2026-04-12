package main

import (
	"notes-api/config"

	"notes-api/handlers"
	"notes-api/repository"
	"notes-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Repositories
	userRepo := repository.NewUserRepository(config.DB)
	noteRepo := repository.NewNoteRepository(config.DB)

	// Handlers
	authHandler := handlers.NewAuthHandler(userRepo)
	noteHandler := handlers.NewNoteHandler(noteRepo)

	// Routes
	r := gin.Default()
	routes.SetupRoutes(r, authHandler, noteHandler)

	r.Run(":8080")
}
