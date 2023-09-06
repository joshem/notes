package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Note is the note structure that the user will write to the system.
type Note struct {
	Content   string    `json:"content" gorm:"primary_key"`
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create is the API endpoint to create a new note.
//
// Returns:
//  - JSOn format of the Note.
func (h *Handler) Create(g *gin.Context) {
	var newNote Note
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

// Update is the endpoint to update the existing Note belonging to the User.
//
// Returns:
//  - JSON of updated Note.
//  - error if no Note exists for this User.
func (h *Handler) Update(g *gin.Context) {
	noteId := g.Param("id")
	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})

	}

	var note Note
	h.db.Where("id = ? AND user_id = ?", noteId, userId).First(&note)

	if note.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"err": "Note not found"})
		return
	}

	var updatedNote Note
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

// Delete will remove the Note belonging to this User from storage.
//
// Returns:
//  - An http.StatusOK.
//  - An error if no Note belongs to this User.
func (h *Handler) Delete(g *gin.Context) {
	noteId := g.Param("id")
	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})

	}

	var note Note
	h.db.Where("id = ? AND user_id = ?", noteId, userId).First(&note)

	if note.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"err": "Note not found"})
		return
	}

	h.db.Delete(&note)

	g.JSON(http.StatusOK, nil)
}
