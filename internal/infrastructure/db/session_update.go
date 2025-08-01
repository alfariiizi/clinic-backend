// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/session"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIPAddress sets the "ip_address" field.
func (su *SessionUpdate) SetIPAddress(s string) *SessionUpdate {
	su.mutation.SetIPAddress(s)
	return su
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (su *SessionUpdate) SetNillableIPAddress(s *string) *SessionUpdate {
	if s != nil {
		su.SetIPAddress(*s)
	}
	return su
}

// ClearIPAddress clears the value of the "ip_address" field.
func (su *SessionUpdate) ClearIPAddress() *SessionUpdate {
	su.mutation.ClearIPAddress()
	return su
}

// SetUserAgent sets the "user_agent" field.
func (su *SessionUpdate) SetUserAgent(s string) *SessionUpdate {
	su.mutation.SetUserAgent(s)
	return su
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserAgent(s *string) *SessionUpdate {
	if s != nil {
		su.SetUserAgent(*s)
	}
	return su
}

// ClearUserAgent clears the value of the "user_agent" field.
func (su *SessionUpdate) ClearUserAgent() *SessionUpdate {
	su.mutation.ClearUserAgent()
	return su
}

// SetDeviceID sets the "device_id" field.
func (su *SessionUpdate) SetDeviceID(s string) *SessionUpdate {
	su.mutation.SetDeviceID(s)
	return su
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableDeviceID(s *string) *SessionUpdate {
	if s != nil {
		su.SetDeviceID(*s)
	}
	return su
}

// ClearDeviceID clears the value of the "device_id" field.
func (su *SessionUpdate) ClearDeviceID() *SessionUpdate {
	su.mutation.ClearDeviceID()
	return su
}

// SetNumberOfUses sets the "number_of_uses" field.
func (su *SessionUpdate) SetNumberOfUses(u uint64) *SessionUpdate {
	su.mutation.ResetNumberOfUses()
	su.mutation.SetNumberOfUses(u)
	return su
}

// SetNillableNumberOfUses sets the "number_of_uses" field if the given value is not nil.
func (su *SessionUpdate) SetNillableNumberOfUses(u *uint64) *SessionUpdate {
	if u != nil {
		su.SetNumberOfUses(*u)
	}
	return su
}

// AddNumberOfUses adds u to the "number_of_uses" field.
func (su *SessionUpdate) AddNumberOfUses(u int64) *SessionUpdate {
	su.mutation.AddNumberOfUses(u)
	return su
}

// SetLastUsedAt sets the "last_used_at" field.
func (su *SessionUpdate) SetLastUsedAt(t time.Time) *SessionUpdate {
	su.mutation.SetLastUsedAt(t)
	return su
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableLastUsedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetLastUsedAt(*t)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetRevokedAt sets the "revoked_at" field.
func (su *SessionUpdate) SetRevokedAt(t time.Time) *SessionUpdate {
	su.mutation.SetRevokedAt(t)
	return su
}

// SetNillableRevokedAt sets the "revoked_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableRevokedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetRevokedAt(*t)
	}
	return su
}

// ClearRevokedAt clears the value of the "revoked_at" field.
func (su *SessionUpdate) ClearRevokedAt() *SessionUpdate {
	su.mutation.ClearRevokedAt()
	return su
}

// SetUserID sets the "user_id" field.
func (su *SessionUpdate) SetUserID(u uuid.UUID) *SessionUpdate {
	su.mutation.SetUserID(u)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(u *uuid.UUID) *SessionUpdate {
	if u != nil {
		su.SetUserID(*u)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SessionUpdate) check() error {
	if v, ok := su.mutation.NumberOfUses(); ok {
		if err := session.NumberOfUsesValidator(v); err != nil {
			return &ValidationError{Name: "number_of_uses", err: fmt.Errorf(`db: validator failed for field "Session.number_of_uses": %w`, err)}
		}
	}
	if su.mutation.UserCleared() && len(su.mutation.UserIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "Session.user"`)
	}
	return nil
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.IPAddress(); ok {
		_spec.SetField(session.FieldIPAddress, field.TypeString, value)
	}
	if su.mutation.IPAddressCleared() {
		_spec.ClearField(session.FieldIPAddress, field.TypeString)
	}
	if value, ok := su.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if su.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := su.mutation.DeviceID(); ok {
		_spec.SetField(session.FieldDeviceID, field.TypeString, value)
	}
	if su.mutation.DeviceIDCleared() {
		_spec.ClearField(session.FieldDeviceID, field.TypeString)
	}
	if value, ok := su.mutation.NumberOfUses(); ok {
		_spec.SetField(session.FieldNumberOfUses, field.TypeUint64, value)
	}
	if value, ok := su.mutation.AddedNumberOfUses(); ok {
		_spec.AddField(session.FieldNumberOfUses, field.TypeUint64, value)
	}
	if value, ok := su.mutation.LastUsedAt(); ok {
		_spec.SetField(session.FieldLastUsedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.RevokedAt(); ok {
		_spec.SetField(session.FieldRevokedAt, field.TypeTime, value)
	}
	if su.mutation.RevokedAtCleared() {
		_spec.ClearField(session.FieldRevokedAt, field.TypeTime)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetIPAddress sets the "ip_address" field.
func (suo *SessionUpdateOne) SetIPAddress(s string) *SessionUpdateOne {
	suo.mutation.SetIPAddress(s)
	return suo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableIPAddress(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetIPAddress(*s)
	}
	return suo
}

// ClearIPAddress clears the value of the "ip_address" field.
func (suo *SessionUpdateOne) ClearIPAddress() *SessionUpdateOne {
	suo.mutation.ClearIPAddress()
	return suo
}

// SetUserAgent sets the "user_agent" field.
func (suo *SessionUpdateOne) SetUserAgent(s string) *SessionUpdateOne {
	suo.mutation.SetUserAgent(s)
	return suo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserAgent(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUserAgent(*s)
	}
	return suo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (suo *SessionUpdateOne) ClearUserAgent() *SessionUpdateOne {
	suo.mutation.ClearUserAgent()
	return suo
}

// SetDeviceID sets the "device_id" field.
func (suo *SessionUpdateOne) SetDeviceID(s string) *SessionUpdateOne {
	suo.mutation.SetDeviceID(s)
	return suo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableDeviceID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetDeviceID(*s)
	}
	return suo
}

// ClearDeviceID clears the value of the "device_id" field.
func (suo *SessionUpdateOne) ClearDeviceID() *SessionUpdateOne {
	suo.mutation.ClearDeviceID()
	return suo
}

// SetNumberOfUses sets the "number_of_uses" field.
func (suo *SessionUpdateOne) SetNumberOfUses(u uint64) *SessionUpdateOne {
	suo.mutation.ResetNumberOfUses()
	suo.mutation.SetNumberOfUses(u)
	return suo
}

// SetNillableNumberOfUses sets the "number_of_uses" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableNumberOfUses(u *uint64) *SessionUpdateOne {
	if u != nil {
		suo.SetNumberOfUses(*u)
	}
	return suo
}

// AddNumberOfUses adds u to the "number_of_uses" field.
func (suo *SessionUpdateOne) AddNumberOfUses(u int64) *SessionUpdateOne {
	suo.mutation.AddNumberOfUses(u)
	return suo
}

// SetLastUsedAt sets the "last_used_at" field.
func (suo *SessionUpdateOne) SetLastUsedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetLastUsedAt(t)
	return suo
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableLastUsedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetLastUsedAt(*t)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetRevokedAt sets the "revoked_at" field.
func (suo *SessionUpdateOne) SetRevokedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetRevokedAt(t)
	return suo
}

// SetNillableRevokedAt sets the "revoked_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableRevokedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetRevokedAt(*t)
	}
	return suo
}

// ClearRevokedAt clears the value of the "revoked_at" field.
func (suo *SessionUpdateOne) ClearRevokedAt() *SessionUpdateOne {
	suo.mutation.ClearRevokedAt()
	return suo
}

// SetUserID sets the "user_id" field.
func (suo *SessionUpdateOne) SetUserID(u uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetUserID(u)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(u *uuid.UUID) *SessionUpdateOne {
	if u != nil {
		suo.SetUserID(*u)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SessionUpdateOne) check() error {
	if v, ok := suo.mutation.NumberOfUses(); ok {
		if err := session.NumberOfUsesValidator(v); err != nil {
			return &ValidationError{Name: "number_of_uses", err: fmt.Errorf(`db: validator failed for field "Session.number_of_uses": %w`, err)}
		}
	}
	if suo.mutation.UserCleared() && len(suo.mutation.UserIDs()) > 0 {
		return errors.New(`db: clearing a required unique edge "Session.user"`)
	}
	return nil
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.IPAddress(); ok {
		_spec.SetField(session.FieldIPAddress, field.TypeString, value)
	}
	if suo.mutation.IPAddressCleared() {
		_spec.ClearField(session.FieldIPAddress, field.TypeString)
	}
	if value, ok := suo.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if suo.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := suo.mutation.DeviceID(); ok {
		_spec.SetField(session.FieldDeviceID, field.TypeString, value)
	}
	if suo.mutation.DeviceIDCleared() {
		_spec.ClearField(session.FieldDeviceID, field.TypeString)
	}
	if value, ok := suo.mutation.NumberOfUses(); ok {
		_spec.SetField(session.FieldNumberOfUses, field.TypeUint64, value)
	}
	if value, ok := suo.mutation.AddedNumberOfUses(); ok {
		_spec.AddField(session.FieldNumberOfUses, field.TypeUint64, value)
	}
	if value, ok := suo.mutation.LastUsedAt(); ok {
		_spec.SetField(session.FieldLastUsedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.RevokedAt(); ok {
		_spec.SetField(session.FieldRevokedAt, field.TypeTime, value)
	}
	if suo.mutation.RevokedAtCleared() {
		_spec.ClearField(session.FieldRevokedAt, field.TypeTime)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
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
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
