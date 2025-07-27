package product_handler

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

type ListProductInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListProductOutput types.OutputResponseData[ListProductData]

type ListProductData []*db.Product

type ListProductHandler model.HTTPHandler[ListProductInput, ListProductOutput]

type listProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListProductHandler {
	h := &listProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listProduct) RegisterRoutes() {
	api := h.api
	method.GET(api, "/products", method.Operation{
		Summary:     "List Products",
		Description: "Retrieve a list of products",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listProduct) Handler(ctx context.Context, input *ListProductInput) (*ListProductOutput, error) {
	items, err := h.client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListProductOutput)(types.GenerateOutputResponseData(ListProductData(items))), nil
}
