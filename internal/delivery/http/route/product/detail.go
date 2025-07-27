package product_handler

import (
	"context"

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

type DetailProductInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailProductOutput types.OutputResponseData[DetailProductData]

type DetailProductData *db.Product

type DetailProductHandler model.HTTPHandler[DetailProductInput, DetailProductOutput]

type detailProduct struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailProduct(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailProductHandler {
	h := &detailProduct{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailProduct) RegisterRoutes() {
	api := h.api
	method.GET(api, "/products/{id}", method.Operation{
		Summary:     "Detail Product",
		Description: "Retrieve details of a product by ID",
		Tags:        []string{"Product"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailProduct) Handler(ctx context.Context, input *DetailProductInput) (*DetailProductOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Product.Query().
		Where(product.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailProductOutput)(types.GenerateOutputResponseData(DetailProductData(item))), nil
}
