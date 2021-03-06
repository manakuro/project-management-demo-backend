// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/deletedtaskactivitytask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeletedTaskActivityTaskUpdate is the builder for updating DeletedTaskActivityTask entities.
type DeletedTaskActivityTaskUpdate struct {
	config
	hooks    []Hook
	mutation *DeletedTaskActivityTaskMutation
}

// Where appends a list predicates to the DeletedTaskActivityTaskUpdate builder.
func (dtatu *DeletedTaskActivityTaskUpdate) Where(ps ...predicate.DeletedTaskActivityTask) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.Where(ps...)
	return dtatu
}

// SetTaskActivityID sets the "task_activity_id" field.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTaskActivityID(u ulid.ID) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.SetTaskActivityID(u)
	return dtatu
}

// SetTaskID sets the "task_id" field.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTaskID(u ulid.ID) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.SetTaskID(u)
	return dtatu
}

// SetTaskActivityTaskID sets the "task_activity_task_id" field.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTaskActivityTaskID(u ulid.ID) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.SetTaskActivityTaskID(u)
	return dtatu
}

// SetTaskActivityTaskCreatedAt sets the "task_activity_task_created_at" field.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTaskActivityTaskCreatedAt(t time.Time) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.SetTaskActivityTaskCreatedAt(t)
	return dtatu
}

// SetTaskActivityTaskUpdatedAt sets the "task_activity_task_updated_at" field.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTaskActivityTaskUpdatedAt(t time.Time) *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.SetTaskActivityTaskUpdatedAt(t)
	return dtatu
}

// SetTask sets the "task" edge to the Task entity.
func (dtatu *DeletedTaskActivityTaskUpdate) SetTask(t *Task) *DeletedTaskActivityTaskUpdate {
	return dtatu.SetTaskID(t.ID)
}

// Mutation returns the DeletedTaskActivityTaskMutation object of the builder.
func (dtatu *DeletedTaskActivityTaskUpdate) Mutation() *DeletedTaskActivityTaskMutation {
	return dtatu.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (dtatu *DeletedTaskActivityTaskUpdate) ClearTask() *DeletedTaskActivityTaskUpdate {
	dtatu.mutation.ClearTask()
	return dtatu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dtatu *DeletedTaskActivityTaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dtatu.hooks) == 0 {
		if err = dtatu.check(); err != nil {
			return 0, err
		}
		affected, err = dtatu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtatu.check(); err != nil {
				return 0, err
			}
			dtatu.mutation = mutation
			affected, err = dtatu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dtatu.hooks) - 1; i >= 0; i-- {
			if dtatu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dtatu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtatu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtatu *DeletedTaskActivityTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := dtatu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtatu *DeletedTaskActivityTaskUpdate) Exec(ctx context.Context) error {
	_, err := dtatu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtatu *DeletedTaskActivityTaskUpdate) ExecX(ctx context.Context) {
	if err := dtatu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtatu *DeletedTaskActivityTaskUpdate) check() error {
	if _, ok := dtatu.mutation.TaskID(); dtatu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTaskActivityTask.task"`)
	}
	return nil
}

func (dtatu *DeletedTaskActivityTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedtaskactivitytask.Table,
			Columns: deletedtaskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedtaskactivitytask.FieldID,
			},
		},
	}
	if ps := dtatu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtatu.mutation.TaskActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityID,
		})
	}
	if value, ok := dtatu.mutation.TaskActivityTaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskID,
		})
	}
	if value, ok := dtatu.mutation.TaskActivityTaskCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskCreatedAt,
		})
	}
	if value, ok := dtatu.mutation.TaskActivityTaskUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskUpdatedAt,
		})
	}
	if dtatu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtaskactivitytask.TaskTable,
			Columns: []string{deletedtaskactivitytask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtatu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtaskactivitytask.TaskTable,
			Columns: []string{deletedtaskactivitytask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtatu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedtaskactivitytask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeletedTaskActivityTaskUpdateOne is the builder for updating a single DeletedTaskActivityTask entity.
type DeletedTaskActivityTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeletedTaskActivityTaskMutation
}

// SetTaskActivityID sets the "task_activity_id" field.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTaskActivityID(u ulid.ID) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.SetTaskActivityID(u)
	return dtatuo
}

// SetTaskID sets the "task_id" field.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTaskID(u ulid.ID) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.SetTaskID(u)
	return dtatuo
}

// SetTaskActivityTaskID sets the "task_activity_task_id" field.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTaskActivityTaskID(u ulid.ID) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.SetTaskActivityTaskID(u)
	return dtatuo
}

// SetTaskActivityTaskCreatedAt sets the "task_activity_task_created_at" field.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTaskActivityTaskCreatedAt(t time.Time) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.SetTaskActivityTaskCreatedAt(t)
	return dtatuo
}

// SetTaskActivityTaskUpdatedAt sets the "task_activity_task_updated_at" field.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTaskActivityTaskUpdatedAt(t time.Time) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.SetTaskActivityTaskUpdatedAt(t)
	return dtatuo
}

// SetTask sets the "task" edge to the Task entity.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SetTask(t *Task) *DeletedTaskActivityTaskUpdateOne {
	return dtatuo.SetTaskID(t.ID)
}

// Mutation returns the DeletedTaskActivityTaskMutation object of the builder.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) Mutation() *DeletedTaskActivityTaskMutation {
	return dtatuo.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) ClearTask() *DeletedTaskActivityTaskUpdateOne {
	dtatuo.mutation.ClearTask()
	return dtatuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) Select(field string, fields ...string) *DeletedTaskActivityTaskUpdateOne {
	dtatuo.fields = append([]string{field}, fields...)
	return dtatuo
}

// Save executes the query and returns the updated DeletedTaskActivityTask entity.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) Save(ctx context.Context) (*DeletedTaskActivityTask, error) {
	var (
		err  error
		node *DeletedTaskActivityTask
	)
	if len(dtatuo.hooks) == 0 {
		if err = dtatuo.check(); err != nil {
			return nil, err
		}
		node, err = dtatuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtatuo.check(); err != nil {
				return nil, err
			}
			dtatuo.mutation = mutation
			node, err = dtatuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dtatuo.hooks) - 1; i >= 0; i-- {
			if dtatuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dtatuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtatuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) SaveX(ctx context.Context) *DeletedTaskActivityTask {
	node, err := dtatuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := dtatuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) ExecX(ctx context.Context) {
	if err := dtatuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtatuo *DeletedTaskActivityTaskUpdateOne) check() error {
	if _, ok := dtatuo.mutation.TaskID(); dtatuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTaskActivityTask.task"`)
	}
	return nil
}

func (dtatuo *DeletedTaskActivityTaskUpdateOne) sqlSave(ctx context.Context) (_node *DeletedTaskActivityTask, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedtaskactivitytask.Table,
			Columns: deletedtaskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedtaskactivitytask.FieldID,
			},
		},
	}
	id, ok := dtatuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeletedTaskActivityTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dtatuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deletedtaskactivitytask.FieldID)
		for _, f := range fields {
			if !deletedtaskactivitytask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deletedtaskactivitytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dtatuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtatuo.mutation.TaskActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityID,
		})
	}
	if value, ok := dtatuo.mutation.TaskActivityTaskID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskID,
		})
	}
	if value, ok := dtatuo.mutation.TaskActivityTaskCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskCreatedAt,
		})
	}
	if value, ok := dtatuo.mutation.TaskActivityTaskUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtaskactivitytask.FieldTaskActivityTaskUpdatedAt,
		})
	}
	if dtatuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtaskactivitytask.TaskTable,
			Columns: []string{deletedtaskactivitytask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtatuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtaskactivitytask.TaskTable,
			Columns: []string{deletedtaskactivitytask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeletedTaskActivityTask{config: dtatuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dtatuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedtaskactivitytask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
