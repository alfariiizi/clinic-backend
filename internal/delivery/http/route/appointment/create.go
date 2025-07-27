package appointment_handler

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

type CreateAppointmentPayload struct {
	// TODO: Add fields
}

type CreateAppointmentInput struct {
	model.HTTPInputHeaderTenant
	Body CreateAppointmentPayload `json:"body" contentType:"application/json"`
}

type CreateAppointmentOutput types.OutputResponseData[CreateAppointmentData]

type CreateAppointmentData *db.Appointment

type CreateAppointmentHandler model.HTTPHandler[CreateAppointmentInput, CreateAppointmentOutput]

type createAppointment struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateAppointment(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateAppointmentHandler {
	h := &createAppointment{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createAppointment) RegisterRoutes() {
	api := h.api
	method.POST(api, "/appointments", method.Operation{
		Summary:     "Create Appointment",
		Description: "Create a new appointment",
		Tags:        []string{"Appointment"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createAppointment) Handler(ctx context.Context, input *CreateAppointmentInput) (*CreateAppointmentOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Appointment.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateAppointmentOutput)(types.GenerateOutputResponseData(CreateAppointmentData(item))), nil
}
