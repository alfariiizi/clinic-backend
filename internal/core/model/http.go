// Package model provides interfaces for HTTP handlers in the service.
package model

import (
	"context"

	"github.com/alfariiizi/vandor/internal/types"
)

type HTTPHandler[I, O any] interface {
	RegisterRoutes()
	Handler(context.Context, *I) (*O, error)
}

type HTTPHandlerData[I, O any] interface {
	RegisterRoutes()
	Handler(context.Context, *I) (*types.OutputResponseData[O], error)
}

type HTTPHandlerMessage[I any] interface {
	RegisterRoutes()
	Handler(context.Context, *I) (*types.OutputResponseMessage, error)
}

type HTTPInputHeaderTenant struct {
	TenantID string `header:"X-Tenant-ID" required:"true" doc:"The tenant ID for scoping the request"`
}
