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
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/queueentry"
	"github.com/google/uuid"
)

// QueueEntryQuery is the builder for querying QueueEntry entities.
type QueueEntryQuery struct {
	config
	ctx        *QueryContext
	order      []queueentry.OrderOption
	inters     []Interceptor
	predicates []predicate.QueueEntry
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QueueEntryQuery builder.
func (qeq *QueueEntryQuery) Where(ps ...predicate.QueueEntry) *QueueEntryQuery {
	qeq.predicates = append(qeq.predicates, ps...)
	return qeq
}

// Limit the number of records to be returned by this query.
func (qeq *QueueEntryQuery) Limit(limit int) *QueueEntryQuery {
	qeq.ctx.Limit = &limit
	return qeq
}

// Offset to start from.
func (qeq *QueueEntryQuery) Offset(offset int) *QueueEntryQuery {
	qeq.ctx.Offset = &offset
	return qeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qeq *QueueEntryQuery) Unique(unique bool) *QueueEntryQuery {
	qeq.ctx.Unique = &unique
	return qeq
}

// Order specifies how the records should be ordered.
func (qeq *QueueEntryQuery) Order(o ...queueentry.OrderOption) *QueueEntryQuery {
	qeq.order = append(qeq.order, o...)
	return qeq
}

// First returns the first QueueEntry entity from the query.
// Returns a *NotFoundError when no QueueEntry was found.
func (qeq *QueueEntryQuery) First(ctx context.Context) (*QueueEntry, error) {
	nodes, err := qeq.Limit(1).All(setContextOp(ctx, qeq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{queueentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qeq *QueueEntryQuery) FirstX(ctx context.Context) *QueueEntry {
	node, err := qeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first QueueEntry ID from the query.
// Returns a *NotFoundError when no QueueEntry ID was found.
func (qeq *QueueEntryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qeq.Limit(1).IDs(setContextOp(ctx, qeq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{queueentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qeq *QueueEntryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := qeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single QueueEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one QueueEntry entity is found.
// Returns a *NotFoundError when no QueueEntry entities are found.
func (qeq *QueueEntryQuery) Only(ctx context.Context) (*QueueEntry, error) {
	nodes, err := qeq.Limit(2).All(setContextOp(ctx, qeq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{queueentry.Label}
	default:
		return nil, &NotSingularError{queueentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qeq *QueueEntryQuery) OnlyX(ctx context.Context) *QueueEntry {
	node, err := qeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only QueueEntry ID in the query.
// Returns a *NotSingularError when more than one QueueEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (qeq *QueueEntryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qeq.Limit(2).IDs(setContextOp(ctx, qeq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{queueentry.Label}
	default:
		err = &NotSingularError{queueentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qeq *QueueEntryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := qeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QueueEntries.
func (qeq *QueueEntryQuery) All(ctx context.Context) ([]*QueueEntry, error) {
	ctx = setContextOp(ctx, qeq.ctx, ent.OpQueryAll)
	if err := qeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*QueueEntry, *QueueEntryQuery]()
	return withInterceptors[[]*QueueEntry](ctx, qeq, qr, qeq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qeq *QueueEntryQuery) AllX(ctx context.Context) []*QueueEntry {
	nodes, err := qeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of QueueEntry IDs.
func (qeq *QueueEntryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if qeq.ctx.Unique == nil && qeq.path != nil {
		qeq.Unique(true)
	}
	ctx = setContextOp(ctx, qeq.ctx, ent.OpQueryIDs)
	if err = qeq.Select(queueentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qeq *QueueEntryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := qeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qeq *QueueEntryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qeq.ctx, ent.OpQueryCount)
	if err := qeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qeq, querierCount[*QueueEntryQuery](), qeq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qeq *QueueEntryQuery) CountX(ctx context.Context) int {
	count, err := qeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qeq *QueueEntryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qeq.ctx, ent.OpQueryExist)
	switch _, err := qeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qeq *QueueEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := qeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QueueEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qeq *QueueEntryQuery) Clone() *QueueEntryQuery {
	if qeq == nil {
		return nil
	}
	return &QueueEntryQuery{
		config:     qeq.config,
		ctx:        qeq.ctx.Clone(),
		order:      append([]queueentry.OrderOption{}, qeq.order...),
		inters:     append([]Interceptor{}, qeq.inters...),
		predicates: append([]predicate.QueueEntry{}, qeq.predicates...),
		// clone intermediate query.
		sql:  qeq.sql.Clone(),
		path: qeq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClinicID string `json:"clinic_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QueueEntry.Query().
//		GroupBy(queueentry.FieldClinicID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (qeq *QueueEntryQuery) GroupBy(field string, fields ...string) *QueueEntryGroupBy {
	qeq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QueueEntryGroupBy{build: qeq}
	grbuild.flds = &qeq.ctx.Fields
	grbuild.label = queueentry.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClinicID string `json:"clinic_id,omitempty"`
//	}
//
//	client.QueueEntry.Query().
//		Select(queueentry.FieldClinicID).
//		Scan(ctx, &v)
func (qeq *QueueEntryQuery) Select(fields ...string) *QueueEntrySelect {
	qeq.ctx.Fields = append(qeq.ctx.Fields, fields...)
	sbuild := &QueueEntrySelect{QueueEntryQuery: qeq}
	sbuild.label = queueentry.Label
	sbuild.flds, sbuild.scan = &qeq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QueueEntrySelect configured with the given aggregations.
func (qeq *QueueEntryQuery) Aggregate(fns ...AggregateFunc) *QueueEntrySelect {
	return qeq.Select().Aggregate(fns...)
}

func (qeq *QueueEntryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qeq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qeq); err != nil {
				return err
			}
		}
	}
	for _, f := range qeq.ctx.Fields {
		if !queueentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if qeq.path != nil {
		prev, err := qeq.path(ctx)
		if err != nil {
			return err
		}
		qeq.sql = prev
	}
	return nil
}

func (qeq *QueueEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*QueueEntry, error) {
	var (
		nodes = []*QueueEntry{}
		_spec = qeq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*QueueEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &QueueEntry{config: qeq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (qeq *QueueEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qeq.querySpec()
	_spec.Node.Columns = qeq.ctx.Fields
	if len(qeq.ctx.Fields) > 0 {
		_spec.Unique = qeq.ctx.Unique != nil && *qeq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qeq.driver, _spec)
}

func (qeq *QueueEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(queueentry.Table, queueentry.Columns, sqlgraph.NewFieldSpec(queueentry.FieldID, field.TypeUUID))
	_spec.From = qeq.sql
	if unique := qeq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qeq.path != nil {
		_spec.Unique = true
	}
	if fields := qeq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, queueentry.FieldID)
		for i := range fields {
			if fields[i] != queueentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qeq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qeq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qeq *QueueEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qeq.driver.Dialect())
	t1 := builder.Table(queueentry.Table)
	columns := qeq.ctx.Fields
	if len(columns) == 0 {
		columns = queueentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qeq.sql != nil {
		selector = qeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qeq.ctx.Unique != nil && *qeq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qeq.predicates {
		p(selector)
	}
	for _, p := range qeq.order {
		p(selector)
	}
	if offset := qeq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qeq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QueueEntryGroupBy is the group-by builder for QueueEntry entities.
type QueueEntryGroupBy struct {
	selector
	build *QueueEntryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qegb *QueueEntryGroupBy) Aggregate(fns ...AggregateFunc) *QueueEntryGroupBy {
	qegb.fns = append(qegb.fns, fns...)
	return qegb
}

// Scan applies the selector query and scans the result into the given value.
func (qegb *QueueEntryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qegb.build.ctx, ent.OpQueryGroupBy)
	if err := qegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QueueEntryQuery, *QueueEntryGroupBy](ctx, qegb.build, qegb, qegb.build.inters, v)
}

func (qegb *QueueEntryGroupBy) sqlScan(ctx context.Context, root *QueueEntryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qegb.fns))
	for _, fn := range qegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qegb.flds)+len(qegb.fns))
		for _, f := range *qegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QueueEntrySelect is the builder for selecting fields of QueueEntry entities.
type QueueEntrySelect struct {
	*QueueEntryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qes *QueueEntrySelect) Aggregate(fns ...AggregateFunc) *QueueEntrySelect {
	qes.fns = append(qes.fns, fns...)
	return qes
}

// Scan applies the selector query and scans the result into the given value.
func (qes *QueueEntrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qes.ctx, ent.OpQuerySelect)
	if err := qes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QueueEntryQuery, *QueueEntrySelect](ctx, qes.QueueEntryQuery, qes, qes.inters, v)
}

func (qes *QueueEntrySelect) sqlScan(ctx context.Context, root *QueueEntryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qes.fns))
	for _, fn := range qes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
