package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ClinicUser holds the schema definition for the ClinicUser entity.
type ClinicUser struct {
	ent.Schema
}

// Fields of the ClinicUser.
func (ClinicUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("clinic_id", uuid.UUID{}).
			Comment("ID of the clinic this user belongs to"),
		field.UUID("user_id", uuid.UUID{}).
			Comment("ID of the user in the system"),
	}
}

// Edges of the ClinicUser.
func (ClinicUser) Edges() []ent.Edge {
	return []ent.Edge{
		// Define edges if needed, e.g., to Clinic or User entities
		edge.From("clinic", Clinic.Type).
			Ref("clinic_users").
			Required().
			Unique().
			Field("clinic_id"),
		edge.From("user", User.Type).
			Ref("clinic_users").
			Required().
			Unique().
			Field("user_id"),
	}
}
