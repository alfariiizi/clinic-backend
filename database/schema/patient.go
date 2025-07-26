package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Patient struct {
	ent.Schema
}

func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("whatsapp_number").Comment("patient's whatsapp number"),
		field.String("name").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Time("birth_date").Optional(),
		field.Enum("gender").Values("MALE", "FEMALE", "OTHER").Optional(),
		field.Text("address").Optional(),
		field.JSON("medical_history", map[string]interface{}{}).Optional(),
		field.JSON("allergies", []string{}).Optional(),
		field.String("emergency_contact_name").Optional(),
		field.String("emergency_contact_phone").Optional(),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("patients").Unique(),
		edge.To("appointments", Appointment.Type),
		edge.To("chat_threads", ChatThread.Type),
		edge.To("documents", Document.Type),
		edge.To("billing_records", BillingRecord.Type),
		edge.To("orders", Order.Type),
	}
}

func (Patient) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("whatsapp_number"),
		index.Fields("email"),
	}
}
