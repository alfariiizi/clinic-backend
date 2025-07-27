package order_handler

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

type ListOrderInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListOrderOutput types.OutputResponseData[ListOrderData]

type ListOrderData []*db.Order

type ListOrderHandler model.HTTPHandler[ListOrderInput, ListOrderOutput]

type listOrder struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListOrder(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListOrderHandler {
	h := &listOrder{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listOrder) RegisterRoutes() {
	api := h.api
	method.GET(api, "/orders", method.Operation{
		Summary:     "List Orders",
		Description: "Retrieve a list of orders",
		Tags:        []string{"Order"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listOrder) Handler(ctx context.Context, input *ListOrderInput) (*ListOrderOutput, error) {
	items, err := h.client.Order.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListOrderOutput)(types.GenerateOutputResponseData(ListOrderData(items))), nil
}
