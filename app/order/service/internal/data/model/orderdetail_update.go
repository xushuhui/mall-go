// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/order/service/internal/data/model/orderdetail"
	"mall-go/app/order/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderDetailUpdate is the builder for updating OrderDetail entities.
type OrderDetailUpdate struct {
	config
	hooks    []Hook
	mutation *OrderDetailMutation
}

// Where appends a list predicates to the OrderDetailUpdate builder.
func (odu *OrderDetailUpdate) Where(ps ...predicate.OrderDetail) *OrderDetailUpdate {
	odu.mutation.Where(ps...)
	return odu
}

// SetUpdateTime sets the "update_time" field.
func (odu *OrderDetailUpdate) SetUpdateTime(t time.Time) *OrderDetailUpdate {
	odu.mutation.SetUpdateTime(t)
	return odu
}

// SetDeleteTime sets the "delete_time" field.
func (odu *OrderDetailUpdate) SetDeleteTime(t time.Time) *OrderDetailUpdate {
	odu.mutation.SetDeleteTime(t)
	return odu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillableDeleteTime(t *time.Time) *OrderDetailUpdate {
	if t != nil {
		odu.SetDeleteTime(*t)
	}
	return odu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (odu *OrderDetailUpdate) ClearDeleteTime() *OrderDetailUpdate {
	odu.mutation.ClearDeleteTime()
	return odu
}

// SetUserID sets the "user_id" field.
func (odu *OrderDetailUpdate) SetUserID(i int64) *OrderDetailUpdate {
	odu.mutation.ResetUserID()
	odu.mutation.SetUserID(i)
	return odu
}

// AddUserID adds i to the "user_id" field.
func (odu *OrderDetailUpdate) AddUserID(i int64) *OrderDetailUpdate {
	odu.mutation.AddUserID(i)
	return odu
}

// SetPayWay sets the "pay_way" field.
func (odu *OrderDetailUpdate) SetPayWay(i int) *OrderDetailUpdate {
	odu.mutation.ResetPayWay()
	odu.mutation.SetPayWay(i)
	return odu
}

// SetNillablePayWay sets the "pay_way" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillablePayWay(i *int) *OrderDetailUpdate {
	if i != nil {
		odu.SetPayWay(*i)
	}
	return odu
}

// AddPayWay adds i to the "pay_way" field.
func (odu *OrderDetailUpdate) AddPayWay(i int) *OrderDetailUpdate {
	odu.mutation.AddPayWay(i)
	return odu
}

// SetClientType sets the "client_type" field.
func (odu *OrderDetailUpdate) SetClientType(i int) *OrderDetailUpdate {
	odu.mutation.ResetClientType()
	odu.mutation.SetClientType(i)
	return odu
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (odu *OrderDetailUpdate) SetNillableClientType(i *int) *OrderDetailUpdate {
	if i != nil {
		odu.SetClientType(*i)
	}
	return odu
}

// AddClientType adds i to the "client_type" field.
func (odu *OrderDetailUpdate) AddClientType(i int) *OrderDetailUpdate {
	odu.mutation.AddClientType(i)
	return odu
}

// SetShipNo sets the "ship_no" field.
func (odu *OrderDetailUpdate) SetShipNo(s string) *OrderDetailUpdate {
	odu.mutation.SetShipNo(s)
	return odu
}

// Mutation returns the OrderDetailMutation object of the builder.
func (odu *OrderDetailUpdate) Mutation() *OrderDetailMutation {
	return odu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odu *OrderDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	odu.defaults()
	if len(odu.hooks) == 0 {
		affected, err = odu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odu.mutation = mutation
			affected, err = odu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odu.hooks) - 1; i >= 0; i-- {
			if odu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = odu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (odu *OrderDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := odu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odu *OrderDetailUpdate) Exec(ctx context.Context) error {
	_, err := odu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odu *OrderDetailUpdate) ExecX(ctx context.Context) {
	if err := odu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (odu *OrderDetailUpdate) defaults() {
	if _, ok := odu.mutation.UpdateTime(); !ok {
		v := orderdetail.UpdateDefaultUpdateTime()
		odu.mutation.SetUpdateTime(v)
	}
}

func (odu *OrderDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		},
	}
	if ps := odu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdateTime,
		})
	}
	if value, ok := odu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldDeleteTime,
		})
	}
	if odu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: orderdetail.FieldDeleteTime,
		})
	}
	if value, ok := odu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderdetail.FieldUserID,
		})
	}
	if value, ok := odu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderdetail.FieldUserID,
		})
	}
	if value, ok := odu.mutation.PayWay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldPayWay,
		})
	}
	if value, ok := odu.mutation.AddedPayWay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldPayWay,
		})
	}
	if value, ok := odu.mutation.ClientType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldClientType,
		})
	}
	if value, ok := odu.mutation.AddedClientType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldClientType,
		})
	}
	if value, ok := odu.mutation.ShipNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderdetail.FieldShipNo,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, odu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderDetailUpdateOne is the builder for updating a single OrderDetail entity.
type OrderDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderDetailMutation
}

// SetUpdateTime sets the "update_time" field.
func (oduo *OrderDetailUpdateOne) SetUpdateTime(t time.Time) *OrderDetailUpdateOne {
	oduo.mutation.SetUpdateTime(t)
	return oduo
}

// SetDeleteTime sets the "delete_time" field.
func (oduo *OrderDetailUpdateOne) SetDeleteTime(t time.Time) *OrderDetailUpdateOne {
	oduo.mutation.SetDeleteTime(t)
	return oduo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillableDeleteTime(t *time.Time) *OrderDetailUpdateOne {
	if t != nil {
		oduo.SetDeleteTime(*t)
	}
	return oduo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (oduo *OrderDetailUpdateOne) ClearDeleteTime() *OrderDetailUpdateOne {
	oduo.mutation.ClearDeleteTime()
	return oduo
}

// SetUserID sets the "user_id" field.
func (oduo *OrderDetailUpdateOne) SetUserID(i int64) *OrderDetailUpdateOne {
	oduo.mutation.ResetUserID()
	oduo.mutation.SetUserID(i)
	return oduo
}

// AddUserID adds i to the "user_id" field.
func (oduo *OrderDetailUpdateOne) AddUserID(i int64) *OrderDetailUpdateOne {
	oduo.mutation.AddUserID(i)
	return oduo
}

// SetPayWay sets the "pay_way" field.
func (oduo *OrderDetailUpdateOne) SetPayWay(i int) *OrderDetailUpdateOne {
	oduo.mutation.ResetPayWay()
	oduo.mutation.SetPayWay(i)
	return oduo
}

// SetNillablePayWay sets the "pay_way" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillablePayWay(i *int) *OrderDetailUpdateOne {
	if i != nil {
		oduo.SetPayWay(*i)
	}
	return oduo
}

// AddPayWay adds i to the "pay_way" field.
func (oduo *OrderDetailUpdateOne) AddPayWay(i int) *OrderDetailUpdateOne {
	oduo.mutation.AddPayWay(i)
	return oduo
}

// SetClientType sets the "client_type" field.
func (oduo *OrderDetailUpdateOne) SetClientType(i int) *OrderDetailUpdateOne {
	oduo.mutation.ResetClientType()
	oduo.mutation.SetClientType(i)
	return oduo
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (oduo *OrderDetailUpdateOne) SetNillableClientType(i *int) *OrderDetailUpdateOne {
	if i != nil {
		oduo.SetClientType(*i)
	}
	return oduo
}

// AddClientType adds i to the "client_type" field.
func (oduo *OrderDetailUpdateOne) AddClientType(i int) *OrderDetailUpdateOne {
	oduo.mutation.AddClientType(i)
	return oduo
}

// SetShipNo sets the "ship_no" field.
func (oduo *OrderDetailUpdateOne) SetShipNo(s string) *OrderDetailUpdateOne {
	oduo.mutation.SetShipNo(s)
	return oduo
}

// Mutation returns the OrderDetailMutation object of the builder.
func (oduo *OrderDetailUpdateOne) Mutation() *OrderDetailMutation {
	return oduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oduo *OrderDetailUpdateOne) Select(field string, fields ...string) *OrderDetailUpdateOne {
	oduo.fields = append([]string{field}, fields...)
	return oduo
}

// Save executes the query and returns the updated OrderDetail entity.
func (oduo *OrderDetailUpdateOne) Save(ctx context.Context) (*OrderDetail, error) {
	var (
		err  error
		node *OrderDetail
	)
	oduo.defaults()
	if len(oduo.hooks) == 0 {
		node, err = oduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oduo.mutation = mutation
			node, err = oduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oduo.hooks) - 1; i >= 0; i-- {
			if oduo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = oduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oduo *OrderDetailUpdateOne) SaveX(ctx context.Context) *OrderDetail {
	node, err := oduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oduo *OrderDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := oduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oduo *OrderDetailUpdateOne) ExecX(ctx context.Context) {
	if err := oduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oduo *OrderDetailUpdateOne) defaults() {
	if _, ok := oduo.mutation.UpdateTime(); !ok {
		v := orderdetail.UpdateDefaultUpdateTime()
		oduo.mutation.SetUpdateTime(v)
	}
}

func (oduo *OrderDetailUpdateOne) sqlSave(ctx context.Context) (_node *OrderDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		},
	}
	id, ok := oduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "OrderDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderdetail.FieldID)
		for _, f := range fields {
			if !orderdetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != orderdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oduo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdateTime,
		})
	}
	if value, ok := oduo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldDeleteTime,
		})
	}
	if oduo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: orderdetail.FieldDeleteTime,
		})
	}
	if value, ok := oduo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderdetail.FieldUserID,
		})
	}
	if value, ok := oduo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderdetail.FieldUserID,
		})
	}
	if value, ok := oduo.mutation.PayWay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldPayWay,
		})
	}
	if value, ok := oduo.mutation.AddedPayWay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldPayWay,
		})
	}
	if value, ok := oduo.mutation.ClientType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldClientType,
		})
	}
	if value, ok := oduo.mutation.AddedClientType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldClientType,
		})
	}
	if value, ok := oduo.mutation.ShipNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderdetail.FieldShipNo,
		})
	}
	_node = &OrderDetail{config: oduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
