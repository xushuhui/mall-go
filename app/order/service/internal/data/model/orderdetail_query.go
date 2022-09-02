// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/app/order/service/internal/data/model/orderdetail"
	"mall-go/app/order/service/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderDetailQuery is the builder for querying OrderDetail entities.
type OrderDetailQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderDetail
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderDetailQuery builder.
func (odq *OrderDetailQuery) Where(ps ...predicate.OrderDetail) *OrderDetailQuery {
	odq.predicates = append(odq.predicates, ps...)
	return odq
}

// Limit adds a limit step to the query.
func (odq *OrderDetailQuery) Limit(limit int) *OrderDetailQuery {
	odq.limit = &limit
	return odq
}

// Offset adds an offset step to the query.
func (odq *OrderDetailQuery) Offset(offset int) *OrderDetailQuery {
	odq.offset = &offset
	return odq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (odq *OrderDetailQuery) Unique(unique bool) *OrderDetailQuery {
	odq.unique = &unique
	return odq
}

// Order adds an order step to the query.
func (odq *OrderDetailQuery) Order(o ...OrderFunc) *OrderDetailQuery {
	odq.order = append(odq.order, o...)
	return odq
}

// First returns the first OrderDetail entity from the query.
// Returns a *NotFoundError when no OrderDetail was found.
func (odq *OrderDetailQuery) First(ctx context.Context) (*OrderDetail, error) {
	nodes, err := odq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderdetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (odq *OrderDetailQuery) FirstX(ctx context.Context) *OrderDetail {
	node, err := odq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderDetail ID from the query.
// Returns a *NotFoundError when no OrderDetail ID was found.
func (odq *OrderDetailQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = odq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (odq *OrderDetailQuery) FirstIDX(ctx context.Context) int64 {
	id, err := odq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderDetail entity is found.
// Returns a *NotFoundError when no OrderDetail entities are found.
func (odq *OrderDetailQuery) Only(ctx context.Context) (*OrderDetail, error) {
	nodes, err := odq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderdetail.Label}
	default:
		return nil, &NotSingularError{orderdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (odq *OrderDetailQuery) OnlyX(ctx context.Context) *OrderDetail {
	node, err := odq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderDetail ID in the query.
// Returns a *NotSingularError when more than one OrderDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (odq *OrderDetailQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = odq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = &NotSingularError{orderdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (odq *OrderDetailQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := odq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderDetails.
func (odq *OrderDetailQuery) All(ctx context.Context) ([]*OrderDetail, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return odq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (odq *OrderDetailQuery) AllX(ctx context.Context) []*OrderDetail {
	nodes, err := odq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderDetail IDs.
func (odq *OrderDetailQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := odq.Select(orderdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (odq *OrderDetailQuery) IDsX(ctx context.Context) []int64 {
	ids, err := odq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (odq *OrderDetailQuery) Count(ctx context.Context) (int, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return odq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (odq *OrderDetailQuery) CountX(ctx context.Context) int {
	count, err := odq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (odq *OrderDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return odq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (odq *OrderDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := odq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (odq *OrderDetailQuery) Clone() *OrderDetailQuery {
	if odq == nil {
		return nil
	}
	return &OrderDetailQuery{
		config:     odq.config,
		limit:      odq.limit,
		offset:     odq.offset,
		order:      append([]OrderFunc{}, odq.order...),
		predicates: append([]predicate.OrderDetail{}, odq.predicates...),
		// clone intermediate query.
		sql:    odq.sql.Clone(),
		path:   odq.path,
		unique: odq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderDetail.Query().
//		GroupBy(orderdetail.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (odq *OrderDetailQuery) GroupBy(field string, fields ...string) *OrderDetailGroupBy {
	grbuild := &OrderDetailGroupBy{config: odq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odq.sqlQuery(ctx), nil
	}
	grbuild.label = orderdetail.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.OrderDetail.Query().
//		Select(orderdetail.FieldCreateTime).
//		Scan(ctx, &v)
//
func (odq *OrderDetailQuery) Select(fields ...string) *OrderDetailSelect {
	odq.fields = append(odq.fields, fields...)
	selbuild := &OrderDetailSelect{OrderDetailQuery: odq}
	selbuild.label = orderdetail.Label
	selbuild.flds, selbuild.scan = &odq.fields, selbuild.Scan
	return selbuild
}

func (odq *OrderDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range odq.fields {
		if !orderdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if odq.path != nil {
		prev, err := odq.path(ctx)
		if err != nil {
			return err
		}
		odq.sql = prev
	}
	return nil
}

func (odq *OrderDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderDetail, error) {
	var (
		nodes = []*OrderDetail{}
		_spec = odq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrderDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrderDetail{config: odq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, odq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (odq *OrderDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := odq.querySpec()
	_spec.Node.Columns = odq.fields
	if len(odq.fields) > 0 {
		_spec.Unique = odq.unique != nil && *odq.unique
	}
	return sqlgraph.CountNodes(ctx, odq.driver, _spec)
}

func (odq *OrderDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := odq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (odq *OrderDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		},
		From:   odq.sql,
		Unique: true,
	}
	if unique := odq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := odq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderdetail.FieldID)
		for i := range fields {
			if fields[i] != orderdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := odq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := odq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := odq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := odq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (odq *OrderDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(odq.driver.Dialect())
	t1 := builder.Table(orderdetail.Table)
	columns := odq.fields
	if len(columns) == 0 {
		columns = orderdetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if odq.sql != nil {
		selector = odq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if odq.unique != nil && *odq.unique {
		selector.Distinct()
	}
	for _, p := range odq.predicates {
		p(selector)
	}
	for _, p := range odq.order {
		p(selector)
	}
	if offset := odq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := odq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderDetailGroupBy is the group-by builder for OrderDetail entities.
type OrderDetailGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (odgb *OrderDetailGroupBy) Aggregate(fns ...AggregateFunc) *OrderDetailGroupBy {
	odgb.fns = append(odgb.fns, fns...)
	return odgb
}

// Scan applies the group-by query and scans the result into the given value.
func (odgb *OrderDetailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := odgb.path(ctx)
	if err != nil {
		return err
	}
	odgb.sql = query
	return odgb.sqlScan(ctx, v)
}

func (odgb *OrderDetailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range odgb.fields {
		if !orderdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := odgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := odgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (odgb *OrderDetailGroupBy) sqlQuery() *sql.Selector {
	selector := odgb.sql.Select()
	aggregation := make([]string, 0, len(odgb.fns))
	for _, fn := range odgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(odgb.fields)+len(odgb.fns))
		for _, f := range odgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(odgb.fields...)...)
}

// OrderDetailSelect is the builder for selecting fields of OrderDetail entities.
type OrderDetailSelect struct {
	*OrderDetailQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ods *OrderDetailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ods.prepareQuery(ctx); err != nil {
		return err
	}
	ods.sql = ods.OrderDetailQuery.sqlQuery(ctx)
	return ods.sqlScan(ctx, v)
}

func (ods *OrderDetailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ods.sql.Query()
	if err := ods.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
