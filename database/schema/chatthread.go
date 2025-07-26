package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ChatThread struct {
	ent.Schema
}

func (ChatThread) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("whatsapp_thread_id").Optional().Comment("external thread ID from WhatsApp"),
		field.Enum("status").Values("ACTIVE", "CLOSED", "ARCHIVED").Default("ACTIVE"),
		field.JSON("context", map[string]interface{}{}).Optional().Comment("AI context data"),
		field.Time("last_message_at").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (ChatThread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("chat_threads").Unique(),
		edge.From("patient", Patient.Type).Ref("chat_threads").Unique(),
		edge.To("messages", ChatMessage.Type),
	}
}

func (ChatThread) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("whatsapp_thread_id"),
		index.Fields("status"),
		index.Fields("last_message_at"),
	}
}
