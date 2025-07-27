package service_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
)

type ListServiceInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListServiceOutput types.OutputResponseData[ListServiceData]

type ListServiceData []*db.Service

type ListServiceHandler model.HTTPHandler[ListServiceInput, ListServiceOutput]

type listService struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListService(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListServiceHandler {
	h := &listService{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listService) RegisterRoutes() {
	api := h.api
	method.GET(api, "/services", method.Operation{
		Summary:     "List Services",
		Description: "Retrieve a list of services",
		Tags:        []string{"Service"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listService) Handler(ctx context.Context, input *ListServiceInput) (*ListServiceOutput, error) {
	items, err := h.client.Service.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListServiceOutput)(types.GenerateOutputResponseData(ListServiceData(items))), nil
}
