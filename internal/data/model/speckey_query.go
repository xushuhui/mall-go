// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/speckey"
	"mall-go/internal/data/model/spu"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpecKeyQuery is the builder for querying SpecKey entities.
type SpecKeyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SpecKey
	// eager-loading edges.
	withSpu *SpuQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpecKeyQuery builder.
func (skq *SpecKeyQuery) Where(ps ...predicate.SpecKey) *SpecKeyQuery {
	skq.predicates = append(skq.predicates, ps...)
	return skq
}

// Limit adds a limit step to the query.
func (skq *SpecKeyQuery) Limit(limit int) *SpecKeyQuery {
	skq.limit = &limit
	return skq
}

// Offset adds an offset step to the query.
func (skq *SpecKeyQuery) Offset(offset int) *SpecKeyQuery {
	skq.offset = &offset
	return skq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (skq *SpecKeyQuery) Unique(unique bool) *SpecKeyQuery {
	skq.unique = &unique
	return skq
}

// Order adds an order step to the query.
func (skq *SpecKeyQuery) Order(o ...OrderFunc) *SpecKeyQuery {
	skq.order = append(skq.order, o...)
	return skq
}

// QuerySpu chains the current query on the "spu" edge.
func (skq *SpecKeyQuery) QuerySpu() *SpuQuery {
	query := &SpuQuery{config: skq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := skq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := skq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(speckey.Table, speckey.FieldID, selector),
			sqlgraph.To(spu.Table, spu.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, speckey.SpuTable, speckey.SpuPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(skq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SpecKey entity from the query.
// Returns a *NotFoundError when no SpecKey was found.
func (skq *SpecKeyQuery) First(ctx context.Context) (*SpecKey, error) {
	nodes, err := skq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{speckey.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (skq *SpecKeyQuery) FirstX(ctx context.Context) *SpecKey {
	node, err := skq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SpecKey ID from the query.
// Returns a *NotFoundError when no SpecKey ID was found.
func (skq *SpecKeyQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = skq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{speckey.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (skq *SpecKeyQuery) FirstIDX(ctx context.Context) int64 {
	id, err := skq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last SpecKey entity from the query.
// Returns a *NotFoundError when no SpecKey was found.
func (skq *SpecKeyQuery) Last(ctx context.Context) (*SpecKey, error) {
	nodes, err := skq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{speckey.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (skq *SpecKeyQuery) LastX(ctx context.Context) *SpecKey {
	node, err := skq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last SpecKey ID from the query.
// Returns a *NotFoundError when no SpecKey ID was found.
func (skq *SpecKeyQuery) LastID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = skq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{speckey.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (skq *SpecKeyQuery) LastIDX(ctx context.Context) int64 {
	id, err := skq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SpecKey entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SpecKey entity is not found.
// Returns a *NotFoundError when no SpecKey entities are found.
func (skq *SpecKeyQuery) Only(ctx context.Context) (*SpecKey, error) {
	nodes, err := skq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{speckey.Label}
	default:
		return nil, &NotSingularError{speckey.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (skq *SpecKeyQuery) OnlyX(ctx context.Context) *SpecKey {
	node, err := skq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SpecKey ID in the query.
// Returns a *NotSingularError when exactly one SpecKey ID is not found.
// Returns a *NotFoundError when no entities are found.
func (skq *SpecKeyQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = skq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = &NotSingularError{speckey.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (skq *SpecKeyQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := skq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SpecKeys.
func (skq *SpecKeyQuery) All(ctx context.Context) ([]*SpecKey, error) {
	if err := skq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return skq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (skq *SpecKeyQuery) AllX(ctx context.Context) []*SpecKey {
	nodes, err := skq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SpecKey IDs.
func (skq *SpecKeyQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := skq.Select(speckey.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (skq *SpecKeyQuery) IDsX(ctx context.Context) []int64 {
	ids, err := skq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (skq *SpecKeyQuery) Count(ctx context.Context) (int, error) {
	if err := skq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return skq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (skq *SpecKeyQuery) CountX(ctx context.Context) int {
	count, err := skq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (skq *SpecKeyQuery) Exist(ctx context.Context) (bool, error) {
	if err := skq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return skq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (skq *SpecKeyQuery) ExistX(ctx context.Context) bool {
	exist, err := skq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecKeyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (skq *SpecKeyQuery) Clone() *SpecKeyQuery {
	if skq == nil {
		return nil
	}
	return &SpecKeyQuery{
		config:     skq.config,
		limit:      skq.limit,
		offset:     skq.offset,
		order:      append([]OrderFunc{}, skq.order...),
		predicates: append([]predicate.SpecKey{}, skq.predicates...),
		withSpu:    skq.withSpu.Clone(),
		// clone intermediate query.
		sql:  skq.sql.Clone(),
		path: skq.path,
	}
}

// WithSpu tells the query-builder to eager-load the nodes that are connected to
// the "spu" edge. The optional arguments are used to configure the query builder of the edge.
func (skq *SpecKeyQuery) WithSpu(opts ...func(*SpuQuery)) *SpecKeyQuery {
	query := &SpuQuery{config: skq.config}
	for _, opt := range opts {
		opt(query)
	}
	skq.withSpu = query
	return skq
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
//	client.SpecKey.Query().
//		GroupBy(speckey.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (skq *SpecKeyQuery) GroupBy(field string, fields ...string) *SpecKeyGroupBy {
	group := &SpecKeyGroupBy{config: skq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := skq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return skq.sqlQuery(ctx), nil
	}
	return group
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
//	client.SpecKey.Query().
//		Select(speckey.FieldCreateTime).
//		Scan(ctx, &v)
//
func (skq *SpecKeyQuery) Select(fields ...string) *SpecKeySelect {
	skq.fields = append(skq.fields, fields...)
	return &SpecKeySelect{SpecKeyQuery: skq}
}

func (skq *SpecKeyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range skq.fields {
		if !speckey.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if skq.path != nil {
		prev, err := skq.path(ctx)
		if err != nil {
			return err
		}
		skq.sql = prev
	}
	return nil
}

func (skq *SpecKeyQuery) sqlAll(ctx context.Context) ([]*SpecKey, error) {
	var (
		nodes       = []*SpecKey{}
		_spec       = skq.querySpec()
		loadedTypes = [1]bool{
			skq.withSpu != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SpecKey{config: skq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, skq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := skq.withSpu; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*SpecKey, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Spu = []*Spu{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*SpecKey)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   speckey.SpuTable,
				Columns: speckey.SpuPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(speckey.SpuPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, skq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "spu": %w`, err)
		}
		query.Where(spu.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "spu" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Spu = append(nodes[i].Edges.Spu, n)
			}
		}
	}

	return nodes, nil
}

func (skq *SpecKeyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := skq.querySpec()
	return sqlgraph.CountNodes(ctx, skq.driver, _spec)
}

func (skq *SpecKeyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := skq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (skq *SpecKeyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   speckey.Table,
			Columns: speckey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: speckey.FieldID,
			},
		},
		From:   skq.sql,
		Unique: true,
	}
	if unique := skq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := skq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, speckey.FieldID)
		for i := range fields {
			if fields[i] != speckey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := skq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := skq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := skq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := skq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (skq *SpecKeyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(skq.driver.Dialect())
	t1 := builder.Table(speckey.Table)
	columns := skq.fields
	if len(columns) == 0 {
		columns = speckey.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if skq.sql != nil {
		selector = skq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range skq.predicates {
		p(selector)
	}
	for _, p := range skq.order {
		p(selector)
	}
	if offset := skq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := skq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SpecKeyGroupBy is the group-by builder for SpecKey entities.
type SpecKeyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (skgb *SpecKeyGroupBy) Aggregate(fns ...AggregateFunc) *SpecKeyGroupBy {
	skgb.fns = append(skgb.fns, fns...)
	return skgb
}

// Scan applies the group-by query and scans the result into the given value.
func (skgb *SpecKeyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := skgb.path(ctx)
	if err != nil {
		return err
	}
	skgb.sql = query
	return skgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := skgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(skgb.fields) > 1 {
		return nil, errors.New("model: SpecKeyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := skgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) StringsX(ctx context.Context) []string {
	v, err := skgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = skgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) StringX(ctx context.Context) string {
	v, err := skgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(skgb.fields) > 1 {
		return nil, errors.New("model: SpecKeyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := skgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) IntsX(ctx context.Context) []int {
	v, err := skgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = skgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) IntX(ctx context.Context) int {
	v, err := skgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(skgb.fields) > 1 {
		return nil, errors.New("model: SpecKeyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := skgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := skgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = skgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := skgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(skgb.fields) > 1 {
		return nil, errors.New("model: SpecKeyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := skgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := skgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (skgb *SpecKeyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = skgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (skgb *SpecKeyGroupBy) BoolX(ctx context.Context) bool {
	v, err := skgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (skgb *SpecKeyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range skgb.fields {
		if !speckey.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := skgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := skgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (skgb *SpecKeyGroupBy) sqlQuery() *sql.Selector {
	selector := skgb.sql.Select()
	aggregation := make([]string, 0, len(skgb.fns))
	for _, fn := range skgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(skgb.fields)+len(skgb.fns))
		for _, f := range skgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(skgb.fields...)...)
}

// SpecKeySelect is the builder for selecting fields of SpecKey entities.
type SpecKeySelect struct {
	*SpecKeyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sks *SpecKeySelect) Scan(ctx context.Context, v interface{}) error {
	if err := sks.prepareQuery(ctx); err != nil {
		return err
	}
	sks.sql = sks.SpecKeyQuery.sqlQuery(ctx)
	return sks.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sks *SpecKeySelect) ScanX(ctx context.Context, v interface{}) {
	if err := sks.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Strings(ctx context.Context) ([]string, error) {
	if len(sks.fields) > 1 {
		return nil, errors.New("model: SpecKeySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sks *SpecKeySelect) StringsX(ctx context.Context) []string {
	v, err := sks.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sks.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sks *SpecKeySelect) StringX(ctx context.Context) string {
	v, err := sks.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Ints(ctx context.Context) ([]int, error) {
	if len(sks.fields) > 1 {
		return nil, errors.New("model: SpecKeySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sks *SpecKeySelect) IntsX(ctx context.Context) []int {
	v, err := sks.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sks.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sks *SpecKeySelect) IntX(ctx context.Context) int {
	v, err := sks.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sks.fields) > 1 {
		return nil, errors.New("model: SpecKeySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sks *SpecKeySelect) Float64sX(ctx context.Context) []float64 {
	v, err := sks.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sks.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sks *SpecKeySelect) Float64X(ctx context.Context) float64 {
	v, err := sks.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sks.fields) > 1 {
		return nil, errors.New("model: SpecKeySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sks *SpecKeySelect) BoolsX(ctx context.Context) []bool {
	v, err := sks.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sks *SpecKeySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sks.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{speckey.Label}
	default:
		err = fmt.Errorf("model: SpecKeySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sks *SpecKeySelect) BoolX(ctx context.Context) bool {
	v, err := sks.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sks *SpecKeySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sks.sql.Query()
	if err := sks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}