package appointment_handler

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

type DeleteAppointmentInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteAppointmentOutput types.OutputResponseMessage

type DeleteAppointmentHandler model.HTTPHandler[DeleteAppointmentInput, DeleteAppointmentOutput]

type deleteAppointment struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteAppointment(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteAppointmentHandler {
	h := &deleteAppointment{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteAppointment) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/appointments", method.Operation{
		Summary:     "Delete Appointment",
		Description: "Delete a appointment by ID",
		Tags:        []string{"Appointment"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteAppointment) Handler(ctx context.Context, input *DeleteAppointmentInput) (*DeleteAppointmentOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid appointment ID: %w", err)
	}
	err = h.client.Appointment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete appointment: %w", err)
	}

	return (*DeleteAppointmentOutput)(types.GenerateOutputResponseMessage("Successfully deleted appointment")), nil
}
