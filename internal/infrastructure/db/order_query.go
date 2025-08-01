// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/order"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderitem"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/orderstatushistory"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/patient"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/google/uuid"
)

// OrderQuery is the builder for querying Order entities.
type OrderQuery struct {
	config
	ctx                    *QueryContext
	order                  []order.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Order
	withClinic             *ClinicQuery
	withPatient            *PatientQuery
	withOrderItems         *OrderItemQuery
	withOrderStatusHistory *OrderStatusHistoryQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderQuery builder.
func (oq *OrderQuery) Where(ps ...predicate.Order) *OrderQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OrderQuery) Limit(limit int) *OrderQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OrderQuery) Offset(offset int) *OrderQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OrderQuery) Unique(unique bool) *OrderQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OrderQuery) Order(o ...order.OrderOption) *OrderQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryClinic chains the current query on the "clinic" edge.
func (oq *OrderQuery) QueryClinic() *ClinicQuery {
	query := (&ClinicClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(clinic.Table, clinic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.ClinicTable, order.ClinicColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the "patient" edge.
func (oq *OrderQuery) QueryPatient() *PatientQuery {
	query := (&PatientClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.PatientTable, order.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItems chains the current query on the "order_items" edge.
func (oq *OrderQuery) QueryOrderItems() *OrderItemQuery {
	query := (&OrderItemClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderItemsTable, order.OrderItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderStatusHistory chains the current query on the "order_status_history" edge.
func (oq *OrderQuery) QueryOrderStatusHistory() *OrderStatusHistoryQuery {
	query := (&OrderStatusHistoryClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(orderstatushistory.Table, orderstatushistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderStatusHistoryTable, order.OrderStatusHistoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Order entity from the query.
// Returns a *NotFoundError when no Order was found.
func (oq *OrderQuery) First(ctx context.Context) (*Order, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{order.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrderQuery) FirstX(ctx context.Context) *Order {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Order ID from the query.
// Returns a *NotFoundError when no Order ID was found.
func (oq *OrderQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{order.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OrderQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Order entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Order entity is found.
// Returns a *NotFoundError when no Order entities are found.
func (oq *OrderQuery) Only(ctx context.Context) (*Order, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{order.Label}
	default:
		return nil, &NotSingularError{order.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrderQuery) OnlyX(ctx context.Context) *Order {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Order ID in the query.
// Returns a *NotSingularError when more than one Order ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OrderQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{order.Label}
	default:
		err = &NotSingularError{order.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OrderQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Orders.
func (oq *OrderQuery) All(ctx context.Context) ([]*Order, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Order, *OrderQuery]()
	return withInterceptors[[]*Order](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrderQuery) AllX(ctx context.Context) []*Order {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Order IDs.
func (oq *OrderQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(order.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrderQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OrderQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrderQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrderQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrderQuery) Clone() *OrderQuery {
	if oq == nil {
		return nil
	}
	return &OrderQuery{
		config:                 oq.config,
		ctx:                    oq.ctx.Clone(),
		order:                  append([]order.OrderOption{}, oq.order...),
		inters:                 append([]Interceptor{}, oq.inters...),
		predicates:             append([]predicate.Order{}, oq.predicates...),
		withClinic:             oq.withClinic.Clone(),
		withPatient:            oq.withPatient.Clone(),
		withOrderItems:         oq.withOrderItems.Clone(),
		withOrderStatusHistory: oq.withOrderStatusHistory.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithClinic tells the query-builder to eager-load the nodes that are connected to
// the "clinic" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithClinic(opts ...func(*ClinicQuery)) *OrderQuery {
	query := (&ClinicClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withClinic = query
	return oq
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithPatient(opts ...func(*PatientQuery)) *OrderQuery {
	query := (&PatientClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withPatient = query
	return oq
}

// WithOrderItems tells the query-builder to eager-load the nodes that are connected to
// the "order_items" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithOrderItems(opts ...func(*OrderItemQuery)) *OrderQuery {
	query := (&OrderItemClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withOrderItems = query
	return oq
}

// WithOrderStatusHistory tells the query-builder to eager-load the nodes that are connected to
// the "order_status_history" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithOrderStatusHistory(opts ...func(*OrderStatusHistoryQuery)) *OrderQuery {
	query := (&OrderStatusHistoryClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withOrderStatusHistory = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderNumber string `json:"order_number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Order.Query().
//		GroupBy(order.FieldOrderNumber).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (oq *OrderQuery) GroupBy(field string, fields ...string) *OrderGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = order.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderNumber string `json:"order_number,omitempty"`
//	}
//
//	client.Order.Query().
//		Select(order.FieldOrderNumber).
//		Scan(ctx, &v)
func (oq *OrderQuery) Select(fields ...string) *OrderSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OrderSelect{OrderQuery: oq}
	sbuild.label = order.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderSelect configured with the given aggregations.
func (oq *OrderQuery) Aggregate(fns ...AggregateFunc) *OrderSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !order.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Order, error) {
	var (
		nodes       = []*Order{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [4]bool{
			oq.withClinic != nil,
			oq.withPatient != nil,
			oq.withOrderItems != nil,
			oq.withOrderStatusHistory != nil,
		}
	)
	if oq.withClinic != nil || oq.withPatient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, order.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Order).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Order{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withClinic; query != nil {
		if err := oq.loadClinic(ctx, query, nodes, nil,
			func(n *Order, e *Clinic) { n.Edges.Clinic = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withPatient; query != nil {
		if err := oq.loadPatient(ctx, query, nodes, nil,
			func(n *Order, e *Patient) { n.Edges.Patient = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withOrderItems; query != nil {
		if err := oq.loadOrderItems(ctx, query, nodes,
			func(n *Order) { n.Edges.OrderItems = []*OrderItem{} },
			func(n *Order, e *OrderItem) { n.Edges.OrderItems = append(n.Edges.OrderItems, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withOrderStatusHistory; query != nil {
		if err := oq.loadOrderStatusHistory(ctx, query, nodes,
			func(n *Order) { n.Edges.OrderStatusHistory = []*OrderStatusHistory{} },
			func(n *Order, e *OrderStatusHistory) {
				n.Edges.OrderStatusHistory = append(n.Edges.OrderStatusHistory, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OrderQuery) loadClinic(ctx context.Context, query *ClinicQuery, nodes []*Order, init func(*Order), assign func(*Order, *Clinic)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Order)
	for i := range nodes {
		if nodes[i].clinic_orders == nil {
			continue
		}
		fk := *nodes[i].clinic_orders
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(clinic.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "clinic_orders" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OrderQuery) loadPatient(ctx context.Context, query *PatientQuery, nodes []*Order, init func(*Order), assign func(*Order, *Patient)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Order)
	for i := range nodes {
		if nodes[i].patient_orders == nil {
			continue
		}
		fk := *nodes[i].patient_orders
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(patient.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "patient_orders" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OrderQuery) loadOrderItems(ctx context.Context, query *OrderItemQuery, nodes []*Order, init func(*Order), assign func(*Order, *OrderItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(order.OrderItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.order_order_items
		if fk == nil {
			return fmt.Errorf(`foreign-key "order_order_items" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "order_order_items" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrderQuery) loadOrderStatusHistory(ctx context.Context, query *OrderStatusHistoryQuery, nodes []*Order, init func(*Order), assign func(*Order, *OrderStatusHistory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderStatusHistory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(order.OrderStatusHistoryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.order_order_status_history
		if fk == nil {
			return fmt.Errorf(`foreign-key "order_order_status_history" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "order_order_status_history" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for i := range fields {
			if fields[i] != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(order.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = order.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderGroupBy is the group-by builder for Order entities.
type OrderGroupBy struct {
	selector
	build *OrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrderGroupBy) Aggregate(fns ...AggregateFunc) *OrderGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderQuery, *OrderGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OrderGroupBy) sqlScan(ctx context.Context, root *OrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderSelect is the builder for selecting fields of Order entities.
type OrderSelect struct {
	*OrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OrderSelect) Aggregate(fns ...AggregateFunc) *OrderSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderQuery, *OrderSelect](ctx, os.OrderQuery, os, os.inters, v)
}

func (os *OrderSelect) sqlScan(ctx context.Context, root *OrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
