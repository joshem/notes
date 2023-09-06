package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/jwalterweatherman"
	"notes/backend/routes"
)

func main() {
	router := gin.Default()
	handler := routes.NewHandler()

	// Define your routes here
	SetNoteRoutes(router, handler)

	err := router.Run(":8080") // Run the server on port 8080}
	if err != nil {
		jwalterweatherman.FATAL.Panicf(err.Error())
	}
}

func SetNoteRoutes(router *gin.Engine, handler *routes.Handler) {
	notes := router.Group("/api/notes")
	notes.POST("/", handler.Create)
	notes.POST("/", handler.Create)
	notes.DELETE("/:id", handler.Delete)
}

func SetAuthRoutes(router *gin.Engine, handler *routes.Handler) {
	auth := router.Group("/api/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
}
