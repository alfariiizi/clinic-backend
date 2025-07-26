package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Service struct {
	ent.Schema
}

func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.Text("description").Optional(),
		field.String("category").Optional(),
		field.Float("price").Default(0),
		field.Int("duration").Default(30).Comment("in minutes"),
		field.JSON("requirements", []string{}).Optional().Comment("prerequisites for service"),
		field.Bool("requires_appointment").Default(true),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("services").Unique(),
		edge.To("appointments", Appointment.Type),
		edge.To("order_items", OrderItem.Type),
	}
}
