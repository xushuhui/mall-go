// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/internal/data/model/banner"
	"mall-go/internal/data/model/banneritem"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerItemCreate is the builder for creating a BannerItem entity.
type BannerItemCreate struct {
	config
	mutation *BannerItemMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (bic *BannerItemCreate) SetCreateTime(t time.Time) *BannerItemCreate {
	bic.mutation.SetCreateTime(t)
	return bic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (bic *BannerItemCreate) SetNillableCreateTime(t *time.Time) *BannerItemCreate {
	if t != nil {
		bic.SetCreateTime(*t)
	}
	return bic
}

// SetUpdateTime sets the "update_time" field.
func (bic *BannerItemCreate) SetUpdateTime(t time.Time) *BannerItemCreate {
	bic.mutation.SetUpdateTime(t)
	return bic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (bic *BannerItemCreate) SetNillableUpdateTime(t *time.Time) *BannerItemCreate {
	if t != nil {
		bic.SetUpdateTime(*t)
	}
	return bic
}

// SetDeleteTime sets the "delete_time" field.
func (bic *BannerItemCreate) SetDeleteTime(t time.Time) *BannerItemCreate {
	bic.mutation.SetDeleteTime(t)
	return bic
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (bic *BannerItemCreate) SetNillableDeleteTime(t *time.Time) *BannerItemCreate {
	if t != nil {
		bic.SetDeleteTime(*t)
	}
	return bic
}

// SetImg sets the "img" field.
func (bic *BannerItemCreate) SetImg(s string) *BannerItemCreate {
	bic.mutation.SetImg(s)
	return bic
}

// SetKeyword sets the "keyword" field.
func (bic *BannerItemCreate) SetKeyword(s string) *BannerItemCreate {
	bic.mutation.SetKeyword(s)
	return bic
}

// SetType sets the "type" field.
func (bic *BannerItemCreate) SetType(i int) *BannerItemCreate {
	bic.mutation.SetType(i)
	return bic
}

// SetBannerID sets the "banner_id" field.
func (bic *BannerItemCreate) SetBannerID(i int64) *BannerItemCreate {
	bic.mutation.SetBannerID(i)
	return bic
}

// SetNillableBannerID sets the "banner_id" field if the given value is not nil.
func (bic *BannerItemCreate) SetNillableBannerID(i *int64) *BannerItemCreate {
	if i != nil {
		bic.SetBannerID(*i)
	}
	return bic
}

// SetName sets the "name" field.
func (bic *BannerItemCreate) SetName(s string) *BannerItemCreate {
	bic.mutation.SetName(s)
	return bic
}

// SetBanner sets the "banner" edge to the Banner entity.
func (bic *BannerItemCreate) SetBanner(b *Banner) *BannerItemCreate {
	return bic.SetBannerID(b.ID)
}

// Mutation returns the BannerItemMutation object of the builder.
func (bic *BannerItemCreate) Mutation() *BannerItemMutation {
	return bic.mutation
}

// Save creates the BannerItem in the database.
func (bic *BannerItemCreate) Save(ctx context.Context) (*BannerItem, error) {
	var (
		err  error
		node *BannerItem
	)
	bic.defaults()
	if len(bic.hooks) == 0 {
		if err = bic.check(); err != nil {
			return nil, err
		}
		node, err = bic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BannerItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bic.check(); err != nil {
				return nil, err
			}
			bic.mutation = mutation
			if node, err = bic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bic.hooks) - 1; i >= 0; i-- {
			if bic.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = bic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bic *BannerItemCreate) SaveX(ctx context.Context) *BannerItem {
	v, err := bic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bic *BannerItemCreate) Exec(ctx context.Context) error {
	_, err := bic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bic *BannerItemCreate) ExecX(ctx context.Context) {
	if err := bic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bic *BannerItemCreate) defaults() {
	if _, ok := bic.mutation.CreateTime(); !ok {
		v := banneritem.DefaultCreateTime()
		bic.mutation.SetCreateTime(v)
	}
	if _, ok := bic.mutation.UpdateTime(); !ok {
		v := banneritem.DefaultUpdateTime()
		bic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bic *BannerItemCreate) check() error {
	if _, ok := bic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := bic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := bic.mutation.Img(); !ok {
		return &ValidationError{Name: "img", err: errors.New(`model: missing required field "img"`)}
	}
	if _, ok := bic.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`model: missing required field "keyword"`)}
	}
	if _, ok := bic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "type"`)}
	}
	if _, ok := bic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "name"`)}
	}
	return nil
}

func (bic *BannerItemCreate) sqlSave(ctx context.Context) (*BannerItem, error) {
	_node, _spec := bic.createSpec()
	if err := sqlgraph.CreateNode(ctx, bic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (bic *BannerItemCreate) createSpec() (*BannerItem, *sqlgraph.CreateSpec) {
	var (
		_node = &BannerItem{config: bic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: banneritem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: banneritem.FieldID,
			},
		}
	)
	if value, ok := bic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := bic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := bic.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := bic.mutation.Img(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldImg,
		})
		_node.Img = value
	}
	if value, ok := bic.mutation.Keyword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldKeyword,
		})
		_node.Keyword = value
	}
	if value, ok := bic.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: banneritem.FieldType,
		})
		_node.Type = value
	}
	if value, ok := bic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldName,
		})
		_node.Name = value
	}
	if nodes := bic.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   banneritem.BannerTable,
			Columns: []string{banneritem.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: banner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BannerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BannerItemCreateBulk is the builder for creating many BannerItem entities in bulk.
type BannerItemCreateBulk struct {
	config
	builders []*BannerItemCreate
}

// Save creates the BannerItem entities in the database.
func (bicb *BannerItemCreateBulk) Save(ctx context.Context) ([]*BannerItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bicb.builders))
	nodes := make([]*BannerItem, len(bicb.builders))
	mutators := make([]Mutator, len(bicb.builders))
	for i := range bicb.builders {
		func(i int, root context.Context) {
			builder := bicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BannerItemMutation)
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
					_, err = mutators[i+1].Mutate(root, bicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bicb *BannerItemCreateBulk) SaveX(ctx context.Context) []*BannerItem {
	v, err := bicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bicb *BannerItemCreateBulk) Exec(ctx context.Context) error {
	_, err := bicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bicb *BannerItemCreateBulk) ExecX(ctx context.Context) {
	if err := bicb.Exec(ctx); err != nil {
		panic(err)
	}
}
