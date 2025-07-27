package feature_handler

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

type CreateFeaturePayload struct {
	// TODO: Add fields
}

type CreateFeatureInput struct {
	model.HTTPInputHeaderTenant
	Body CreateFeaturePayload `json:"body" contentType:"application/json"`
}

type CreateFeatureOutput types.OutputResponseData[CreateFeatureData]

type CreateFeatureData *db.Feature

type CreateFeatureHandler model.HTTPHandler[CreateFeatureInput, CreateFeatureOutput]

type createFeature struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateFeature(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateFeatureHandler {
	h := &createFeature{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createFeature) RegisterRoutes() {
	api := h.api
	method.POST(api, "/features", method.Operation{
		Summary:     "Create Feature",
		Description: "Create a new feature",
		Tags:        []string{"Feature"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createFeature) Handler(ctx context.Context, input *CreateFeatureInput) (*CreateFeatureOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Feature.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateFeatureOutput)(types.GenerateOutputResponseData(CreateFeatureData(item))), nil
}
