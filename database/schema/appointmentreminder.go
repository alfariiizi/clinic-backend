package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type AppointmentReminder struct {
	ent.Schema
}

func (AppointmentReminder) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Enum("type").Values("SMS", "WHATSAPP", "EMAIL"),
		field.Time("scheduled_time"),
		field.String("message"),
		field.Enum("status").Values("PENDING", "SENT", "FAILED").Default("PENDING"),
		field.Time("sent_at").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (AppointmentReminder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appointment", Appointment.Type).Ref("reminders").Unique(),
	}
}
