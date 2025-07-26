package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("whatsapp_message_id").Optional().Comment("external message ID from WhatsApp"),
		field.Enum("sender_type").Values("PATIENT", "AI", "STAFF", "DOCTOR"),
		field.Enum("message_type").Values("TEXT", "IMAGE", "DOCUMENT", "AUDIO", "LOCATION").Default("TEXT"),
		field.Text("content"),
		field.JSON("metadata", map[string]interface{}{}).Optional().Comment("file info, location data, etc"),
		field.String("ai_tool_call").Optional().Comment("tool call made by AI"),
		field.JSON("ai_tool_result", map[string]interface{}{}).Optional(),
		field.Bool("is_read").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread", ChatThread.Type).Ref("messages").Unique(),
	}
}

func (ChatMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("whatsapp_message_id"),
		index.Fields("sender_type"),
		index.Fields("created_at"),
	}
}
