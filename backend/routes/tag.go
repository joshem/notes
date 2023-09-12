package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notes/backend/models"
)

// CreateTag handles creating a new Tag.
func (h *Handler) CreateTag(g *gin.Context) {
	var newTag models.Tag
	if err := g.ShouldBindJSON(&newTag); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.AddTag(newTag); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tag"})
		return
	}

	g.JSON(http.StatusOK, newTag)
}

// UpdateTag handles updating a Tag.
func (h *Handler) UpdateTag(g *gin.Context) {
	tagId := g.Param("id")

	old, err := h.db.GetTag(tagId)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	var updated models.Tag
	if err := g.ShouldBindJSON(&updated); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.UpdateTag(old, updated); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}

	g.JSON(http.StatusOK, updated)
}

// DeleteTag handles deleting a tag
func (h *Handler) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")

	tag, err := h.db.GetTag(tagID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
	}

	if err := h.db.DeleteTag(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	// Return a success response
	c.JSON(http.StatusNoContent, nil)
}
