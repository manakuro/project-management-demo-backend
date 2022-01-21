// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskColumnUpdate is the builder for updating TaskColumn entities.
type TaskColumnUpdate struct {
	config
	hooks    []Hook
	mutation *TaskColumnMutation
}

// Where appends a list predicates to the TaskColumnUpdate builder.
func (tcu *TaskColumnUpdate) Where(ps ...predicate.TaskColumn) *TaskColumnUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetName sets the "name" field.
func (tcu *TaskColumnUpdate) SetName(s string) *TaskColumnUpdate {
	tcu.mutation.SetName(s)
	return tcu
}

// SetType sets the "type" field.
func (tcu *TaskColumnUpdate) SetType(t taskcolumn.Type) *TaskColumnUpdate {
	tcu.mutation.SetType(t)
	return tcu
}

// Mutation returns the TaskColumnMutation object of the builder.
func (tcu *TaskColumnUpdate) Mutation() *TaskColumnMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TaskColumnUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tcu.hooks) == 0 {
		if err = tcu.check(); err != nil {
			return 0, err
		}
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcu.check(); err != nil {
				return 0, err
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TaskColumnUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TaskColumnUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TaskColumnUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcu *TaskColumnUpdate) check() error {
	if v, ok := tcu.mutation.Name(); ok {
		if err := taskcolumn.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tcu.mutation.GetType(); ok {
		if err := taskcolumn.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (tcu *TaskColumnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskcolumn.Table,
			Columns: taskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcolumn.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskcolumn.FieldName,
		})
	}
	if value, ok := tcu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: taskcolumn.FieldType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskColumnUpdateOne is the builder for updating a single TaskColumn entity.
type TaskColumnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskColumnMutation
}

// SetName sets the "name" field.
func (tcuo *TaskColumnUpdateOne) SetName(s string) *TaskColumnUpdateOne {
	tcuo.mutation.SetName(s)
	return tcuo
}

// SetType sets the "type" field.
func (tcuo *TaskColumnUpdateOne) SetType(t taskcolumn.Type) *TaskColumnUpdateOne {
	tcuo.mutation.SetType(t)
	return tcuo
}

// Mutation returns the TaskColumnMutation object of the builder.
func (tcuo *TaskColumnUpdateOne) Mutation() *TaskColumnMutation {
	return tcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TaskColumnUpdateOne) Select(field string, fields ...string) *TaskColumnUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TaskColumn entity.
func (tcuo *TaskColumnUpdateOne) Save(ctx context.Context) (*TaskColumn, error) {
	var (
		err  error
		node *TaskColumn
	)
	if len(tcuo.hooks) == 0 {
		if err = tcuo.check(); err != nil {
			return nil, err
		}
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcuo.check(); err != nil {
				return nil, err
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TaskColumnUpdateOne) SaveX(ctx context.Context) *TaskColumn {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TaskColumnUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TaskColumnUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcuo *TaskColumnUpdateOne) check() error {
	if v, ok := tcuo.mutation.Name(); ok {
		if err := taskcolumn.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tcuo.mutation.GetType(); ok {
		if err := taskcolumn.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (tcuo *TaskColumnUpdateOne) sqlSave(ctx context.Context) (_node *TaskColumn, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taskcolumn.Table,
			Columns: taskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcolumn.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TaskColumn.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskcolumn.FieldID)
		for _, f := range fields {
			if !taskcolumn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskcolumn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: taskcolumn.FieldName,
		})
	}
	if value, ok := tcuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: taskcolumn.FieldType,
		})
	}
	_node = &TaskColumn{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}