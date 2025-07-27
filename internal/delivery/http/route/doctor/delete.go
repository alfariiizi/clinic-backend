package doctor_handler

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

type DeleteDoctorInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteDoctorOutput types.OutputResponseMessage

type DeleteDoctorHandler model.HTTPHandler[DeleteDoctorInput, DeleteDoctorOutput]

type deleteDoctor struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteDoctor(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteDoctorHandler {
	h := &deleteDoctor{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteDoctor) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/doctors", method.Operation{
		Summary:     "Delete Doctor",
		Description: "Delete a doctor by ID",
		Tags:        []string{"Doctor"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteDoctor) Handler(ctx context.Context, input *DeleteDoctorInput) (*DeleteDoctorOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid doctor ID: %w", err)
	}
	err = h.client.Doctor.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete doctor: %w", err)
	}

	return (*DeleteDoctorOutput)(types.GenerateOutputResponseMessage("Successfully deleted doctor")), nil
}
