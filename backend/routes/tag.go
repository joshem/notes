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

	h.db.Create(&newTag)

	g.JSON(http.StatusOK, newTag)
}

// UpdateTag handles updating a Tag.
func (h *Handler) UpdateTag(g *gin.Context) {
	tagId := g.Param("id")

	var tag models.Tag
	h.db.Where("id = ?", tagId).First(&tag)

	if tag.Id == 0 {
		g.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	var updated models.Tag
	if err := g.ShouldBindJSON(&updated); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.Name = updated.Name

	h.db.Save(&tag)

	g.JSON(http.StatusOK, updated)
}

// DeleteTag handles deleting a tag
func (h *Handler) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")

	var tag models.Tag
	h.db.Where("id = ?", tagID).First(&tag)

	if tag.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	h.db.Delete(&tag)

	// Return a success response
	c.JSON(http.StatusNoContent, nil)
}
