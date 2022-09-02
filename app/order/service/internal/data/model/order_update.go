// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/order/service/internal/data/model/order"
	"mall-go/app/order/service/internal/data/model/ordersnap"
	"mall-go/app/order/service/internal/data/model/ordersub"
	"mall-go/app/order/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdateTime sets the "update_time" field.
func (ou *OrderUpdate) SetUpdateTime(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdateTime(t)
	return ou
}

// SetDeleteTime sets the "delete_time" field.
func (ou *OrderUpdate) SetDeleteTime(t time.Time) *OrderUpdate {
	ou.mutation.SetDeleteTime(t)
	return ou
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeleteTime(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetDeleteTime(*t)
	}
	return ou
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ou *OrderUpdate) ClearDeleteTime() *OrderUpdate {
	ou.mutation.ClearDeleteTime()
	return ou
}

// SetOrderNo sets the "order_no" field.
func (ou *OrderUpdate) SetOrderNo(s string) *OrderUpdate {
	ou.mutation.SetOrderNo(s)
	return ou
}

// SetTransactionID sets the "transaction_id" field.
func (ou *OrderUpdate) SetTransactionID(s string) *OrderUpdate {
	ou.mutation.SetTransactionID(s)
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(i int64) *OrderUpdate {
	ou.mutation.ResetUserID()
	ou.mutation.SetUserID(i)
	return ou
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(i *int64) *OrderUpdate {
	if i != nil {
		ou.SetUserID(*i)
	}
	return ou
}

// AddUserID adds i to the "user_id" field.
func (ou *OrderUpdate) AddUserID(i int64) *OrderUpdate {
	ou.mutation.AddUserID(i)
	return ou
}

// ClearUserID clears the value of the "user_id" field.
func (ou *OrderUpdate) ClearUserID() *OrderUpdate {
	ou.mutation.ClearUserID()
	return ou
}

// SetTotalPrice sets the "total_price" field.
func (ou *OrderUpdate) SetTotalPrice(f float64) *OrderUpdate {
	ou.mutation.ResetTotalPrice()
	ou.mutation.SetTotalPrice(f)
	return ou
}

// AddTotalPrice adds f to the "total_price" field.
func (ou *OrderUpdate) AddTotalPrice(f float64) *OrderUpdate {
	ou.mutation.AddTotalPrice(f)
	return ou
}

// SetTotalCount sets the "total_count" field.
func (ou *OrderUpdate) SetTotalCount(i int) *OrderUpdate {
	ou.mutation.ResetTotalCount()
	ou.mutation.SetTotalCount(i)
	return ou
}

// AddTotalCount adds i to the "total_count" field.
func (ou *OrderUpdate) AddTotalCount(i int) *OrderUpdate {
	ou.mutation.AddTotalCount(i)
	return ou
}

// SetFinalTotalPrice sets the "final_total_price" field.
func (ou *OrderUpdate) SetFinalTotalPrice(f float64) *OrderUpdate {
	ou.mutation.ResetFinalTotalPrice()
	ou.mutation.SetFinalTotalPrice(f)
	return ou
}

// AddFinalTotalPrice adds f to the "final_total_price" field.
func (ou *OrderUpdate) AddFinalTotalPrice(f float64) *OrderUpdate {
	ou.mutation.AddFinalTotalPrice(f)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(i int) *OrderUpdate {
	ou.mutation.ResetStatus()
	ou.mutation.SetStatus(i)
	return ou
}

// AddStatus adds i to the "status" field.
func (ou *OrderUpdate) AddStatus(i int) *OrderUpdate {
	ou.mutation.AddStatus(i)
	return ou
}

// AddOrderSnapIDs adds the "order_snap" edge to the OrderSnap entity by IDs.
func (ou *OrderUpdate) AddOrderSnapIDs(ids ...int64) *OrderUpdate {
	ou.mutation.AddOrderSnapIDs(ids...)
	return ou
}

// AddOrderSnap adds the "order_snap" edges to the OrderSnap entity.
func (ou *OrderUpdate) AddOrderSnap(o ...*OrderSnap) *OrderUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddOrderSnapIDs(ids...)
}

// AddOrderSubIDs adds the "order_sub" edge to the OrderSub entity by IDs.
func (ou *OrderUpdate) AddOrderSubIDs(ids ...int64) *OrderUpdate {
	ou.mutation.AddOrderSubIDs(ids...)
	return ou
}

// AddOrderSub adds the "order_sub" edges to the OrderSub entity.
func (ou *OrderUpdate) AddOrderSub(o ...*OrderSub) *OrderUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddOrderSubIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearOrderSnap clears all "order_snap" edges to the OrderSnap entity.
func (ou *OrderUpdate) ClearOrderSnap() *OrderUpdate {
	ou.mutation.ClearOrderSnap()
	return ou
}

// RemoveOrderSnapIDs removes the "order_snap" edge to OrderSnap entities by IDs.
func (ou *OrderUpdate) RemoveOrderSnapIDs(ids ...int64) *OrderUpdate {
	ou.mutation.RemoveOrderSnapIDs(ids...)
	return ou
}

// RemoveOrderSnap removes "order_snap" edges to OrderSnap entities.
func (ou *OrderUpdate) RemoveOrderSnap(o ...*OrderSnap) *OrderUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveOrderSnapIDs(ids...)
}

// ClearOrderSub clears all "order_sub" edges to the OrderSub entity.
func (ou *OrderUpdate) ClearOrderSub() *OrderUpdate {
	ou.mutation.ClearOrderSub()
	return ou
}

// RemoveOrderSubIDs removes the "order_sub" edge to OrderSub entities by IDs.
func (ou *OrderUpdate) RemoveOrderSubIDs(ids ...int64) *OrderUpdate {
	ou.mutation.RemoveOrderSubIDs(ids...)
	return ou
}

// RemoveOrderSub removes "order_sub" edges to OrderSub entities.
func (ou *OrderUpdate) RemoveOrderSub(o ...*OrderSub) *OrderUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveOrderSubIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdateTime(); !ok {
		v := order.UpdateDefaultUpdateTime()
		ou.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.OrderNo(); ok {
		if err := order.OrderNoValidator(v); err != nil {
			return &ValidationError{Name: "order_no", err: fmt.Errorf(`model: validator failed for field "Order.order_no": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdateTime,
		})
	}
	if value, ok := ou.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeleteTime,
		})
	}
	if ou.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: order.FieldDeleteTime,
		})
	}
	if value, ok := ou.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldOrderNo,
		})
	}
	if value, ok := ou.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldTransactionID,
		})
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ou.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if ou.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ou.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalPrice,
		})
	}
	if value, ok := ou.mutation.AddedTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalPrice,
		})
	}
	if value, ok := ou.mutation.TotalCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldTotalCount,
		})
	}
	if value, ok := ou.mutation.AddedTotalCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldTotalCount,
		})
	}
	if value, ok := ou.mutation.FinalTotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldFinalTotalPrice,
		})
	}
	if value, ok := ou.mutation.AddedFinalTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldFinalTotalPrice,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ou.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if ou.mutation.OrderSnapCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedOrderSnapIDs(); len(nodes) > 0 && !ou.mutation.OrderSnapCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderSnapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.OrderSubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedOrderSubIDs(); len(nodes) > 0 && !ou.mutation.OrderSubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderSubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetUpdateTime sets the "update_time" field.
func (ouo *OrderUpdateOne) SetUpdateTime(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdateTime(t)
	return ouo
}

// SetDeleteTime sets the "delete_time" field.
func (ouo *OrderUpdateOne) SetDeleteTime(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetDeleteTime(t)
	return ouo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeleteTime(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetDeleteTime(*t)
	}
	return ouo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ouo *OrderUpdateOne) ClearDeleteTime() *OrderUpdateOne {
	ouo.mutation.ClearDeleteTime()
	return ouo
}

// SetOrderNo sets the "order_no" field.
func (ouo *OrderUpdateOne) SetOrderNo(s string) *OrderUpdateOne {
	ouo.mutation.SetOrderNo(s)
	return ouo
}

// SetTransactionID sets the "transaction_id" field.
func (ouo *OrderUpdateOne) SetTransactionID(s string) *OrderUpdateOne {
	ouo.mutation.SetTransactionID(s)
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(i int64) *OrderUpdateOne {
	ouo.mutation.ResetUserID()
	ouo.mutation.SetUserID(i)
	return ouo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(i *int64) *OrderUpdateOne {
	if i != nil {
		ouo.SetUserID(*i)
	}
	return ouo
}

// AddUserID adds i to the "user_id" field.
func (ouo *OrderUpdateOne) AddUserID(i int64) *OrderUpdateOne {
	ouo.mutation.AddUserID(i)
	return ouo
}

// ClearUserID clears the value of the "user_id" field.
func (ouo *OrderUpdateOne) ClearUserID() *OrderUpdateOne {
	ouo.mutation.ClearUserID()
	return ouo
}

// SetTotalPrice sets the "total_price" field.
func (ouo *OrderUpdateOne) SetTotalPrice(f float64) *OrderUpdateOne {
	ouo.mutation.ResetTotalPrice()
	ouo.mutation.SetTotalPrice(f)
	return ouo
}

// AddTotalPrice adds f to the "total_price" field.
func (ouo *OrderUpdateOne) AddTotalPrice(f float64) *OrderUpdateOne {
	ouo.mutation.AddTotalPrice(f)
	return ouo
}

// SetTotalCount sets the "total_count" field.
func (ouo *OrderUpdateOne) SetTotalCount(i int) *OrderUpdateOne {
	ouo.mutation.ResetTotalCount()
	ouo.mutation.SetTotalCount(i)
	return ouo
}

// AddTotalCount adds i to the "total_count" field.
func (ouo *OrderUpdateOne) AddTotalCount(i int) *OrderUpdateOne {
	ouo.mutation.AddTotalCount(i)
	return ouo
}

// SetFinalTotalPrice sets the "final_total_price" field.
func (ouo *OrderUpdateOne) SetFinalTotalPrice(f float64) *OrderUpdateOne {
	ouo.mutation.ResetFinalTotalPrice()
	ouo.mutation.SetFinalTotalPrice(f)
	return ouo
}

// AddFinalTotalPrice adds f to the "final_total_price" field.
func (ouo *OrderUpdateOne) AddFinalTotalPrice(f float64) *OrderUpdateOne {
	ouo.mutation.AddFinalTotalPrice(f)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(i int) *OrderUpdateOne {
	ouo.mutation.ResetStatus()
	ouo.mutation.SetStatus(i)
	return ouo
}

// AddStatus adds i to the "status" field.
func (ouo *OrderUpdateOne) AddStatus(i int) *OrderUpdateOne {
	ouo.mutation.AddStatus(i)
	return ouo
}

// AddOrderSnapIDs adds the "order_snap" edge to the OrderSnap entity by IDs.
func (ouo *OrderUpdateOne) AddOrderSnapIDs(ids ...int64) *OrderUpdateOne {
	ouo.mutation.AddOrderSnapIDs(ids...)
	return ouo
}

// AddOrderSnap adds the "order_snap" edges to the OrderSnap entity.
func (ouo *OrderUpdateOne) AddOrderSnap(o ...*OrderSnap) *OrderUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddOrderSnapIDs(ids...)
}

// AddOrderSubIDs adds the "order_sub" edge to the OrderSub entity by IDs.
func (ouo *OrderUpdateOne) AddOrderSubIDs(ids ...int64) *OrderUpdateOne {
	ouo.mutation.AddOrderSubIDs(ids...)
	return ouo
}

// AddOrderSub adds the "order_sub" edges to the OrderSub entity.
func (ouo *OrderUpdateOne) AddOrderSub(o ...*OrderSub) *OrderUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddOrderSubIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearOrderSnap clears all "order_snap" edges to the OrderSnap entity.
func (ouo *OrderUpdateOne) ClearOrderSnap() *OrderUpdateOne {
	ouo.mutation.ClearOrderSnap()
	return ouo
}

// RemoveOrderSnapIDs removes the "order_snap" edge to OrderSnap entities by IDs.
func (ouo *OrderUpdateOne) RemoveOrderSnapIDs(ids ...int64) *OrderUpdateOne {
	ouo.mutation.RemoveOrderSnapIDs(ids...)
	return ouo
}

// RemoveOrderSnap removes "order_snap" edges to OrderSnap entities.
func (ouo *OrderUpdateOne) RemoveOrderSnap(o ...*OrderSnap) *OrderUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveOrderSnapIDs(ids...)
}

// ClearOrderSub clears all "order_sub" edges to the OrderSub entity.
func (ouo *OrderUpdateOne) ClearOrderSub() *OrderUpdateOne {
	ouo.mutation.ClearOrderSub()
	return ouo
}

// RemoveOrderSubIDs removes the "order_sub" edge to OrderSub entities by IDs.
func (ouo *OrderUpdateOne) RemoveOrderSubIDs(ids ...int64) *OrderUpdateOne {
	ouo.mutation.RemoveOrderSubIDs(ids...)
	return ouo
}

// RemoveOrderSub removes "order_sub" edges to OrderSub entities.
func (ouo *OrderUpdateOne) RemoveOrderSub(o ...*OrderSub) *OrderUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveOrderSubIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdateTime(); !ok {
		v := order.UpdateDefaultUpdateTime()
		ouo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.OrderNo(); ok {
		if err := order.OrderNoValidator(v); err != nil {
			return &ValidationError{Name: "order_no", err: fmt.Errorf(`model: validator failed for field "Order.order_no": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdateTime,
		})
	}
	if value, ok := ouo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeleteTime,
		})
	}
	if ouo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: order.FieldDeleteTime,
		})
	}
	if value, ok := ouo.mutation.OrderNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldOrderNo,
		})
	}
	if value, ok := ouo.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldTransactionID,
		})
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ouo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if ouo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ouo.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalPrice,
		})
	}
	if value, ok := ouo.mutation.AddedTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalPrice,
		})
	}
	if value, ok := ouo.mutation.TotalCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldTotalCount,
		})
	}
	if value, ok := ouo.mutation.AddedTotalCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldTotalCount,
		})
	}
	if value, ok := ouo.mutation.FinalTotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldFinalTotalPrice,
		})
	}
	if value, ok := ouo.mutation.AddedFinalTotalPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldFinalTotalPrice,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if ouo.mutation.OrderSnapCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedOrderSnapIDs(); len(nodes) > 0 && !ouo.mutation.OrderSnapCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderSnapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.OrderSubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedOrderSubIDs(); len(nodes) > 0 && !ouo.mutation.OrderSubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderSubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
