package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Clinic represents a clinic entity (tenant)
type Clinic struct {
	ent.Schema
}

func (Clinic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.String("type").Comment("teeth, skincare, general, etc"),
		field.String("phone").Optional(),
		field.String("email").Optional(),
		field.Text("address").Optional(),
		field.JSON("business_hours", map[string]interface{}{}).Optional(),
		field.String("whatsapp_number").Optional(),
		field.String("subscription_plan").Default("basic"),
		field.JSON("enabled_features", []string{}).Comment("array of enabled feature names"),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Clinic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("patients", Patient.Type),
		edge.To("doctors", Doctor.Type),
		edge.To("services", Service.Type),
		edge.To("appointments", Appointment.Type),
		edge.To("chat_threads", ChatThread.Type),
		edge.To("knowledge_base", KnowledgeBase.Type),
		edge.To("billing_records", BillingRecord.Type),
		edge.To("documents", Document.Type),
		edge.To("products", Product.Type),
		edge.To("product_categories", ProductCategory.Type),
		edge.To("inventory_movements", InventoryMovement.Type),
		edge.To("orders", Order.Type),
	}
}

func (Clinic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("whatsapp_number"),
		index.Fields("active"),
	}
}
