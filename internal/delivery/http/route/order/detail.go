package order_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/order"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type DetailOrderInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailOrderOutput types.OutputResponseData[DetailOrderData]

type DetailOrderData *db.Order

type DetailOrderHandler model.HTTPHandler[DetailOrderInput, DetailOrderOutput]

type detailOrder struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailOrder(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailOrderHandler {
	h := &detailOrder{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailOrder) RegisterRoutes() {
	api := h.api
	method.GET(api, "/orders/{id}", method.Operation{
		Summary:     "Detail Order",
		Description: "Retrieve details of a order by ID",
		Tags:        []string{"Order"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailOrder) Handler(ctx context.Context, input *DetailOrderInput) (*DetailOrderOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Order.Query().
		Where(order.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailOrderOutput)(types.GenerateOutputResponseData(DetailOrderData(item))), nil
}
