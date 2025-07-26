package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("order_number").Unique(),
		field.Enum("order_type").Values("PRODUCT_ONLY", "SERVICE_ONLY", "MIXED").Default("MIXED"),
		field.Enum("status").Values("DRAFT", "PENDING", "CONFIRMED", "PROCESSING", "READY", "COMPLETED", "CANCELLED").Default("PENDING"),
		field.Float("subtotal").Default(0),
		field.Float("tax_amount").Default(0),
		field.Float("discount_amount").Default(0),
		field.Float("shipping_cost").Default(0),
		field.Float("total_amount").Default(0),
		field.String("currency").Default("IDR"),
		field.Enum("payment_status").Values("PENDING", "PAID", "PARTIAL", "REFUNDED", "CANCELLED").Default("PENDING"),
		field.Enum("payment_method").Values("CASH", "CREDIT_CARD", "DEBIT_CARD", "BANK_TRANSFER", "DIGITAL_WALLET").Optional(),
		field.JSON("shipping_address", map[string]interface{}{}).Optional(),
		field.JSON("billing_address", map[string]interface{}{}).Optional(),
		field.Enum("delivery_method").Values("PICKUP", "DELIVERY", "SHIPPING").Default("PICKUP"),
		field.Time("expected_delivery_date").Optional(),
		field.Time("delivered_at").Optional(),
		field.Text("notes").Optional(),
		field.Text("cancellation_reason").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("orders").Unique(),
		edge.From("patient", Patient.Type).Ref("orders").Unique(),
		edge.To("order_items", OrderItem.Type),
		edge.To("order_status_history", OrderStatusHistory.Type),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_number"),
		index.Fields("status"),
		index.Fields("payment_status"),
		index.Fields("order_type"),
		index.Fields("created_at"),
	}
}
