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
	"github.com/alfariiizi/vandor/internal/infrastructure/db/order"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderitem"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/product"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/service"
	"github.com/google/uuid"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// SetItemType sets the "item_type" field.
func (oic *OrderItemCreate) SetItemType(ot orderitem.ItemType) *OrderItemCreate {
	oic.mutation.SetItemType(ot)
	return oic
}

// SetNillableItemType sets the "item_type" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableItemType(ot *orderitem.ItemType) *OrderItemCreate {
	if ot != nil {
		oic.SetItemType(*ot)
	}
	return oic
}

// SetItemName sets the "item_name" field.
func (oic *OrderItemCreate) SetItemName(s string) *OrderItemCreate {
	oic.mutation.SetItemName(s)
	return oic
}

// SetItemDescription sets the "item_description" field.
func (oic *OrderItemCreate) SetItemDescription(s string) *OrderItemCreate {
	oic.mutation.SetItemDescription(s)
	return oic
}

// SetNillableItemDescription sets the "item_description" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableItemDescription(s *string) *OrderItemCreate {
	if s != nil {
		oic.SetItemDescription(*s)
	}
	return oic
}

// SetQuantity sets the "quantity" field.
func (oic *OrderItemCreate) SetQuantity(i int) *OrderItemCreate {
	oic.mutation.SetQuantity(i)
	return oic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableQuantity(i *int) *OrderItemCreate {
	if i != nil {
		oic.SetQuantity(*i)
	}
	return oic
}

// SetUnitPrice sets the "unit_price" field.
func (oic *OrderItemCreate) SetUnitPrice(f float64) *OrderItemCreate {
	oic.mutation.SetUnitPrice(f)
	return oic
}

// SetDiscountAmount sets the "discount_amount" field.
func (oic *OrderItemCreate) SetDiscountAmount(f float64) *OrderItemCreate {
	oic.mutation.SetDiscountAmount(f)
	return oic
}

// SetNillableDiscountAmount sets the "discount_amount" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableDiscountAmount(f *float64) *OrderItemCreate {
	if f != nil {
		oic.SetDiscountAmount(*f)
	}
	return oic
}

// SetTotalPrice sets the "total_price" field.
func (oic *OrderItemCreate) SetTotalPrice(f float64) *OrderItemCreate {
	oic.mutation.SetTotalPrice(f)
	return oic
}

// SetItemMetadata sets the "item_metadata" field.
func (oic *OrderItemCreate) SetItemMetadata(m map[string]interface{}) *OrderItemCreate {
	oic.mutation.SetItemMetadata(m)
	return oic
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrderItemCreate) SetCreatedAt(t time.Time) *OrderItemCreate {
	oic.mutation.SetCreatedAt(t)
	return oic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableCreatedAt(t *time.Time) *OrderItemCreate {
	if t != nil {
		oic.SetCreatedAt(*t)
	}
	return oic
}

// SetID sets the "id" field.
func (oic *OrderItemCreate) SetID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetID(u)
	return oic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableID(u *uuid.UUID) *OrderItemCreate {
	if u != nil {
		oic.SetID(*u)
	}
	return oic
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oic *OrderItemCreate) SetOrderID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetOrderID(id)
	return oic
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableOrderID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetOrderID(*id)
	}
	return oic
}

// SetOrder sets the "order" edge to the Order entity.
func (oic *OrderItemCreate) SetOrder(o *Order) *OrderItemCreate {
	return oic.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (oic *OrderItemCreate) SetProductID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetProductID(id)
	return oic
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableProductID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetProductID(*id)
	}
	return oic
}

// SetProduct sets the "product" edge to the Product entity.
func (oic *OrderItemCreate) SetProduct(p *Product) *OrderItemCreate {
	return oic.SetProductID(p.ID)
}

// SetServiceID sets the "service" edge to the Service entity by ID.
func (oic *OrderItemCreate) SetServiceID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetServiceID(id)
	return oic
}

// SetNillableServiceID sets the "service" edge to the Service entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableServiceID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetServiceID(*id)
	}
	return oic
}

// SetService sets the "service" edge to the Service entity.
func (oic *OrderItemCreate) SetService(s *Service) *OrderItemCreate {
	return oic.SetServiceID(s.ID)
}

// SetAppointmentID sets the "appointment" edge to the Appointment entity by ID.
func (oic *OrderItemCreate) SetAppointmentID(id uuid.UUID) *OrderItemCreate {
	oic.mutation.SetAppointmentID(id)
	return oic
}

// SetNillableAppointmentID sets the "appointment" edge to the Appointment entity by ID if the given value is not nil.
func (oic *OrderItemCreate) SetNillableAppointmentID(id *uuid.UUID) *OrderItemCreate {
	if id != nil {
		oic = oic.SetAppointmentID(*id)
	}
	return oic
}

// SetAppointment sets the "appointment" edge to the Appointment entity.
func (oic *OrderItemCreate) SetAppointment(a *Appointment) *OrderItemCreate {
	return oic.SetAppointmentID(a.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() {
	if _, ok := oic.mutation.ItemType(); !ok {
		v := orderitem.DefaultItemType
		oic.mutation.SetItemType(v)
	}
	if _, ok := oic.mutation.Quantity(); !ok {
		v := orderitem.DefaultQuantity
		oic.mutation.SetQuantity(v)
	}
	if _, ok := oic.mutation.DiscountAmount(); !ok {
		v := orderitem.DefaultDiscountAmount
		oic.mutation.SetDiscountAmount(v)
	}
	if _, ok := oic.mutation.CreatedAt(); !ok {
		v := orderitem.DefaultCreatedAt()
		oic.mutation.SetCreatedAt(v)
	}
	if _, ok := oic.mutation.ID(); !ok {
		v := orderitem.DefaultID()
		oic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	if _, ok := oic.mutation.ItemType(); !ok {
		return &ValidationError{Name: "item_type", err: errors.New(`db: missing required field "OrderItem.item_type"`)}
	}
	if v, ok := oic.mutation.ItemType(); ok {
		if err := orderitem.ItemTypeValidator(v); err != nil {
			return &ValidationError{Name: "item_type", err: fmt.Errorf(`db: validator failed for field "OrderItem.item_type": %w`, err)}
		}
	}
	if _, ok := oic.mutation.ItemName(); !ok {
		return &ValidationError{Name: "item_name", err: errors.New(`db: missing required field "OrderItem.item_name"`)}
	}
	if _, ok := oic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`db: missing required field "OrderItem.quantity"`)}
	}
	if _, ok := oic.mutation.UnitPrice(); !ok {
		return &ValidationError{Name: "unit_price", err: errors.New(`db: missing required field "OrderItem.unit_price"`)}
	}
	if _, ok := oic.mutation.DiscountAmount(); !ok {
		return &ValidationError{Name: "discount_amount", err: errors.New(`db: missing required field "OrderItem.discount_amount"`)}
	}
	if _, ok := oic.mutation.TotalPrice(); !ok {
		return &ValidationError{Name: "total_price", err: errors.New(`db: missing required field "OrderItem.total_price"`)}
	}
	if _, ok := oic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "OrderItem.created_at"`)}
	}
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
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
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderitem.Table, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID))
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oic.mutation.ItemType(); ok {
		_spec.SetField(orderitem.FieldItemType, field.TypeEnum, value)
		_node.ItemType = value
	}
	if value, ok := oic.mutation.ItemName(); ok {
		_spec.SetField(orderitem.FieldItemName, field.TypeString, value)
		_node.ItemName = value
	}
	if value, ok := oic.mutation.ItemDescription(); ok {
		_spec.SetField(orderitem.FieldItemDescription, field.TypeString, value)
		_node.ItemDescription = value
	}
	if value, ok := oic.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := oic.mutation.UnitPrice(); ok {
		_spec.SetField(orderitem.FieldUnitPrice, field.TypeFloat64, value)
		_node.UnitPrice = value
	}
	if value, ok := oic.mutation.DiscountAmount(); ok {
		_spec.SetField(orderitem.FieldDiscountAmount, field.TypeFloat64, value)
		_node.DiscountAmount = value
	}
	if value, ok := oic.mutation.TotalPrice(); ok {
		_spec.SetField(orderitem.FieldTotalPrice, field.TypeFloat64, value)
		_node.TotalPrice = value
	}
	if value, ok := oic.mutation.ItemMetadata(); ok {
		_spec.SetField(orderitem.FieldItemMetadata, field.TypeJSON, value)
		_node.ItemMetadata = value
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.SetField(orderitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := oic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_order_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.ProductTable,
			Columns: []string{orderitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_order_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.ServiceTable,
			Columns: []string{orderitem.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.service_order_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.AppointmentTable,
			Columns: []string{orderitem.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.appointment_order_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	err      error
	builders []*OrderItemCreate
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
