// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/appointment"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderitem"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/service"
	"github.com/google/uuid"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *ServiceCreate) SetName(s string) *ServiceCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *ServiceCreate) SetDescription(s string) *ServiceCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescription(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetCategory sets the "category" field.
func (sc *ServiceCreate) SetCategory(s string) *ServiceCreate {
	sc.mutation.SetCategory(s)
	return sc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCategory(s *string) *ServiceCreate {
	if s != nil {
		sc.SetCategory(*s)
	}
	return sc
}

// SetPrice sets the "price" field.
func (sc *ServiceCreate) SetPrice(f float64) *ServiceCreate {
	sc.mutation.SetPrice(f)
	return sc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (sc *ServiceCreate) SetNillablePrice(f *float64) *ServiceCreate {
	if f != nil {
		sc.SetPrice(*f)
	}
	return sc
}

// SetDuration sets the "duration" field.
func (sc *ServiceCreate) SetDuration(i int) *ServiceCreate {
	sc.mutation.SetDuration(i)
	return sc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDuration(i *int) *ServiceCreate {
	if i != nil {
		sc.SetDuration(*i)
	}
	return sc
}

// SetRequirements sets the "requirements" field.
func (sc *ServiceCreate) SetRequirements(s []string) *ServiceCreate {
	sc.mutation.SetRequirements(s)
	return sc
}

// SetRequiresAppointment sets the "requires_appointment" field.
func (sc *ServiceCreate) SetRequiresAppointment(b bool) *ServiceCreate {
	sc.mutation.SetRequiresAppointment(b)
	return sc
}

// SetNillableRequiresAppointment sets the "requires_appointment" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableRequiresAppointment(b *bool) *ServiceCreate {
	if b != nil {
		sc.SetRequiresAppointment(*b)
	}
	return sc
}

// SetActive sets the "active" field.
func (sc *ServiceCreate) SetActive(b bool) *ServiceCreate {
	sc.mutation.SetActive(b)
	return sc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableActive(b *bool) *ServiceCreate {
	if b != nil {
		sc.SetActive(*b)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *ServiceCreate) SetCreatedAt(t time.Time) *ServiceCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCreatedAt(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ServiceCreate) SetUpdatedAt(t time.Time) *ServiceCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableUpdatedAt(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ServiceCreate) SetID(u uuid.UUID) *ServiceCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableID(u *uuid.UUID) *ServiceCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetClinicID sets the "clinic" edge to the Clinic entity by ID.
func (sc *ServiceCreate) SetClinicID(id uuid.UUID) *ServiceCreate {
	sc.mutation.SetClinicID(id)
	return sc
}

// SetNillableClinicID sets the "clinic" edge to the Clinic entity by ID if the given value is not nil.
func (sc *ServiceCreate) SetNillableClinicID(id *uuid.UUID) *ServiceCreate {
	if id != nil {
		sc = sc.SetClinicID(*id)
	}
	return sc
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (sc *ServiceCreate) SetClinic(c *Clinic) *ServiceCreate {
	return sc.SetClinicID(c.ID)
}

// AddAppointmentIDs adds the "appointments" edge to the Appointment entity by IDs.
func (sc *ServiceCreate) AddAppointmentIDs(ids ...uuid.UUID) *ServiceCreate {
	sc.mutation.AddAppointmentIDs(ids...)
	return sc
}

// AddAppointments adds the "appointments" edges to the Appointment entity.
func (sc *ServiceCreate) AddAppointments(a ...*Appointment) *ServiceCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAppointmentIDs(ids...)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (sc *ServiceCreate) AddOrderItemIDs(ids ...uuid.UUID) *ServiceCreate {
	sc.mutation.AddOrderItemIDs(ids...)
	return sc
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (sc *ServiceCreate) AddOrderItems(o ...*OrderItem) *ServiceCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return sc.AddOrderItemIDs(ids...)
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.Price(); !ok {
		v := service.DefaultPrice
		sc.mutation.SetPrice(v)
	}
	if _, ok := sc.mutation.Duration(); !ok {
		v := service.DefaultDuration
		sc.mutation.SetDuration(v)
	}
	if _, ok := sc.mutation.RequiresAppointment(); !ok {
		v := service.DefaultRequiresAppointment
		sc.mutation.SetRequiresAppointment(v)
	}
	if _, ok := sc.mutation.Active(); !ok {
		v := service.DefaultActive
		sc.mutation.SetActive(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := service.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := service.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := service.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "Service.name"`)}
	}
	if _, ok := sc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`db: missing required field "Service.price"`)}
	}
	if _, ok := sc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`db: missing required field "Service.duration"`)}
	}
	if _, ok := sc.mutation.RequiresAppointment(); !ok {
		return &ValidationError{Name: "requires_appointment", err: errors.New(`db: missing required field "Service.requires_appointment"`)}
	}
	if _, ok := sc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`db: missing required field "Service.active"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "Service.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "Service.updated_at"`)}
	}
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(service.Table, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(service.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Category(); ok {
		_spec.SetField(service.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := sc.mutation.Price(); ok {
		_spec.SetField(service.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.SetField(service.FieldDuration, field.TypeInt, value)
		_node.Duration = value
	}
	if value, ok := sc.mutation.Requirements(); ok {
		_spec.SetField(service.FieldRequirements, field.TypeJSON, value)
		_node.Requirements = value
	}
	if value, ok := sc.mutation.RequiresAppointment(); ok {
		_spec.SetField(service.FieldRequiresAppointment, field.TypeBool, value)
		_node.RequiresAppointment = value
	}
	if value, ok := sc.mutation.Active(); ok {
		_spec.SetField(service.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(service.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(service.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.ClinicTable,
			Columns: []string{service.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.clinic_services = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.AppointmentsTable,
			Columns: []string{service.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.OrderItemsTable,
			Columns: []string{service.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	err      error
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
