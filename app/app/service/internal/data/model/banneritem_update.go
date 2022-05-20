// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/app/service/internal/data/model/banner"
	"mall-go/app/app/service/internal/data/model/banneritem"
	"mall-go/app/app/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerItemUpdate is the builder for updating BannerItem entities.
type BannerItemUpdate struct {
	config
	hooks    []Hook
	mutation *BannerItemMutation
}

// Where appends a list predicates to the BannerItemUpdate builder.
func (biu *BannerItemUpdate) Where(ps ...predicate.BannerItem) *BannerItemUpdate {
	biu.mutation.Where(ps...)
	return biu
}

// SetUpdateTime sets the "update_time" field.
func (biu *BannerItemUpdate) SetUpdateTime(t time.Time) *BannerItemUpdate {
	biu.mutation.SetUpdateTime(t)
	return biu
}

// SetDeleteTime sets the "delete_time" field.
func (biu *BannerItemUpdate) SetDeleteTime(t time.Time) *BannerItemUpdate {
	biu.mutation.SetDeleteTime(t)
	return biu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (biu *BannerItemUpdate) SetNillableDeleteTime(t *time.Time) *BannerItemUpdate {
	if t != nil {
		biu.SetDeleteTime(*t)
	}
	return biu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (biu *BannerItemUpdate) ClearDeleteTime() *BannerItemUpdate {
	biu.mutation.ClearDeleteTime()
	return biu
}

// SetImg sets the "img" field.
func (biu *BannerItemUpdate) SetImg(s string) *BannerItemUpdate {
	biu.mutation.SetImg(s)
	return biu
}

// SetKeyword sets the "keyword" field.
func (biu *BannerItemUpdate) SetKeyword(s string) *BannerItemUpdate {
	biu.mutation.SetKeyword(s)
	return biu
}

// SetType sets the "type" field.
func (biu *BannerItemUpdate) SetType(i int) *BannerItemUpdate {
	biu.mutation.ResetType()
	biu.mutation.SetType(i)
	return biu
}

// AddType adds i to the "type" field.
func (biu *BannerItemUpdate) AddType(i int) *BannerItemUpdate {
	biu.mutation.AddType(i)
	return biu
}

// SetBannerID sets the "banner_id" field.
func (biu *BannerItemUpdate) SetBannerID(i int64) *BannerItemUpdate {
	biu.mutation.SetBannerID(i)
	return biu
}

// SetNillableBannerID sets the "banner_id" field if the given value is not nil.
func (biu *BannerItemUpdate) SetNillableBannerID(i *int64) *BannerItemUpdate {
	if i != nil {
		biu.SetBannerID(*i)
	}
	return biu
}

// ClearBannerID clears the value of the "banner_id" field.
func (biu *BannerItemUpdate) ClearBannerID() *BannerItemUpdate {
	biu.mutation.ClearBannerID()
	return biu
}

// SetName sets the "name" field.
func (biu *BannerItemUpdate) SetName(s string) *BannerItemUpdate {
	biu.mutation.SetName(s)
	return biu
}

// SetBanner sets the "banner" edge to the Banner entity.
func (biu *BannerItemUpdate) SetBanner(b *Banner) *BannerItemUpdate {
	return biu.SetBannerID(b.ID)
}

// Mutation returns the BannerItemMutation object of the builder.
func (biu *BannerItemUpdate) Mutation() *BannerItemMutation {
	return biu.mutation
}

// ClearBanner clears the "banner" edge to the Banner entity.
func (biu *BannerItemUpdate) ClearBanner() *BannerItemUpdate {
	biu.mutation.ClearBanner()
	return biu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (biu *BannerItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	biu.defaults()
	if len(biu.hooks) == 0 {
		affected, err = biu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BannerItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			biu.mutation = mutation
			affected, err = biu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(biu.hooks) - 1; i >= 0; i-- {
			if biu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = biu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, biu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (biu *BannerItemUpdate) SaveX(ctx context.Context) int {
	affected, err := biu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (biu *BannerItemUpdate) Exec(ctx context.Context) error {
	_, err := biu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biu *BannerItemUpdate) ExecX(ctx context.Context) {
	if err := biu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (biu *BannerItemUpdate) defaults() {
	if _, ok := biu.mutation.UpdateTime(); !ok {
		v := banneritem.UpdateDefaultUpdateTime()
		biu.mutation.SetUpdateTime(v)
	}
}

func (biu *BannerItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banneritem.Table,
			Columns: banneritem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: banneritem.FieldID,
			},
		},
	}
	if ps := biu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := biu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldUpdateTime,
		})
	}
	if value, ok := biu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldDeleteTime,
		})
	}
	if biu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: banneritem.FieldDeleteTime,
		})
	}
	if value, ok := biu.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldImg,
		})
	}
	if value, ok := biu.mutation.Keyword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldKeyword,
		})
	}
	if value, ok := biu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: banneritem.FieldType,
		})
	}
	if value, ok := biu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: banneritem.FieldType,
		})
	}
	if value, ok := biu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldName,
		})
	}
	if biu.mutation.BannerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := biu.mutation.BannerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, biu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{banneritem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BannerItemUpdateOne is the builder for updating a single BannerItem entity.
type BannerItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BannerItemMutation
}

// SetUpdateTime sets the "update_time" field.
func (biuo *BannerItemUpdateOne) SetUpdateTime(t time.Time) *BannerItemUpdateOne {
	biuo.mutation.SetUpdateTime(t)
	return biuo
}

// SetDeleteTime sets the "delete_time" field.
func (biuo *BannerItemUpdateOne) SetDeleteTime(t time.Time) *BannerItemUpdateOne {
	biuo.mutation.SetDeleteTime(t)
	return biuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (biuo *BannerItemUpdateOne) SetNillableDeleteTime(t *time.Time) *BannerItemUpdateOne {
	if t != nil {
		biuo.SetDeleteTime(*t)
	}
	return biuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (biuo *BannerItemUpdateOne) ClearDeleteTime() *BannerItemUpdateOne {
	biuo.mutation.ClearDeleteTime()
	return biuo
}

// SetImg sets the "img" field.
func (biuo *BannerItemUpdateOne) SetImg(s string) *BannerItemUpdateOne {
	biuo.mutation.SetImg(s)
	return biuo
}

// SetKeyword sets the "keyword" field.
func (biuo *BannerItemUpdateOne) SetKeyword(s string) *BannerItemUpdateOne {
	biuo.mutation.SetKeyword(s)
	return biuo
}

// SetType sets the "type" field.
func (biuo *BannerItemUpdateOne) SetType(i int) *BannerItemUpdateOne {
	biuo.mutation.ResetType()
	biuo.mutation.SetType(i)
	return biuo
}

// AddType adds i to the "type" field.
func (biuo *BannerItemUpdateOne) AddType(i int) *BannerItemUpdateOne {
	biuo.mutation.AddType(i)
	return biuo
}

// SetBannerID sets the "banner_id" field.
func (biuo *BannerItemUpdateOne) SetBannerID(i int64) *BannerItemUpdateOne {
	biuo.mutation.SetBannerID(i)
	return biuo
}

// SetNillableBannerID sets the "banner_id" field if the given value is not nil.
func (biuo *BannerItemUpdateOne) SetNillableBannerID(i *int64) *BannerItemUpdateOne {
	if i != nil {
		biuo.SetBannerID(*i)
	}
	return biuo
}

// ClearBannerID clears the value of the "banner_id" field.
func (biuo *BannerItemUpdateOne) ClearBannerID() *BannerItemUpdateOne {
	biuo.mutation.ClearBannerID()
	return biuo
}

// SetName sets the "name" field.
func (biuo *BannerItemUpdateOne) SetName(s string) *BannerItemUpdateOne {
	biuo.mutation.SetName(s)
	return biuo
}

// SetBanner sets the "banner" edge to the Banner entity.
func (biuo *BannerItemUpdateOne) SetBanner(b *Banner) *BannerItemUpdateOne {
	return biuo.SetBannerID(b.ID)
}

// Mutation returns the BannerItemMutation object of the builder.
func (biuo *BannerItemUpdateOne) Mutation() *BannerItemMutation {
	return biuo.mutation
}

// ClearBanner clears the "banner" edge to the Banner entity.
func (biuo *BannerItemUpdateOne) ClearBanner() *BannerItemUpdateOne {
	biuo.mutation.ClearBanner()
	return biuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (biuo *BannerItemUpdateOne) Select(field string, fields ...string) *BannerItemUpdateOne {
	biuo.fields = append([]string{field}, fields...)
	return biuo
}

// Save executes the query and returns the updated BannerItem entity.
func (biuo *BannerItemUpdateOne) Save(ctx context.Context) (*BannerItem, error) {
	var (
		err  error
		node *BannerItem
	)
	biuo.defaults()
	if len(biuo.hooks) == 0 {
		node, err = biuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BannerItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			biuo.mutation = mutation
			node, err = biuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(biuo.hooks) - 1; i >= 0; i-- {
			if biuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = biuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, biuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (biuo *BannerItemUpdateOne) SaveX(ctx context.Context) *BannerItem {
	node, err := biuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (biuo *BannerItemUpdateOne) Exec(ctx context.Context) error {
	_, err := biuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biuo *BannerItemUpdateOne) ExecX(ctx context.Context) {
	if err := biuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (biuo *BannerItemUpdateOne) defaults() {
	if _, ok := biuo.mutation.UpdateTime(); !ok {
		v := banneritem.UpdateDefaultUpdateTime()
		biuo.mutation.SetUpdateTime(v)
	}
}

func (biuo *BannerItemUpdateOne) sqlSave(ctx context.Context) (_node *BannerItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   banneritem.Table,
			Columns: banneritem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: banneritem.FieldID,
			},
		},
	}
	id, ok := biuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "BannerItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := biuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banneritem.FieldID)
		for _, f := range fields {
			if !banneritem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != banneritem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := biuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := biuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldUpdateTime,
		})
	}
	if value, ok := biuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: banneritem.FieldDeleteTime,
		})
	}
	if biuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: banneritem.FieldDeleteTime,
		})
	}
	if value, ok := biuo.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldImg,
		})
	}
	if value, ok := biuo.mutation.Keyword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldKeyword,
		})
	}
	if value, ok := biuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: banneritem.FieldType,
		})
	}
	if value, ok := biuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: banneritem.FieldType,
		})
	}
	if value, ok := biuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: banneritem.FieldName,
		})
	}
	if biuo.mutation.BannerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := biuo.mutation.BannerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BannerItem{config: biuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, biuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{banneritem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
