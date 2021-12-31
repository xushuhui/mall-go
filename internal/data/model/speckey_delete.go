// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/speckey"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpecKeyDelete is the builder for deleting a SpecKey entity.
type SpecKeyDelete struct {
	config
	hooks    []Hook
	mutation *SpecKeyMutation
}

// Where appends a list predicates to the SpecKeyDelete builder.
func (skd *SpecKeyDelete) Where(ps ...predicate.SpecKey) *SpecKeyDelete {
	skd.mutation.Where(ps...)
	return skd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (skd *SpecKeyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(skd.hooks) == 0 {
		affected, err = skd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			skd.mutation = mutation
			affected, err = skd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(skd.hooks) - 1; i >= 0; i-- {
			if skd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = skd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, skd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (skd *SpecKeyDelete) ExecX(ctx context.Context) int {
	n, err := skd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (skd *SpecKeyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: speckey.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: speckey.FieldID,
			},
		},
	}
	if ps := skd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, skd.driver, _spec)
}

// SpecKeyDeleteOne is the builder for deleting a single SpecKey entity.
type SpecKeyDeleteOne struct {
	skd *SpecKeyDelete
}

// Exec executes the deletion query.
func (skdo *SpecKeyDeleteOne) Exec(ctx context.Context) error {
	n, err := skdo.skd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{speckey.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (skdo *SpecKeyDeleteOne) ExecX(ctx context.Context) {
	skdo.skd.ExecX(ctx)
}
