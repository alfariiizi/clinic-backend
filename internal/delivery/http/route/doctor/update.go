package doctor_handler

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/doctor"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type UpdateDoctorPayload struct {
	// TODO: Add fields
}

type UpdateDoctorInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateDoctorPayload `json:"body" contentType:"application/json"`
}

type UpdateDoctorOutput types.OutputResponseMessage

type UpdateDoctorHandler model.HTTPHandler[UpdateDoctorInput, UpdateDoctorOutput]

type updateDoctor struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateDoctor(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateDoctorHandler {
	h := &updateDoctor{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateDoctor) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/doctors", method.Operation{
		Summary:     "Update doctor",
		Description: "Update an existing doctor by ID",
		Tags:        []string{"Doctor"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateDoctor) Handler(ctx context.Context, input *UpdateDoctorInput) (*UpdateDoctorOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Doctor.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find doctor with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("doctor with ID %s not found", id)
	}

	exec := h.client.Doctor.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(doctor.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update doctor with ID %s: %w", id, err)
	}

	return (*UpdateDoctorOutput)(types.GenerateOutputResponseMessage("Successfully updated doctor")), nil
}
