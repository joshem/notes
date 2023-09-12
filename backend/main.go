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
	SetAuthRoutes(router, handler)
	SetTagRoutes(router, handler)

	err := router.Run(":8080") // Run the server on port 8080}
	if err != nil {
		jwalterweatherman.FATAL.Panicf(err.Error())
	}
}

func SetTagRoutes(router *gin.Engine, handler *routes.Handler) {
	tags := router.Group("/api/tag")
	tags.POST("/", handler.CreateTag)
	tags.PUT("/:id", handler.UpdateTag)
	tags.DELETE("/:id", handler.DeleteTag)

}

func SetNoteRoutes(router *gin.Engine, handler *routes.Handler) {
	notes := router.Group("/api/notes")
	notes.POST("/", handler.CreateNote)
	notes.POST("/", handler.CreateNote)
	notes.DELETE("/:id", handler.DeleteNote)
}

func SetAuthRoutes(router *gin.Engine, handler *routes.Handler) {
	auth := router.Group("/api/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
}
