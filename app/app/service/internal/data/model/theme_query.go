// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"mall-go/app/app/service/internal/data/model/predicate"
	"mall-go/app/app/service/internal/data/model/theme"
	"mall-go/app/app/service/internal/data/model/themespu"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThemeQuery is the builder for querying Theme entities.
type ThemeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Theme
	// eager-loading edges.
	withThemeSpu *ThemeSpuQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThemeQuery builder.
func (tq *ThemeQuery) Where(ps ...predicate.Theme) *ThemeQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *ThemeQuery) Limit(limit int) *ThemeQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *ThemeQuery) Offset(offset int) *ThemeQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *ThemeQuery) Unique(unique bool) *ThemeQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *ThemeQuery) Order(o ...OrderFunc) *ThemeQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryThemeSpu chains the current query on the "theme_spu" edge.
func (tq *ThemeQuery) QueryThemeSpu() *ThemeSpuQuery {
	query := &ThemeSpuQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(theme.Table, theme.FieldID, selector),
			sqlgraph.To(themespu.Table, themespu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, theme.ThemeSpuTable, theme.ThemeSpuColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Theme entity from the query.
// Returns a *NotFoundError when no Theme was found.
func (tq *ThemeQuery) First(ctx context.Context) (*Theme, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{theme.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *ThemeQuery) FirstX(ctx context.Context) *Theme {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Theme ID from the query.
// Returns a *NotFoundError when no Theme ID was found.
func (tq *ThemeQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{theme.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *ThemeQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Theme entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Theme entity is found.
// Returns a *NotFoundError when no Theme entities are found.
func (tq *ThemeQuery) Only(ctx context.Context) (*Theme, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{theme.Label}
	default:
		return nil, &NotSingularError{theme.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *ThemeQuery) OnlyX(ctx context.Context) *Theme {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Theme ID in the query.
// Returns a *NotSingularError when more than one Theme ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *ThemeQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{theme.Label}
	default:
		err = &NotSingularError{theme.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *ThemeQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Themes.
func (tq *ThemeQuery) All(ctx context.Context) ([]*Theme, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *ThemeQuery) AllX(ctx context.Context) []*Theme {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Theme IDs.
func (tq *ThemeQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := tq.Select(theme.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *ThemeQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *ThemeQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *ThemeQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *ThemeQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *ThemeQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThemeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *ThemeQuery) Clone() *ThemeQuery {
	if tq == nil {
		return nil
	}
	return &ThemeQuery{
		config:       tq.config,
		limit:        tq.limit,
		offset:       tq.offset,
		order:        append([]OrderFunc{}, tq.order...),
		predicates:   append([]predicate.Theme{}, tq.predicates...),
		withThemeSpu: tq.withThemeSpu.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithThemeSpu tells the query-builder to eager-load the nodes that are connected to
// the "theme_spu" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThemeQuery) WithThemeSpu(opts ...func(*ThemeSpuQuery)) *ThemeQuery {
	query := &ThemeSpuQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withThemeSpu = query
	return tq
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
//	client.Theme.Query().
//		GroupBy(theme.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (tq *ThemeQuery) GroupBy(field string, fields ...string) *ThemeGroupBy {
	grbuild := &ThemeGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = theme.Label
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
//	client.Theme.Query().
//		Select(theme.FieldCreateTime).
//		Scan(ctx, &v)
//
func (tq *ThemeQuery) Select(fields ...string) *ThemeSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &ThemeSelect{ThemeQuery: tq}
	selbuild.label = theme.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *ThemeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !theme.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *ThemeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Theme, error) {
	var (
		nodes       = []*Theme{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withThemeSpu != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Theme).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Theme{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withThemeSpu; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*Theme)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ThemeSpu = []*ThemeSpu{}
		}
		query.Where(predicate.ThemeSpu(func(s *sql.Selector) {
			s.Where(sql.InValues(theme.ThemeSpuColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ThemeID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "theme_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ThemeSpu = append(node.Edges.ThemeSpu, n)
		}
	}

	return nodes, nil
}

func (tq *ThemeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *ThemeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *ThemeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   theme.Table,
			Columns: theme.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: theme.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, theme.FieldID)
		for i := range fields {
			if fields[i] != theme.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *ThemeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(theme.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = theme.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThemeGroupBy is the group-by builder for Theme entities.
type ThemeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *ThemeGroupBy) Aggregate(fns ...AggregateFunc) *ThemeGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *ThemeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *ThemeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !theme.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *ThemeGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// ThemeSelect is the builder for selecting fields of Theme entities.
type ThemeSelect struct {
	*ThemeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *ThemeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.ThemeQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *ThemeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
