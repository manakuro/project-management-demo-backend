// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/ent/teammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskLikeUpdate is the builder for updating TaskLike entities.
type TaskLikeUpdate struct {
	config
	hooks    []Hook
	mutation *TaskLikeMutation
}

// Where appends a list predicates to the TaskLikeUpdate builder.
func (tlu *TaskLikeUpdate) Where(ps ...predicate.TaskLike) *TaskLikeUpdate {
	tlu.mutation.Where(ps...)
	return tlu
}

// SetTaskID sets the "task_id" field.
func (tlu *TaskLikeUpdate) SetTaskID(u ulid.ID) *TaskLikeUpdate {
	tlu.mutation.SetTaskID(u)
	return tlu
}

// SetTeammateID sets the "teammate_id" field.
func (tlu *TaskLikeUpdate) SetTeammateID(u ulid.ID) *TaskLikeUpdate {
	tlu.mutation.SetTeammateID(u)
	return tlu
}

// SetTask sets the "task" edge to the Task entity.
func (tlu *TaskLikeUpdate) SetTask(t *Task) *TaskLikeUpdate {
	return tlu.SetTaskID(t.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tlu *TaskLikeUpdate) SetTeammate(t *Teammate) *TaskLikeUpdate {
	return tlu.SetTeammateID(t.ID)
}

// Mutation returns the TaskLikeMutation object of the builder.
func (tlu *TaskLikeUpdate) Mutation() *TaskLikeMutation {
	return tlu.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (tlu *TaskLikeUpdate) ClearTask() *TaskLikeUpdate {
	tlu.mutation.ClearTask()
	return tlu
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (tlu *TaskLikeUpdate) ClearTeammate() *TaskLikeUpdate {
	tlu.mutation.ClearTeammate()
	return tlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlu *TaskLikeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlu.hooks) == 0 {
		if err = tlu.check(); err != nil {
			return 0, err
		}
		affected, err = tlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlu.check(); err != nil {
				return 0, err
			}
			tlu.mutation = mutation
			affected, err = tlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlu.hooks) - 1; i >= 0; i-- {
			if tlu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlu *TaskLikeUpdate) SaveX(ctx context.Context) int {
	affected, err := tlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlu *TaskLikeUpdate) Exec(ctx context.Context) error {
	_, err := tlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlu *TaskLikeUpdate) ExecX(ctx context.Context) {
	if err := tlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlu *TaskLikeUpdate) check() error {
	if _, ok := tlu.mutation.TaskID(); tlu.mutation.TaskCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task\"")
	}
	if _, ok := tlu.mutation.TeammateID(); tlu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (tlu *TaskLikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklike.Table,
			Columns: tasklike.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklike.FieldID,
			},
		},
	}
	if ps := tlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tlu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TaskTable,
			Columns: []string{tasklike.TaskColumn},
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
	if nodes := tlu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TaskTable,
			Columns: []string{tasklike.TaskColumn},
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
	if tlu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TeammateTable,
			Columns: []string{tasklike.TeammateColumn},
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
	if nodes := tlu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TeammateTable,
			Columns: []string{tasklike.TeammateColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskLikeUpdateOne is the builder for updating a single TaskLike entity.
type TaskLikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskLikeMutation
}

// SetTaskID sets the "task_id" field.
func (tluo *TaskLikeUpdateOne) SetTaskID(u ulid.ID) *TaskLikeUpdateOne {
	tluo.mutation.SetTaskID(u)
	return tluo
}

// SetTeammateID sets the "teammate_id" field.
func (tluo *TaskLikeUpdateOne) SetTeammateID(u ulid.ID) *TaskLikeUpdateOne {
	tluo.mutation.SetTeammateID(u)
	return tluo
}

// SetTask sets the "task" edge to the Task entity.
func (tluo *TaskLikeUpdateOne) SetTask(t *Task) *TaskLikeUpdateOne {
	return tluo.SetTaskID(t.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tluo *TaskLikeUpdateOne) SetTeammate(t *Teammate) *TaskLikeUpdateOne {
	return tluo.SetTeammateID(t.ID)
}

// Mutation returns the TaskLikeMutation object of the builder.
func (tluo *TaskLikeUpdateOne) Mutation() *TaskLikeMutation {
	return tluo.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (tluo *TaskLikeUpdateOne) ClearTask() *TaskLikeUpdateOne {
	tluo.mutation.ClearTask()
	return tluo
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (tluo *TaskLikeUpdateOne) ClearTeammate() *TaskLikeUpdateOne {
	tluo.mutation.ClearTeammate()
	return tluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tluo *TaskLikeUpdateOne) Select(field string, fields ...string) *TaskLikeUpdateOne {
	tluo.fields = append([]string{field}, fields...)
	return tluo
}

// Save executes the query and returns the updated TaskLike entity.
func (tluo *TaskLikeUpdateOne) Save(ctx context.Context) (*TaskLike, error) {
	var (
		err  error
		node *TaskLike
	)
	if len(tluo.hooks) == 0 {
		if err = tluo.check(); err != nil {
			return nil, err
		}
		node, err = tluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tluo.check(); err != nil {
				return nil, err
			}
			tluo.mutation = mutation
			node, err = tluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tluo.hooks) - 1; i >= 0; i-- {
			if tluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tluo *TaskLikeUpdateOne) SaveX(ctx context.Context) *TaskLike {
	node, err := tluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tluo *TaskLikeUpdateOne) Exec(ctx context.Context) error {
	_, err := tluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tluo *TaskLikeUpdateOne) ExecX(ctx context.Context) {
	if err := tluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tluo *TaskLikeUpdateOne) check() error {
	if _, ok := tluo.mutation.TaskID(); tluo.mutation.TaskCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task\"")
	}
	if _, ok := tluo.mutation.TeammateID(); tluo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (tluo *TaskLikeUpdateOne) sqlSave(ctx context.Context) (_node *TaskLike, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklike.Table,
			Columns: tasklike.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklike.FieldID,
			},
		},
	}
	id, ok := tluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TaskLike.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklike.FieldID)
		for _, f := range fields {
			if !tasklike.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tasklike.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tluo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TaskTable,
			Columns: []string{tasklike.TaskColumn},
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
	if nodes := tluo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TaskTable,
			Columns: []string{tasklike.TaskColumn},
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
	if tluo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TeammateTable,
			Columns: []string{tasklike.TeammateColumn},
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
	if nodes := tluo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasklike.TeammateTable,
			Columns: []string{tasklike.TeammateColumn},
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
	_node = &TaskLike{config: tluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklike.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
