package patient_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/patient"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type DetailPatientInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailPatientOutput types.OutputResponseData[DetailPatientData]

type DetailPatientData *db.Patient

type DetailPatientHandler model.HTTPHandler[DetailPatientInput, DetailPatientOutput]

type detailPatient struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailPatient(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailPatientHandler {
	h := &detailPatient{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailPatient) RegisterRoutes() {
	api := h.api
	method.GET(api, "/patients/{id}", method.Operation{
		Summary:     "Detail Patient",
		Description: "Retrieve details of a patient by ID",
		Tags:        []string{"Patient"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailPatient) Handler(ctx context.Context, input *DetailPatientInput) (*DetailPatientOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Patient.Query().
		Where(patient.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailPatientOutput)(types.GenerateOutputResponseData(DetailPatientData(item))), nil
}
