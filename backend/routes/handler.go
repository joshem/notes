package routes

import "gorm.io/gorm"

// Handler is the main structure to access all API calls within this package.
type Handler struct {
	db         *gorm.DB
	privateKey []byte
}

// NewHandler is a constructor to initialize the Handler.
func NewHandler() *Handler {
	return &Handler{
		db:         nil,
		privateKey: nil,
	}
}
