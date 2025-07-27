package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Doctor struct {
	ent.Schema
}

func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.String("specialization"),
		field.String("license_number").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Text("bio").Optional(),
		field.JSON("qualifications", []string{}).Optional(),
		field.JSON("availability", map[string]any{}).Comment("weekly schedule"),
		field.Int("consultation_duration").Default(30).Comment("in minutes"),
		field.Float("consultation_fee").Default(0),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("doctors").Unique(),
		edge.To("appointments", Appointment.Type),
		edge.To("schedules", DoctorSchedule.Type),
	}
}
