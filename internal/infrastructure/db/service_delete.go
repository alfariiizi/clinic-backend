// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/service"
)

// ServiceDelete is the builder for deleting a Service entity.
type ServiceDelete struct {
	config
	hooks    []Hook
	mutation *ServiceMutation
}

// Where appends a list predicates to the ServiceDelete builder.
func (sd *ServiceDelete) Where(ps ...predicate.Service) *ServiceDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ServiceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ServiceDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ServiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(service.Table, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ServiceDeleteOne is the builder for deleting a single Service entity.
type ServiceDeleteOne struct {
	sd *ServiceDelete
}

// Where appends a list predicates to the ServiceDelete builder.
func (sdo *ServiceDeleteOne) Where(ps ...predicate.Service) *ServiceDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *ServiceDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{service.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ServiceDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
