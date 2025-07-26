package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Appointment struct {
	ent.Schema
}

func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Time("appointment_date"),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Enum("status").Values("SCHEDULED", "CONFIRMED", "IN_PROGRESS", "COMPLETED", "CANCELLED", "NO_SHOW").Default("SCHEDULED"),
		field.Text("notes").Optional(),
		field.Text("symptoms").Optional(),
		field.Text("diagnosis").Optional(),
		field.Text("treatment_plan").Optional(),
		field.JSON("prescriptions", []map[string]interface{}{}).Optional(),
		field.Float("total_cost").Optional(),
		field.Enum("payment_status").Values("PENDING", "PAID", "PARTIAL", "CANCELLED").Default("PENDING"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("appointments").Unique(),
		edge.From("patient", Patient.Type).Ref("appointments").Unique(),
		edge.From("doctor", Doctor.Type).Ref("appointments").Unique(),
		edge.From("service", Service.Type).Ref("appointments").Unique().Required(),
		edge.To("reminders", AppointmentReminder.Type),
		edge.To("order_items", OrderItem.Type),
	}
}

func (Appointment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("appointment_date", "start_time"),
		index.Fields("status"),
	}
}
