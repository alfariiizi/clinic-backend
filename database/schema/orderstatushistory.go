package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type OrderStatusHistory struct {
	ent.Schema
}

func (OrderStatusHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("status"),
		field.Text("notes").Optional(),
		field.String("changed_by").Optional().Comment("user who changed status"),
		field.Time("created_at").Default(time.Now),
	}
}

func (OrderStatusHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_status_history").Unique(),
	}
}

func (OrderStatusHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
		index.Fields("created_at"),
	}
}
