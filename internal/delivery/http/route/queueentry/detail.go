package queueentry_handler

import (
	"context"

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

type DetailQueueEntryInput struct {
	model.HTTPInputHeaderTenant
	ID string `path:"id" doc:"ID of the item" example:"123"`
}

type DetailQueueEntryOutput types.OutputResponseData[DetailQueueEntryData]

type DetailQueueEntryData *db.QueueEntry

type DetailQueueEntryHandler model.HTTPHandler[DetailQueueEntryInput, DetailQueueEntryOutput]

type detailQueueEntry struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewDetailQueueEntry(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) DetailQueueEntryHandler {
	h := &detailQueueEntry{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *detailQueueEntry) RegisterRoutes() {
	api := h.api
	method.GET(api, "/queue-entrys/{id}", method.Operation{
		Summary:     "Detail QueueEntry",
		Description: "Retrieve details of a queueentry by ID",
		Tags:        []string{"QueueEntry"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *detailQueueEntry) Handler(ctx context.Context, input *DetailQueueEntryInput) (*DetailQueueEntryOutput, error) {
	id, err := utils.IDParser(input.ID)
	if err != nil {
		return nil, err
	}
	item, err := h.client.QueueEntry.Query().
		Where(queueentry.ID(*id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return (*DetailQueueEntryOutput)(types.GenerateOutputResponseData(DetailQueueEntryData(item))), nil
}
