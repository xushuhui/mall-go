// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/userwalletdetail"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserWalletDetailDelete is the builder for deleting a UserWalletDetail entity.
type UserWalletDetailDelete struct {
	config
	hooks    []Hook
	mutation *UserWalletDetailMutation
}

// Where appends a list predicates to the UserWalletDetailDelete builder.
func (uwdd *UserWalletDetailDelete) Where(ps ...predicate.UserWalletDetail) *UserWalletDetailDelete {
	uwdd.mutation.Where(ps...)
	return uwdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uwdd *UserWalletDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uwdd.hooks) == 0 {
		affected, err = uwdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWalletDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwdd.mutation = mutation
			affected, err = uwdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwdd.hooks) - 1; i >= 0; i-- {
			if uwdd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = uwdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwdd *UserWalletDetailDelete) ExecX(ctx context.Context) int {
	n, err := uwdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uwdd *UserWalletDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userwalletdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userwalletdetail.FieldID,
			},
		},
	}
	if ps := uwdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uwdd.driver, _spec)
}

// UserWalletDetailDeleteOne is the builder for deleting a single UserWalletDetail entity.
type UserWalletDetailDeleteOne struct {
	uwdd *UserWalletDetailDelete
}

// Exec executes the deletion query.
func (uwddo *UserWalletDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := uwddo.uwdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userwalletdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uwddo *UserWalletDetailDeleteOne) ExecX(ctx context.Context) {
	uwddo.uwdd.ExecX(ctx)
}
