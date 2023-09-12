package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notes/backend/models"
	"time"
)

// CreateNote is the API endpoint to create a new note.
//
// Returns:
//   - JSOn format of the Note.
func (h *Handler) CreateNote(g *gin.Context) {
	var newNote models.Note
	if err := g.ShouldBindJSON(&newNote); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})
	}
	newNote.UserID = userId

	g.JSON(http.StatusOK, newNote)

	h.db.Create(newNote)

}

// UpdateNote is the endpoint to update the existing Note belonging to the User.
//
// Returns:
//   - JSON of updated Note.
//   - error if no Note exists for this User.
func (h *Handler) UpdateNote(g *gin.Context) {
	noteId := g.Param("id")
	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})

	}

	var note models.Note
	h.db.Where("id = ? AND user_id = ?", noteId, userId).First(&note)

	if note.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"err": "Note not found"})
		return
	}

	var updatedNote models.Note
	if err := g.ShouldBindJSON(&updatedNote); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedNote.Content != "" {
		note.Content = updatedNote.Content
	}

	if updatedNote.Title != "" {
		note.Content = updatedNote.Title
	}

	updatedNote.UpdatedAt = time.Now()

	h.db.Save(&updatedNote)

	g.JSON(http.StatusOK, note)
}

// DeleteNote will remove the Note belonging to this User from storage.
//
// Returns:
//   - An http.StatusOK.
//   - An error if no Note belongs to this User.
func (h *Handler) DeleteNote(g *gin.Context) {
	noteId := g.Param("id")
	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})
		return
	}

	var note models.Note
	h.db.Where("id = ? AND user_id = ?", noteId, userId).First(&note)

	if note.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"err": "Note not found"})
		return
	}

	h.db.Delete(&note)

	g.JSON(http.StatusOK, nil)
}
