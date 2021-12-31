// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/coupontype"
	"mall-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponTypeDelete is the builder for deleting a CouponType entity.
type CouponTypeDelete struct {
	config
	hooks    []Hook
	mutation *CouponTypeMutation
}

// Where appends a list predicates to the CouponTypeDelete builder.
func (ctd *CouponTypeDelete) Where(ps ...predicate.CouponType) *CouponTypeDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CouponTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctd.hooks) == 0 {
		affected, err = ctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctd.mutation = mutation
			affected, err = ctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctd.hooks) - 1; i >= 0; i-- {
			if ctd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CouponTypeDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CouponTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coupontype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontype.FieldID,
			},
		},
	}
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
}

// CouponTypeDeleteOne is the builder for deleting a single CouponType entity.
type CouponTypeDeleteOne struct {
	ctd *CouponTypeDelete
}

// Exec executes the deletion query.
func (ctdo *CouponTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coupontype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CouponTypeDeleteOne) ExecX(ctx context.Context) {
	ctdo.ctd.ExecX(ctx)
}
