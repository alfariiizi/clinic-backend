package patient_handler

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

type CreatePatientPayload struct {
	// TODO: Add fields
}

type CreatePatientInput struct {
	model.HTTPInputHeaderTenant
	Body CreatePatientPayload `json:"body" contentType:"application/json"`
}

type CreatePatientOutput types.OutputResponseData[CreatePatientData]

type CreatePatientData *db.Patient

type CreatePatientHandler model.HTTPHandler[CreatePatientInput, CreatePatientOutput]

type createPatient struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreatePatient(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreatePatientHandler {
	h := &createPatient{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createPatient) RegisterRoutes() {
	api := h.api
	method.POST(api, "/patients", method.Operation{
		Summary:     "Create Patient",
		Description: "Create a new patient",
		Tags:        []string{"Patient"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createPatient) Handler(ctx context.Context, input *CreatePatientInput) (*CreatePatientOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Patient.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreatePatientOutput)(types.GenerateOutputResponseData(CreatePatientData(item))), nil
}
