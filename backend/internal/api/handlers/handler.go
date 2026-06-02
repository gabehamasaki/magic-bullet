package handlers

import (
	"github.com/danielgtaylor/huma/v2"
)

type Handler struct {
	API huma.API
}

func NewHandler(api huma.API) *Handler {
	return &Handler{
		API: api,
	}
}

func (h *Handler) RegisterRoutes() {
	h.registerHealthRoutes()
	h.registerGetUserRoute()
}
