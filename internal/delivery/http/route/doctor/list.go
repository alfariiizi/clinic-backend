package doctor_handler

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

type ListDoctorInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListDoctorOutput types.OutputResponseData[ListDoctorData]

type ListDoctorData []*db.Doctor

type ListDoctorHandler model.HTTPHandler[ListDoctorInput, ListDoctorOutput]

type listDoctor struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListDoctor(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListDoctorHandler {
	h := &listDoctor{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listDoctor) RegisterRoutes() {
	api := h.api
	method.GET(api, "/doctors", method.Operation{
		Summary:     "List Doctors",
		Description: "Retrieve a list of doctors",
		Tags:        []string{"Doctor"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listDoctor) Handler(ctx context.Context, input *ListDoctorInput) (*ListDoctorOutput, error) {
	items, err := h.client.Doctor.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListDoctorOutput)(types.GenerateOutputResponseData(ListDoctorData(items))), nil
}
