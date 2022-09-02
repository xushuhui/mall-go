// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"mall-go/app/app/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GridCategoryUpdate is the builder for updating GridCategory entities.
type GridCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *GridCategoryMutation
}

// Where appends a list predicates to the GridCategoryUpdate builder.
func (gcu *GridCategoryUpdate) Where(ps ...predicate.GridCategory) *GridCategoryUpdate {
	gcu.mutation.Where(ps...)
	return gcu
}

// SetUpdateTime sets the "update_time" field.
func (gcu *GridCategoryUpdate) SetUpdateTime(t time.Time) *GridCategoryUpdate {
	gcu.mutation.SetUpdateTime(t)
	return gcu
}

// SetDeleteTime sets the "delete_time" field.
func (gcu *GridCategoryUpdate) SetDeleteTime(t time.Time) *GridCategoryUpdate {
	gcu.mutation.SetDeleteTime(t)
	return gcu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (gcu *GridCategoryUpdate) SetNillableDeleteTime(t *time.Time) *GridCategoryUpdate {
	if t != nil {
		gcu.SetDeleteTime(*t)
	}
	return gcu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (gcu *GridCategoryUpdate) ClearDeleteTime() *GridCategoryUpdate {
	gcu.mutation.ClearDeleteTime()
	return gcu
}

// SetTitle sets the "title" field.
func (gcu *GridCategoryUpdate) SetTitle(s string) *GridCategoryUpdate {
	gcu.mutation.SetTitle(s)
	return gcu
}

// SetImg sets the "img" field.
func (gcu *GridCategoryUpdate) SetImg(s string) *GridCategoryUpdate {
	gcu.mutation.SetImg(s)
	return gcu
}

// SetName sets the "name" field.
func (gcu *GridCategoryUpdate) SetName(s string) *GridCategoryUpdate {
	gcu.mutation.SetName(s)
	return gcu
}

// SetCategoryID sets the "category_id" field.
func (gcu *GridCategoryUpdate) SetCategoryID(i int) *GridCategoryUpdate {
	gcu.mutation.ResetCategoryID()
	gcu.mutation.SetCategoryID(i)
	return gcu
}

// AddCategoryID adds i to the "category_id" field.
func (gcu *GridCategoryUpdate) AddCategoryID(i int) *GridCategoryUpdate {
	gcu.mutation.AddCategoryID(i)
	return gcu
}

// SetRootCategoryID sets the "root_category_id" field.
func (gcu *GridCategoryUpdate) SetRootCategoryID(i int) *GridCategoryUpdate {
	gcu.mutation.ResetRootCategoryID()
	gcu.mutation.SetRootCategoryID(i)
	return gcu
}

// AddRootCategoryID adds i to the "root_category_id" field.
func (gcu *GridCategoryUpdate) AddRootCategoryID(i int) *GridCategoryUpdate {
	gcu.mutation.AddRootCategoryID(i)
	return gcu
}

// Mutation returns the GridCategoryMutation object of the builder.
func (gcu *GridCategoryUpdate) Mutation() *GridCategoryMutation {
	return gcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcu *GridCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gcu.defaults()
	if len(gcu.hooks) == 0 {
		affected, err = gcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GridCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcu.mutation = mutation
			affected, err = gcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcu.hooks) - 1; i >= 0; i-- {
			if gcu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = gcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcu *GridCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := gcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcu *GridCategoryUpdate) Exec(ctx context.Context) error {
	_, err := gcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcu *GridCategoryUpdate) ExecX(ctx context.Context) {
	if err := gcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcu *GridCategoryUpdate) defaults() {
	if _, ok := gcu.mutation.UpdateTime(); !ok {
		v := gridcategory.UpdateDefaultUpdateTime()
		gcu.mutation.SetUpdateTime(v)
	}
}

func (gcu *GridCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gridcategory.Table,
			Columns: gridcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gridcategory.FieldID,
			},
		},
	}
	if ps := gcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldUpdateTime,
		})
	}
	if value, ok := gcu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldDeleteTime,
		})
	}
	if gcu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: gridcategory.FieldDeleteTime,
		})
	}
	if value, ok := gcu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldTitle,
		})
	}
	if value, ok := gcu.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldImg,
		})
	}
	if value, ok := gcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldName,
		})
	}
	if value, ok := gcu.mutation.CategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldCategoryID,
		})
	}
	if value, ok := gcu.mutation.AddedCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldCategoryID,
		})
	}
	if value, ok := gcu.mutation.RootCategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldRootCategoryID,
		})
	}
	if value, ok := gcu.mutation.AddedRootCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldRootCategoryID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gridcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GridCategoryUpdateOne is the builder for updating a single GridCategory entity.
type GridCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GridCategoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (gcuo *GridCategoryUpdateOne) SetUpdateTime(t time.Time) *GridCategoryUpdateOne {
	gcuo.mutation.SetUpdateTime(t)
	return gcuo
}

// SetDeleteTime sets the "delete_time" field.
func (gcuo *GridCategoryUpdateOne) SetDeleteTime(t time.Time) *GridCategoryUpdateOne {
	gcuo.mutation.SetDeleteTime(t)
	return gcuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (gcuo *GridCategoryUpdateOne) SetNillableDeleteTime(t *time.Time) *GridCategoryUpdateOne {
	if t != nil {
		gcuo.SetDeleteTime(*t)
	}
	return gcuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (gcuo *GridCategoryUpdateOne) ClearDeleteTime() *GridCategoryUpdateOne {
	gcuo.mutation.ClearDeleteTime()
	return gcuo
}

// SetTitle sets the "title" field.
func (gcuo *GridCategoryUpdateOne) SetTitle(s string) *GridCategoryUpdateOne {
	gcuo.mutation.SetTitle(s)
	return gcuo
}

// SetImg sets the "img" field.
func (gcuo *GridCategoryUpdateOne) SetImg(s string) *GridCategoryUpdateOne {
	gcuo.mutation.SetImg(s)
	return gcuo
}

// SetName sets the "name" field.
func (gcuo *GridCategoryUpdateOne) SetName(s string) *GridCategoryUpdateOne {
	gcuo.mutation.SetName(s)
	return gcuo
}

// SetCategoryID sets the "category_id" field.
func (gcuo *GridCategoryUpdateOne) SetCategoryID(i int) *GridCategoryUpdateOne {
	gcuo.mutation.ResetCategoryID()
	gcuo.mutation.SetCategoryID(i)
	return gcuo
}

// AddCategoryID adds i to the "category_id" field.
func (gcuo *GridCategoryUpdateOne) AddCategoryID(i int) *GridCategoryUpdateOne {
	gcuo.mutation.AddCategoryID(i)
	return gcuo
}

// SetRootCategoryID sets the "root_category_id" field.
func (gcuo *GridCategoryUpdateOne) SetRootCategoryID(i int) *GridCategoryUpdateOne {
	gcuo.mutation.ResetRootCategoryID()
	gcuo.mutation.SetRootCategoryID(i)
	return gcuo
}

// AddRootCategoryID adds i to the "root_category_id" field.
func (gcuo *GridCategoryUpdateOne) AddRootCategoryID(i int) *GridCategoryUpdateOne {
	gcuo.mutation.AddRootCategoryID(i)
	return gcuo
}

// Mutation returns the GridCategoryMutation object of the builder.
func (gcuo *GridCategoryUpdateOne) Mutation() *GridCategoryMutation {
	return gcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcuo *GridCategoryUpdateOne) Select(field string, fields ...string) *GridCategoryUpdateOne {
	gcuo.fields = append([]string{field}, fields...)
	return gcuo
}

// Save executes the query and returns the updated GridCategory entity.
func (gcuo *GridCategoryUpdateOne) Save(ctx context.Context) (*GridCategory, error) {
	var (
		err  error
		node *GridCategory
	)
	gcuo.defaults()
	if len(gcuo.hooks) == 0 {
		node, err = gcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GridCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcuo.mutation = mutation
			node, err = gcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gcuo.hooks) - 1; i >= 0; i-- {
			if gcuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = gcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GridCategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GridCategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcuo *GridCategoryUpdateOne) SaveX(ctx context.Context) *GridCategory {
	node, err := gcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcuo *GridCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := gcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcuo *GridCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := gcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcuo *GridCategoryUpdateOne) defaults() {
	if _, ok := gcuo.mutation.UpdateTime(); !ok {
		v := gridcategory.UpdateDefaultUpdateTime()
		gcuo.mutation.SetUpdateTime(v)
	}
}

func (gcuo *GridCategoryUpdateOne) sqlSave(ctx context.Context) (_node *GridCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gridcategory.Table,
			Columns: gridcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gridcategory.FieldID,
			},
		},
	}
	id, ok := gcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "GridCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gridcategory.FieldID)
		for _, f := range fields {
			if !gridcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != gridcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldUpdateTime,
		})
	}
	if value, ok := gcuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldDeleteTime,
		})
	}
	if gcuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: gridcategory.FieldDeleteTime,
		})
	}
	if value, ok := gcuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldTitle,
		})
	}
	if value, ok := gcuo.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldImg,
		})
	}
	if value, ok := gcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldName,
		})
	}
	if value, ok := gcuo.mutation.CategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldCategoryID,
		})
	}
	if value, ok := gcuo.mutation.AddedCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldCategoryID,
		})
	}
	if value, ok := gcuo.mutation.RootCategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldRootCategoryID,
		})
	}
	if value, ok := gcuo.mutation.AddedRootCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldRootCategoryID,
		})
	}
	_node = &GridCategory{config: gcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gridcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
