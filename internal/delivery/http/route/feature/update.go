package feature_handler

import (
	"context"
	"fmt"

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

type UpdateFeaturePayload struct {
	// TODO: Add fields
}

type UpdateFeatureInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateFeaturePayload `json:"body" contentType:"application/json"`
}

type UpdateFeatureOutput types.OutputResponseMessage

type UpdateFeatureHandler model.HTTPHandler[UpdateFeatureInput, UpdateFeatureOutput]

type updateFeature struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateFeature(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateFeatureHandler {
	h := &updateFeature{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateFeature) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/features", method.Operation{
		Summary:     "Update feature",
		Description: "Update an existing feature by ID",
		Tags:        []string{"Feature"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateFeature) Handler(ctx context.Context, input *UpdateFeatureInput) (*UpdateFeatureOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Feature.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find feature with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("feature with ID %s not found", id)
	}

	exec := h.client.Feature.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(feature.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update feature with ID %s: %w", id, err)
	}

	return (*UpdateFeatureOutput)(types.GenerateOutputResponseMessage("Successfully updated feature")), nil
}
