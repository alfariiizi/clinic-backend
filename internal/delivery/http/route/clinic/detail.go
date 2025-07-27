package clinic_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type DetailClinicInput struct {
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailClinicOutput types.OutputResponseData[DetailClinicData]

type DetailClinicData *db.Clinic

type DetailClinicHandler model.HTTPHandler[DetailClinicInput, DetailClinicOutput]

type detailClinic struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailClinic(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailClinicHandler {
	h := &detailClinic{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailClinic) RegisterRoutes() {
	api := h.api
	method.GET(api, "/clinics/{id}", method.Operation{
		Summary:     "Detail Clinic",
		Description: "Retrieve details of a clinic by ID",
		Tags:        []string{"Clinic"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailClinic) Handler(ctx context.Context, input *DetailClinicInput) (*DetailClinicOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.Clinic.Query().
		Where(clinic.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailClinicOutput)(types.GenerateOutputResponseData(DetailClinicData(item))), nil
}
