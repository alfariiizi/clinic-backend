package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Feature struct {
	ent.Schema
}

func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").Unique(),
		field.String("display_name"),
		field.Text("description").Optional(),
		field.Enum("category").Values("CORE", "AI", "BILLING", "COMMUNICATION", "ANALYTICS").Default("CORE"),
		field.Float("monthly_price").Default(0),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now),
	}
}
