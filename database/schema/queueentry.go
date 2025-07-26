package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type QueueEntry struct {
	ent.Schema
}

func (QueueEntry) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("clinic_id"),
		field.String("patient_id"),
		field.String("doctor_id").Optional(),
		field.String("service_id").Optional(),
		field.Int("queue_number"),
		field.Enum("status").Values("WAITING", "CALLED", "IN_PROGRESS", "COMPLETED", "CANCELLED").Default("WAITING"),
		field.Time("estimated_time").Optional(),
		field.Time("called_at").Optional(),
		field.Time("completed_at").Optional(),
		field.Text("notes").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (QueueEntry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("clinic_id", "queue_number"),
		index.Fields("status"),
		index.Fields("created_at"),
	}
}
