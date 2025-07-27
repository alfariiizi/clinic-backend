package doctor_handler

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

type CreateDoctorPayload struct {
	// TODO: Add fields
}

type CreateDoctorInput struct {
	model.HTTPInputHeaderTenant
	Body CreateDoctorPayload `json:"body" contentType:"application/json"`
}

type CreateDoctorOutput types.OutputResponseData[CreateDoctorData]

type CreateDoctorData *db.Doctor

type CreateDoctorHandler model.HTTPHandler[CreateDoctorInput, CreateDoctorOutput]

type createDoctor struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateDoctor(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateDoctorHandler {
	h := &createDoctor{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createDoctor) RegisterRoutes() {
	api := h.api
	method.POST(api, "/doctors", method.Operation{
		Summary:     "Create Doctor",
		Description: "Create a new doctor",
		Tags:        []string{"Doctor"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createDoctor) Handler(ctx context.Context, input *CreateDoctorInput) (*CreateDoctorOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Doctor.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateDoctorOutput)(types.GenerateOutputResponseData(CreateDoctorData(item))), nil
}
