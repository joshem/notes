package routes

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth is middleware which assigns a JWT.
func (h *Handler) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("Auth")
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return h.privateKey, nil
		})

		claim, _ := token.Claims.(jwt.MapClaims)
		userID := claim["sub"].(uint)

		context.Set("userId", userID)

		context.Next()
	}
}

func (h *Handler) generateJwtToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"iss": "insert", // todo: specify user
	})

	return token.SignedString(h.privateKey)
}

func getUserIdFromToken(g *gin.Context) (uint, error) {
	userId, exists := g.Get("userId")
	if !exists {
		return 0, errors.New("failed to retrieve token")
	}

	return userId.(uint), nil
}
