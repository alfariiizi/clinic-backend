package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type AIInteraction struct {
	ent.Schema
}

func (AIInteraction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("clinic_id"),
		field.String("patient_whatsapp").Optional(),
		field.Enum("interaction_type").Values("CHAT", "TOOL_CALL", "RECOMMENDATION").Default("CHAT"),
		field.JSON("request_payload", map[string]interface{}{}),
		field.JSON("response_payload", map[string]interface{}{}),
		field.String("ai_model").Optional(),
		field.Int("response_time_ms").Optional(),
		field.Enum("status").Values("SUCCESS", "ERROR", "TIMEOUT").Default("SUCCESS"),
		field.Text("error_message").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (AIInteraction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("clinic_id"),
		index.Fields("interaction_type"),
		index.Fields("status"),
		index.Fields("created_at"),
	}
}
