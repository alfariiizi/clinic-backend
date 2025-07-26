package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("sku").Comment("Stock Keeping Unit"),
		field.String("name"),
		field.Text("description").Optional(),
		field.Text("short_description").Optional(),
		field.String("brand").Optional(),
		field.JSON("images", []string{}).Optional().Comment("array of image URLs"),
		field.Float("purchase_price").Default(0).Comment("cost price from supplier"),
		field.Float("selling_price").Default(0).Comment("price to sell to customer"),
		field.Float("discount_price").Optional().Comment("discounted price if any"),
		field.String("unit").Default("pcs").Comment("unit of measurement: pcs, ml, gram, etc"),
		field.Int("min_stock_level").Default(0).Comment("minimum stock before alert"),
		field.Int("current_stock").Default(0),
		field.Bool("track_inventory").Default(true),
		field.Bool("prescription_required").Default(false),
		field.JSON("specifications", map[string]interface{}{}).Optional().Comment("product specs like ingredients, size, etc"),
		field.JSON("usage_instructions", []string{}).Optional(),
		field.JSON("warnings", []string{}).Optional(),
		field.Time("expiry_date").Optional(),
		field.String("batch_number").Optional(),
		field.Enum("status").Values("ACTIVE", "INACTIVE", "OUT_OF_STOCK", "DISCONTINUED").Default("ACTIVE"),
		field.JSON("tags", []string{}).Optional().Comment("for search and filtering"),
		field.Float("weight").Optional().Comment("for shipping calculation"),
		field.JSON("dimensions", map[string]float64{}).Optional().Comment("length, width, height"),
		field.Bool("featured").Default(false).Comment("featured product"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clinic", Clinic.Type).Ref("products").Unique(),
		edge.From("category", ProductCategory.Type).Ref("products").Unique().Required(),
		edge.To("inventory_movements", InventoryMovement.Type),
		edge.To("order_items", OrderItem.Type),
		// edge.To("product_reviews", ProductReview.Type),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sku"),
		index.Fields("status"),
		index.Fields("current_stock"),
		index.Fields("featured"),
		index.Fields("prescription_required"),
	}
}
