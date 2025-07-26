package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type InventoryMovement struct {
	ent.Schema
}

func (InventoryMovement) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Enum("type").Values("PURCHASE", "SALE", "ADJUSTMENT", "RETURN", "DAMAGE", "EXPIRED").Default("SALE"),
		field.Int("quantity"),
		field.String("reference_id").Optional().Comment("order ID, adjustment ID, etc"),
		field.Text("notes").Optional(),
		field.String("performed_by").Optional().Comment("user who made the movement"),
		field.Time("movement_date").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

func (InventoryMovement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("inventory_movements").Unique(),
		edge.From("product", Product.Type).Ref("inventory_movements").Unique(),
	}
}

func (InventoryMovement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("movement_date"),
		index.Fields("reference_id"),
	}
}
