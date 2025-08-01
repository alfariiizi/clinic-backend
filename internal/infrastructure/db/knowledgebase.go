// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/knowledgebase"
	"github.com/google/uuid"
)

// KnowledgeBase is the model entity for the KnowledgeBase schema.
type KnowledgeBase struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Category holds the value of the "category" field.
	Category knowledgebase.Category `json:"category,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// higher number = higher priority
	Priority int `json:"priority,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KnowledgeBaseQuery when eager-loading is set.
	Edges                 KnowledgeBaseEdges `json:"edges"`
	clinic_knowledge_base *uuid.UUID
	selectValues          sql.SelectValues
}

// KnowledgeBaseEdges holds the relations/edges for other nodes in the graph.
type KnowledgeBaseEdges struct {
	// Clinic holds the value of the clinic edge.
	Clinic *Clinic `json:"clinic,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClinicOrErr returns the Clinic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KnowledgeBaseEdges) ClinicOrErr() (*Clinic, error) {
	if e.Clinic != nil {
		return e.Clinic, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: clinic.Label}
	}
	return nil, &NotLoadedError{edge: "clinic"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KnowledgeBase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case knowledgebase.FieldTags, knowledgebase.FieldMetadata:
			values[i] = new([]byte)
		case knowledgebase.FieldActive:
			values[i] = new(sql.NullBool)
		case knowledgebase.FieldPriority:
			values[i] = new(sql.NullInt64)
		case knowledgebase.FieldTitle, knowledgebase.FieldCategory, knowledgebase.FieldContent:
			values[i] = new(sql.NullString)
		case knowledgebase.FieldCreatedAt, knowledgebase.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case knowledgebase.FieldID:
			values[i] = new(uuid.UUID)
		case knowledgebase.ForeignKeys[0]: // clinic_knowledge_base
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KnowledgeBase fields.
func (kb *KnowledgeBase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case knowledgebase.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				kb.ID = *value
			}
		case knowledgebase.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				kb.Title = value.String
			}
		case knowledgebase.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				kb.Category = knowledgebase.Category(value.String)
			}
		case knowledgebase.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				kb.Content = value.String
			}
		case knowledgebase.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &kb.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case knowledgebase.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &kb.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case knowledgebase.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				kb.Active = value.Bool
			}
		case knowledgebase.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				kb.Priority = int(value.Int64)
			}
		case knowledgebase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kb.CreatedAt = value.Time
			}
		case knowledgebase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kb.UpdatedAt = value.Time
			}
		case knowledgebase.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field clinic_knowledge_base", values[i])
			} else if value.Valid {
				kb.clinic_knowledge_base = new(uuid.UUID)
				*kb.clinic_knowledge_base = *value.S.(*uuid.UUID)
			}
		default:
			kb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the KnowledgeBase.
// This includes values selected through modifiers, order, etc.
func (kb *KnowledgeBase) Value(name string) (ent.Value, error) {
	return kb.selectValues.Get(name)
}

// QueryClinic queries the "clinic" edge of the KnowledgeBase entity.
func (kb *KnowledgeBase) QueryClinic() *ClinicQuery {
	return NewKnowledgeBaseClient(kb.config).QueryClinic(kb)
}

// Update returns a builder for updating this KnowledgeBase.
// Note that you need to call KnowledgeBase.Unwrap() before calling this method if this KnowledgeBase
// was returned from a transaction, and the transaction was committed or rolled back.
func (kb *KnowledgeBase) Update() *KnowledgeBaseUpdateOne {
	return NewKnowledgeBaseClient(kb.config).UpdateOne(kb)
}

// Unwrap unwraps the KnowledgeBase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kb *KnowledgeBase) Unwrap() *KnowledgeBase {
	_tx, ok := kb.config.driver.(*txDriver)
	if !ok {
		panic("db: KnowledgeBase is not a transactional entity")
	}
	kb.config.driver = _tx.drv
	return kb
}

// String implements the fmt.Stringer.
func (kb *KnowledgeBase) String() string {
	var builder strings.Builder
	builder.WriteString("KnowledgeBase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", kb.ID))
	builder.WriteString("title=")
	builder.WriteString(kb.Title)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", kb.Category))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(kb.Content)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", kb.Tags))
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", kb.Metadata))
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", kb.Active))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", kb.Priority))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(kb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(kb.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// KnowledgeBases is a parsable slice of KnowledgeBase.
type KnowledgeBases []*KnowledgeBase
