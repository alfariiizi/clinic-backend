package feature_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/feature"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type DetailFeatureInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailFeatureOutput types.OutputResponseData[DetailFeatureData]

type DetailFeatureData *db.Feature

type DetailFeatureHandler model.HTTPHandler[DetailFeatureInput, DetailFeatureOutput]

type detailFeature struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailFeature(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailFeatureHandler {
	h := &detailFeature{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailFeature) RegisterRoutes() {
	api := h.api
	method.GET(api, "/features/{id}", method.Operation{
		Summary:     "Detail Feature",
		Description: "Retrieve details of a feature by ID",
		Tags:        []string{"Feature"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailFeature) Handler(ctx context.Context, input *DetailFeatureInput) (*DetailFeatureOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Feature.Query().
		Where(feature.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailFeatureOutput)(types.GenerateOutputResponseData(DetailFeatureData(item))), nil
}
