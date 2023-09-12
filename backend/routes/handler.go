package routes

import (
	"notes/backend/models"
)

// Handler is the main structure to access all API calls within this package.
type Handler struct {
	db         models.Storage
	privateKey []byte
}

// NewHandler is a constructor to initialize the Handler.
func NewHandler() *Handler {
	return &Handler{
		db:         nil,
		privateKey: nil,
	}
}
