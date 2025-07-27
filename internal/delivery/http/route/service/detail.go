package service_handler

import (
	"context"

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

type DetailServiceInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailServiceOutput types.OutputResponseData[DetailServiceData]

type DetailServiceData *db.Service

type DetailServiceHandler model.HTTPHandler[DetailServiceInput, DetailServiceOutput]

type detailService struct {
	api     huma.API
	service *coreservice.Services
	client  *db.Client
}

func NewDetailService(
	api *api.HttpApi,
	service *coreservice.Services,
	client *db.Client,
) DetailServiceHandler {
	h := &detailService{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailService) RegisterRoutes() {
	api := h.api
	method.GET(api, "/services/{id}", method.Operation{
		Summary:     "Detail Service",
		Description: "Retrieve details of a service by ID",
		Tags:        []string{"Service"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailService) Handler(ctx context.Context, input *DetailServiceInput) (*DetailServiceOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Service.Query().
		Where(service.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailServiceOutput)(types.GenerateOutputResponseData(DetailServiceData(item))), nil
}
