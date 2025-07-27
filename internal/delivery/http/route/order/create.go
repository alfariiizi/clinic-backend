package order_handler

import (
	"context"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
)

type CreateOrderPayload struct {
	// TODO: Add fields
}

type CreateOrderInput struct {
	model.HTTPInputHeaderTenant
	Body CreateOrderPayload `json:"body" contentType:"application/json"`
}

type CreateOrderOutput types.OutputResponseData[CreateOrderData]

type CreateOrderData *db.Order

type CreateOrderHandler model.HTTPHandler[CreateOrderInput, CreateOrderOutput]

type createOrder struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateOrder(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateOrderHandler {
	h := &createOrder{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createOrder) RegisterRoutes() {
	api := h.api
	method.POST(api, "/orders", method.Operation{
		Summary:     "Create Order",
		Description: "Create a new order",
		Tags:        []string{"Order"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createOrder) Handler(ctx context.Context, input *CreateOrderInput) (*CreateOrderOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Order.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateOrderOutput)(types.GenerateOutputResponseData(CreateOrderData(item))), nil
}
