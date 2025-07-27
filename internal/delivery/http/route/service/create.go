package service_handler

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

type CreateServicePayload struct {
	// TODO: Add fields
}

type CreateServiceInput struct {
	model.HTTPInputHeaderTenant
	Body CreateServicePayload `json:"body" contentType:"application/json"`
}

type CreateServiceOutput types.OutputResponseData[CreateServiceData]

type CreateServiceData *db.Service

type CreateServiceHandler model.HTTPHandler[CreateServiceInput, CreateServiceOutput]

type createService struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateService(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateServiceHandler {
	h := &createService{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createService) RegisterRoutes() {
	api := h.api
	method.POST(api, "/services", method.Operation{
		Summary:     "Create Service",
		Description: "Create a new service",
		Tags:        []string{"Service"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createService) Handler(ctx context.Context, input *CreateServiceInput) (*CreateServiceOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Service.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateServiceOutput)(types.GenerateOutputResponseData(CreateServiceData(item))), nil
}
