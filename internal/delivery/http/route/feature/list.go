package feature_handler

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

type ListFeatureInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListFeatureOutput types.OutputResponseData[ListFeatureData]

type ListFeatureData []*db.Feature

type ListFeatureHandler model.HTTPHandler[ListFeatureInput, ListFeatureOutput]

type listFeature struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListFeature(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListFeatureHandler {
	h := &listFeature{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listFeature) RegisterRoutes() {
	api := h.api
	method.GET(api, "/features", method.Operation{
		Summary:     "List Features",
		Description: "Retrieve a list of features",
		Tags:        []string{"Feature"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listFeature) Handler(ctx context.Context, input *ListFeatureInput) (*ListFeatureOutput, error) {
	items, err := h.client.Feature.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListFeatureOutput)(types.GenerateOutputResponseData(ListFeatureData(items))), nil
}
