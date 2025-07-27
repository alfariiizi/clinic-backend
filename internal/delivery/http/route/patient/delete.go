package patient_handler

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

type DeletePatientInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeletePatientOutput types.OutputResponseMessage

type DeletePatientHandler model.HTTPHandler[DeletePatientInput, DeletePatientOutput]

type deletePatient struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeletePatient(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeletePatientHandler {
	h := &deletePatient{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deletePatient) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/patients", method.Operation{
		Summary:     "Delete Patient",
		Description: "Delete a patient by ID",
		Tags:        []string{"Patient"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deletePatient) Handler(ctx context.Context, input *DeletePatientInput) (*DeletePatientOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid patient ID: %w", err)
	}
	err = h.client.Patient.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete patient: %w", err)
	}

	return (*DeletePatientOutput)(types.GenerateOutputResponseMessage("Successfully deleted patient")), nil
}
