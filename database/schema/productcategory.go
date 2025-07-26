package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ProductCategory struct {
	ent.Schema
}

func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.Text("description").Optional(),
		field.String("image_url").Optional(),
		field.Bool("active").Default(true),
		field.Int("sort_order").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("product_categories").Unique(),
		edge.To("products", Product.Type),
	}
}

func (ProductCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active"),
		index.Fields("sort_order"),
	}
}
