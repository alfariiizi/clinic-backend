package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Document struct {
	ent.Schema
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.Enum("type").Values("LAB_RESULT", "PRESCRIPTION", "MEDICAL_RECORD", "IMAGE", "CONSENT_FORM", "INSURANCE").Default("MEDICAL_RECORD"),
		field.String("file_path"),
		field.String("file_type").Comment("pdf, jpg, png, etc"),
		field.Int64("file_size"),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Time("document_date").Optional(),
		field.Bool("is_confidential").Default(true),
		field.Time("created_at").Default(time.Now),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("documents").Unique(),
		edge.From("patient", Patient.Type).Ref("documents").Unique(),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("document_date"),
	}
}
