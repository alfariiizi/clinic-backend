// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/adminauditlog"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/google/uuid"
)

// AdminAuditLogQuery is the builder for querying AdminAuditLog entities.
type AdminAuditLogQuery struct {
	config
	ctx        *QueryContext
	order      []adminauditlog.OrderOption
	inters     []Interceptor
	predicates []predicate.AdminAuditLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminAuditLogQuery builder.
func (aalq *AdminAuditLogQuery) Where(ps ...predicate.AdminAuditLog) *AdminAuditLogQuery {
	aalq.predicates = append(aalq.predicates, ps...)
	return aalq
}

// Limit the number of records to be returned by this query.
func (aalq *AdminAuditLogQuery) Limit(limit int) *AdminAuditLogQuery {
	aalq.ctx.Limit = &limit
	return aalq
}

// Offset to start from.
func (aalq *AdminAuditLogQuery) Offset(offset int) *AdminAuditLogQuery {
	aalq.ctx.Offset = &offset
	return aalq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aalq *AdminAuditLogQuery) Unique(unique bool) *AdminAuditLogQuery {
	aalq.ctx.Unique = &unique
	return aalq
}

// Order specifies how the records should be ordered.
func (aalq *AdminAuditLogQuery) Order(o ...adminauditlog.OrderOption) *AdminAuditLogQuery {
	aalq.order = append(aalq.order, o...)
	return aalq
}

// First returns the first AdminAuditLog entity from the query.
// Returns a *NotFoundError when no AdminAuditLog was found.
func (aalq *AdminAuditLogQuery) First(ctx context.Context) (*AdminAuditLog, error) {
	nodes, err := aalq.Limit(1).All(setContextOp(ctx, aalq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminauditlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) FirstX(ctx context.Context) *AdminAuditLog {
	node, err := aalq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminAuditLog ID from the query.
// Returns a *NotFoundError when no AdminAuditLog ID was found.
func (aalq *AdminAuditLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aalq.Limit(1).IDs(setContextOp(ctx, aalq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminauditlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aalq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminAuditLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminAuditLog entity is found.
// Returns a *NotFoundError when no AdminAuditLog entities are found.
func (aalq *AdminAuditLogQuery) Only(ctx context.Context) (*AdminAuditLog, error) {
	nodes, err := aalq.Limit(2).All(setContextOp(ctx, aalq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminauditlog.Label}
	default:
		return nil, &NotSingularError{adminauditlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) OnlyX(ctx context.Context) *AdminAuditLog {
	node, err := aalq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminAuditLog ID in the query.
// Returns a *NotSingularError when more than one AdminAuditLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (aalq *AdminAuditLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aalq.Limit(2).IDs(setContextOp(ctx, aalq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminauditlog.Label}
	default:
		err = &NotSingularError{adminauditlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aalq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminAuditLogs.
func (aalq *AdminAuditLogQuery) All(ctx context.Context) ([]*AdminAuditLog, error) {
	ctx = setContextOp(ctx, aalq.ctx, ent.OpQueryAll)
	if err := aalq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminAuditLog, *AdminAuditLogQuery]()
	return withInterceptors[[]*AdminAuditLog](ctx, aalq, qr, aalq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) AllX(ctx context.Context) []*AdminAuditLog {
	nodes, err := aalq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminAuditLog IDs.
func (aalq *AdminAuditLogQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aalq.ctx.Unique == nil && aalq.path != nil {
		aalq.Unique(true)
	}
	ctx = setContextOp(ctx, aalq.ctx, ent.OpQueryIDs)
	if err = aalq.Select(adminauditlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aalq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aalq *AdminAuditLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aalq.ctx, ent.OpQueryCount)
	if err := aalq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aalq, querierCount[*AdminAuditLogQuery](), aalq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) CountX(ctx context.Context) int {
	count, err := aalq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aalq *AdminAuditLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aalq.ctx, ent.OpQueryExist)
	switch _, err := aalq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aalq *AdminAuditLogQuery) ExistX(ctx context.Context) bool {
	exist, err := aalq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminAuditLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aalq *AdminAuditLogQuery) Clone() *AdminAuditLogQuery {
	if aalq == nil {
		return nil
	}
	return &AdminAuditLogQuery{
		config:     aalq.config,
		ctx:        aalq.ctx.Clone(),
		order:      append([]adminauditlog.OrderOption{}, aalq.order...),
		inters:     append([]Interceptor{}, aalq.inters...),
		predicates: append([]predicate.AdminAuditLog{}, aalq.predicates...),
		// clone intermediate query.
		sql:  aalq.sql.Clone(),
		path: aalq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserEmail string `json:"user_email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminAuditLog.Query().
//		GroupBy(adminauditlog.FieldUserEmail).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (aalq *AdminAuditLogQuery) GroupBy(field string, fields ...string) *AdminAuditLogGroupBy {
	aalq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminAuditLogGroupBy{build: aalq}
	grbuild.flds = &aalq.ctx.Fields
	grbuild.label = adminauditlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserEmail string `json:"user_email,omitempty"`
//	}
//
//	client.AdminAuditLog.Query().
//		Select(adminauditlog.FieldUserEmail).
//		Scan(ctx, &v)
func (aalq *AdminAuditLogQuery) Select(fields ...string) *AdminAuditLogSelect {
	aalq.ctx.Fields = append(aalq.ctx.Fields, fields...)
	sbuild := &AdminAuditLogSelect{AdminAuditLogQuery: aalq}
	sbuild.label = adminauditlog.Label
	sbuild.flds, sbuild.scan = &aalq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminAuditLogSelect configured with the given aggregations.
func (aalq *AdminAuditLogQuery) Aggregate(fns ...AggregateFunc) *AdminAuditLogSelect {
	return aalq.Select().Aggregate(fns...)
}

func (aalq *AdminAuditLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aalq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aalq); err != nil {
				return err
			}
		}
	}
	for _, f := range aalq.ctx.Fields {
		if !adminauditlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if aalq.path != nil {
		prev, err := aalq.path(ctx)
		if err != nil {
			return err
		}
		aalq.sql = prev
	}
	return nil
}

func (aalq *AdminAuditLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminAuditLog, error) {
	var (
		nodes = []*AdminAuditLog{}
		_spec = aalq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminAuditLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminAuditLog{config: aalq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aalq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aalq *AdminAuditLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aalq.querySpec()
	_spec.Node.Columns = aalq.ctx.Fields
	if len(aalq.ctx.Fields) > 0 {
		_spec.Unique = aalq.ctx.Unique != nil && *aalq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aalq.driver, _spec)
}

func (aalq *AdminAuditLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminauditlog.Table, adminauditlog.Columns, sqlgraph.NewFieldSpec(adminauditlog.FieldID, field.TypeUUID))
	_spec.From = aalq.sql
	if unique := aalq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aalq.path != nil {
		_spec.Unique = true
	}
	if fields := aalq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminauditlog.FieldID)
		for i := range fields {
			if fields[i] != adminauditlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aalq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aalq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aalq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aalq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aalq *AdminAuditLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aalq.driver.Dialect())
	t1 := builder.Table(adminauditlog.Table)
	columns := aalq.ctx.Fields
	if len(columns) == 0 {
		columns = adminauditlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aalq.sql != nil {
		selector = aalq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aalq.ctx.Unique != nil && *aalq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aalq.predicates {
		p(selector)
	}
	for _, p := range aalq.order {
		p(selector)
	}
	if offset := aalq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aalq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminAuditLogGroupBy is the group-by builder for AdminAuditLog entities.
type AdminAuditLogGroupBy struct {
	selector
	build *AdminAuditLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aalgb *AdminAuditLogGroupBy) Aggregate(fns ...AggregateFunc) *AdminAuditLogGroupBy {
	aalgb.fns = append(aalgb.fns, fns...)
	return aalgb
}

// Scan applies the selector query and scans the result into the given value.
func (aalgb *AdminAuditLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aalgb.build.ctx, ent.OpQueryGroupBy)
	if err := aalgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminAuditLogQuery, *AdminAuditLogGroupBy](ctx, aalgb.build, aalgb, aalgb.build.inters, v)
}

func (aalgb *AdminAuditLogGroupBy) sqlScan(ctx context.Context, root *AdminAuditLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aalgb.fns))
	for _, fn := range aalgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aalgb.flds)+len(aalgb.fns))
		for _, f := range *aalgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aalgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aalgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminAuditLogSelect is the builder for selecting fields of AdminAuditLog entities.
type AdminAuditLogSelect struct {
	*AdminAuditLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aals *AdminAuditLogSelect) Aggregate(fns ...AggregateFunc) *AdminAuditLogSelect {
	aals.fns = append(aals.fns, fns...)
	return aals
}

// Scan applies the selector query and scans the result into the given value.
func (aals *AdminAuditLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aals.ctx, ent.OpQuerySelect)
	if err := aals.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminAuditLogQuery, *AdminAuditLogSelect](ctx, aals.AdminAuditLogQuery, aals, aals.inters, v)
}

func (aals *AdminAuditLogSelect) sqlScan(ctx context.Context, root *AdminAuditLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aals.fns))
	for _, fn := range aals.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aals.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aals.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
