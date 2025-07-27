package appointment_handler

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/appointment"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type UpdateAppointmentPayload struct {
	// TODO: Add fields
}

type UpdateAppointmentInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateAppointmentPayload `json:"body" contentType:"application/json"`
}

type UpdateAppointmentOutput types.OutputResponseMessage

type UpdateAppointmentHandler model.HTTPHandler[UpdateAppointmentInput, UpdateAppointmentOutput]

type updateAppointment struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateAppointment(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateAppointmentHandler {
	h := &updateAppointment{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateAppointment) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/appointments", method.Operation{
		Summary:     "Update appointment",
		Description: "Update an existing appointment by ID",
		Tags:        []string{"Appointment"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateAppointment) Handler(ctx context.Context, input *UpdateAppointmentInput) (*UpdateAppointmentOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.Appointment.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find appointment with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("appointment with ID %s not found", id)
	}

	exec := h.client.Appointment.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(appointment.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update appointment with ID %s: %w", id, err)
	}

	return (*UpdateAppointmentOutput)(types.GenerateOutputResponseMessage("Successfully updated appointment")), nil
}
