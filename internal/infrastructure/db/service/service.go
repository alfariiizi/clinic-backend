// Code generated by ent, DO NOT EDIT.

package service

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldRequirements holds the string denoting the requirements field in the database.
	FieldRequirements = "requirements"
	// FieldRequiresAppointment holds the string denoting the requires_appointment field in the database.
	FieldRequiresAppointment = "requires_appointment"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeClinic holds the string denoting the clinic edge name in mutations.
	EdgeClinic = "clinic"
	// EdgeAppointments holds the string denoting the appointments edge name in mutations.
	EdgeAppointments = "appointments"
	// EdgeOrderItems holds the string denoting the order_items edge name in mutations.
	EdgeOrderItems = "order_items"
	// Table holds the table name of the service in the database.
	Table = "services"
	// ClinicTable is the table that holds the clinic relation/edge.
	ClinicTable = "services"
	// ClinicInverseTable is the table name for the Clinic entity.
	// It exists in this package in order to avoid circular dependency with the "clinic" package.
	ClinicInverseTable = "clinics"
	// ClinicColumn is the table column denoting the clinic relation/edge.
	ClinicColumn = "clinic_services"
	// AppointmentsTable is the table that holds the appointments relation/edge.
	AppointmentsTable = "appointments"
	// AppointmentsInverseTable is the table name for the Appointment entity.
	// It exists in this package in order to avoid circular dependency with the "appointment" package.
	AppointmentsInverseTable = "appointments"
	// AppointmentsColumn is the table column denoting the appointments relation/edge.
	AppointmentsColumn = "service_appointments"
	// OrderItemsTable is the table that holds the order_items relation/edge.
	OrderItemsTable = "order_items"
	// OrderItemsInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemsInverseTable = "order_items"
	// OrderItemsColumn is the table column denoting the order_items relation/edge.
	OrderItemsColumn = "service_order_items"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCategory,
	FieldPrice,
	FieldDuration,
	FieldRequirements,
	FieldRequiresAppointment,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "services"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"clinic_services",
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
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice float64
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultRequiresAppointment holds the default value on creation for the "requires_appointment" field.
	DefaultRequiresAppointment bool
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Service queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByRequiresAppointment orders the results by the requires_appointment field.
func ByRequiresAppointment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiresAppointment, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByClinicField orders the results by clinic field.
func ByClinicField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClinicStep(), sql.OrderByField(field, opts...))
	}
}

// ByAppointmentsCount orders the results by appointments count.
func ByAppointmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppointmentsStep(), opts...)
	}
}

// ByAppointments orders the results by appointments terms.
func ByAppointments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppointmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderItemsCount orders the results by order_items count.
func ByOrderItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderItemsStep(), opts...)
	}
}

// ByOrderItems orders the results by order_items terms.
func ByOrderItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newClinicStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClinicInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClinicTable, ClinicColumn),
	)
}
func newAppointmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppointmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppointmentsTable, AppointmentsColumn),
	)
}
func newOrderItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderItemsTable, OrderItemsColumn),
	)
}
