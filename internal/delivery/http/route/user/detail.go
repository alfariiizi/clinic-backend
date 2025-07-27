package user_handler

import (
	"context"

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

type DetailUserInput struct {
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailUserOutput types.OutputResponseData[DetailUserData]

type DetailUserData *db.User

type DetailUserHandler model.HTTPHandler[DetailUserInput, DetailUserOutput]

type detailUser struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailUser(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailUserHandler {
	h := &detailUser{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailUser) RegisterRoutes() {
	api := h.api
	method.GET(api, "/users/{id}", method.Operation{
		Summary:     "Detail User",
		Description: "Retrieve details of a user by ID",
		Tags:        []string{"User"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailUser) Handler(ctx context.Context, input *DetailUserInput) (*DetailUserOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.User.Query().
		Where(user.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailUserOutput)(types.GenerateOutputResponseData(DetailUserData(item))), nil
}
