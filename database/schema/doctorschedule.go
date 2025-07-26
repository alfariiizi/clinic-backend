package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ent.Schema
}

func (DoctorSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Time("date"),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Bool("available").Default(true),
		field.String("notes").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (DoctorSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).Ref("schedules").Unique(),
	}
}

func (DoctorSchedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("date", "start_time"),
	}
}
