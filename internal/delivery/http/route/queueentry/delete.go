package queueentry_handler

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

type DeleteQueueEntryInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item to delete" example:"123"`
}

type DeleteQueueEntryOutput types.OutputResponseMessage

type DeleteQueueEntryHandler model.HTTPHandler[DeleteQueueEntryInput, DeleteQueueEntryOutput]

type deleteQueueEntry struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDeleteQueueEntry(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DeleteQueueEntryHandler {
	h := &deleteQueueEntry{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *deleteQueueEntry) RegisterRoutes() {
	api := h.api
	method.DELETE(api, "/queue-entrys", method.Operation{
		Summary:     "Delete QueueEntry",
		Description: "Delete a queueentry by ID",
		Tags:        []string{"QueueEntry"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *deleteQueueEntry) Handler(ctx context.Context, input *DeleteQueueEntryInput) (*DeleteQueueEntryOutput, error) {
	id, err := utils.GetData(utils.IDParser(input.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid queueentry ID: %w", err)
	}
	err = h.client.QueueEntry.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete queueentry: %w", err)
	}

	return (*DeleteQueueEntryOutput)(types.GenerateOutputResponseMessage("Successfully deleted queueentry")), nil
}
