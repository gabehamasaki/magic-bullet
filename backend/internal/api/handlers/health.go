package handlers

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type PingInput struct {
}

type PingOutput struct {
	Body string `json:"body" example:"pong" doc:"Ping response"`
}

func (h *Handler) registerHealthRoutes() {
	huma.Register(h.API, huma.Operation{
		OperationID: "ping",
		Method:      http.MethodGet,
		Path:        "/ping",
		Summary:     "Ping the server",
		Description: "Check the server's health status.",
		Tags:        []string{"Health"},
	}, h.ping)
}

func (h *Handler) ping(_ context.Context, input *PingInput) (*PingOutput, error) {
	return &PingOutput{
		Body: "pong",
	}, nil
}
