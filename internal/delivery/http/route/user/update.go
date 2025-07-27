package user_handler

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type UpdateUserPayload struct {
	// TODO: Add fields
}

type UpdateUserInput struct {
	ID   string            `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateUserPayload `json:"body" contentType:"application/json"`
}

type UpdateUserOutput types.OutputResponseMessage

type UpdateUserHandler model.HTTPHandler[UpdateUserInput, UpdateUserOutput]

type updateUser struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateUser(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateUserHandler {
	h := &updateUser{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateUser) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/users", method.Operation{
		Summary:     "Update user",
		Description: "Update an existing user by ID",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateUser) Handler(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	exec := h.client.User.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(user.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update user with ID %s: %w", id, err)
	}

	return (*UpdateUserOutput)(types.GenerateOutputResponseMessage("Successfully updated user")), nil
}
