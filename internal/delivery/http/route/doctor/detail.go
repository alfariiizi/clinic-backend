package doctor_handler

import (
	"context"

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

type DetailDoctorInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailDoctorOutput types.OutputResponseData[DetailDoctorData]

type DetailDoctorData *db.Doctor

type DetailDoctorHandler model.HTTPHandler[DetailDoctorInput, DetailDoctorOutput]

type detailDoctor struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailDoctor(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailDoctorHandler {
	h := &detailDoctor{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailDoctor) RegisterRoutes() {
	api := h.api
	method.GET(api, "/doctors/{id}", method.Operation{
		Summary:     "Detail Doctor",
		Description: "Retrieve details of a doctor by ID",
		Tags:        []string{"Doctor"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailDoctor) Handler(ctx context.Context, input *DetailDoctorInput) (*DetailDoctorOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Doctor.Query().
		Where(doctor.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailDoctorOutput)(types.GenerateOutputResponseData(DetailDoctorData(item))), nil
}
