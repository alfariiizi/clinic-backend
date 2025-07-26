package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Enum("item_type").Values("PRODUCT", "SERVICE").Default("PRODUCT"),
		field.String("item_name").Comment("snapshot of name at time of order"),
		field.Text("item_description").Optional(),
		field.Int("quantity").Default(1),
		field.Float("unit_price"),
		field.Float("discount_amount").Default(0),
		field.Float("total_price"),
		field.JSON("item_metadata", map[string]any{}).Optional().Comment("snapshot of product/service data"),
		field.Time("created_at").Default(time.Now),
	}
}

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_items").Unique(),
		edge.From("product", Product.Type).Ref("order_items").Unique(),
		edge.From("service", Service.Type).Ref("order_items").Unique(),
		edge.From("appointment", Appointment.Type).Ref("order_items").Unique(),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("item_type"),
	}
}
