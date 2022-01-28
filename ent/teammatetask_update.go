// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskUpdate is the builder for updating TeammateTask entities.
type TeammateTaskUpdate struct {
	config
	hooks    []Hook
	mutation *TeammateTaskMutation
}

// Where appends a list predicates to the TeammateTaskUpdate builder.
func (ttu *TeammateTaskUpdate) Where(ps ...predicate.TeammateTask) *TeammateTaskUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetTeammateID sets the "teammate_id" field.
func (ttu *TeammateTaskUpdate) SetTeammateID(u ulid.ID) *TeammateTaskUpdate {
	ttu.mutation.SetTeammateID(u)
	return ttu
}

// SetTaskID sets the "task_id" field.
func (ttu *TeammateTaskUpdate) SetTaskID(u ulid.ID) *TeammateTaskUpdate {
	ttu.mutation.SetTaskID(u)
	return ttu
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (ttu *TeammateTaskUpdate) SetTeammateTaskSectionID(u ulid.ID) *TeammateTaskUpdate {
	ttu.mutation.SetTeammateTaskSectionID(u)
	return ttu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttu *TeammateTaskUpdate) SetTeammate(t *Teammate) *TeammateTaskUpdate {
	return ttu.SetTeammateID(t.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ttu *TeammateTaskUpdate) SetTask(t *Task) *TeammateTaskUpdate {
	return ttu.SetTaskID(t.ID)
}

// SetTeammateTaskSection sets the "teammate_task_section" edge to the TeammateTaskSection entity.
func (ttu *TeammateTaskUpdate) SetTeammateTaskSection(t *TeammateTaskSection) *TeammateTaskUpdate {
	return ttu.SetTeammateTaskSectionID(t.ID)
}

// Mutation returns the TeammateTaskMutation object of the builder.
func (ttu *TeammateTaskUpdate) Mutation() *TeammateTaskMutation {
	return ttu.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (ttu *TeammateTaskUpdate) ClearTeammate() *TeammateTaskUpdate {
	ttu.mutation.ClearTeammate()
	return ttu
}

// ClearTask clears the "task" edge to the Task entity.
func (ttu *TeammateTaskUpdate) ClearTask() *TeammateTaskUpdate {
	ttu.mutation.ClearTask()
	return ttu
}

// ClearTeammateTaskSection clears the "teammate_task_section" edge to the TeammateTaskSection entity.
func (ttu *TeammateTaskUpdate) ClearTeammateTaskSection() *TeammateTaskUpdate {
	ttu.mutation.ClearTeammateTaskSection()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TeammateTaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttu.hooks) == 0 {
		if err = ttu.check(); err != nil {
			return 0, err
		}
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttu.check(); err != nil {
				return 0, err
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			if ttu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TeammateTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TeammateTaskUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TeammateTaskUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TeammateTaskUpdate) check() error {
	if _, ok := ttu.mutation.TeammateID(); ttu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	if _, ok := ttu.mutation.TaskID(); ttu.mutation.TaskCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task\"")
	}
	if _, ok := ttu.mutation.TeammateTaskSectionID(); ttu.mutation.TeammateTaskSectionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate_task_section\"")
	}
	return nil
}

func (ttu *TeammateTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teammatetask.Table,
			Columns: teammatetask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetask.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ttu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTable,
			Columns: []string{teammatetask.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTable,
			Columns: []string{teammatetask.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TaskTable,
			Columns: []string{teammatetask.TaskColumn},
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
	if nodes := ttu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TaskTable,
			Columns: []string{teammatetask.TaskColumn},
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
	if ttu.mutation.TeammateTaskSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTaskSectionTable,
			Columns: []string{teammatetask.TeammateTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasksection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TeammateTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTaskSectionTable,
			Columns: []string{teammatetask.TeammateTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasksection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teammatetask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TeammateTaskUpdateOne is the builder for updating a single TeammateTask entity.
type TeammateTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeammateTaskMutation
}

// SetTeammateID sets the "teammate_id" field.
func (ttuo *TeammateTaskUpdateOne) SetTeammateID(u ulid.ID) *TeammateTaskUpdateOne {
	ttuo.mutation.SetTeammateID(u)
	return ttuo
}

// SetTaskID sets the "task_id" field.
func (ttuo *TeammateTaskUpdateOne) SetTaskID(u ulid.ID) *TeammateTaskUpdateOne {
	ttuo.mutation.SetTaskID(u)
	return ttuo
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (ttuo *TeammateTaskUpdateOne) SetTeammateTaskSectionID(u ulid.ID) *TeammateTaskUpdateOne {
	ttuo.mutation.SetTeammateTaskSectionID(u)
	return ttuo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttuo *TeammateTaskUpdateOne) SetTeammate(t *Teammate) *TeammateTaskUpdateOne {
	return ttuo.SetTeammateID(t.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (ttuo *TeammateTaskUpdateOne) SetTask(t *Task) *TeammateTaskUpdateOne {
	return ttuo.SetTaskID(t.ID)
}

// SetTeammateTaskSection sets the "teammate_task_section" edge to the TeammateTaskSection entity.
func (ttuo *TeammateTaskUpdateOne) SetTeammateTaskSection(t *TeammateTaskSection) *TeammateTaskUpdateOne {
	return ttuo.SetTeammateTaskSectionID(t.ID)
}

// Mutation returns the TeammateTaskMutation object of the builder.
func (ttuo *TeammateTaskUpdateOne) Mutation() *TeammateTaskMutation {
	return ttuo.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (ttuo *TeammateTaskUpdateOne) ClearTeammate() *TeammateTaskUpdateOne {
	ttuo.mutation.ClearTeammate()
	return ttuo
}

// ClearTask clears the "task" edge to the Task entity.
func (ttuo *TeammateTaskUpdateOne) ClearTask() *TeammateTaskUpdateOne {
	ttuo.mutation.ClearTask()
	return ttuo
}

// ClearTeammateTaskSection clears the "teammate_task_section" edge to the TeammateTaskSection entity.
func (ttuo *TeammateTaskUpdateOne) ClearTeammateTaskSection() *TeammateTaskUpdateOne {
	ttuo.mutation.ClearTeammateTaskSection()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TeammateTaskUpdateOne) Select(field string, fields ...string) *TeammateTaskUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TeammateTask entity.
func (ttuo *TeammateTaskUpdateOne) Save(ctx context.Context) (*TeammateTask, error) {
	var (
		err  error
		node *TeammateTask
	)
	if len(ttuo.hooks) == 0 {
		if err = ttuo.check(); err != nil {
			return nil, err
		}
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttuo.check(); err != nil {
				return nil, err
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			if ttuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TeammateTaskUpdateOne) SaveX(ctx context.Context) *TeammateTask {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TeammateTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TeammateTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TeammateTaskUpdateOne) check() error {
	if _, ok := ttuo.mutation.TeammateID(); ttuo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	if _, ok := ttuo.mutation.TaskID(); ttuo.mutation.TaskCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task\"")
	}
	if _, ok := ttuo.mutation.TeammateTaskSectionID(); ttuo.mutation.TeammateTaskSectionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate_task_section\"")
	}
	return nil
}

func (ttuo *TeammateTaskUpdateOne) sqlSave(ctx context.Context) (_node *TeammateTask, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teammatetask.Table,
			Columns: teammatetask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetask.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TeammateTask.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teammatetask.FieldID)
		for _, f := range fields {
			if !teammatetask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teammatetask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ttuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTable,
			Columns: []string{teammatetask.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTable,
			Columns: []string{teammatetask.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TaskTable,
			Columns: []string{teammatetask.TaskColumn},
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
	if nodes := ttuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TaskTable,
			Columns: []string{teammatetask.TaskColumn},
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
	if ttuo.mutation.TeammateTaskSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTaskSectionTable,
			Columns: []string{teammatetask.TeammateTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasksection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TeammateTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetask.TeammateTaskSectionTable,
			Columns: []string{teammatetask.TeammateTaskSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasksection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TeammateTask{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teammatetask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
