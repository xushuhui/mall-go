// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/coupon/service/internal/data/model/usercoupon"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCouponCreate is the builder for creating a UserCoupon entity.
type UserCouponCreate struct {
	config
	mutation *UserCouponMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ucc *UserCouponCreate) SetCreateTime(t time.Time) *UserCouponCreate {
	ucc.mutation.SetCreateTime(t)
	return ucc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ucc *UserCouponCreate) SetNillableCreateTime(t *time.Time) *UserCouponCreate {
	if t != nil {
		ucc.SetCreateTime(*t)
	}
	return ucc
}

// SetUpdateTime sets the "update_time" field.
func (ucc *UserCouponCreate) SetUpdateTime(t time.Time) *UserCouponCreate {
	ucc.mutation.SetUpdateTime(t)
	return ucc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ucc *UserCouponCreate) SetNillableUpdateTime(t *time.Time) *UserCouponCreate {
	if t != nil {
		ucc.SetUpdateTime(*t)
	}
	return ucc
}

// SetDeleteTime sets the "delete_time" field.
func (ucc *UserCouponCreate) SetDeleteTime(t time.Time) *UserCouponCreate {
	ucc.mutation.SetDeleteTime(t)
	return ucc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ucc *UserCouponCreate) SetNillableDeleteTime(t *time.Time) *UserCouponCreate {
	if t != nil {
		ucc.SetDeleteTime(*t)
	}
	return ucc
}

// SetUserID sets the "user_id" field.
func (ucc *UserCouponCreate) SetUserID(i int64) *UserCouponCreate {
	ucc.mutation.SetUserID(i)
	return ucc
}

// SetCouponID sets the "coupon_id" field.
func (ucc *UserCouponCreate) SetCouponID(i int64) *UserCouponCreate {
	ucc.mutation.SetCouponID(i)
	return ucc
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ucc *UserCouponCreate) SetNillableCouponID(i *int64) *UserCouponCreate {
	if i != nil {
		ucc.SetCouponID(*i)
	}
	return ucc
}

// SetStatus sets the "status" field.
func (ucc *UserCouponCreate) SetStatus(i int) *UserCouponCreate {
	ucc.mutation.SetStatus(i)
	return ucc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ucc *UserCouponCreate) SetNillableStatus(i *int) *UserCouponCreate {
	if i != nil {
		ucc.SetStatus(*i)
	}
	return ucc
}

// SetOrderID sets the "order_id" field.
func (ucc *UserCouponCreate) SetOrderID(i int) *UserCouponCreate {
	ucc.mutation.SetOrderID(i)
	return ucc
}

// Mutation returns the UserCouponMutation object of the builder.
func (ucc *UserCouponCreate) Mutation() *UserCouponMutation {
	return ucc.mutation
}

// Save creates the UserCoupon in the database.
func (ucc *UserCouponCreate) Save(ctx context.Context) (*UserCoupon, error) {
	var (
		err  error
		node *UserCoupon
	)
	ucc.defaults()
	if len(ucc.hooks) == 0 {
		if err = ucc.check(); err != nil {
			return nil, err
		}
		node, err = ucc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucc.check(); err != nil {
				return nil, err
			}
			ucc.mutation = mutation
			if node, err = ucc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucc.hooks) - 1; i >= 0; i-- {
			if ucc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ucc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UserCouponCreate) SaveX(ctx context.Context) *UserCoupon {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UserCouponCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UserCouponCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucc *UserCouponCreate) defaults() {
	if _, ok := ucc.mutation.CreateTime(); !ok {
		v := usercoupon.DefaultCreateTime()
		ucc.mutation.SetCreateTime(v)
	}
	if _, ok := ucc.mutation.UpdateTime(); !ok {
		v := usercoupon.DefaultUpdateTime()
		ucc.mutation.SetUpdateTime(v)
	}
	if _, ok := ucc.mutation.Status(); !ok {
		v := usercoupon.DefaultStatus
		ucc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UserCouponCreate) check() error {
	if _, ok := ucc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "UserCoupon.create_time"`)}
	}
	if _, ok := ucc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "UserCoupon.update_time"`)}
	}
	if _, ok := ucc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`model: missing required field "UserCoupon.user_id"`)}
	}
	if _, ok := ucc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`model: missing required field "UserCoupon.status"`)}
	}
	if _, ok := ucc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`model: missing required field "UserCoupon.order_id"`)}
	}
	return nil
}

func (ucc *UserCouponCreate) sqlSave(ctx context.Context) (*UserCoupon, error) {
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (ucc *UserCouponCreate) createSpec() (*UserCoupon, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCoupon{config: ucc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usercoupon.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: usercoupon.FieldID,
			},
		}
	)
	if value, ok := ucc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ucc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ucc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoupon.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := ucc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ucc.mutation.CouponID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: usercoupon.FieldCouponID,
		})
		_node.CouponID = value
	}
	if value, ok := ucc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ucc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usercoupon.FieldOrderID,
		})
		_node.OrderID = value
	}
	return _node, _spec
}

// UserCouponCreateBulk is the builder for creating many UserCoupon entities in bulk.
type UserCouponCreateBulk struct {
	config
	builders []*UserCouponCreate
}

// Save creates the UserCoupon entities in the database.
func (uccb *UserCouponCreateBulk) Save(ctx context.Context) ([]*UserCoupon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UserCoupon, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCouponMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UserCouponCreateBulk) SaveX(ctx context.Context) []*UserCoupon {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UserCouponCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UserCouponCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}
