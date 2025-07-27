package queueentry_handler

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

type CreateQueueEntryPayload struct {
	// TODO: Add fields
}

type CreateQueueEntryInput struct {
	model.HTTPInputHeaderTenant
	Body CreateQueueEntryPayload `json:"body" contentType:"application/json"`
}

type CreateQueueEntryOutput types.OutputResponseData[CreateQueueEntryData]

type CreateQueueEntryData *db.QueueEntry

type CreateQueueEntryHandler model.HTTPHandler[CreateQueueEntryInput, CreateQueueEntryOutput]

type createQueueEntry struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewCreateQueueEntry(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) CreateQueueEntryHandler {
	h := &createQueueEntry{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *createQueueEntry) RegisterRoutes() {
	api := h.api
	method.POST(api, "/queue-entrys", method.Operation{
		Summary:     "Create QueueEntry",
		Description: "Create a new queueentry",
		Tags:        []string{"QueueEntry"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *createQueueEntry) Handler(ctx context.Context, input *CreateQueueEntryInput) (*CreateQueueEntryOutput, error) {
	// TODO: Add field mapping
	item, err := h.client.QueueEntry.Create().
		// Example: SetField(input.Body.Field).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*CreateQueueEntryOutput)(types.GenerateOutputResponseData(CreateQueueEntryData(item))), nil
}
