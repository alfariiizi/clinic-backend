package queueentry_handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/alfariiizi/vandor/internal/core/model"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/core/service"
	"github.com/alfariiizi/vandor/internal/delivery/http/api"
	"github.com/alfariiizi/vandor/internal/delivery/http/method"
	"github.com/alfariiizi/vandor/internal/types"
)

type ListQueueEntryInput struct {
	model.HTTPInputHeaderTenant
	Query string `query:"q" doc:"Query parameter for filtering" example:"search term"`
}

type ListQueueEntryOutput types.OutputResponseData[ListQueueEntryData]

type ListQueueEntryData []*db.QueueEntry

type ListQueueEntryHandler model.HTTPHandler[ListQueueEntryInput, ListQueueEntryOutput]

type listQueueEntry struct {
	api     huma.API
	service *service.Services
	client  *db.Client
}

func NewListQueueEntry(
	api *api.HttpApi,
	service *service.Services,
	client *db.Client,
) ListQueueEntryHandler {
	h := &listQueueEntry{
		api:     api.BaseAPI,
		service: service,
		client:  client,
	}
	h.RegisterRoutes()
	return h
}

func (h *listQueueEntry) RegisterRoutes() {
	api := h.api
	method.GET(api, "/queue-entrys", method.Operation{
		Summary:     "List QueueEntrys",
		Description: "Retrieve a list of queueentrys",
		Tags:        []string{"QueueEntry"},
		BearerAuth:  true,
	}, h.Handler)
}

func (h *listQueueEntry) Handler(ctx context.Context, input *ListQueueEntryInput) (*ListQueueEntryOutput, error) {
	items, err := h.client.QueueEntry.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return (*ListQueueEntryOutput)(types.GenerateOutputResponseData(ListQueueEntryData(items))), nil
}
