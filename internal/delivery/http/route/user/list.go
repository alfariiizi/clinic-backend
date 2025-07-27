package user_handler

import (
	"context"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

type ListUserInput struct {
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListUserOutput types.OutputResponseData[ListUserData]

type ListUserData []*db.User

type ListUserHandler model.HTTPHandler[ListUserInput, ListUserOutput]

type listUser struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListUser(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListUserHandler {
	h := &listUser{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listUser) RegisterRoutes() {
	api := h.api
	method.GET(api, "/users", method.Operation{
		Summary:     "List Users",
		Description: "Retrieve a list of users",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listUser) Handler(ctx context.Context, input *ListUserInput) (*ListUserOutput, error) {
	items, err := h.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListUserOutput)(types.GenerateOutputResponseData(ListUserData(items))), nil
}
