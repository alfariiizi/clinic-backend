package clinic_handler

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

type ListClinicInput struct {
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListClinicOutput types.OutputResponseData[ListClinicData]

type ListClinicData []*db.Clinic

type ListClinicHandler model.HTTPHandler[ListClinicInput, ListClinicOutput]

type listClinic struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListClinic(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListClinicHandler {
	h := &listClinic{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listClinic) RegisterRoutes() {
	api := h.api
	method.GET(api, "/clinics", method.Operation{
		Summary:     "List Clinics",
		Description: "Retrieve a list of clinics",
		Tags:        []string{"Clinic"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listClinic) Handler(ctx context.Context, input *ListClinicInput) (*ListClinicOutput, error) {
	items, err := h.client.Clinic.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListClinicOutput)(types.GenerateOutputResponseData(ListClinicData(items))), nil
}
