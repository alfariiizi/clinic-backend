package user_handler

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type DeleteUserInput struct {
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteUserOutput types.OutputResponseMessage

type DeleteUserHandler model.HTTPHandler[DeleteUserInput, DeleteUserOutput]

type deleteUser struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteUser(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteUserHandler {
	h := &deleteUser{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteUser) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/users", method.Operation{
		Summary:     "Delete User",
		Description: "Delete a user by ID",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteUser) Handler(ctx context.Context, input *DeleteUserInput) (*DeleteUserOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	err = h.client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}

	return (*DeleteUserOutput)(types.GenerateOutputResponseMessage("Successfully deleted user")), nil
}
