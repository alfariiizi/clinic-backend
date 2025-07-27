package service_handler

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/core/model"
	coreservice "github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/service"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type UpdateServicePayload struct {
	// TODO: Add fields
}

type UpdateServiceInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateServicePayload `json:"body" contentType:"application/json"`
}

type UpdateServiceOutput types.OutputResponseMessage

type UpdateServiceHandler model.HTTPHandler[UpdateServiceInput, UpdateServiceOutput]

type updateService struct {
	api     huma.API
	service *coreservice.Services
	client  *db.Client
}

func NewUpdateService(
	api *api.HttpApi,
	service *coreservice.Services,
	client *db.Client,
) UpdateServiceHandler {
	h := &updateService{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateService) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/services", method.Operation{
		Summary:     "Update service",
		Description: "Update an existing service by ID",
		Tags:        []string{"Service"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateService) Handler(ctx context.Context, input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Service.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find service with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("service with ID %s not found", id)
	}

	exec := h.client.Service.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(service.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update service with ID %s: %w", id, err)
	}

	return (*UpdateServiceOutput)(types.GenerateOutputResponseMessage("Successfully updated service")), nil
}
