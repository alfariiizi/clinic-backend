package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type KnowledgeBase struct {
	ent.Schema
}

func (KnowledgeBase) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("title"),
		field.Enum("category").Values("FAQ", "SERVICE", "POLICY", "PROCEDURE", "PRODUCT").Default("FAQ"),
		field.Text("content"),
		field.JSON("tags", []string{}).Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Bool("active").Default(true),
		field.Int("priority").Default(0).Comment("higher number = higher priority"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (KnowledgeBase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("knowledge_base").Unique(),
	}
}

func (KnowledgeBase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
		index.Fields("active"),
		index.Fields("priority"),
	}
}
