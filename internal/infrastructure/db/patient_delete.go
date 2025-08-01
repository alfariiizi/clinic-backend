// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/patient"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
)

// PatientDelete is the builder for deleting a Patient entity.
type PatientDelete struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// Where appends a list predicates to the PatientDelete builder.
func (pd *PatientDelete) Where(ps ...predicate.Patient) *PatientDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PatientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PatientDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PatientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(patient.Table, sqlgraph.NewFieldSpec(patient.FieldID, field.TypeUUID))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PatientDeleteOne is the builder for deleting a single Patient entity.
type PatientDeleteOne struct {
	pd *PatientDelete
}

// Where appends a list predicates to the PatientDelete builder.
func (pdo *PatientDeleteOne) Where(ps ...predicate.Patient) *PatientDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PatientDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{patient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PatientDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
