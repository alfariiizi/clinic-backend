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
	"github.com/alfariiizi/vandor/internal/infrastructure/db/chatmessage"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/chatthread"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/google/uuid"
)

// ChatMessageQuery is the builder for querying ChatMessage entities.
type ChatMessageQuery struct {
	config
	ctx        *QueryContext
	order      []chatmessage.OrderOption
	inters     []Interceptor
	predicates []predicate.ChatMessage
	withThread *ChatThreadQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatMessageQuery builder.
func (cmq *ChatMessageQuery) Where(ps ...predicate.ChatMessage) *ChatMessageQuery {
	cmq.predicates = append(cmq.predicates, ps...)
	return cmq
}

// Limit the number of records to be returned by this query.
func (cmq *ChatMessageQuery) Limit(limit int) *ChatMessageQuery {
	cmq.ctx.Limit = &limit
	return cmq
}

// Offset to start from.
func (cmq *ChatMessageQuery) Offset(offset int) *ChatMessageQuery {
	cmq.ctx.Offset = &offset
	return cmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmq *ChatMessageQuery) Unique(unique bool) *ChatMessageQuery {
	cmq.ctx.Unique = &unique
	return cmq
}

// Order specifies how the records should be ordered.
func (cmq *ChatMessageQuery) Order(o ...chatmessage.OrderOption) *ChatMessageQuery {
	cmq.order = append(cmq.order, o...)
	return cmq
}

// QueryThread chains the current query on the "thread" edge.
func (cmq *ChatMessageQuery) QueryThread() *ChatThreadQuery {
	query := (&ChatThreadClient{config: cmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatmessage.Table, chatmessage.FieldID, selector),
			sqlgraph.To(chatthread.Table, chatthread.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chatmessage.ThreadTable, chatmessage.ThreadColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChatMessage entity from the query.
// Returns a *NotFoundError when no ChatMessage was found.
func (cmq *ChatMessageQuery) First(ctx context.Context) (*ChatMessage, error) {
	nodes, err := cmq.Limit(1).All(setContextOp(ctx, cmq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmq *ChatMessageQuery) FirstX(ctx context.Context) *ChatMessage {
	node, err := cmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatMessage ID from the query.
// Returns a *NotFoundError when no ChatMessage ID was found.
func (cmq *ChatMessageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmq.Limit(1).IDs(setContextOp(ctx, cmq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmq *ChatMessageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChatMessage entity is found.
// Returns a *NotFoundError when no ChatMessage entities are found.
func (cmq *ChatMessageQuery) Only(ctx context.Context) (*ChatMessage, error) {
	nodes, err := cmq.Limit(2).All(setContextOp(ctx, cmq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatmessage.Label}
	default:
		return nil, &NotSingularError{chatmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmq *ChatMessageQuery) OnlyX(ctx context.Context) *ChatMessage {
	node, err := cmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatMessage ID in the query.
// Returns a *NotSingularError when more than one ChatMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (cmq *ChatMessageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmq.Limit(2).IDs(setContextOp(ctx, cmq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatmessage.Label}
	default:
		err = &NotSingularError{chatmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmq *ChatMessageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatMessages.
func (cmq *ChatMessageQuery) All(ctx context.Context) ([]*ChatMessage, error) {
	ctx = setContextOp(ctx, cmq.ctx, ent.OpQueryAll)
	if err := cmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChatMessage, *ChatMessageQuery]()
	return withInterceptors[[]*ChatMessage](ctx, cmq, qr, cmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cmq *ChatMessageQuery) AllX(ctx context.Context) []*ChatMessage {
	nodes, err := cmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatMessage IDs.
func (cmq *ChatMessageQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cmq.ctx.Unique == nil && cmq.path != nil {
		cmq.Unique(true)
	}
	ctx = setContextOp(ctx, cmq.ctx, ent.OpQueryIDs)
	if err = cmq.Select(chatmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmq *ChatMessageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmq *ChatMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cmq.ctx, ent.OpQueryCount)
	if err := cmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cmq, querierCount[*ChatMessageQuery](), cmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cmq *ChatMessageQuery) CountX(ctx context.Context) int {
	count, err := cmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmq *ChatMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cmq.ctx, ent.OpQueryExist)
	switch _, err := cmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cmq *ChatMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := cmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmq *ChatMessageQuery) Clone() *ChatMessageQuery {
	if cmq == nil {
		return nil
	}
	return &ChatMessageQuery{
		config:     cmq.config,
		ctx:        cmq.ctx.Clone(),
		order:      append([]chatmessage.OrderOption{}, cmq.order...),
		inters:     append([]Interceptor{}, cmq.inters...),
		predicates: append([]predicate.ChatMessage{}, cmq.predicates...),
		withThread: cmq.withThread.Clone(),
		// clone intermediate query.
		sql:  cmq.sql.Clone(),
		path: cmq.path,
	}
}

// WithThread tells the query-builder to eager-load the nodes that are connected to
// the "thread" edge. The optional arguments are used to configure the query builder of the edge.
func (cmq *ChatMessageQuery) WithThread(opts ...func(*ChatThreadQuery)) *ChatMessageQuery {
	query := (&ChatThreadClient{config: cmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmq.withThread = query
	return cmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WhatsappMessageID string `json:"whatsapp_message_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatMessage.Query().
//		GroupBy(chatmessage.FieldWhatsappMessageID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (cmq *ChatMessageQuery) GroupBy(field string, fields ...string) *ChatMessageGroupBy {
	cmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChatMessageGroupBy{build: cmq}
	grbuild.flds = &cmq.ctx.Fields
	grbuild.label = chatmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WhatsappMessageID string `json:"whatsapp_message_id,omitempty"`
//	}
//
//	client.ChatMessage.Query().
//		Select(chatmessage.FieldWhatsappMessageID).
//		Scan(ctx, &v)
func (cmq *ChatMessageQuery) Select(fields ...string) *ChatMessageSelect {
	cmq.ctx.Fields = append(cmq.ctx.Fields, fields...)
	sbuild := &ChatMessageSelect{ChatMessageQuery: cmq}
	sbuild.label = chatmessage.Label
	sbuild.flds, sbuild.scan = &cmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChatMessageSelect configured with the given aggregations.
func (cmq *ChatMessageQuery) Aggregate(fns ...AggregateFunc) *ChatMessageSelect {
	return cmq.Select().Aggregate(fns...)
}

func (cmq *ChatMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cmq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cmq); err != nil {
				return err
			}
		}
	}
	for _, f := range cmq.ctx.Fields {
		if !chatmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if cmq.path != nil {
		prev, err := cmq.path(ctx)
		if err != nil {
			return err
		}
		cmq.sql = prev
	}
	return nil
}

func (cmq *ChatMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChatMessage, error) {
	var (
		nodes       = []*ChatMessage{}
		withFKs     = cmq.withFKs
		_spec       = cmq.querySpec()
		loadedTypes = [1]bool{
			cmq.withThread != nil,
		}
	)
	if cmq.withThread != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, chatmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChatMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChatMessage{config: cmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cmq.withThread; query != nil {
		if err := cmq.loadThread(ctx, query, nodes, nil,
			func(n *ChatMessage, e *ChatThread) { n.Edges.Thread = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cmq *ChatMessageQuery) loadThread(ctx context.Context, query *ChatThreadQuery, nodes []*ChatMessage, init func(*ChatMessage), assign func(*ChatMessage, *ChatThread)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ChatMessage)
	for i := range nodes {
		if nodes[i].chat_thread_messages == nil {
			continue
		}
		fk := *nodes[i].chat_thread_messages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(chatthread.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chat_thread_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cmq *ChatMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmq.querySpec()
	_spec.Node.Columns = cmq.ctx.Fields
	if len(cmq.ctx.Fields) > 0 {
		_spec.Unique = cmq.ctx.Unique != nil && *cmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cmq.driver, _spec)
}

func (cmq *ChatMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chatmessage.Table, chatmessage.Columns, sqlgraph.NewFieldSpec(chatmessage.FieldID, field.TypeUUID))
	_spec.From = cmq.sql
	if unique := cmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cmq.path != nil {
		_spec.Unique = true
	}
	if fields := cmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatmessage.FieldID)
		for i := range fields {
			if fields[i] != chatmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmq *ChatMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmq.driver.Dialect())
	t1 := builder.Table(chatmessage.Table)
	columns := cmq.ctx.Fields
	if len(columns) == 0 {
		columns = chatmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmq.sql != nil {
		selector = cmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cmq.ctx.Unique != nil && *cmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cmq.predicates {
		p(selector)
	}
	for _, p := range cmq.order {
		p(selector)
	}
	if offset := cmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatMessageGroupBy is the group-by builder for ChatMessage entities.
type ChatMessageGroupBy struct {
	selector
	build *ChatMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmgb *ChatMessageGroupBy) Aggregate(fns ...AggregateFunc) *ChatMessageGroupBy {
	cmgb.fns = append(cmgb.fns, fns...)
	return cmgb
}

// Scan applies the selector query and scans the result into the given value.
func (cmgb *ChatMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cmgb.build.ctx, ent.OpQueryGroupBy)
	if err := cmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatMessageQuery, *ChatMessageGroupBy](ctx, cmgb.build, cmgb, cmgb.build.inters, v)
}

func (cmgb *ChatMessageGroupBy) sqlScan(ctx context.Context, root *ChatMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cmgb.fns))
	for _, fn := range cmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cmgb.flds)+len(cmgb.fns))
		for _, f := range *cmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChatMessageSelect is the builder for selecting fields of ChatMessage entities.
type ChatMessageSelect struct {
	*ChatMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cms *ChatMessageSelect) Aggregate(fns ...AggregateFunc) *ChatMessageSelect {
	cms.fns = append(cms.fns, fns...)
	return cms
}

// Scan applies the selector query and scans the result into the given value.
func (cms *ChatMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cms.ctx, ent.OpQuerySelect)
	if err := cms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatMessageQuery, *ChatMessageSelect](ctx, cms.ChatMessageQuery, cms, cms.inters, v)
}

func (cms *ChatMessageSelect) sqlScan(ctx context.Context, root *ChatMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cms.fns))
	for _, fn := range cms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
