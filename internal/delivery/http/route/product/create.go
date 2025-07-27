package product_handler

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

type CreateProductPayload struct {
	// TODO: Add fields
}

type CreateProductInput struct {
	model.HTTPInputHeaderTenant
	Body CreateProductPayload `json:"body" contentType:"application/json"`
}

type CreateProductOutput types.OutputResponseData[CreateProductData]

type CreateProductData *db.Product

type CreateProductHandler model.HTTPHandler[CreateProductInput, CreateProductOutput]

type createProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateProductHandler {
	h := &createProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createProduct) RegisterRoutes() {
	api := h.api
	method.POST(api, "/products", method.Operation{
		Summary:     "Create Product",
		Description: "Create a new product",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createProduct) Handler(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Product.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateProductOutput)(types.GenerateOutputResponseData(CreateProductData(item))), nil
}
