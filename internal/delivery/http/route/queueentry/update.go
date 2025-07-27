package queueentry_handler

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/queueentry"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
	"github.com/alfariiizi/vandor/internal/utils"
)

type UpdateQueueEntryPayload struct {
	// TODO: Add fields
}

type UpdateQueueEntryInput struct {
	model.HTTPInputHeaderTenant
	ID   string               `path:"id" doc:"ID of the item to update" example:"123"`
	Body UpdateQueueEntryPayload `json:"body" contentType:"application/json"`
}

type UpdateQueueEntryOutput types.OutputResponseMessage

type UpdateQueueEntryHandler model.HTTPHandler[UpdateQueueEntryInput, UpdateQueueEntryOutput]

type updateQueueEntry struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewUpdateQueueEntry(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) UpdateQueueEntryHandler {
	h := &updateQueueEntry{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *updateQueueEntry) RegisterRoutes() {
	api := h.api
	method.PATCH(api, "/queue-entrys", method.Operation{
		Summary:     "Update queueentry",
		Description: "Update an existing queueentry by ID",
		Tags:        []string{"QueueEntry"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *updateQueueEntry) Handler(ctx context.Context, input *UpdateQueueEntryInput) (*UpdateQueueEntryOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	// payload := input.Body

	existing, err := h.client.QueueEntry.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find queueentry with ID %s: %w", id, err)
	}
	if existing == nil {
		return nil, fmt.Errorf("queueentry with ID %s not found", id)
	}

	exec := h.client.QueueEntry.Update()
	// TODO: Set fields like:
	// if payload.Field != nil {
	//     exec.SetField(*payload.Field)
	// }

	_, err = exec.Where(queueentry.ID(id)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update queueentry with ID %s: %w", id, err)
	}

	return (*UpdateQueueEntryOutput)(types.GenerateOutputResponseMessage("Successfully updated queueentry")), nil
}
