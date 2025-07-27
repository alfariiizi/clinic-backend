package clinic_handler

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

type DeleteClinicInput struct {
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteClinicOutput types.OutputResponseMessage

type DeleteClinicHandler model.HTTPHandler[DeleteClinicInput, DeleteClinicOutput]

type deleteClinic struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteClinic(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteClinicHandler {
	h := &deleteClinic{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteClinic) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/clinics", method.Operation{
		Summary:     "Delete Clinic",
		Description: "Delete a clinic by ID",
		Tags:        []string{"Clinic"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteClinic) Handler(ctx context.Context, input *DeleteClinicInput) (*DeleteClinicOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid clinic ID: %w", err)
	}
	err = h.client.Clinic.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete clinic: %w", err)
	}

	return (*DeleteClinicOutput)(types.GenerateOutputResponseMessage("Successfully deleted clinic")), nil
}
