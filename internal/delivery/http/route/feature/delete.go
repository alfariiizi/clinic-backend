package feature_handler

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

type DeleteFeatureInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteFeatureOutput types.OutputResponseMessage

type DeleteFeatureHandler model.HTTPHandler[DeleteFeatureInput, DeleteFeatureOutput]

type deleteFeature struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteFeature(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteFeatureHandler {
	h := &deleteFeature{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteFeature) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/features", method.Operation{
		Summary:     "Delete Feature",
		Description: "Delete a feature by ID",
		Tags:        []string{"Feature"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteFeature) Handler(ctx context.Context, input *DeleteFeatureInput) (*DeleteFeatureOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid feature ID: %w", err)
	}
	err = h.client.Feature.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete feature: %w", err)
	}

	return (*DeleteFeatureOutput)(types.GenerateOutputResponseMessage("Successfully deleted feature")), nil
}
