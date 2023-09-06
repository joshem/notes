package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// User defines a user who may write a Note to the system.
type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register adds a User to the system.
func (h *Handler) Register(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registration successful"})
}

// Login will initialize a session for the User, providing them with a token.
func (h *Handler) Login(c *gin.Context) {
	var input, user User
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Where("username = ?", input.Username).First(&user)

	token, err := h.generateJwtToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
