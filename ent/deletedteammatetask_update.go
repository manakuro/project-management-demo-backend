// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/deletedteammatetask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeletedTeammateTaskUpdate is the builder for updating DeletedTeammateTask entities.
type DeletedTeammateTaskUpdate struct {
	config
	hooks    []Hook
	mutation *DeletedTeammateTaskMutation
}

// Where appends a list predicates to the DeletedTeammateTaskUpdate builder.
func (dttu *DeletedTeammateTaskUpdate) Where(ps ...predicate.DeletedTeammateTask) *DeletedTeammateTaskUpdate {
	dttu.mutation.Where(ps...)
	return dttu
}

// SetTeammateID sets the "teammate_id" field.
func (dttu *DeletedTeammateTaskUpdate) SetTeammateID(u ulid.ID) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetTeammateID(u)
	return dttu
}

// SetTaskID sets the "task_id" field.
func (dttu *DeletedTeammateTaskUpdate) SetTaskID(u ulid.ID) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetTaskID(u)
	return dttu
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (dttu *DeletedTeammateTaskUpdate) SetTeammateTaskSectionID(u ulid.ID) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetTeammateTaskSectionID(u)
	return dttu
}

// SetWorkspaceID sets the "workspace_id" field.
func (dttu *DeletedTeammateTaskUpdate) SetWorkspaceID(u ulid.ID) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetWorkspaceID(u)
	return dttu
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (dttu *DeletedTeammateTaskUpdate) SetTeammateTaskCreatedAt(t time.Time) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetTeammateTaskCreatedAt(t)
	return dttu
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (dttu *DeletedTeammateTaskUpdate) SetTeammateTaskUpdatedAt(t time.Time) *DeletedTeammateTaskUpdate {
	dttu.mutation.SetTeammateTaskUpdatedAt(t)
	return dttu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (dttu *DeletedTeammateTaskUpdate) SetTeammate(t *Teammate) *DeletedTeammateTaskUpdate {
	return dttu.SetTeammateID(t.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (dttu *DeletedTeammateTaskUpdate) SetTask(t *Task) *DeletedTeammateTaskUpdate {
	return dttu.SetTaskID(t.ID)
}

// SetTeammateTaskSection sets the "teammateTaskSection" edge to the TeammateTaskSection entity.
func (dttu *DeletedTeammateTaskUpdate) SetTeammateTaskSection(t *TeammateTaskSection) *DeletedTeammateTaskUpdate {
	return dttu.SetTeammateTaskSectionID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (dttu *DeletedTeammateTaskUpdate) SetWorkspace(w *Workspace) *DeletedTeammateTaskUpdate {
	return dttu.SetWorkspaceID(w.ID)
}

// Mutation returns the DeletedTeammateTaskMutation object of the builder.
func (dttu *DeletedTeammateTaskUpdate) Mutation() *DeletedTeammateTaskMutation {
	return dttu.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (dttu *DeletedTeammateTaskUpdate) ClearTeammate() *DeletedTeammateTaskUpdate {
	dttu.mutation.ClearTeammate()
	return dttu
}

// ClearTask clears the "task" edge to the Task entity.
func (dttu *DeletedTeammateTaskUpdate) ClearTask() *DeletedTeammateTaskUpdate {
	dttu.mutation.ClearTask()
	return dttu
}

// ClearTeammateTaskSection clears the "teammateTaskSection" edge to the TeammateTaskSection entity.
func (dttu *DeletedTeammateTaskUpdate) ClearTeammateTaskSection() *DeletedTeammateTaskUpdate {
	dttu.mutation.ClearTeammateTaskSection()
	return dttu
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (dttu *DeletedTeammateTaskUpdate) ClearWorkspace() *DeletedTeammateTaskUpdate {
	dttu.mutation.ClearWorkspace()
	return dttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dttu *DeletedTeammateTaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dttu.hooks) == 0 {
		if err = dttu.check(); err != nil {
			return 0, err
		}
		affected, err = dttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTeammateTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dttu.check(); err != nil {
				return 0, err
			}
			dttu.mutation = mutation
			affected, err = dttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dttu.hooks) - 1; i >= 0; i-- {
			if dttu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dttu *DeletedTeammateTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := dttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dttu *DeletedTeammateTaskUpdate) Exec(ctx context.Context) error {
	_, err := dttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dttu *DeletedTeammateTaskUpdate) ExecX(ctx context.Context) {
	if err := dttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dttu *DeletedTeammateTaskUpdate) check() error {
	if _, ok := dttu.mutation.TeammateID(); dttu.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.teammate"`)
	}
	if _, ok := dttu.mutation.TaskID(); dttu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.task"`)
	}
	if _, ok := dttu.mutation.TeammateTaskSectionID(); dttu.mutation.TeammateTaskSectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.teammateTaskSection"`)
	}
	if _, ok := dttu.mutation.WorkspaceID(); dttu.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.workspace"`)
	}
	return nil
}

func (dttu *DeletedTeammateTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedteammatetask.Table,
			Columns: deletedteammatetask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedteammatetask.FieldID,
			},
		},
	}
	if ps := dttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dttu.mutation.TeammateTaskCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskCreatedAt,
		})
	}
	if value, ok := dttu.mutation.TeammateTaskUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskUpdatedAt,
		})
	}
	if dttu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTable,
			Columns: []string{deletedteammatetask.TeammateColumn},
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
	if nodes := dttu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTable,
			Columns: []string{deletedteammatetask.TeammateColumn},
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
	if dttu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TaskTable,
			Columns: []string{deletedteammatetask.TaskColumn},
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
	if nodes := dttu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TaskTable,
			Columns: []string{deletedteammatetask.TaskColumn},
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
	if dttu.mutation.TeammateTaskSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTaskSectionTable,
			Columns: []string{deletedteammatetask.TeammateTaskSectionColumn},
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
	if nodes := dttu.mutation.TeammateTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTaskSectionTable,
			Columns: []string{deletedteammatetask.TeammateTaskSectionColumn},
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
	if dttu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.WorkspaceTable,
			Columns: []string{deletedteammatetask.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dttu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.WorkspaceTable,
			Columns: []string{deletedteammatetask.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedteammatetask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DeletedTeammateTaskUpdateOne is the builder for updating a single DeletedTeammateTask entity.
type DeletedTeammateTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeletedTeammateTaskMutation
}

// SetTeammateID sets the "teammate_id" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammateID(u ulid.ID) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetTeammateID(u)
	return dttuo
}

// SetTaskID sets the "task_id" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTaskID(u ulid.ID) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetTaskID(u)
	return dttuo
}

// SetTeammateTaskSectionID sets the "teammate_task_section_id" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammateTaskSectionID(u ulid.ID) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetTeammateTaskSectionID(u)
	return dttuo
}

// SetWorkspaceID sets the "workspace_id" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetWorkspaceID(u ulid.ID) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetWorkspaceID(u)
	return dttuo
}

// SetTeammateTaskCreatedAt sets the "teammate_task_created_at" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammateTaskCreatedAt(t time.Time) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetTeammateTaskCreatedAt(t)
	return dttuo
}

// SetTeammateTaskUpdatedAt sets the "teammate_task_updated_at" field.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammateTaskUpdatedAt(t time.Time) *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.SetTeammateTaskUpdatedAt(t)
	return dttuo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammate(t *Teammate) *DeletedTeammateTaskUpdateOne {
	return dttuo.SetTeammateID(t.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTask(t *Task) *DeletedTeammateTaskUpdateOne {
	return dttuo.SetTaskID(t.ID)
}

// SetTeammateTaskSection sets the "teammateTaskSection" edge to the TeammateTaskSection entity.
func (dttuo *DeletedTeammateTaskUpdateOne) SetTeammateTaskSection(t *TeammateTaskSection) *DeletedTeammateTaskUpdateOne {
	return dttuo.SetTeammateTaskSectionID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (dttuo *DeletedTeammateTaskUpdateOne) SetWorkspace(w *Workspace) *DeletedTeammateTaskUpdateOne {
	return dttuo.SetWorkspaceID(w.ID)
}

// Mutation returns the DeletedTeammateTaskMutation object of the builder.
func (dttuo *DeletedTeammateTaskUpdateOne) Mutation() *DeletedTeammateTaskMutation {
	return dttuo.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (dttuo *DeletedTeammateTaskUpdateOne) ClearTeammate() *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.ClearTeammate()
	return dttuo
}

// ClearTask clears the "task" edge to the Task entity.
func (dttuo *DeletedTeammateTaskUpdateOne) ClearTask() *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.ClearTask()
	return dttuo
}

// ClearTeammateTaskSection clears the "teammateTaskSection" edge to the TeammateTaskSection entity.
func (dttuo *DeletedTeammateTaskUpdateOne) ClearTeammateTaskSection() *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.ClearTeammateTaskSection()
	return dttuo
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (dttuo *DeletedTeammateTaskUpdateOne) ClearWorkspace() *DeletedTeammateTaskUpdateOne {
	dttuo.mutation.ClearWorkspace()
	return dttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dttuo *DeletedTeammateTaskUpdateOne) Select(field string, fields ...string) *DeletedTeammateTaskUpdateOne {
	dttuo.fields = append([]string{field}, fields...)
	return dttuo
}

// Save executes the query and returns the updated DeletedTeammateTask entity.
func (dttuo *DeletedTeammateTaskUpdateOne) Save(ctx context.Context) (*DeletedTeammateTask, error) {
	var (
		err  error
		node *DeletedTeammateTask
	)
	if len(dttuo.hooks) == 0 {
		if err = dttuo.check(); err != nil {
			return nil, err
		}
		node, err = dttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTeammateTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dttuo.check(); err != nil {
				return nil, err
			}
			dttuo.mutation = mutation
			node, err = dttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dttuo.hooks) - 1; i >= 0; i-- {
			if dttuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dttuo *DeletedTeammateTaskUpdateOne) SaveX(ctx context.Context) *DeletedTeammateTask {
	node, err := dttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dttuo *DeletedTeammateTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := dttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dttuo *DeletedTeammateTaskUpdateOne) ExecX(ctx context.Context) {
	if err := dttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dttuo *DeletedTeammateTaskUpdateOne) check() error {
	if _, ok := dttuo.mutation.TeammateID(); dttuo.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.teammate"`)
	}
	if _, ok := dttuo.mutation.TaskID(); dttuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.task"`)
	}
	if _, ok := dttuo.mutation.TeammateTaskSectionID(); dttuo.mutation.TeammateTaskSectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.teammateTaskSection"`)
	}
	if _, ok := dttuo.mutation.WorkspaceID(); dttuo.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DeletedTeammateTask.workspace"`)
	}
	return nil
}

func (dttuo *DeletedTeammateTaskUpdateOne) sqlSave(ctx context.Context) (_node *DeletedTeammateTask, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deletedteammatetask.Table,
			Columns: deletedteammatetask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedteammatetask.FieldID,
			},
		},
	}
	id, ok := dttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeletedTeammateTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deletedteammatetask.FieldID)
		for _, f := range fields {
			if !deletedteammatetask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deletedteammatetask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dttuo.mutation.TeammateTaskCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskCreatedAt,
		})
	}
	if value, ok := dttuo.mutation.TeammateTaskUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedteammatetask.FieldTeammateTaskUpdatedAt,
		})
	}
	if dttuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTable,
			Columns: []string{deletedteammatetask.TeammateColumn},
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
	if nodes := dttuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTable,
			Columns: []string{deletedteammatetask.TeammateColumn},
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
	if dttuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TaskTable,
			Columns: []string{deletedteammatetask.TaskColumn},
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
	if nodes := dttuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TaskTable,
			Columns: []string{deletedteammatetask.TaskColumn},
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
	if dttuo.mutation.TeammateTaskSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTaskSectionTable,
			Columns: []string{deletedteammatetask.TeammateTaskSectionColumn},
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
	if nodes := dttuo.mutation.TeammateTaskSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.TeammateTaskSectionTable,
			Columns: []string{deletedteammatetask.TeammateTaskSectionColumn},
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
	if dttuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.WorkspaceTable,
			Columns: []string{deletedteammatetask.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dttuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedteammatetask.WorkspaceTable,
			Columns: []string{deletedteammatetask.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeletedTeammateTask{config: dttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deletedteammatetask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
