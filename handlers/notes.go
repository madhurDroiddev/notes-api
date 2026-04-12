package handlers

import (
	"net/http"
	"notes-api/models"
	"notes-api/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	repo repository.NoteRepository
}

func NewNoteHandler(repo repository.NoteRepository) *NoteHandler {
	return &NoteHandler{repo: repo}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}
	userID := c.GetInt("user_id")
	note, err := h.repo.Create(models.Note{
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

func (h *NoteHandler) GetNotes(c *gin.Context) {
	userID := c.GetInt("user_id")
	notes, err := h.repo.GetByUser(userID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not fetch notes")
		return
	}
	successResponse(c, notes)
}

func (h *NoteHandler) GetNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	note, err := h.repo.GetByID(id, userID)
	if err != nil {
		errorResponse(c, http.StatusNotFound, "Note not found")
		return
	}
	successResponse(c, note)
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}
	err := h.repo.Update(models.Note{
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

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	if err := h.repo.Delete(id, userID); err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not delete note")
		return
	}
	successResponse(c, "Note deleted")
}

func (h *NoteHandler) SearchNotes(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Search query is required")
		return
	}

	userID := c.GetInt("user_id")
	notes, err := h.repo.Search(userID, query)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Search failed")
		return
	}

	successResponse(c, notes)
}
