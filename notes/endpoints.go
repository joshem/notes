package notes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Handler struct {
	db         *gorm.DB
	privateKey []byte
}

func (h *Handler) Create(g *gin.Context) {
	var NewNote Note
	if err := g.ShouldBindJSON(&NewNote); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := getUserIdFromToken(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse token"})
	}
	NewNote.UserID = userId

	g.JSON(http.StatusOK, NewNote)

	h.db.Create(NewNote)

}

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
