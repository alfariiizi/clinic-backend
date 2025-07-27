package patient_handler

import (
	"context"
	"fmt"

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

type UpdatePatientPayload struct {
	// TODO: Add fields
}

type UpdatePatientInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdatePatientPayload `json:"body" contentType:"application/json"`
}

type UpdatePatientOutput types.OutputResponseMessage

type UpdatePatientHandler model.HTTPHandler[UpdatePatientInput, UpdatePatientOutput]

type updatePatient struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdatePatient(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdatePatientHandler {
	h := &updatePatient{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updatePatient) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/patients", method.Operation{
		Summary:     "Update patient",
		Description: "Update an existing patient by ID",
		Tags:        []string{"Patient"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updatePatient) Handler(ctx context.Context, input *UpdatePatientInput) (*UpdatePatientOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Patient.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find patient with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("patient with ID %s not found", id)
	}

	exec := h.client.Patient.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(patient.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update patient with ID %s: %w", id, err)
	}

	return (*UpdatePatientOutput)(types.GenerateOutputResponseMessage("Successfully updated patient")), nil
}
