// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/order/service/internal/data/model/order"
	"mall-go/app/order/service/internal/data/model/ordersnap"
	"mall-go/app/order/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSnapUpdate is the builder for updating OrderSnap entities.
type OrderSnapUpdate struct {
	config
	hooks    []Hook
	mutation *OrderSnapMutation
}

// Where appends a list predicates to the OrderSnapUpdate builder.
func (osu *OrderSnapUpdate) Where(ps ...predicate.OrderSnap) *OrderSnapUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetSnapImg sets the "snap_img" field.
func (osu *OrderSnapUpdate) SetSnapImg(s string) *OrderSnapUpdate {
	osu.mutation.SetSnapImg(s)
	return osu
}

// SetSnapTitle sets the "snap_title" field.
func (osu *OrderSnapUpdate) SetSnapTitle(s string) *OrderSnapUpdate {
	osu.mutation.SetSnapTitle(s)
	return osu
}

// SetSnapItems sets the "snap_items" field.
func (osu *OrderSnapUpdate) SetSnapItems(s string) *OrderSnapUpdate {
	osu.mutation.SetSnapItems(s)
	return osu
}

// SetSnapAddress sets the "snap_address" field.
func (osu *OrderSnapUpdate) SetSnapAddress(s string) *OrderSnapUpdate {
	osu.mutation.SetSnapAddress(s)
	return osu
}

// SetOrderID sets the "order_id" field.
func (osu *OrderSnapUpdate) SetOrderID(i int64) *OrderSnapUpdate {
	osu.mutation.SetOrderID(i)
	return osu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osu *OrderSnapUpdate) SetNillableOrderID(i *int64) *OrderSnapUpdate {
	if i != nil {
		osu.SetOrderID(*i)
	}
	return osu
}

// ClearOrderID clears the value of the "order_id" field.
func (osu *OrderSnapUpdate) ClearOrderID() *OrderSnapUpdate {
	osu.mutation.ClearOrderID()
	return osu
}

// SetOrder sets the "order" edge to the Order entity.
func (osu *OrderSnapUpdate) SetOrder(o *Order) *OrderSnapUpdate {
	return osu.SetOrderID(o.ID)
}

// Mutation returns the OrderSnapMutation object of the builder.
func (osu *OrderSnapUpdate) Mutation() *OrderSnapMutation {
	return osu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osu *OrderSnapUpdate) ClearOrder() *OrderSnapUpdate {
	osu.mutation.ClearOrder()
	return osu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrderSnapUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(osu.hooks) == 0 {
		affected, err = osu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderSnapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osu.mutation = mutation
			affected, err = osu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osu.hooks) - 1; i >= 0; i-- {
			if osu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = osu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrderSnapUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrderSnapUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrderSnapUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (osu *OrderSnapUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordersnap.Table,
			Columns: ordersnap.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersnap.FieldID,
			},
		},
	}
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.SnapImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapImg,
		})
	}
	if value, ok := osu.mutation.SnapTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapTitle,
		})
	}
	if value, ok := osu.mutation.SnapItems(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapItems,
		})
	}
	if value, ok := osu.mutation.SnapAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapAddress,
		})
	}
	if osu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersnap.OrderTable,
			Columns: []string{ordersnap.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersnap.OrderTable,
			Columns: []string{ordersnap.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersnap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderSnapUpdateOne is the builder for updating a single OrderSnap entity.
type OrderSnapUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderSnapMutation
}

// SetSnapImg sets the "snap_img" field.
func (osuo *OrderSnapUpdateOne) SetSnapImg(s string) *OrderSnapUpdateOne {
	osuo.mutation.SetSnapImg(s)
	return osuo
}

// SetSnapTitle sets the "snap_title" field.
func (osuo *OrderSnapUpdateOne) SetSnapTitle(s string) *OrderSnapUpdateOne {
	osuo.mutation.SetSnapTitle(s)
	return osuo
}

// SetSnapItems sets the "snap_items" field.
func (osuo *OrderSnapUpdateOne) SetSnapItems(s string) *OrderSnapUpdateOne {
	osuo.mutation.SetSnapItems(s)
	return osuo
}

// SetSnapAddress sets the "snap_address" field.
func (osuo *OrderSnapUpdateOne) SetSnapAddress(s string) *OrderSnapUpdateOne {
	osuo.mutation.SetSnapAddress(s)
	return osuo
}

// SetOrderID sets the "order_id" field.
func (osuo *OrderSnapUpdateOne) SetOrderID(i int64) *OrderSnapUpdateOne {
	osuo.mutation.SetOrderID(i)
	return osuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osuo *OrderSnapUpdateOne) SetNillableOrderID(i *int64) *OrderSnapUpdateOne {
	if i != nil {
		osuo.SetOrderID(*i)
	}
	return osuo
}

// ClearOrderID clears the value of the "order_id" field.
func (osuo *OrderSnapUpdateOne) ClearOrderID() *OrderSnapUpdateOne {
	osuo.mutation.ClearOrderID()
	return osuo
}

// SetOrder sets the "order" edge to the Order entity.
func (osuo *OrderSnapUpdateOne) SetOrder(o *Order) *OrderSnapUpdateOne {
	return osuo.SetOrderID(o.ID)
}

// Mutation returns the OrderSnapMutation object of the builder.
func (osuo *OrderSnapUpdateOne) Mutation() *OrderSnapMutation {
	return osuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osuo *OrderSnapUpdateOne) ClearOrder() *OrderSnapUpdateOne {
	osuo.mutation.ClearOrder()
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrderSnapUpdateOne) Select(field string, fields ...string) *OrderSnapUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrderSnap entity.
func (osuo *OrderSnapUpdateOne) Save(ctx context.Context) (*OrderSnap, error) {
	var (
		err  error
		node *OrderSnap
	)
	if len(osuo.hooks) == 0 {
		node, err = osuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderSnapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osuo.mutation = mutation
			node, err = osuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osuo.hooks) - 1; i >= 0; i-- {
			if osuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = osuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderSnap)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderSnapMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrderSnapUpdateOne) SaveX(ctx context.Context) *OrderSnap {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrderSnapUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrderSnapUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (osuo *OrderSnapUpdateOne) sqlSave(ctx context.Context) (_node *OrderSnap, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordersnap.Table,
			Columns: ordersnap.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersnap.FieldID,
			},
		},
	}
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "OrderSnap.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordersnap.FieldID)
		for _, f := range fields {
			if !ordersnap.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != ordersnap.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.SnapImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapImg,
		})
	}
	if value, ok := osuo.mutation.SnapTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapTitle,
		})
	}
	if value, ok := osuo.mutation.SnapItems(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapItems,
		})
	}
	if value, ok := osuo.mutation.SnapAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordersnap.FieldSnapAddress,
		})
	}
	if osuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersnap.OrderTable,
			Columns: []string{ordersnap.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersnap.OrderTable,
			Columns: []string{ordersnap.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderSnap{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersnap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
