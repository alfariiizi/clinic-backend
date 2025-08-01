// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the orderitem type in the database.
	Label = "order_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldItemType holds the string denoting the item_type field in the database.
	FieldItemType = "item_type"
	// FieldItemName holds the string denoting the item_name field in the database.
	FieldItemName = "item_name"
	// FieldItemDescription holds the string denoting the item_description field in the database.
	FieldItemDescription = "item_description"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnitPrice holds the string denoting the unit_price field in the database.
	FieldUnitPrice = "unit_price"
	// FieldDiscountAmount holds the string denoting the discount_amount field in the database.
	FieldDiscountAmount = "discount_amount"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// FieldItemMetadata holds the string denoting the item_metadata field in the database.
	FieldItemMetadata = "item_metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeAppointment holds the string denoting the appointment edge name in mutations.
	EdgeAppointment = "appointment"
	// Table holds the table name of the orderitem in the database.
	Table = "order_items"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "order_items"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_order_items"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "order_items"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_order_items"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "order_items"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_order_items"
	// AppointmentTable is the table that holds the appointment relation/edge.
	AppointmentTable = "order_items"
	// AppointmentInverseTable is the table name for the Appointment entity.
	// It exists in this package in order to avoid circular dependency with the "appointment" package.
	AppointmentInverseTable = "appointments"
	// AppointmentColumn is the table column denoting the appointment relation/edge.
	AppointmentColumn = "appointment_order_items"
)

// Columns holds all SQL columns for orderitem fields.
var Columns = []string{
	FieldID,
	FieldItemType,
	FieldItemName,
	FieldItemDescription,
	FieldQuantity,
	FieldUnitPrice,
	FieldDiscountAmount,
	FieldTotalPrice,
	FieldItemMetadata,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"appointment_order_items",
	"order_order_items",
	"product_order_items",
	"service_order_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultQuantity holds the default value on creation for the "quantity" field.
	DefaultQuantity int
	// DefaultDiscountAmount holds the default value on creation for the "discount_amount" field.
	DefaultDiscountAmount float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// ItemType defines the type for the "item_type" enum field.
type ItemType string

// ItemTypePRODUCT is the default value of the ItemType enum.
const DefaultItemType = ItemTypePRODUCT

// ItemType values.
const (
	ItemTypePRODUCT ItemType = "PRODUCT"
	ItemTypeSERVICE ItemType = "SERVICE"
)

func (it ItemType) String() string {
	return string(it)
}

// ItemTypeValidator is a validator for the "item_type" field enum values. It is called by the builders before save.
func ItemTypeValidator(it ItemType) error {
	switch it {
	case ItemTypePRODUCT, ItemTypeSERVICE:
		return nil
	default:
		return fmt.Errorf("orderitem: invalid enum value for item_type field: %q", it)
	}
}

// OrderOption defines the ordering options for the OrderItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByItemType orders the results by the item_type field.
func ByItemType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItemType, opts...).ToFunc()
}

// ByItemName orders the results by the item_name field.
func ByItemName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItemName, opts...).ToFunc()
}

// ByItemDescription orders the results by the item_description field.
func ByItemDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItemDescription, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUnitPrice orders the results by the unit_price field.
func ByUnitPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitPrice, opts...).ToFunc()
}

// ByDiscountAmount orders the results by the discount_amount field.
func ByDiscountAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscountAmount, opts...).ToFunc()
}

// ByTotalPrice orders the results by the total_price field.
func ByTotalPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPrice, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductField orders the results by product field.
func ByProductField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), sql.OrderByField(field, opts...))
	}
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByAppointmentField orders the results by appointment field.
func ByAppointmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppointmentStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
	)
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
	)
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
	)
}
func newAppointmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppointmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AppointmentTable, AppointmentColumn),
	)
}
