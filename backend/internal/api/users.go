package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UserProfile struct {
	ID    string `json:"id" example:"usr_9988" doc:"Unique user identifier"`
	Email string `json:"email" example:"developer@domain.com" doc:"Primary email address"`
	Role  string `json:"role" example:"admin" doc:"Role assigned to the user"`
}

type GetUserInput struct {
	ID string `path:"id" example:"usr_9988" doc:"User identifier"`
}

type GetUserOutput struct {
	Body UserProfile
}

func registerUserRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        "/users/{id}",
		Summary:     "Fetch user profile",
		Description: "Get a user by their unique ID.",
		Tags:        []string{"Users"},
	}, getUser)
}

func getUser(_ context.Context, input *GetUserInput) (*GetUserOutput, error) {
	return &GetUserOutput{
		Body: UserProfile{
			ID:    input.ID,
			Email: "developer@domain.com",
			Role:  "admin",
		},
	}, nil
}
