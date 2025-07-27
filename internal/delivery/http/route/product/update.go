package product_handler

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/product"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type UpdateProductPayload struct {
	// TODO: Add fields
}

type UpdateProductInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateProductPayload `json:"body" contentType:"application/json"`
}

type UpdateProductOutput types.OutputResponseMessage

type UpdateProductHandler model.HTTPHandler[UpdateProductInput, UpdateProductOutput]

type updateProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateProductHandler {
	h := &updateProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateProduct) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/products", method.Operation{
		Summary:     "Update product",
		Description: "Update an existing product by ID",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateProduct) Handler(ctx context.Context, input *UpdateProductInput) (*UpdateProductOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Product.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find product with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("product with ID %s not found", id)
	}

	exec := h.client.Product.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(product.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update product with ID %s: %w", id, err)
	}

	return (*UpdateProductOutput)(types.GenerateOutputResponseMessage("Successfully updated product")), nil
}
