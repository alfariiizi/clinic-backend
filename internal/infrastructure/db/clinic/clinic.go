// Code generated by ent, DO NOT EDIT.

package clinic

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the clinic type in the database.
	Label = "clinic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldBusinessHours holds the string denoting the business_hours field in the database.
	FieldBusinessHours = "business_hours"
	// FieldWhatsappNumber holds the string denoting the whatsapp_number field in the database.
	FieldWhatsappNumber = "whatsapp_number"
	// FieldSubscriptionPlan holds the string denoting the subscription_plan field in the database.
	FieldSubscriptionPlan = "subscription_plan"
	// FieldEnabledFeatures holds the string denoting the enabled_features field in the database.
	FieldEnabledFeatures = "enabled_features"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeClinicUsers holds the string denoting the clinic_users edge name in mutations.
	EdgeClinicUsers = "clinic_users"
	// EdgePatients holds the string denoting the patients edge name in mutations.
	EdgePatients = "patients"
	// EdgeDoctors holds the string denoting the doctors edge name in mutations.
	EdgeDoctors = "doctors"
	// EdgeServices holds the string denoting the services edge name in mutations.
	EdgeServices = "services"
	// EdgeAppointments holds the string denoting the appointments edge name in mutations.
	EdgeAppointments = "appointments"
	// EdgeChatThreads holds the string denoting the chat_threads edge name in mutations.
	EdgeChatThreads = "chat_threads"
	// EdgeKnowledgeBase holds the string denoting the knowledge_base edge name in mutations.
	EdgeKnowledgeBase = "knowledge_base"
	// EdgeBillingRecords holds the string denoting the billing_records edge name in mutations.
	EdgeBillingRecords = "billing_records"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeProductCategories holds the string denoting the product_categories edge name in mutations.
	EdgeProductCategories = "product_categories"
	// EdgeInventoryMovements holds the string denoting the inventory_movements edge name in mutations.
	EdgeInventoryMovements = "inventory_movements"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the clinic in the database.
	Table = "clinics"
	// ClinicUsersTable is the table that holds the clinic_users relation/edge.
	ClinicUsersTable = "clinic_users"
	// ClinicUsersInverseTable is the table name for the ClinicUser entity.
	// It exists in this package in order to avoid circular dependency with the "clinicuser" package.
	ClinicUsersInverseTable = "clinic_users"
	// ClinicUsersColumn is the table column denoting the clinic_users relation/edge.
	ClinicUsersColumn = "clinic_id"
	// PatientsTable is the table that holds the patients relation/edge.
	PatientsTable = "patients"
	// PatientsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientsInverseTable = "patients"
	// PatientsColumn is the table column denoting the patients relation/edge.
	PatientsColumn = "clinic_patients"
	// DoctorsTable is the table that holds the doctors relation/edge.
	DoctorsTable = "doctors"
	// DoctorsInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorsInverseTable = "doctors"
	// DoctorsColumn is the table column denoting the doctors relation/edge.
	DoctorsColumn = "clinic_doctors"
	// ServicesTable is the table that holds the services relation/edge.
	ServicesTable = "services"
	// ServicesInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServicesInverseTable = "services"
	// ServicesColumn is the table column denoting the services relation/edge.
	ServicesColumn = "clinic_services"
	// AppointmentsTable is the table that holds the appointments relation/edge.
	AppointmentsTable = "appointments"
	// AppointmentsInverseTable is the table name for the Appointment entity.
	// It exists in this package in order to avoid circular dependency with the "appointment" package.
	AppointmentsInverseTable = "appointments"
	// AppointmentsColumn is the table column denoting the appointments relation/edge.
	AppointmentsColumn = "clinic_appointments"
	// ChatThreadsTable is the table that holds the chat_threads relation/edge.
	ChatThreadsTable = "chat_threads"
	// ChatThreadsInverseTable is the table name for the ChatThread entity.
	// It exists in this package in order to avoid circular dependency with the "chatthread" package.
	ChatThreadsInverseTable = "chat_threads"
	// ChatThreadsColumn is the table column denoting the chat_threads relation/edge.
	ChatThreadsColumn = "clinic_chat_threads"
	// KnowledgeBaseTable is the table that holds the knowledge_base relation/edge.
	KnowledgeBaseTable = "knowledge_bases"
	// KnowledgeBaseInverseTable is the table name for the KnowledgeBase entity.
	// It exists in this package in order to avoid circular dependency with the "knowledgebase" package.
	KnowledgeBaseInverseTable = "knowledge_bases"
	// KnowledgeBaseColumn is the table column denoting the knowledge_base relation/edge.
	KnowledgeBaseColumn = "clinic_knowledge_base"
	// BillingRecordsTable is the table that holds the billing_records relation/edge.
	BillingRecordsTable = "billing_records"
	// BillingRecordsInverseTable is the table name for the BillingRecord entity.
	// It exists in this package in order to avoid circular dependency with the "billingrecord" package.
	BillingRecordsInverseTable = "billing_records"
	// BillingRecordsColumn is the table column denoting the billing_records relation/edge.
	BillingRecordsColumn = "clinic_billing_records"
	// DocumentsTable is the table that holds the documents relation/edge.
	DocumentsTable = "documents"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// DocumentsColumn is the table column denoting the documents relation/edge.
	DocumentsColumn = "clinic_documents"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "clinic_products"
	// ProductCategoriesTable is the table that holds the product_categories relation/edge.
	ProductCategoriesTable = "product_categories"
	// ProductCategoriesInverseTable is the table name for the ProductCategory entity.
	// It exists in this package in order to avoid circular dependency with the "productcategory" package.
	ProductCategoriesInverseTable = "product_categories"
	// ProductCategoriesColumn is the table column denoting the product_categories relation/edge.
	ProductCategoriesColumn = "clinic_product_categories"
	// InventoryMovementsTable is the table that holds the inventory_movements relation/edge.
	InventoryMovementsTable = "inventory_movements"
	// InventoryMovementsInverseTable is the table name for the InventoryMovement entity.
	// It exists in this package in order to avoid circular dependency with the "inventorymovement" package.
	InventoryMovementsInverseTable = "inventory_movements"
	// InventoryMovementsColumn is the table column denoting the inventory_movements relation/edge.
	InventoryMovementsColumn = "clinic_inventory_movements"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "clinic_orders"
)

// Columns holds all SQL columns for clinic fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldPhone,
	FieldEmail,
	FieldAddress,
	FieldBusinessHours,
	FieldWhatsappNumber,
	FieldSubscriptionPlan,
	FieldEnabledFeatures,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSubscriptionPlan holds the default value on creation for the "subscription_plan" field.
	DefaultSubscriptionPlan string
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

// OrderOption defines the ordering options for the Clinic queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByWhatsappNumber orders the results by the whatsapp_number field.
func ByWhatsappNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWhatsappNumber, opts...).ToFunc()
}

// BySubscriptionPlan orders the results by the subscription_plan field.
func BySubscriptionPlan(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionPlan, opts...).ToFunc()
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

// ByClinicUsersCount orders the results by clinic_users count.
func ByClinicUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClinicUsersStep(), opts...)
	}
}

// ByClinicUsers orders the results by clinic_users terms.
func ByClinicUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClinicUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPatientsCount orders the results by patients count.
func ByPatientsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPatientsStep(), opts...)
	}
}

// ByPatients orders the results by patients terms.
func ByPatients(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPatientsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDoctorsCount orders the results by doctors count.
func ByDoctorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDoctorsStep(), opts...)
	}
}

// ByDoctors orders the results by doctors terms.
func ByDoctors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDoctorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServicesCount orders the results by services count.
func ByServicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServicesStep(), opts...)
	}
}

// ByServices orders the results by services terms.
func ByServices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServicesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByChatThreadsCount orders the results by chat_threads count.
func ByChatThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChatThreadsStep(), opts...)
	}
}

// ByChatThreads orders the results by chat_threads terms.
func ByChatThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKnowledgeBaseCount orders the results by knowledge_base count.
func ByKnowledgeBaseCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKnowledgeBaseStep(), opts...)
	}
}

// ByKnowledgeBase orders the results by knowledge_base terms.
func ByKnowledgeBase(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKnowledgeBaseStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBillingRecordsCount orders the results by billing_records count.
func ByBillingRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBillingRecordsStep(), opts...)
	}
}

// ByBillingRecords orders the results by billing_records terms.
func ByBillingRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillingRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentsCount orders the results by documents count.
func ByDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentsStep(), opts...)
	}
}

// ByDocuments orders the results by documents terms.
func ByDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductCategoriesCount orders the results by product_categories count.
func ByProductCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductCategoriesStep(), opts...)
	}
}

// ByProductCategories orders the results by product_categories terms.
func ByProductCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInventoryMovementsCount orders the results by inventory_movements count.
func ByInventoryMovementsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInventoryMovementsStep(), opts...)
	}
}

// ByInventoryMovements orders the results by inventory_movements terms.
func ByInventoryMovements(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInventoryMovementsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newClinicUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClinicUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ClinicUsersTable, ClinicUsersColumn),
	)
}
func newPatientsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PatientsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PatientsTable, PatientsColumn),
	)
}
func newDoctorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DoctorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DoctorsTable, DoctorsColumn),
	)
}
func newServicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServicesTable, ServicesColumn),
	)
}
func newAppointmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppointmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppointmentsTable, AppointmentsColumn),
	)
}
func newChatThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChatThreadsTable, ChatThreadsColumn),
	)
}
func newKnowledgeBaseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KnowledgeBaseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, KnowledgeBaseTable, KnowledgeBaseColumn),
	)
}
func newBillingRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillingRecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BillingRecordsTable, BillingRecordsColumn),
	)
}
func newDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DocumentsTable, DocumentsColumn),
	)
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
	)
}
func newProductCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductCategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductCategoriesTable, ProductCategoriesColumn),
	)
}
func newInventoryMovementsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InventoryMovementsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InventoryMovementsTable, InventoryMovementsColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}
