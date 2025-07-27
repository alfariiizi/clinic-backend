package user_handler

import (
	"context"
	"time"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

type CreateUserPayload struct {
	// TODO: Add fields
}

type CreateUserInput struct {
	Body CreateUserPayload `json:"body" contentType:"application/json"`
}

type CreateUserOutput types.OutputResponseData[CreateUserData]

type CreateUserData *db.User

type CreateUserHandler model.HTTPHandler[CreateUserInput, CreateUserOutput]

type createUser struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateUser(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateUserHandler {
	h := &createUser{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createUser) RegisterRoutes() {
	api := h.api
	method.POST(api, "/users", method.Operation{
		Summary:     "Create User",
		Description: "Create a new user",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createUser) Handler(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.User.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateUserOutput)(types.GenerateOutputResponseData(CreateUserData(item))), nil
}
