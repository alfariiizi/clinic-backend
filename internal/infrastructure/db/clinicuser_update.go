// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinicuser"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/google/uuid"
)

// ClinicUserUpdate is the builder for updating ClinicUser entities.
type ClinicUserUpdate struct {
	config
	hooks    []Hook
	mutation *ClinicUserMutation
}

// Where appends a list predicates to the ClinicUserUpdate builder.
func (cuu *ClinicUserUpdate) Where(ps ...predicate.ClinicUser) *ClinicUserUpdate {
	cuu.mutation.Where(ps...)
	return cuu
}

// SetClinicID sets the "clinic_id" field.
func (cuu *ClinicUserUpdate) SetClinicID(u uuid.UUID) *ClinicUserUpdate {
	cuu.mutation.SetClinicID(u)
	return cuu
}

// SetNillableClinicID sets the "clinic_id" field if the given value is not nil.
func (cuu *ClinicUserUpdate) SetNillableClinicID(u *uuid.UUID) *ClinicUserUpdate {
	if u != nil {
		cuu.SetClinicID(*u)
	}
	return cuu
}

// SetUserID sets the "user_id" field.
func (cuu *ClinicUserUpdate) SetUserID(u uuid.UUID) *ClinicUserUpdate {
	cuu.mutation.SetUserID(u)
	return cuu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuu *ClinicUserUpdate) SetNillableUserID(u *uuid.UUID) *ClinicUserUpdate {
	if u != nil {
		cuu.SetUserID(*u)
	}
	return cuu
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (cuu *ClinicUserUpdate) SetClinic(c *Clinic) *ClinicUserUpdate {
	return cuu.SetClinicID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cuu *ClinicUserUpdate) SetUser(u *User) *ClinicUserUpdate {
	return cuu.SetUserID(u.ID)
}

// Mutation returns the ClinicUserMutation object of the builder.
func (cuu *ClinicUserUpdate) Mutation() *ClinicUserMutation {
	return cuu.mutation
}

// ClearClinic clears the "clinic" edge to the Clinic entity.
func (cuu *ClinicUserUpdate) ClearClinic() *ClinicUserUpdate {
	cuu.mutation.ClearClinic()
	return cuu
}

// ClearUser clears the "user" edge to the User entity.
func (cuu *ClinicUserUpdate) ClearUser() *ClinicUserUpdate {
	cuu.mutation.ClearUser()
	return cuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cuu *ClinicUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cuu.sqlSave, cuu.mutation, cuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuu *ClinicUserUpdate) SaveX(ctx context.Context) int {
	affected, err := cuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cuu *ClinicUserUpdate) Exec(ctx context.Context) error {
	_, err := cuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuu *ClinicUserUpdate) ExecX(ctx context.Context) {
	if err := cuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuu *ClinicUserUpdate) check() error {
	if cuu.mutation.ClinicCleared() && len(cuu.mutation.ClinicIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "ClinicUser.clinic"`)
	}
	if cuu.mutation.UserCleared() && len(cuu.mutation.UserIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "ClinicUser.user"`)
	}
	return nil
}

func (cuu *ClinicUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cuu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(clinicuser.Table, clinicuser.Columns, sqlgraph.NewFieldSpec(clinicuser.FieldID, field.TypeUUID))
	if ps := cuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuu.mutation.ClinicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.ClinicTable,
			Columns: []string{clinicuser.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.ClinicTable,
			Columns: []string{clinicuser.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.UserTable,
			Columns: []string{clinicuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.UserTable,
			Columns: []string{clinicuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cuu.mutation.done = true
	return n, nil
}

// ClinicUserUpdateOne is the builder for updating a single ClinicUser entity.
type ClinicUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClinicUserMutation
}

// SetClinicID sets the "clinic_id" field.
func (cuuo *ClinicUserUpdateOne) SetClinicID(u uuid.UUID) *ClinicUserUpdateOne {
	cuuo.mutation.SetClinicID(u)
	return cuuo
}

// SetNillableClinicID sets the "clinic_id" field if the given value is not nil.
func (cuuo *ClinicUserUpdateOne) SetNillableClinicID(u *uuid.UUID) *ClinicUserUpdateOne {
	if u != nil {
		cuuo.SetClinicID(*u)
	}
	return cuuo
}

// SetUserID sets the "user_id" field.
func (cuuo *ClinicUserUpdateOne) SetUserID(u uuid.UUID) *ClinicUserUpdateOne {
	cuuo.mutation.SetUserID(u)
	return cuuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuuo *ClinicUserUpdateOne) SetNillableUserID(u *uuid.UUID) *ClinicUserUpdateOne {
	if u != nil {
		cuuo.SetUserID(*u)
	}
	return cuuo
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (cuuo *ClinicUserUpdateOne) SetClinic(c *Clinic) *ClinicUserUpdateOne {
	return cuuo.SetClinicID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cuuo *ClinicUserUpdateOne) SetUser(u *User) *ClinicUserUpdateOne {
	return cuuo.SetUserID(u.ID)
}

// Mutation returns the ClinicUserMutation object of the builder.
func (cuuo *ClinicUserUpdateOne) Mutation() *ClinicUserMutation {
	return cuuo.mutation
}

// ClearClinic clears the "clinic" edge to the Clinic entity.
func (cuuo *ClinicUserUpdateOne) ClearClinic() *ClinicUserUpdateOne {
	cuuo.mutation.ClearClinic()
	return cuuo
}

// ClearUser clears the "user" edge to the User entity.
func (cuuo *ClinicUserUpdateOne) ClearUser() *ClinicUserUpdateOne {
	cuuo.mutation.ClearUser()
	return cuuo
}

// Where appends a list predicates to the ClinicUserUpdate builder.
func (cuuo *ClinicUserUpdateOne) Where(ps ...predicate.ClinicUser) *ClinicUserUpdateOne {
	cuuo.mutation.Where(ps...)
	return cuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuuo *ClinicUserUpdateOne) Select(field string, fields ...string) *ClinicUserUpdateOne {
	cuuo.fields = append([]string{field}, fields...)
	return cuuo
}

// Save executes the query and returns the updated ClinicUser entity.
func (cuuo *ClinicUserUpdateOne) Save(ctx context.Context) (*ClinicUser, error) {
	return withHooks(ctx, cuuo.sqlSave, cuuo.mutation, cuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuuo *ClinicUserUpdateOne) SaveX(ctx context.Context) *ClinicUser {
	node, err := cuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuuo *ClinicUserUpdateOne) Exec(ctx context.Context) error {
	_, err := cuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuuo *ClinicUserUpdateOne) ExecX(ctx context.Context) {
	if err := cuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuuo *ClinicUserUpdateOne) check() error {
	if cuuo.mutation.ClinicCleared() && len(cuuo.mutation.ClinicIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "ClinicUser.clinic"`)
	}
	if cuuo.mutation.UserCleared() && len(cuuo.mutation.UserIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "ClinicUser.user"`)
	}
	return nil
}

func (cuuo *ClinicUserUpdateOne) sqlSave(ctx context.Context) (_node *ClinicUser, err error) {
	if err := cuuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(clinicuser.Table, clinicuser.Columns, sqlgraph.NewFieldSpec(clinicuser.FieldID, field.TypeUUID))
	id, ok := cuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "ClinicUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clinicuser.FieldID)
		for _, f := range fields {
			if !clinicuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != clinicuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuuo.mutation.ClinicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.ClinicTable,
			Columns: []string{clinicuser.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.ClinicTable,
			Columns: []string{clinicuser.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.UserTable,
			Columns: []string{clinicuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clinicuser.UserTable,
			Columns: []string{clinicuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClinicUser{config: cuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clinicuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuuo.mutation.done = true
	return _node, nil
}
