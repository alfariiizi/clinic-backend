package appointment_handler

import (
	"context"

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

type DetailAppointmentInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailAppointmentOutput types.OutputResponseData[DetailAppointmentData]

type DetailAppointmentData *db.Appointment

type DetailAppointmentHandler model.HTTPHandler[DetailAppointmentInput, DetailAppointmentOutput]

type detailAppointment struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailAppointment(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailAppointmentHandler {
	h := &detailAppointment{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailAppointment) RegisterRoutes() {
	api := h.api
	method.GET(api, "/appointments/{id}", method.Operation{
		Summary:     "Detail Appointment",
		Description: "Retrieve details of a appointment by ID",
		Tags:        []string{"Appointment"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailAppointment) Handler(ctx context.Context, input *DetailAppointmentInput) (*DetailAppointmentOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Appointment.Query().
		Where(appointment.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailAppointmentOutput)(types.GenerateOutputResponseData(DetailAppointmentData(item))), nil
}
