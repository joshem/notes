package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notes/backend/models"
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

	if err := h.db.AddNote(newNote); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add note"})
		return
	}

	g.JSON(http.StatusOK, newNote)

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

	note, err := h.db.GetNote(noteId, userId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedNote models.Note
	if err := g.ShouldBindJSON(&updatedNote); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.UpdateNote(note, updatedNote); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	note, err := h.db.GetNote(noteId, userId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if note.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"err": "Note not found"})
		return
	}

	g.JSON(http.StatusOK, nil)
}

func (h *Handler) LinkTagToNote(c *gin.Context) {
	noteId := c.Param("note_id")
	tagId := c.Param("tag_id")
	userId, err := getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})
	}

	note, err := h.db.GetNote(noteId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	tag, err := h.db.GetTag(tagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	noteTag, err := h.db.LinkTagToNote(note, tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link note and tag"})
		return
	}

	c.JSON(http.StatusOK, noteTag)

}
