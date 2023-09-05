package users

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	db         *gorm.DB
	privateKey string // todo: parse private key
}

func (h *Handler) Register(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registration successful"})
}

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

func (h *Handler) generateJwtToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"iss": "insert", // todo: specify user
	})

	return token.SignedString(h.privateKey)
}
