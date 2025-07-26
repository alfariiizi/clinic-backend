package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type BillingRecord struct {
	ent.Schema
}

func (BillingRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("invoice_number").Unique(),
		field.Float("amount"),
		field.Float("tax_amount").Default(0),
		field.Float("discount_amount").Default(0),
		field.Float("total_amount"),
		field.String("currency").Default("IDR"),
		field.Enum("payment_method").Values("CASH", "CREDIT_CARD", "DEBIT_CARD", "BANK_TRANSFER", "DIGITAL_WALLET").Optional(),
		field.Enum("payment_status").Values("PENDING", "PAID", "PARTIAL", "OVERDUE", "CANCELLED").Default("PENDING"),
		field.JSON("line_items", []map[string]interface{}{}).Optional().Comment("services and costs breakdown"),
		field.Time("due_date").Optional(),
		field.Time("paid_at").Optional(),
		field.Text("notes").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BillingRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("billing_records").Unique(),
		edge.From("patient", Patient.Type).Ref("billing_records").Unique(),
	}
}

func (BillingRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("invoice_number"),
		index.Fields("payment_status"),
		index.Fields("due_date"),
	}
}
