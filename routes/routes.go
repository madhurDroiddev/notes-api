package routes

import (
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, noteHandler *handlers.NoteHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	notes := r.Group("/notes")
	notes.Use(middleware.AuthMiddleware())
	{
		notes.POST("/", noteHandler.CreateNote)
		notes.GET("/", noteHandler.GetNotes)
		notes.GET("/search", noteHandler.SearchNotes)
		notes.GET("/:id", noteHandler.GetNote)
		notes.PUT("/:id", noteHandler.UpdateNote)
		notes.DELETE("/:id", noteHandler.DeleteNote)
	}
}
