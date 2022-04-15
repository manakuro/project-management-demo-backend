// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/archivedtaskactivity"
	"project-management-demo-backend/ent/archivedtaskactivitytask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedTaskActivityTaskUpdate is the builder for updating ArchivedTaskActivityTask entities.
type ArchivedTaskActivityTaskUpdate struct {
	config
	hooks    []Hook
	mutation *ArchivedTaskActivityTaskMutation
}

// Where appends a list predicates to the ArchivedTaskActivityTaskUpdate builder.
func (atatu *ArchivedTaskActivityTaskUpdate) Where(ps ...predicate.ArchivedTaskActivityTask) *ArchivedTaskActivityTaskUpdate {
	atatu.mutation.Where(ps...)
	return atatu
}

// SetArchivedTaskActivityID sets the "archived_task_activity_id" field.
func (atatu *ArchivedTaskActivityTaskUpdate) SetArchivedTaskActivityID(u ulid.ID) *ArchivedTaskActivityTaskUpdate {
	atatu.mutation.SetArchivedTaskActivityID(u)
	return atatu
}

// SetTaskID sets the "task_id" field.
func (atatu *ArchivedTaskActivityTaskUpdate) SetTaskID(u ulid.ID) *ArchivedTaskActivityTaskUpdate {
	atatu.mutation.SetTaskID(u)
	return atatu
}

// SetTask sets the "task" edge to the Task entity.
func (atatu *ArchivedTaskActivityTaskUpdate) SetTask(t *Task) *ArchivedTaskActivityTaskUpdate {
	return atatu.SetTaskID(t.ID)
}

// SetArchivedTaskActivity sets the "archivedTaskActivity" edge to the ArchivedTaskActivity entity.
func (atatu *ArchivedTaskActivityTaskUpdate) SetArchivedTaskActivity(a *ArchivedTaskActivity) *ArchivedTaskActivityTaskUpdate {
	return atatu.SetArchivedTaskActivityID(a.ID)
}

// Mutation returns the ArchivedTaskActivityTaskMutation object of the builder.
func (atatu *ArchivedTaskActivityTaskUpdate) Mutation() *ArchivedTaskActivityTaskMutation {
	return atatu.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (atatu *ArchivedTaskActivityTaskUpdate) ClearTask() *ArchivedTaskActivityTaskUpdate {
	atatu.mutation.ClearTask()
	return atatu
}

// ClearArchivedTaskActivity clears the "archivedTaskActivity" edge to the ArchivedTaskActivity entity.
func (atatu *ArchivedTaskActivityTaskUpdate) ClearArchivedTaskActivity() *ArchivedTaskActivityTaskUpdate {
	atatu.mutation.ClearArchivedTaskActivity()
	return atatu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atatu *ArchivedTaskActivityTaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atatu.hooks) == 0 {
		if err = atatu.check(); err != nil {
			return 0, err
		}
		affected, err = atatu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedTaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atatu.check(); err != nil {
				return 0, err
			}
			atatu.mutation = mutation
			affected, err = atatu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atatu.hooks) - 1; i >= 0; i-- {
			if atatu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atatu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atatu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atatu *ArchivedTaskActivityTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := atatu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atatu *ArchivedTaskActivityTaskUpdate) Exec(ctx context.Context) error {
	_, err := atatu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atatu *ArchivedTaskActivityTaskUpdate) ExecX(ctx context.Context) {
	if err := atatu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atatu *ArchivedTaskActivityTaskUpdate) check() error {
	if _, ok := atatu.mutation.TaskID(); atatu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedTaskActivityTask.task"`)
	}
	if _, ok := atatu.mutation.ArchivedTaskActivityID(); atatu.mutation.ArchivedTaskActivityCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedTaskActivityTask.archivedTaskActivity"`)
	}
	return nil
}

func (atatu *ArchivedTaskActivityTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedtaskactivitytask.Table,
			Columns: archivedtaskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedtaskactivitytask.FieldID,
			},
		},
	}
	if ps := atatu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atatu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskTable,
			Columns: []string{archivedtaskactivitytask.TaskColumn},
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
	if nodes := atatu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskTable,
			Columns: []string{archivedtaskactivitytask.TaskColumn},
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
	if atatu.mutation.ArchivedTaskActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.ArchivedTaskActivityTable,
			Columns: []string{archivedtaskactivitytask.ArchivedTaskActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atatu.mutation.ArchivedTaskActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.ArchivedTaskActivityTable,
			Columns: []string{archivedtaskactivitytask.ArchivedTaskActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atatu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{archivedtaskactivitytask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ArchivedTaskActivityTaskUpdateOne is the builder for updating a single ArchivedTaskActivityTask entity.
type ArchivedTaskActivityTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArchivedTaskActivityTaskMutation
}

// SetArchivedTaskActivityID sets the "archived_task_activity_id" field.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) SetArchivedTaskActivityID(u ulid.ID) *ArchivedTaskActivityTaskUpdateOne {
	atatuo.mutation.SetArchivedTaskActivityID(u)
	return atatuo
}

// SetTaskID sets the "task_id" field.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) SetTaskID(u ulid.ID) *ArchivedTaskActivityTaskUpdateOne {
	atatuo.mutation.SetTaskID(u)
	return atatuo
}

// SetTask sets the "task" edge to the Task entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) SetTask(t *Task) *ArchivedTaskActivityTaskUpdateOne {
	return atatuo.SetTaskID(t.ID)
}

// SetArchivedTaskActivity sets the "archivedTaskActivity" edge to the ArchivedTaskActivity entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) SetArchivedTaskActivity(a *ArchivedTaskActivity) *ArchivedTaskActivityTaskUpdateOne {
	return atatuo.SetArchivedTaskActivityID(a.ID)
}

// Mutation returns the ArchivedTaskActivityTaskMutation object of the builder.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) Mutation() *ArchivedTaskActivityTaskMutation {
	return atatuo.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) ClearTask() *ArchivedTaskActivityTaskUpdateOne {
	atatuo.mutation.ClearTask()
	return atatuo
}

// ClearArchivedTaskActivity clears the "archivedTaskActivity" edge to the ArchivedTaskActivity entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) ClearArchivedTaskActivity() *ArchivedTaskActivityTaskUpdateOne {
	atatuo.mutation.ClearArchivedTaskActivity()
	return atatuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) Select(field string, fields ...string) *ArchivedTaskActivityTaskUpdateOne {
	atatuo.fields = append([]string{field}, fields...)
	return atatuo
}

// Save executes the query and returns the updated ArchivedTaskActivityTask entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) Save(ctx context.Context) (*ArchivedTaskActivityTask, error) {
	var (
		err  error
		node *ArchivedTaskActivityTask
	)
	if len(atatuo.hooks) == 0 {
		if err = atatuo.check(); err != nil {
			return nil, err
		}
		node, err = atatuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedTaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atatuo.check(); err != nil {
				return nil, err
			}
			atatuo.mutation = mutation
			node, err = atatuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atatuo.hooks) - 1; i >= 0; i-- {
			if atatuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atatuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atatuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) SaveX(ctx context.Context) *ArchivedTaskActivityTask {
	node, err := atatuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := atatuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) ExecX(ctx context.Context) {
	if err := atatuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atatuo *ArchivedTaskActivityTaskUpdateOne) check() error {
	if _, ok := atatuo.mutation.TaskID(); atatuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedTaskActivityTask.task"`)
	}
	if _, ok := atatuo.mutation.ArchivedTaskActivityID(); atatuo.mutation.ArchivedTaskActivityCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedTaskActivityTask.archivedTaskActivity"`)
	}
	return nil
}

func (atatuo *ArchivedTaskActivityTaskUpdateOne) sqlSave(ctx context.Context) (_node *ArchivedTaskActivityTask, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedtaskactivitytask.Table,
			Columns: archivedtaskactivitytask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedtaskactivitytask.FieldID,
			},
		},
	}
	id, ok := atatuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArchivedTaskActivityTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atatuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, archivedtaskactivitytask.FieldID)
		for _, f := range fields {
			if !archivedtaskactivitytask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != archivedtaskactivitytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atatuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atatuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskTable,
			Columns: []string{archivedtaskactivitytask.TaskColumn},
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
	if nodes := atatuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskTable,
			Columns: []string{archivedtaskactivitytask.TaskColumn},
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
	if atatuo.mutation.ArchivedTaskActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.ArchivedTaskActivityTable,
			Columns: []string{archivedtaskactivitytask.ArchivedTaskActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atatuo.mutation.ArchivedTaskActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.ArchivedTaskActivityTable,
			Columns: []string{archivedtaskactivitytask.ArchivedTaskActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ArchivedTaskActivityTask{config: atatuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atatuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{archivedtaskactivitytask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
