// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/knowledgebase"
	"github.com/google/uuid"
)

// KnowledgeBaseCreate is the builder for creating a KnowledgeBase entity.
type KnowledgeBaseCreate struct {
	config
	mutation *KnowledgeBaseMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (kbc *KnowledgeBaseCreate) SetTitle(s string) *KnowledgeBaseCreate {
	kbc.mutation.SetTitle(s)
	return kbc
}

// SetCategory sets the "category" field.
func (kbc *KnowledgeBaseCreate) SetCategory(k knowledgebase.Category) *KnowledgeBaseCreate {
	kbc.mutation.SetCategory(k)
	return kbc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableCategory(k *knowledgebase.Category) *KnowledgeBaseCreate {
	if k != nil {
		kbc.SetCategory(*k)
	}
	return kbc
}

// SetContent sets the "content" field.
func (kbc *KnowledgeBaseCreate) SetContent(s string) *KnowledgeBaseCreate {
	kbc.mutation.SetContent(s)
	return kbc
}

// SetTags sets the "tags" field.
func (kbc *KnowledgeBaseCreate) SetTags(s []string) *KnowledgeBaseCreate {
	kbc.mutation.SetTags(s)
	return kbc
}

// SetMetadata sets the "metadata" field.
func (kbc *KnowledgeBaseCreate) SetMetadata(m map[string]interface{}) *KnowledgeBaseCreate {
	kbc.mutation.SetMetadata(m)
	return kbc
}

// SetActive sets the "active" field.
func (kbc *KnowledgeBaseCreate) SetActive(b bool) *KnowledgeBaseCreate {
	kbc.mutation.SetActive(b)
	return kbc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableActive(b *bool) *KnowledgeBaseCreate {
	if b != nil {
		kbc.SetActive(*b)
	}
	return kbc
}

// SetPriority sets the "priority" field.
func (kbc *KnowledgeBaseCreate) SetPriority(i int) *KnowledgeBaseCreate {
	kbc.mutation.SetPriority(i)
	return kbc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillablePriority(i *int) *KnowledgeBaseCreate {
	if i != nil {
		kbc.SetPriority(*i)
	}
	return kbc
}

// SetCreatedAt sets the "created_at" field.
func (kbc *KnowledgeBaseCreate) SetCreatedAt(t time.Time) *KnowledgeBaseCreate {
	kbc.mutation.SetCreatedAt(t)
	return kbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableCreatedAt(t *time.Time) *KnowledgeBaseCreate {
	if t != nil {
		kbc.SetCreatedAt(*t)
	}
	return kbc
}

// SetUpdatedAt sets the "updated_at" field.
func (kbc *KnowledgeBaseCreate) SetUpdatedAt(t time.Time) *KnowledgeBaseCreate {
	kbc.mutation.SetUpdatedAt(t)
	return kbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableUpdatedAt(t *time.Time) *KnowledgeBaseCreate {
	if t != nil {
		kbc.SetUpdatedAt(*t)
	}
	return kbc
}

// SetID sets the "id" field.
func (kbc *KnowledgeBaseCreate) SetID(u uuid.UUID) *KnowledgeBaseCreate {
	kbc.mutation.SetID(u)
	return kbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableID(u *uuid.UUID) *KnowledgeBaseCreate {
	if u != nil {
		kbc.SetID(*u)
	}
	return kbc
}

// SetClinicID sets the "clinic" edge to the Clinic entity by ID.
func (kbc *KnowledgeBaseCreate) SetClinicID(id uuid.UUID) *KnowledgeBaseCreate {
	kbc.mutation.SetClinicID(id)
	return kbc
}

// SetNillableClinicID sets the "clinic" edge to the Clinic entity by ID if the given value is not nil.
func (kbc *KnowledgeBaseCreate) SetNillableClinicID(id *uuid.UUID) *KnowledgeBaseCreate {
	if id != nil {
		kbc = kbc.SetClinicID(*id)
	}
	return kbc
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (kbc *KnowledgeBaseCreate) SetClinic(c *Clinic) *KnowledgeBaseCreate {
	return kbc.SetClinicID(c.ID)
}

// Mutation returns the KnowledgeBaseMutation object of the builder.
func (kbc *KnowledgeBaseCreate) Mutation() *KnowledgeBaseMutation {
	return kbc.mutation
}

// Save creates the KnowledgeBase in the database.
func (kbc *KnowledgeBaseCreate) Save(ctx context.Context) (*KnowledgeBase, error) {
	kbc.defaults()
	return withHooks(ctx, kbc.sqlSave, kbc.mutation, kbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kbc *KnowledgeBaseCreate) SaveX(ctx context.Context) *KnowledgeBase {
	v, err := kbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kbc *KnowledgeBaseCreate) Exec(ctx context.Context) error {
	_, err := kbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kbc *KnowledgeBaseCreate) ExecX(ctx context.Context) {
	if err := kbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kbc *KnowledgeBaseCreate) defaults() {
	if _, ok := kbc.mutation.Category(); !ok {
		v := knowledgebase.DefaultCategory
		kbc.mutation.SetCategory(v)
	}
	if _, ok := kbc.mutation.Active(); !ok {
		v := knowledgebase.DefaultActive
		kbc.mutation.SetActive(v)
	}
	if _, ok := kbc.mutation.Priority(); !ok {
		v := knowledgebase.DefaultPriority
		kbc.mutation.SetPriority(v)
	}
	if _, ok := kbc.mutation.CreatedAt(); !ok {
		v := knowledgebase.DefaultCreatedAt()
		kbc.mutation.SetCreatedAt(v)
	}
	if _, ok := kbc.mutation.UpdatedAt(); !ok {
		v := knowledgebase.DefaultUpdatedAt()
		kbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kbc.mutation.ID(); !ok {
		v := knowledgebase.DefaultID()
		kbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kbc *KnowledgeBaseCreate) check() error {
	if _, ok := kbc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`db: missing required field "KnowledgeBase.title"`)}
	}
	if _, ok := kbc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`db: missing required field "KnowledgeBase.category"`)}
	}
	if v, ok := kbc.mutation.Category(); ok {
		if err := knowledgebase.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`db: validator failed for field "KnowledgeBase.category": %w`, err)}
		}
	}
	if _, ok := kbc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`db: missing required field "KnowledgeBase.content"`)}
	}
	if _, ok := kbc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`db: missing required field "KnowledgeBase.active"`)}
	}
	if _, ok := kbc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`db: missing required field "KnowledgeBase.priority"`)}
	}
	if _, ok := kbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "KnowledgeBase.created_at"`)}
	}
	if _, ok := kbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "KnowledgeBase.updated_at"`)}
	}
	return nil
}

func (kbc *KnowledgeBaseCreate) sqlSave(ctx context.Context) (*KnowledgeBase, error) {
	if err := kbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	kbc.mutation.id = &_node.ID
	kbc.mutation.done = true
	return _node, nil
}

func (kbc *KnowledgeBaseCreate) createSpec() (*KnowledgeBase, *sqlgraph.CreateSpec) {
	var (
		_node = &KnowledgeBase{config: kbc.config}
		_spec = sqlgraph.NewCreateSpec(knowledgebase.Table, sqlgraph.NewFieldSpec(knowledgebase.FieldID, field.TypeUUID))
	)
	if id, ok := kbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kbc.mutation.Title(); ok {
		_spec.SetField(knowledgebase.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := kbc.mutation.Category(); ok {
		_spec.SetField(knowledgebase.FieldCategory, field.TypeEnum, value)
		_node.Category = value
	}
	if value, ok := kbc.mutation.Content(); ok {
		_spec.SetField(knowledgebase.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := kbc.mutation.Tags(); ok {
		_spec.SetField(knowledgebase.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := kbc.mutation.Metadata(); ok {
		_spec.SetField(knowledgebase.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := kbc.mutation.Active(); ok {
		_spec.SetField(knowledgebase.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := kbc.mutation.Priority(); ok {
		_spec.SetField(knowledgebase.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if value, ok := kbc.mutation.CreatedAt(); ok {
		_spec.SetField(knowledgebase.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := kbc.mutation.UpdatedAt(); ok {
		_spec.SetField(knowledgebase.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := kbc.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   knowledgebase.ClinicTable,
			Columns: []string{knowledgebase.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.clinic_knowledge_base = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KnowledgeBaseCreateBulk is the builder for creating many KnowledgeBase entities in bulk.
type KnowledgeBaseCreateBulk struct {
	config
	err      error
	builders []*KnowledgeBaseCreate
}

// Save creates the KnowledgeBase entities in the database.
func (kbcb *KnowledgeBaseCreateBulk) Save(ctx context.Context) ([]*KnowledgeBase, error) {
	if kbcb.err != nil {
		return nil, kbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kbcb.builders))
	nodes := make([]*KnowledgeBase, len(kbcb.builders))
	mutators := make([]Mutator, len(kbcb.builders))
	for i := range kbcb.builders {
		func(i int, root context.Context) {
			builder := kbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KnowledgeBaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kbcb *KnowledgeBaseCreateBulk) SaveX(ctx context.Context) []*KnowledgeBase {
	v, err := kbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kbcb *KnowledgeBaseCreateBulk) Exec(ctx context.Context) error {
	_, err := kbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kbcb *KnowledgeBaseCreateBulk) ExecX(ctx context.Context) {
	if err := kbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
