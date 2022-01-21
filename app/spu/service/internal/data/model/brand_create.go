// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/spu/service/internal/data/model/brand"
	"mall-go/app/spu/service/internal/data/model/spu"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BrandCreate is the builder for creating a Brand entity.
type BrandCreate struct {
	config
	mutation *BrandMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (bc *BrandCreate) SetCreateTime(t time.Time) *BrandCreate {
	bc.mutation.SetCreateTime(t)
	return bc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (bc *BrandCreate) SetNillableCreateTime(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetCreateTime(*t)
	}
	return bc
}

// SetUpdateTime sets the "update_time" field.
func (bc *BrandCreate) SetUpdateTime(t time.Time) *BrandCreate {
	bc.mutation.SetUpdateTime(t)
	return bc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (bc *BrandCreate) SetNillableUpdateTime(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetUpdateTime(*t)
	}
	return bc
}

// SetDeleteTime sets the "delete_time" field.
func (bc *BrandCreate) SetDeleteTime(t time.Time) *BrandCreate {
	bc.mutation.SetDeleteTime(t)
	return bc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (bc *BrandCreate) SetNillableDeleteTime(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetDeleteTime(*t)
	}
	return bc
}

// SetName sets the "name" field.
func (bc *BrandCreate) SetName(s string) *BrandCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BrandCreate) SetDescription(s string) *BrandCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (bc *BrandCreate) AddSpuIDs(ids ...int64) *BrandCreate {
	bc.mutation.AddSpuIDs(ids...)
	return bc
}

// AddSpu adds the "spu" edges to the Spu entity.
func (bc *BrandCreate) AddSpu(s ...*Spu) *BrandCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bc.AddSpuIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bc *BrandCreate) Mutation() *BrandMutation {
	return bc.mutation
}

// Save creates the Brand in the database.
func (bc *BrandCreate) Save(ctx context.Context) (*Brand, error) {
	var (
		err  error
		node *Brand
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrandCreate) SaveX(ctx context.Context) *Brand {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BrandCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BrandCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BrandCreate) defaults() {
	if _, ok := bc.mutation.CreateTime(); !ok {
		v := brand.DefaultCreateTime()
		bc.mutation.SetCreateTime(v)
	}
	if _, ok := bc.mutation.UpdateTime(); !ok {
		v := brand.DefaultUpdateTime()
		bc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BrandCreate) check() error {
	if _, ok := bc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := bc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "name"`)}
	}
	if _, ok := bc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "description"`)}
	}
	return nil
}

func (bc *BrandCreate) sqlSave(ctx context.Context) (*Brand, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (bc *BrandCreate) createSpec() (*Brand, *sqlgraph.CreateSpec) {
	var (
		_node = &Brand{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: brand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: brand.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := bc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := bc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := bc.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
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

// BrandCreateBulk is the builder for creating many Brand entities in bulk.
type BrandCreateBulk struct {
	config
	builders []*BrandCreate
}

// Save creates the Brand entities in the database.
func (bcb *BrandCreateBulk) Save(ctx context.Context) ([]*Brand, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Brand, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BrandMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BrandCreateBulk) SaveX(ctx context.Context) []*Brand {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BrandCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BrandCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}