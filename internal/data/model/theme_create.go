// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/internal/data/model/spu"
	"mall-go/internal/data/model/theme"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThemeCreate is the builder for creating a Theme entity.
type ThemeCreate struct {
	config
	mutation *ThemeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *ThemeCreate) SetCreateTime(t time.Time) *ThemeCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *ThemeCreate) SetNillableCreateTime(t *time.Time) *ThemeCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *ThemeCreate) SetUpdateTime(t time.Time) *ThemeCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *ThemeCreate) SetNillableUpdateTime(t *time.Time) *ThemeCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetDeleteTime sets the "delete_time" field.
func (tc *ThemeCreate) SetDeleteTime(t time.Time) *ThemeCreate {
	tc.mutation.SetDeleteTime(t)
	return tc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (tc *ThemeCreate) SetNillableDeleteTime(t *time.Time) *ThemeCreate {
	if t != nil {
		tc.SetDeleteTime(*t)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *ThemeCreate) SetTitle(s string) *ThemeCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *ThemeCreate) SetDescription(s string) *ThemeCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetName sets the "name" field.
func (tc *ThemeCreate) SetName(s string) *ThemeCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetTplName sets the "tpl_name" field.
func (tc *ThemeCreate) SetTplName(s string) *ThemeCreate {
	tc.mutation.SetTplName(s)
	return tc
}

// SetEntranceImg sets the "entrance_img" field.
func (tc *ThemeCreate) SetEntranceImg(s string) *ThemeCreate {
	tc.mutation.SetEntranceImg(s)
	return tc
}

// SetExtend sets the "extend" field.
func (tc *ThemeCreate) SetExtend(s string) *ThemeCreate {
	tc.mutation.SetExtend(s)
	return tc
}

// SetInternalTopImg sets the "internal_top_img" field.
func (tc *ThemeCreate) SetInternalTopImg(s string) *ThemeCreate {
	tc.mutation.SetInternalTopImg(s)
	return tc
}

// SetTitleImg sets the "title_img" field.
func (tc *ThemeCreate) SetTitleImg(s string) *ThemeCreate {
	tc.mutation.SetTitleImg(s)
	return tc
}

// SetOnline sets the "online" field.
func (tc *ThemeCreate) SetOnline(i int8) *ThemeCreate {
	tc.mutation.SetOnline(i)
	return tc
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (tc *ThemeCreate) AddSpuIDs(ids ...int64) *ThemeCreate {
	tc.mutation.AddSpuIDs(ids...)
	return tc
}

// AddSpu adds the "spu" edges to the Spu entity.
func (tc *ThemeCreate) AddSpu(s ...*Spu) *ThemeCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddSpuIDs(ids...)
}

// Mutation returns the ThemeMutation object of the builder.
func (tc *ThemeCreate) Mutation() *ThemeMutation {
	return tc.mutation
}

// Save creates the Theme in the database.
func (tc *ThemeCreate) Save(ctx context.Context) (*Theme, error) {
	var (
		err  error
		node *Theme
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThemeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ThemeCreate) SaveX(ctx context.Context) *Theme {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ThemeCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ThemeCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ThemeCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := theme.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := theme.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ThemeCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "title"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "description"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "name"`)}
	}
	if _, ok := tc.mutation.TplName(); !ok {
		return &ValidationError{Name: "tpl_name", err: errors.New(`model: missing required field "tpl_name"`)}
	}
	if _, ok := tc.mutation.EntranceImg(); !ok {
		return &ValidationError{Name: "entrance_img", err: errors.New(`model: missing required field "entrance_img"`)}
	}
	if _, ok := tc.mutation.Extend(); !ok {
		return &ValidationError{Name: "extend", err: errors.New(`model: missing required field "extend"`)}
	}
	if _, ok := tc.mutation.InternalTopImg(); !ok {
		return &ValidationError{Name: "internal_top_img", err: errors.New(`model: missing required field "internal_top_img"`)}
	}
	if _, ok := tc.mutation.TitleImg(); !ok {
		return &ValidationError{Name: "title_img", err: errors.New(`model: missing required field "title_img"`)}
	}
	if _, ok := tc.mutation.Online(); !ok {
		return &ValidationError{Name: "online", err: errors.New(`model: missing required field "online"`)}
	}
	return nil
}

func (tc *ThemeCreate) sqlSave(ctx context.Context) (*Theme, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (tc *ThemeCreate) createSpec() (*Theme, *sqlgraph.CreateSpec) {
	var (
		_node = &Theme{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: theme.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: theme.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.TplName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldTplName,
		})
		_node.TplName = value
	}
	if value, ok := tc.mutation.EntranceImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldEntranceImg,
		})
		_node.EntranceImg = value
	}
	if value, ok := tc.mutation.Extend(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldExtend,
		})
		_node.Extend = value
	}
	if value, ok := tc.mutation.InternalTopImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldInternalTopImg,
		})
		_node.InternalTopImg = value
	}
	if value, ok := tc.mutation.TitleImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldTitleImg,
		})
		_node.TitleImg = value
	}
	if value, ok := tc.mutation.Online(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: theme.FieldOnline,
		})
		_node.Online = value
	}
	if nodes := tc.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   theme.SpuTable,
			Columns: theme.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThemeCreateBulk is the builder for creating many Theme entities in bulk.
type ThemeCreateBulk struct {
	config
	builders []*ThemeCreate
}

// Save creates the Theme entities in the database.
func (tcb *ThemeCreateBulk) Save(ctx context.Context) ([]*Theme, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Theme, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThemeMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ThemeCreateBulk) SaveX(ctx context.Context) []*Theme {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ThemeCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ThemeCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}