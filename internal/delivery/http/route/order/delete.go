package order_handler

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type DeleteOrderInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteOrderOutput types.OutputResponseMessage

type DeleteOrderHandler model.HTTPHandler[DeleteOrderInput, DeleteOrderOutput]

type deleteOrder struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteOrder(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteOrderHandler {
	h := &deleteOrder{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteOrder) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/orders", method.Operation{
		Summary:     "Delete Order",
		Description: "Delete a order by ID",
		Tags:        []string{"Order"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteOrder) Handler(ctx context.Context, input *DeleteOrderInput) (*DeleteOrderOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid order ID: %w", err)
	}
	err = h.client.Order.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete order: %w", err)
	}

	return (*DeleteOrderOutput)(types.GenerateOutputResponseMessage("Successfully deleted order")), nil
}
