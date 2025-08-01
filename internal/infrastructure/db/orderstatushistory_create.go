// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/order"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderstatushistory"
	"github.com/google/uuid"
)

// OrderStatusHistoryCreate is the builder for creating a OrderStatusHistory entity.
type OrderStatusHistoryCreate struct {
	config
	mutation *OrderStatusHistoryMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (oshc *OrderStatusHistoryCreate) SetStatus(s string) *OrderStatusHistoryCreate {
	oshc.mutation.SetStatus(s)
	return oshc
}

// SetNotes sets the "notes" field.
func (oshc *OrderStatusHistoryCreate) SetNotes(s string) *OrderStatusHistoryCreate {
	oshc.mutation.SetNotes(s)
	return oshc
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (oshc *OrderStatusHistoryCreate) SetNillableNotes(s *string) *OrderStatusHistoryCreate {
	if s != nil {
		oshc.SetNotes(*s)
	}
	return oshc
}

// SetChangedBy sets the "changed_by" field.
func (oshc *OrderStatusHistoryCreate) SetChangedBy(s string) *OrderStatusHistoryCreate {
	oshc.mutation.SetChangedBy(s)
	return oshc
}

// SetNillableChangedBy sets the "changed_by" field if the given value is not nil.
func (oshc *OrderStatusHistoryCreate) SetNillableChangedBy(s *string) *OrderStatusHistoryCreate {
	if s != nil {
		oshc.SetChangedBy(*s)
	}
	return oshc
}

// SetCreatedAt sets the "created_at" field.
func (oshc *OrderStatusHistoryCreate) SetCreatedAt(t time.Time) *OrderStatusHistoryCreate {
	oshc.mutation.SetCreatedAt(t)
	return oshc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oshc *OrderStatusHistoryCreate) SetNillableCreatedAt(t *time.Time) *OrderStatusHistoryCreate {
	if t != nil {
		oshc.SetCreatedAt(*t)
	}
	return oshc
}

// SetID sets the "id" field.
func (oshc *OrderStatusHistoryCreate) SetID(u uuid.UUID) *OrderStatusHistoryCreate {
	oshc.mutation.SetID(u)
	return oshc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oshc *OrderStatusHistoryCreate) SetNillableID(u *uuid.UUID) *OrderStatusHistoryCreate {
	if u != nil {
		oshc.SetID(*u)
	}
	return oshc
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oshc *OrderStatusHistoryCreate) SetOrderID(id uuid.UUID) *OrderStatusHistoryCreate {
	oshc.mutation.SetOrderID(id)
	return oshc
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (oshc *OrderStatusHistoryCreate) SetNillableOrderID(id *uuid.UUID) *OrderStatusHistoryCreate {
	if id != nil {
		oshc = oshc.SetOrderID(*id)
	}
	return oshc
}

// SetOrder sets the "order" edge to the Order entity.
func (oshc *OrderStatusHistoryCreate) SetOrder(o *Order) *OrderStatusHistoryCreate {
	return oshc.SetOrderID(o.ID)
}

// Mutation returns the OrderStatusHistoryMutation object of the builder.
func (oshc *OrderStatusHistoryCreate) Mutation() *OrderStatusHistoryMutation {
	return oshc.mutation
}

// Save creates the OrderStatusHistory in the database.
func (oshc *OrderStatusHistoryCreate) Save(ctx context.Context) (*OrderStatusHistory, error) {
	oshc.defaults()
	return withHooks(ctx, oshc.sqlSave, oshc.mutation, oshc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oshc *OrderStatusHistoryCreate) SaveX(ctx context.Context) *OrderStatusHistory {
	v, err := oshc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oshc *OrderStatusHistoryCreate) Exec(ctx context.Context) error {
	_, err := oshc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshc *OrderStatusHistoryCreate) ExecX(ctx context.Context) {
	if err := oshc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oshc *OrderStatusHistoryCreate) defaults() {
	if _, ok := oshc.mutation.CreatedAt(); !ok {
		v := orderstatushistory.DefaultCreatedAt()
		oshc.mutation.SetCreatedAt(v)
	}
	if _, ok := oshc.mutation.ID(); !ok {
		v := orderstatushistory.DefaultID()
		oshc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oshc *OrderStatusHistoryCreate) check() error {
	if _, ok := oshc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "OrderStatusHistory.status"`)}
	}
	if _, ok := oshc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "OrderStatusHistory.created_at"`)}
	}
	return nil
}

func (oshc *OrderStatusHistoryCreate) sqlSave(ctx context.Context) (*OrderStatusHistory, error) {
	if err := oshc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oshc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oshc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oshc.mutation.id = &_node.ID
	oshc.mutation.done = true
	return _node, nil
}

func (oshc *OrderStatusHistoryCreate) createSpec() (*OrderStatusHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderStatusHistory{config: oshc.config}
		_spec = sqlgraph.NewCreateSpec(orderstatushistory.Table, sqlgraph.NewFieldSpec(orderstatushistory.FieldID, field.TypeUUID))
	)
	if id, ok := oshc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oshc.mutation.Status(); ok {
		_spec.SetField(orderstatushistory.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := oshc.mutation.Notes(); ok {
		_spec.SetField(orderstatushistory.FieldNotes, field.TypeString, value)
		_node.Notes = value
	}
	if value, ok := oshc.mutation.ChangedBy(); ok {
		_spec.SetField(orderstatushistory.FieldChangedBy, field.TypeString, value)
		_node.ChangedBy = value
	}
	if value, ok := oshc.mutation.CreatedAt(); ok {
		_spec.SetField(orderstatushistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := oshc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderstatushistory.OrderTable,
			Columns: []string{orderstatushistory.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_order_status_history = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderStatusHistoryCreateBulk is the builder for creating many OrderStatusHistory entities in bulk.
type OrderStatusHistoryCreateBulk struct {
	config
	err      error
	builders []*OrderStatusHistoryCreate
}

// Save creates the OrderStatusHistory entities in the database.
func (oshcb *OrderStatusHistoryCreateBulk) Save(ctx context.Context) ([]*OrderStatusHistory, error) {
	if oshcb.err != nil {
		return nil, oshcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oshcb.builders))
	nodes := make([]*OrderStatusHistory, len(oshcb.builders))
	mutators := make([]Mutator, len(oshcb.builders))
	for i := range oshcb.builders {
		func(i int, root context.Context) {
			builder := oshcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderStatusHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oshcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oshcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oshcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oshcb *OrderStatusHistoryCreateBulk) SaveX(ctx context.Context) []*OrderStatusHistory {
	v, err := oshcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oshcb *OrderStatusHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := oshcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshcb *OrderStatusHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := oshcb.Exec(ctx); err != nil {
		panic(err)
	}
}
