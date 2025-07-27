package appointment_handler

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

type ListAppointmentInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListAppointmentOutput types.OutputResponseData[ListAppointmentData]

type ListAppointmentData []*db.Appointment

type ListAppointmentHandler model.HTTPHandler[ListAppointmentInput, ListAppointmentOutput]

type listAppointment struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListAppointment(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListAppointmentHandler {
	h := &listAppointment{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listAppointment) RegisterRoutes() {
	api := h.api
	method.GET(api, "/appointments", method.Operation{
		Summary:     "List Appointments",
		Description: "Retrieve a list of appointments",
		Tags:        []string{"Appointment"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listAppointment) Handler(ctx context.Context, input *ListAppointmentInput) (*ListAppointmentOutput, error) {
	items, err := h.client.Appointment.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListAppointmentOutput)(types.GenerateOutputResponseData(ListAppointmentData(items))), nil
}
