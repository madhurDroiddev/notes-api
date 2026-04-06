package handlers

import (
	"net/http"
	"strconv"

	"notes-api/config"
	"notes-api/models"
	"notes-api/repository"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}
	userID := c.GetInt("user_id")
	note, err := repository.CreateNote(config.DB, models.Note{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not create note")
		return
	}
	successResponse(c, note)
}

func GetNotes(c *gin.Context) {
	userID := c.GetInt("user_id")
	notes, err := repository.GetNotesByUser(config.DB, userID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not fetch notes")
		return
	}
	successResponse(c, notes)
}

func GetNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	note, err := repository.GetNoteByID(config.DB, id, userID)
	if err != nil {
		errorResponse(c, http.StatusNotFound, "Note not found")
		return
	}
	successResponse(c, note)
}

func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}
	err := repository.UpdateNote(config.DB, models.Note{
		ID:      id,
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not update note")
		return
	}
	successResponse(c, "Note updated")
}

func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	if err := repository.DeleteNote(config.DB, id, userID); err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not delete note")
		return
	}
	successResponse(c, "Note deleted")
}

func SearchNotes(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Search query is required")
		return
	}

	userID := c.GetInt("user_id")
	notes, err := repository.SearchNotes(config.DB, userID, query)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Search failed")
		return
	}

	successResponse(c, notes)
}
