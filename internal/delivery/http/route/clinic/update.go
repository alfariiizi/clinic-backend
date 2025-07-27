package clinic_handler

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
	"github.com/danielgtaylor/huma/v2"
)

type UpdateClinicPayload struct {
	// TODO: Add fields
}

type UpdateClinicInput struct {
	ID   string              `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateClinicPayload `json:"body" contentType:"application/json"`
}

type UpdateClinicOutput types.OutputResponseMessage

type UpdateClinicHandler model.HTTPHandler[UpdateClinicInput, UpdateClinicOutput]

type updateClinic struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateClinic(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateClinicHandler {
	h := &updateClinic{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateClinic) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/clinics", method.Operation{
		Summary:     "Update clinic",
		Description: "Update an existing clinic by ID",
		Tags:        []string{"Clinic"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateClinic) Handler(ctx context.Context, input *UpdateClinicInput) (*UpdateClinicOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Clinic.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find clinic with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("clinic with ID %s not found", id)
	}

	exec := h.client.Clinic.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(clinic.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update clinic with ID %s: %w", id, err)
	}

	return (*UpdateClinicOutput)(types.GenerateOutputResponseMessage("Successfully updated clinic")), nil
}
