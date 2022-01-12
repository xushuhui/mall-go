// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/app/app/service/internal/data/model/coupon"
	"mall-go/app/app/service/internal/data/model/predicate"
	"mall-go/app/app/service/internal/data/model/usercoupon"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCouponUpdate is the builder for updating UserCoupon entities.
type UserCouponUpdate struct {
	config
	hooks    []Hook
	mutation *UserCouponMutation
}

// Where appends a list predicates to the UserCouponUpdate builder.
func (ucu *UserCouponUpdate) Where(ps ...predicate.UserCoupon) *UserCouponUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUpdateTime sets the "update_time" field.
func (ucu *UserCouponUpdate) SetUpdateTime(t time.Time) *UserCouponUpdate {
	ucu.mutation.SetUpdateTime(t)
	return ucu
}

// SetDeleteTime sets the "delete_time" field.
func (ucu *UserCouponUpdate) SetDeleteTime(t time.Time) *UserCouponUpdate {
	ucu.mutation.SetDeleteTime(t)
	return ucu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ucu *UserCouponUpdate) SetNillableDeleteTime(t *time.Time) *UserCouponUpdate {
	if t != nil {
		ucu.SetDeleteTime(*t)
	}
	return ucu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ucu *UserCouponUpdate) ClearDeleteTime() *UserCouponUpdate {
	ucu.mutation.ClearDeleteTime()
	return ucu
}

// SetUserID sets the "user_id" field.
func (ucu *UserCouponUpdate) SetUserID(i int64) *UserCouponUpdate {
	ucu.mutation.ResetUserID()
	ucu.mutation.SetUserID(i)
	return ucu
}

// AddUserID adds i to the "user_id" field.
func (ucu *UserCouponUpdate) AddUserID(i int64) *UserCouponUpdate {
	ucu.mutation.AddUserID(i)
	return ucu
}

// SetCouponID sets the "coupon_id" field.
func (ucu *UserCouponUpdate) SetCouponID(i int64) *UserCouponUpdate {
	ucu.mutation.SetCouponID(i)
	return ucu
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ucu *UserCouponUpdate) SetNillableCouponID(i *int64) *UserCouponUpdate {
	if i != nil {
		ucu.SetCouponID(*i)
	}
	return ucu
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ucu *UserCouponUpdate) ClearCouponID() *UserCouponUpdate {
	ucu.mutation.ClearCouponID()
	return ucu
}

// SetStatus sets the "status" field.
func (ucu *UserCouponUpdate) SetStatus(i int) *UserCouponUpdate {
	ucu.mutation.ResetStatus()
	ucu.mutation.SetStatus(i)
	return ucu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ucu *UserCouponUpdate) SetNillableStatus(i *int) *UserCouponUpdate {
	if i != nil {
		ucu.SetStatus(*i)
	}
	return ucu
}

// AddStatus adds i to the "status" field.
func (ucu *UserCouponUpdate) AddStatus(i int) *UserCouponUpdate {
	ucu.mutation.AddStatus(i)
	return ucu
}

// SetOrderID sets the "order_id" field.
func (ucu *UserCouponUpdate) SetOrderID(i int) *UserCouponUpdate {
	ucu.mutation.ResetOrderID()
	ucu.mutation.SetOrderID(i)
	return ucu
}

// AddOrderID adds i to the "order_id" field.
func (ucu *UserCouponUpdate) AddOrderID(i int) *UserCouponUpdate {
	ucu.mutation.AddOrderID(i)
	return ucu
}

// SetCoupon sets the "coupon" edge to the Coupon entity.
func (ucu *UserCouponUpdate) SetCoupon(c *Coupon) *UserCouponUpdate {
	return ucu.SetCouponID(c.ID)
}

// Mutation returns the UserCouponMutation object of the builder.
func (ucu *UserCouponUpdate) Mutation() *UserCouponMutation {
	return ucu.mutation
}

// ClearCoupon clears the "coupon" edge to the Coupon entity.
func (ucu *UserCouponUpdate) ClearCoupon() *UserCouponUpdate {
	ucu.mutation.ClearCoupon()
	return ucu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserCouponUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ucu.defaults()
	if len(ucu.hooks) == 0 {
		affected, err = ucu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucu.mutation = mutation
			affected, err = ucu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucu.hooks) - 1; i >= 0; i-- {
			if ucu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ucu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserCouponUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserCouponUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserCouponUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucu *UserCouponUpdate) defaults() {
	if _, ok := ucu.mutation.UpdateTime(); !ok {
		v := usercoupon.UpdateDefaultUpdateTime()
		ucu.mutation.SetUpdateTime(v)
	}
}

func (ucu *UserCouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoupon.Table,
			Columns: usercoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercoupon.FieldID,
			},
		},
	}
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldUpdateTime,
		})
	}
	if value, ok := ucu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldDeleteTime,
		})
	}
	if ucu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercoupon.FieldDeleteTime,
		})
	}
	if value, ok := ucu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldUserID,
		})
	}
	if value, ok := ucu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldUserID,
		})
	}
	if value, ok := ucu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldStatus,
		})
	}
	if value, ok := ucu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldStatus,
		})
	}
	if value, ok := ucu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldOrderID,
		})
	}
	if value, ok := ucu.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldOrderID,
		})
	}
	if ucu.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercoupon.CouponTable,
			Columns: []string{usercoupon.CouponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.CouponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercoupon.CouponTable,
			Columns: []string{usercoupon.CouponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserCouponUpdateOne is the builder for updating a single UserCoupon entity.
type UserCouponUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCouponMutation
}

// SetUpdateTime sets the "update_time" field.
func (ucuo *UserCouponUpdateOne) SetUpdateTime(t time.Time) *UserCouponUpdateOne {
	ucuo.mutation.SetUpdateTime(t)
	return ucuo
}

// SetDeleteTime sets the "delete_time" field.
func (ucuo *UserCouponUpdateOne) SetDeleteTime(t time.Time) *UserCouponUpdateOne {
	ucuo.mutation.SetDeleteTime(t)
	return ucuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ucuo *UserCouponUpdateOne) SetNillableDeleteTime(t *time.Time) *UserCouponUpdateOne {
	if t != nil {
		ucuo.SetDeleteTime(*t)
	}
	return ucuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ucuo *UserCouponUpdateOne) ClearDeleteTime() *UserCouponUpdateOne {
	ucuo.mutation.ClearDeleteTime()
	return ucuo
}

// SetUserID sets the "user_id" field.
func (ucuo *UserCouponUpdateOne) SetUserID(i int64) *UserCouponUpdateOne {
	ucuo.mutation.ResetUserID()
	ucuo.mutation.SetUserID(i)
	return ucuo
}

// AddUserID adds i to the "user_id" field.
func (ucuo *UserCouponUpdateOne) AddUserID(i int64) *UserCouponUpdateOne {
	ucuo.mutation.AddUserID(i)
	return ucuo
}

// SetCouponID sets the "coupon_id" field.
func (ucuo *UserCouponUpdateOne) SetCouponID(i int64) *UserCouponUpdateOne {
	ucuo.mutation.SetCouponID(i)
	return ucuo
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ucuo *UserCouponUpdateOne) SetNillableCouponID(i *int64) *UserCouponUpdateOne {
	if i != nil {
		ucuo.SetCouponID(*i)
	}
	return ucuo
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ucuo *UserCouponUpdateOne) ClearCouponID() *UserCouponUpdateOne {
	ucuo.mutation.ClearCouponID()
	return ucuo
}

// SetStatus sets the "status" field.
func (ucuo *UserCouponUpdateOne) SetStatus(i int) *UserCouponUpdateOne {
	ucuo.mutation.ResetStatus()
	ucuo.mutation.SetStatus(i)
	return ucuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ucuo *UserCouponUpdateOne) SetNillableStatus(i *int) *UserCouponUpdateOne {
	if i != nil {
		ucuo.SetStatus(*i)
	}
	return ucuo
}

// AddStatus adds i to the "status" field.
func (ucuo *UserCouponUpdateOne) AddStatus(i int) *UserCouponUpdateOne {
	ucuo.mutation.AddStatus(i)
	return ucuo
}

// SetOrderID sets the "order_id" field.
func (ucuo *UserCouponUpdateOne) SetOrderID(i int) *UserCouponUpdateOne {
	ucuo.mutation.ResetOrderID()
	ucuo.mutation.SetOrderID(i)
	return ucuo
}

// AddOrderID adds i to the "order_id" field.
func (ucuo *UserCouponUpdateOne) AddOrderID(i int) *UserCouponUpdateOne {
	ucuo.mutation.AddOrderID(i)
	return ucuo
}

// SetCoupon sets the "coupon" edge to the Coupon entity.
func (ucuo *UserCouponUpdateOne) SetCoupon(c *Coupon) *UserCouponUpdateOne {
	return ucuo.SetCouponID(c.ID)
}

// Mutation returns the UserCouponMutation object of the builder.
func (ucuo *UserCouponUpdateOne) Mutation() *UserCouponMutation {
	return ucuo.mutation
}

// ClearCoupon clears the "coupon" edge to the Coupon entity.
func (ucuo *UserCouponUpdateOne) ClearCoupon() *UserCouponUpdateOne {
	ucuo.mutation.ClearCoupon()
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserCouponUpdateOne) Select(field string, fields ...string) *UserCouponUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserCoupon entity.
func (ucuo *UserCouponUpdateOne) Save(ctx context.Context) (*UserCoupon, error) {
	var (
		err  error
		node *UserCoupon
	)
	ucuo.defaults()
	if len(ucuo.hooks) == 0 {
		node, err = ucuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucuo.mutation = mutation
			node, err = ucuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucuo.hooks) - 1; i >= 0; i-- {
			if ucuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ucuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserCouponUpdateOne) SaveX(ctx context.Context) *UserCoupon {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserCouponUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserCouponUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucuo *UserCouponUpdateOne) defaults() {
	if _, ok := ucuo.mutation.UpdateTime(); !ok {
		v := usercoupon.UpdateDefaultUpdateTime()
		ucuo.mutation.SetUpdateTime(v)
	}
}

func (ucuo *UserCouponUpdateOne) sqlSave(ctx context.Context) (_node *UserCoupon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoupon.Table,
			Columns: usercoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercoupon.FieldID,
			},
		},
	}
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserCoupon.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercoupon.FieldID)
		for _, f := range fields {
			if !usercoupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != usercoupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldUpdateTime,
		})
	}
	if value, ok := ucuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldDeleteTime,
		})
	}
	if ucuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercoupon.FieldDeleteTime,
		})
	}
	if value, ok := ucuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldUserID,
		})
	}
	if value, ok := ucuo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldUserID,
		})
	}
	if value, ok := ucuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldStatus,
		})
	}
	if value, ok := ucuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldStatus,
		})
	}
	if value, ok := ucuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldOrderID,
		})
	}
	if value, ok := ucuo.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldOrderID,
		})
	}
	if ucuo.mutation.CouponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercoupon.CouponTable,
			Columns: []string{usercoupon.CouponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.CouponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercoupon.CouponTable,
			Columns: []string{usercoupon.CouponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: coupon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCoupon{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
