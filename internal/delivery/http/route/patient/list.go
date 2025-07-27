package patient_handler

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

type ListPatientInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListPatientOutput types.OutputResponseData[ListPatientData]

type ListPatientData []*db.Patient

type ListPatientHandler model.HTTPHandler[ListPatientInput, ListPatientOutput]

type listPatient struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListPatient(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListPatientHandler {
	h := &listPatient{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listPatient) RegisterRoutes() {
	api := h.api
	method.GET(api, "/patients", method.Operation{
		Summary:     "List Patients",
		Description: "Retrieve a list of patients",
		Tags:        []string{"Patient"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listPatient) Handler(ctx context.Context, input *ListPatientInput) (*ListPatientOutput, error) {
	items, err := h.client.Patient.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListPatientOutput)(types.GenerateOutputResponseData(ListPatientData(items))), nil
}
