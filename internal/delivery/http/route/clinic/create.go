package clinic_handler

import (
	"context"
	"time"

	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/danielgtaylor/huma/v2"
)

type CreateClinicPayload struct {
	// TODO: Add fields
}

type CreateClinicInput struct {
	Body CreateClinicPayload `json:"body" contentType:"application/json"`
}

type CreateClinicOutput types.OutputResponseData[CreateClinicData]

type CreateClinicData *db.Clinic

type CreateClinicHandler model.HTTPHandler[CreateClinicInput, CreateClinicOutput]

type createClinic struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateClinic(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateClinicHandler {
	h := &createClinic{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createClinic) RegisterRoutes() {
	api := h.api
	method.POST(api, "/clinics", method.Operation{
		Summary:     "Create Clinic",
		Description: "Create a new clinic",
		Tags:        []string{"Clinic"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createClinic) Handler(ctx context.Context, input *CreateClinicInput) (*CreateClinicOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.Clinic.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateClinicOutput)(types.GenerateOutputResponseData(CreateClinicData(item))), nil
}
