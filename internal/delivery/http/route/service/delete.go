package service_handler

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

type DeleteServiceInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteServiceOutput types.OutputResponseMessage

type DeleteServiceHandler model.HTTPHandler[DeleteServiceInput, DeleteServiceOutput]

type deleteService struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteService(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteServiceHandler {
	h := &deleteService{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteService) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/services", method.Operation{
		Summary:     "Delete Service",
		Description: "Delete a service by ID",
		Tags:        []string{"Service"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteService) Handler(ctx context.Context, input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid service ID: %w", err)
	}
	err = h.client.Service.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete service: %w", err)
	}

	return (*DeleteServiceOutput)(types.GenerateOutputResponseMessage("Successfully deleted service")), nil
}
