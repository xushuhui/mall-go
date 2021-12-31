// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/spuimg"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpuImgDelete is the builder for deleting a SpuImg entity.
type SpuImgDelete struct {
	config
	hooks    []Hook
	mutation *SpuImgMutation
}

// Where appends a list predicates to the SpuImgDelete builder.
func (sid *SpuImgDelete) Where(ps ...predicate.SpuImg) *SpuImgDelete {
	sid.mutation.Where(ps...)
	return sid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sid *SpuImgDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sid.hooks) == 0 {
		affected, err = sid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpuImgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sid.mutation = mutation
			affected, err = sid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sid.hooks) - 1; i >= 0; i-- {
			if sid.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = sid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sid *SpuImgDelete) ExecX(ctx context.Context) int {
	n, err := sid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sid *SpuImgDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: spuimg.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: spuimg.FieldID,
			},
		},
	}
	if ps := sid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sid.driver, _spec)
}

// SpuImgDeleteOne is the builder for deleting a single SpuImg entity.
type SpuImgDeleteOne struct {
	sid *SpuImgDelete
}

// Exec executes the deletion query.
func (sido *SpuImgDeleteOne) Exec(ctx context.Context) error {
	n, err := sido.sid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spuimg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sido *SpuImgDeleteOne) ExecX(ctx context.Context) {
	sido.sid.ExecX(ctx)
}
