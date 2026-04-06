package routes

import (
	"notes-api/handlers"

	"github.com/gin-gonic/gin"

	"notes-api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	notes := r.Group("/notes")
	notes.Use(middleware.AuthMiddleware())
	{
		notes.POST("/", handlers.CreateNote)
		notes.GET("/", handlers.GetNotes)
		notes.GET("/:id", handlers.GetNote)
		notes.PUT("/:id", handlers.UpdateNote)
		notes.DELETE("/:id", handlers.DeleteNote)
		notes.GET("/search", handlers.SearchNotes)
	}
}
