// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/appointment"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/order"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderitem"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/product"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/service"
	"github.com/google/uuid"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetItemType sets the "item_type" field.
func (oiu *OrderItemUpdate) SetItemType(ot orderitem.ItemType) *OrderItemUpdate {
	oiu.mutation.SetItemType(ot)
	return oiu
}

// SetNillableItemType sets the "item_type" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableItemType(ot *orderitem.ItemType) *OrderItemUpdate {
	if ot != nil {
		oiu.SetItemType(*ot)
	}
	return oiu
}

// SetItemName sets the "item_name" field.
func (oiu *OrderItemUpdate) SetItemName(s string) *OrderItemUpdate {
	oiu.mutation.SetItemName(s)
	return oiu
}

// SetNillableItemName sets the "item_name" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableItemName(s *string) *OrderItemUpdate {
	if s != nil {
		oiu.SetItemName(*s)
	}
	return oiu
}

// SetItemDescription sets the "item_description" field.
func (oiu *OrderItemUpdate) SetItemDescription(s string) *OrderItemUpdate {
	oiu.mutation.SetItemDescription(s)
	return oiu
}

// SetNillableItemDescription sets the "item_description" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableItemDescription(s *string) *OrderItemUpdate {
	if s != nil {
		oiu.SetItemDescription(*s)
	}
	return oiu
}

// ClearItemDescription clears the value of the "item_description" field.
func (oiu *OrderItemUpdate) ClearItemDescription() *OrderItemUpdate {
	oiu.mutation.ClearItemDescription()
	return oiu
}

// SetQuantity sets the "quantity" field.
func (oiu *OrderItemUpdate) SetQuantity(i int) *OrderItemUpdate {
	oiu.mutation.ResetQuantity()
	oiu.mutation.SetQuantity(i)
	return oiu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableQuantity(i *int) *OrderItemUpdate {
	if i != nil {
		oiu.SetQuantity(*i)
	}
	return oiu
}

// AddQuantity adds i to the "quantity" field.
func (oiu *OrderItemUpdate) AddQuantity(i int) *OrderItemUpdate {
	oiu.mutation.AddQuantity(i)
	return oiu
}

// SetUnitPrice sets the "unit_price" field.
func (oiu *OrderItemUpdate) SetUnitPrice(f float64) *OrderItemUpdate {
	oiu.mutation.ResetUnitPrice()
	oiu.mutation.SetUnitPrice(f)
	return oiu
}

// SetNillableUnitPrice sets the "unit_price" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableUnitPrice(f *float64) *OrderItemUpdate {
	if f != nil {
		oiu.SetUnitPrice(*f)
	}
	return oiu
}

// AddUnitPrice adds f to the "unit_price" field.
func (oiu *OrderItemUpdate) AddUnitPrice(f float64) *OrderItemUpdate {
	oiu.mutation.AddUnitPrice(f)
	return oiu
}

// SetDiscountAmount sets the "discount_amount" field.
func (oiu *OrderItemUpdate) SetDiscountAmount(f float64) *OrderItemUpdate {
	oiu.mutation.ResetDiscountAmount()
	oiu.mutation.SetDiscountAmount(f)
	return oiu
}

// SetNillableDiscountAmount sets the "discount_amount" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableDiscountAmount(f *float64) *OrderItemUpdate {
	if f != nil {
		oiu.SetDiscountAmount(*f)
	}
	return oiu
}

// AddDiscountAmount adds f to the "discount_amount" field.
func (oiu *OrderItemUpdate) AddDiscountAmount(f float64) *OrderItemUpdate {
	oiu.mutation.AddDiscountAmount(f)
	return oiu
}

// SetTotalPrice sets the "total_price" field.
func (oiu *OrderItemUpdate) SetTotalPrice(f float64) *OrderItemUpdate {
	oiu.mutation.ResetTotalPrice()
	oiu.mutation.SetTotalPrice(f)
	return oiu
}

// SetNillableTotalPrice sets the "total_price" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableTotalPrice(f *float64) *OrderItemUpdate {
	if f != nil {
		oiu.SetTotalPrice(*f)
	}
	return oiu
}

// AddTotalPrice adds f to the "total_price" field.
func (oiu *OrderItemUpdate) AddTotalPrice(f float64) *OrderItemUpdate {
	oiu.mutation.AddTotalPrice(f)
	return oiu
}

// SetItemMetadata sets the "item_metadata" field.
func (oiu *OrderItemUpdate) SetItemMetadata(m map[string]interface{}) *OrderItemUpdate {
	oiu.mutation.SetItemMetadata(m)
	return oiu
}

// ClearItemMetadata clears the value of the "item_metadata" field.
func (oiu *OrderItemUpdate) ClearItemMetadata() *OrderItemUpdate {
	oiu.mutation.ClearItemMetadata()
	return oiu
}

// SetCreatedAt sets the "created_at" field.
func (oiu *OrderItemUpdate) SetCreatedAt(t time.Time) *OrderItemUpdate {
	oiu.mutation.SetCreatedAt(t)
	return oiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableCreatedAt(t *time.Time) *OrderItemUpdate {
	if t != nil {
		oiu.SetCreatedAt(*t)
	}
	return oiu
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oiu *OrderItemUpdate) SetOrderID(id uuid.UUID) *OrderItemUpdate {
	oiu.mutation.SetOrderID(id)
	return oiu
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableOrderID(id *uuid.UUID) *OrderItemUpdate {
	if id != nil {
		oiu = oiu.SetOrderID(*id)
	}
	return oiu
}

// SetOrder sets the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) SetOrder(o *Order) *OrderItemUpdate {
	return oiu.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (oiu *OrderItemUpdate) SetProductID(id uuid.UUID) *OrderItemUpdate {
	oiu.mutation.SetProductID(id)
	return oiu
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableProductID(id *uuid.UUID) *OrderItemUpdate {
	if id != nil {
		oiu = oiu.SetProductID(*id)
	}
	return oiu
}

// SetProduct sets the "product" edge to the Product entity.
func (oiu *OrderItemUpdate) SetProduct(p *Product) *OrderItemUpdate {
	return oiu.SetProductID(p.ID)
}

// SetServiceID sets the "service" edge to the Service entity by ID.
func (oiu *OrderItemUpdate) SetServiceID(id uuid.UUID) *OrderItemUpdate {
	oiu.mutation.SetServiceID(id)
	return oiu
}

// SetNillableServiceID sets the "service" edge to the Service entity by ID if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableServiceID(id *uuid.UUID) *OrderItemUpdate {
	if id != nil {
		oiu = oiu.SetServiceID(*id)
	}
	return oiu
}

// SetService sets the "service" edge to the Service entity.
func (oiu *OrderItemUpdate) SetService(s *Service) *OrderItemUpdate {
	return oiu.SetServiceID(s.ID)
}

// SetAppointmentID sets the "appointment" edge to the Appointment entity by ID.
func (oiu *OrderItemUpdate) SetAppointmentID(id uuid.UUID) *OrderItemUpdate {
	oiu.mutation.SetAppointmentID(id)
	return oiu
}

// SetNillableAppointmentID sets the "appointment" edge to the Appointment entity by ID if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableAppointmentID(id *uuid.UUID) *OrderItemUpdate {
	if id != nil {
		oiu = oiu.SetAppointmentID(*id)
	}
	return oiu
}

// SetAppointment sets the "appointment" edge to the Appointment entity.
func (oiu *OrderItemUpdate) SetAppointment(a *Appointment) *OrderItemUpdate {
	return oiu.SetAppointmentID(a.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) ClearOrder() *OrderItemUpdate {
	oiu.mutation.ClearOrder()
	return oiu
}

// ClearProduct clears the "product" edge to the Product entity.
func (oiu *OrderItemUpdate) ClearProduct() *OrderItemUpdate {
	oiu.mutation.ClearProduct()
	return oiu
}

// ClearService clears the "service" edge to the Service entity.
func (oiu *OrderItemUpdate) ClearService() *OrderItemUpdate {
	oiu.mutation.ClearService()
	return oiu
}

// ClearAppointment clears the "appointment" edge to the Appointment entity.
func (oiu *OrderItemUpdate) ClearAppointment() *OrderItemUpdate {
	oiu.mutation.ClearAppointment()
	return oiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiu *OrderItemUpdate) check() error {
	if v, ok := oiu.mutation.ItemType(); ok {
		if err := orderitem.ItemTypeValidator(v); err != nil {
			return &ValidationError{Name: "item_type", err: fmt.Errorf(`db: validator failed for field "OrderItem.item_type": %w`, err)}
		}
	}
	return nil
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.ItemType(); ok {
		_spec.SetField(orderitem.FieldItemType, field.TypeEnum, value)
	}
	if value, ok := oiu.mutation.ItemName(); ok {
		_spec.SetField(orderitem.FieldItemName, field.TypeString, value)
	}
	if value, ok := oiu.mutation.ItemDescription(); ok {
		_spec.SetField(orderitem.FieldItemDescription, field.TypeString, value)
	}
	if oiu.mutation.ItemDescriptionCleared() {
		_spec.ClearField(orderitem.FieldItemDescription, field.TypeString)
	}
	if value, ok := oiu.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.UnitPrice(); ok {
		_spec.SetField(orderitem.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.AddedUnitPrice(); ok {
		_spec.AddField(orderitem.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.DiscountAmount(); ok {
		_spec.SetField(orderitem.FieldDiscountAmount, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.AddedDiscountAmount(); ok {
		_spec.AddField(orderitem.FieldDiscountAmount, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.TotalPrice(); ok {
		_spec.SetField(orderitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.AddedTotalPrice(); ok {
		_spec.AddField(orderitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.ItemMetadata(); ok {
		_spec.SetField(orderitem.FieldItemMetadata, field.TypeJSON, value)
	}
	if oiu.mutation.ItemMetadataCleared() {
		_spec.ClearField(orderitem.FieldItemMetadata, field.TypeJSON)
	}
	if value, ok := oiu.mutation.CreatedAt(); ok {
		_spec.SetField(orderitem.FieldCreatedAt, field.TypeTime, value)
	}
	if oiu.mutation.OrderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.OrderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiu.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiu.mutation.ServiceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.ServiceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiu.mutation.AppointmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.AppointmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderItemMutation
}

// SetItemType sets the "item_type" field.
func (oiuo *OrderItemUpdateOne) SetItemType(ot orderitem.ItemType) *OrderItemUpdateOne {
	oiuo.mutation.SetItemType(ot)
	return oiuo
}

// SetNillableItemType sets the "item_type" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableItemType(ot *orderitem.ItemType) *OrderItemUpdateOne {
	if ot != nil {
		oiuo.SetItemType(*ot)
	}
	return oiuo
}

// SetItemName sets the "item_name" field.
func (oiuo *OrderItemUpdateOne) SetItemName(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetItemName(s)
	return oiuo
}

// SetNillableItemName sets the "item_name" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableItemName(s *string) *OrderItemUpdateOne {
	if s != nil {
		oiuo.SetItemName(*s)
	}
	return oiuo
}

// SetItemDescription sets the "item_description" field.
func (oiuo *OrderItemUpdateOne) SetItemDescription(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetItemDescription(s)
	return oiuo
}

// SetNillableItemDescription sets the "item_description" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableItemDescription(s *string) *OrderItemUpdateOne {
	if s != nil {
		oiuo.SetItemDescription(*s)
	}
	return oiuo
}

// ClearItemDescription clears the value of the "item_description" field.
func (oiuo *OrderItemUpdateOne) ClearItemDescription() *OrderItemUpdateOne {
	oiuo.mutation.ClearItemDescription()
	return oiuo
}

// SetQuantity sets the "quantity" field.
func (oiuo *OrderItemUpdateOne) SetQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.ResetQuantity()
	oiuo.mutation.SetQuantity(i)
	return oiuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableQuantity(i *int) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetQuantity(*i)
	}
	return oiuo
}

// AddQuantity adds i to the "quantity" field.
func (oiuo *OrderItemUpdateOne) AddQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.AddQuantity(i)
	return oiuo
}

// SetUnitPrice sets the "unit_price" field.
func (oiuo *OrderItemUpdateOne) SetUnitPrice(f float64) *OrderItemUpdateOne {
	oiuo.mutation.ResetUnitPrice()
	oiuo.mutation.SetUnitPrice(f)
	return oiuo
}

// SetNillableUnitPrice sets the "unit_price" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableUnitPrice(f *float64) *OrderItemUpdateOne {
	if f != nil {
		oiuo.SetUnitPrice(*f)
	}
	return oiuo
}

// AddUnitPrice adds f to the "unit_price" field.
func (oiuo *OrderItemUpdateOne) AddUnitPrice(f float64) *OrderItemUpdateOne {
	oiuo.mutation.AddUnitPrice(f)
	return oiuo
}

// SetDiscountAmount sets the "discount_amount" field.
func (oiuo *OrderItemUpdateOne) SetDiscountAmount(f float64) *OrderItemUpdateOne {
	oiuo.mutation.ResetDiscountAmount()
	oiuo.mutation.SetDiscountAmount(f)
	return oiuo
}

// SetNillableDiscountAmount sets the "discount_amount" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableDiscountAmount(f *float64) *OrderItemUpdateOne {
	if f != nil {
		oiuo.SetDiscountAmount(*f)
	}
	return oiuo
}

// AddDiscountAmount adds f to the "discount_amount" field.
func (oiuo *OrderItemUpdateOne) AddDiscountAmount(f float64) *OrderItemUpdateOne {
	oiuo.mutation.AddDiscountAmount(f)
	return oiuo
}

// SetTotalPrice sets the "total_price" field.
func (oiuo *OrderItemUpdateOne) SetTotalPrice(f float64) *OrderItemUpdateOne {
	oiuo.mutation.ResetTotalPrice()
	oiuo.mutation.SetTotalPrice(f)
	return oiuo
}

// SetNillableTotalPrice sets the "total_price" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableTotalPrice(f *float64) *OrderItemUpdateOne {
	if f != nil {
		oiuo.SetTotalPrice(*f)
	}
	return oiuo
}

// AddTotalPrice adds f to the "total_price" field.
func (oiuo *OrderItemUpdateOne) AddTotalPrice(f float64) *OrderItemUpdateOne {
	oiuo.mutation.AddTotalPrice(f)
	return oiuo
}

// SetItemMetadata sets the "item_metadata" field.
func (oiuo *OrderItemUpdateOne) SetItemMetadata(m map[string]interface{}) *OrderItemUpdateOne {
	oiuo.mutation.SetItemMetadata(m)
	return oiuo
}

// ClearItemMetadata clears the value of the "item_metadata" field.
func (oiuo *OrderItemUpdateOne) ClearItemMetadata() *OrderItemUpdateOne {
	oiuo.mutation.ClearItemMetadata()
	return oiuo
}

// SetCreatedAt sets the "created_at" field.
func (oiuo *OrderItemUpdateOne) SetCreatedAt(t time.Time) *OrderItemUpdateOne {
	oiuo.mutation.SetCreatedAt(t)
	return oiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableCreatedAt(t *time.Time) *OrderItemUpdateOne {
	if t != nil {
		oiuo.SetCreatedAt(*t)
	}
	return oiuo
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oiuo *OrderItemUpdateOne) SetOrderID(id uuid.UUID) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderID(id)
	return oiuo
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableOrderID(id *uuid.UUID) *OrderItemUpdateOne {
	if id != nil {
		oiuo = oiuo.SetOrderID(*id)
	}
	return oiuo
}

// SetOrder sets the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) SetOrder(o *Order) *OrderItemUpdateOne {
	return oiuo.SetOrderID(o.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (oiuo *OrderItemUpdateOne) SetProductID(id uuid.UUID) *OrderItemUpdateOne {
	oiuo.mutation.SetProductID(id)
	return oiuo
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableProductID(id *uuid.UUID) *OrderItemUpdateOne {
	if id != nil {
		oiuo = oiuo.SetProductID(*id)
	}
	return oiuo
}

// SetProduct sets the "product" edge to the Product entity.
func (oiuo *OrderItemUpdateOne) SetProduct(p *Product) *OrderItemUpdateOne {
	return oiuo.SetProductID(p.ID)
}

// SetServiceID sets the "service" edge to the Service entity by ID.
func (oiuo *OrderItemUpdateOne) SetServiceID(id uuid.UUID) *OrderItemUpdateOne {
	oiuo.mutation.SetServiceID(id)
	return oiuo
}

// SetNillableServiceID sets the "service" edge to the Service entity by ID if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableServiceID(id *uuid.UUID) *OrderItemUpdateOne {
	if id != nil {
		oiuo = oiuo.SetServiceID(*id)
	}
	return oiuo
}

// SetService sets the "service" edge to the Service entity.
func (oiuo *OrderItemUpdateOne) SetService(s *Service) *OrderItemUpdateOne {
	return oiuo.SetServiceID(s.ID)
}

// SetAppointmentID sets the "appointment" edge to the Appointment entity by ID.
func (oiuo *OrderItemUpdateOne) SetAppointmentID(id uuid.UUID) *OrderItemUpdateOne {
	oiuo.mutation.SetAppointmentID(id)
	return oiuo
}

// SetNillableAppointmentID sets the "appointment" edge to the Appointment entity by ID if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableAppointmentID(id *uuid.UUID) *OrderItemUpdateOne {
	if id != nil {
		oiuo = oiuo.SetAppointmentID(*id)
	}
	return oiuo
}

// SetAppointment sets the "appointment" edge to the Appointment entity.
func (oiuo *OrderItemUpdateOne) SetAppointment(a *Appointment) *OrderItemUpdateOne {
	return oiuo.SetAppointmentID(a.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearOrder() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrder()
	return oiuo
}

// ClearProduct clears the "product" edge to the Product entity.
func (oiuo *OrderItemUpdateOne) ClearProduct() *OrderItemUpdateOne {
	oiuo.mutation.ClearProduct()
	return oiuo
}

// ClearService clears the "service" edge to the Service entity.
func (oiuo *OrderItemUpdateOne) ClearService() *OrderItemUpdateOne {
	oiuo.mutation.ClearService()
	return oiuo
}

// ClearAppointment clears the "appointment" edge to the Appointment entity.
func (oiuo *OrderItemUpdateOne) ClearAppointment() *OrderItemUpdateOne {
	oiuo.mutation.ClearAppointment()
	return oiuo
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiuo *OrderItemUpdateOne) check() error {
	if v, ok := oiuo.mutation.ItemType(); ok {
		if err := orderitem.ItemTypeValidator(v); err != nil {
			return &ValidationError{Name: "item_type", err: fmt.Errorf(`db: validator failed for field "OrderItem.item_type": %w`, err)}
		}
	}
	return nil
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	if err := oiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.ItemType(); ok {
		_spec.SetField(orderitem.FieldItemType, field.TypeEnum, value)
	}
	if value, ok := oiuo.mutation.ItemName(); ok {
		_spec.SetField(orderitem.FieldItemName, field.TypeString, value)
	}
	if value, ok := oiuo.mutation.ItemDescription(); ok {
		_spec.SetField(orderitem.FieldItemDescription, field.TypeString, value)
	}
	if oiuo.mutation.ItemDescriptionCleared() {
		_spec.ClearField(orderitem.FieldItemDescription, field.TypeString)
	}
	if value, ok := oiuo.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.UnitPrice(); ok {
		_spec.SetField(orderitem.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.AddedUnitPrice(); ok {
		_spec.AddField(orderitem.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.DiscountAmount(); ok {
		_spec.SetField(orderitem.FieldDiscountAmount, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.AddedDiscountAmount(); ok {
		_spec.AddField(orderitem.FieldDiscountAmount, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.TotalPrice(); ok {
		_spec.SetField(orderitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.AddedTotalPrice(); ok {
		_spec.AddField(orderitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.ItemMetadata(); ok {
		_spec.SetField(orderitem.FieldItemMetadata, field.TypeJSON, value)
	}
	if oiuo.mutation.ItemMetadataCleared() {
		_spec.ClearField(orderitem.FieldItemMetadata, field.TypeJSON)
	}
	if value, ok := oiuo.mutation.CreatedAt(); ok {
		_spec.SetField(orderitem.FieldCreatedAt, field.TypeTime, value)
	}
	if oiuo.mutation.OrderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.OrderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiuo.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiuo.mutation.ServiceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.ServiceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiuo.mutation.AppointmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.AppointmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}
