package order_handler

import (
	"context"
	"fmt"

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

type UpdateOrderPayload struct {
	// TODO: Add fields
}

type UpdateOrderInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateOrderPayload `json:"body" contentType:"application/json"`
}

type UpdateOrderOutput types.OutputResponseMessage

type UpdateOrderHandler model.HTTPHandler[UpdateOrderInput, UpdateOrderOutput]

type updateOrder struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateOrder(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateOrderHandler {
	h := &updateOrder{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateOrder) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/orders", method.Operation{
		Summary:     "Update order",
		Description: "Update an existing order by ID",
		Tags:        []string{"Order"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateOrder) Handler(ctx context.Context, input *UpdateOrderInput) (*UpdateOrderOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Order.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find order with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}

	exec := h.client.Order.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(order.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update order with ID %s: %w", id, err)
	}

	return (*UpdateOrderOutput)(types.GenerateOutputResponseMessage("Successfully updated order")), nil
}
