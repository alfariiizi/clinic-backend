package product_handler

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

type DeleteProductInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteProductOutput types.OutputResponseMessage

type DeleteProductHandler model.HTTPHandler[DeleteProductInput, DeleteProductOutput]

type deleteProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteProductHandler {
	h := &deleteProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteProduct) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/products", method.Operation{
		Summary:     "Delete Product",
		Description: "Delete a product by ID",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteProduct) Handler(ctx context.Context, input *DeleteProductInput) (*DeleteProductOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %w", err)
	}
	err = h.client.Product.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete product: %w", err)
	}

	return (*DeleteProductOutput)(types.GenerateOutputResponseMessage("Successfully deleted product")), nil
}
