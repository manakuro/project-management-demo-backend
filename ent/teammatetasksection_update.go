// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskSectionUpdate is the builder for updating TeammateTaskSection entities.
type TeammateTaskSectionUpdate struct {
	config
	hooks    []Hook
	mutation *TeammateTaskSectionMutation
}

// Where appends a list predicates to the TeammateTaskSectionUpdate builder.
func (ttsu *TeammateTaskSectionUpdate) Where(ps ...predicate.TeammateTaskSection) *TeammateTaskSectionUpdate {
	ttsu.mutation.Where(ps...)
	return ttsu
}

// SetTeammateID sets the "teammate_id" field.
func (ttsu *TeammateTaskSectionUpdate) SetTeammateID(u ulid.ID) *TeammateTaskSectionUpdate {
	ttsu.mutation.SetTeammateID(u)
	return ttsu
}

// SetWorkspaceID sets the "workspace_id" field.
func (ttsu *TeammateTaskSectionUpdate) SetWorkspaceID(u ulid.ID) *TeammateTaskSectionUpdate {
	ttsu.mutation.SetWorkspaceID(u)
	return ttsu
}

// SetName sets the "name" field.
func (ttsu *TeammateTaskSectionUpdate) SetName(s string) *TeammateTaskSectionUpdate {
	ttsu.mutation.SetName(s)
	return ttsu
}

// SetAssigned sets the "assigned" field.
func (ttsu *TeammateTaskSectionUpdate) SetAssigned(b bool) *TeammateTaskSectionUpdate {
	ttsu.mutation.SetAssigned(b)
	return ttsu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttsu *TeammateTaskSectionUpdate) SetTeammate(t *Teammate) *TeammateTaskSectionUpdate {
	return ttsu.SetTeammateID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (ttsu *TeammateTaskSectionUpdate) SetWorkspace(w *Workspace) *TeammateTaskSectionUpdate {
	return ttsu.SetWorkspaceID(w.ID)
}

// AddTeammateTaskIDs adds the "teammate_tasks" edge to the TeammateTask entity by IDs.
func (ttsu *TeammateTaskSectionUpdate) AddTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionUpdate {
	ttsu.mutation.AddTeammateTaskIDs(ids...)
	return ttsu
}

// AddTeammateTasks adds the "teammate_tasks" edges to the TeammateTask entity.
func (ttsu *TeammateTaskSectionUpdate) AddTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsu.AddTeammateTaskIDs(ids...)
}

// Mutation returns the TeammateTaskSectionMutation object of the builder.
func (ttsu *TeammateTaskSectionUpdate) Mutation() *TeammateTaskSectionMutation {
	return ttsu.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (ttsu *TeammateTaskSectionUpdate) ClearTeammate() *TeammateTaskSectionUpdate {
	ttsu.mutation.ClearTeammate()
	return ttsu
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (ttsu *TeammateTaskSectionUpdate) ClearWorkspace() *TeammateTaskSectionUpdate {
	ttsu.mutation.ClearWorkspace()
	return ttsu
}

// ClearTeammateTasks clears all "teammate_tasks" edges to the TeammateTask entity.
func (ttsu *TeammateTaskSectionUpdate) ClearTeammateTasks() *TeammateTaskSectionUpdate {
	ttsu.mutation.ClearTeammateTasks()
	return ttsu
}

// RemoveTeammateTaskIDs removes the "teammate_tasks" edge to TeammateTask entities by IDs.
func (ttsu *TeammateTaskSectionUpdate) RemoveTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionUpdate {
	ttsu.mutation.RemoveTeammateTaskIDs(ids...)
	return ttsu
}

// RemoveTeammateTasks removes "teammate_tasks" edges to TeammateTask entities.
func (ttsu *TeammateTaskSectionUpdate) RemoveTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsu.RemoveTeammateTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttsu *TeammateTaskSectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttsu.hooks) == 0 {
		if err = ttsu.check(); err != nil {
			return 0, err
		}
		affected, err = ttsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttsu.check(); err != nil {
				return 0, err
			}
			ttsu.mutation = mutation
			affected, err = ttsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttsu.hooks) - 1; i >= 0; i-- {
			if ttsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttsu *TeammateTaskSectionUpdate) SaveX(ctx context.Context) int {
	affected, err := ttsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttsu *TeammateTaskSectionUpdate) Exec(ctx context.Context) error {
	_, err := ttsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttsu *TeammateTaskSectionUpdate) ExecX(ctx context.Context) {
	if err := ttsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttsu *TeammateTaskSectionUpdate) check() error {
	if v, ok := ttsu.mutation.Name(); ok {
		if err := teammatetasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ttsu.mutation.TeammateID(); ttsu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	if _, ok := ttsu.mutation.WorkspaceID(); ttsu.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	return nil
}

func (ttsu *TeammateTaskSectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teammatetasksection.Table,
			Columns: teammatetasksection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasksection.FieldID,
			},
		},
	}
	if ps := ttsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammatetasksection.FieldName,
		})
	}
	if value, ok := ttsu.mutation.Assigned(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetasksection.FieldAssigned,
		})
	}
	if ttsu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if nodes := ttsu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if ttsu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if nodes := ttsu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if ttsu.mutation.TeammateTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttsu.mutation.RemovedTeammateTasksIDs(); len(nodes) > 0 && !ttsu.mutation.TeammateTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttsu.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teammatetasksection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TeammateTaskSectionUpdateOne is the builder for updating a single TeammateTaskSection entity.
type TeammateTaskSectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeammateTaskSectionMutation
}

// SetTeammateID sets the "teammate_id" field.
func (ttsuo *TeammateTaskSectionUpdateOne) SetTeammateID(u ulid.ID) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.SetTeammateID(u)
	return ttsuo
}

// SetWorkspaceID sets the "workspace_id" field.
func (ttsuo *TeammateTaskSectionUpdateOne) SetWorkspaceID(u ulid.ID) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.SetWorkspaceID(u)
	return ttsuo
}

// SetName sets the "name" field.
func (ttsuo *TeammateTaskSectionUpdateOne) SetName(s string) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.SetName(s)
	return ttsuo
}

// SetAssigned sets the "assigned" field.
func (ttsuo *TeammateTaskSectionUpdateOne) SetAssigned(b bool) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.SetAssigned(b)
	return ttsuo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttsuo *TeammateTaskSectionUpdateOne) SetTeammate(t *Teammate) *TeammateTaskSectionUpdateOne {
	return ttsuo.SetTeammateID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (ttsuo *TeammateTaskSectionUpdateOne) SetWorkspace(w *Workspace) *TeammateTaskSectionUpdateOne {
	return ttsuo.SetWorkspaceID(w.ID)
}

// AddTeammateTaskIDs adds the "teammate_tasks" edge to the TeammateTask entity by IDs.
func (ttsuo *TeammateTaskSectionUpdateOne) AddTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.AddTeammateTaskIDs(ids...)
	return ttsuo
}

// AddTeammateTasks adds the "teammate_tasks" edges to the TeammateTask entity.
func (ttsuo *TeammateTaskSectionUpdateOne) AddTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsuo.AddTeammateTaskIDs(ids...)
}

// Mutation returns the TeammateTaskSectionMutation object of the builder.
func (ttsuo *TeammateTaskSectionUpdateOne) Mutation() *TeammateTaskSectionMutation {
	return ttsuo.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (ttsuo *TeammateTaskSectionUpdateOne) ClearTeammate() *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.ClearTeammate()
	return ttsuo
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (ttsuo *TeammateTaskSectionUpdateOne) ClearWorkspace() *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.ClearWorkspace()
	return ttsuo
}

// ClearTeammateTasks clears all "teammate_tasks" edges to the TeammateTask entity.
func (ttsuo *TeammateTaskSectionUpdateOne) ClearTeammateTasks() *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.ClearTeammateTasks()
	return ttsuo
}

// RemoveTeammateTaskIDs removes the "teammate_tasks" edge to TeammateTask entities by IDs.
func (ttsuo *TeammateTaskSectionUpdateOne) RemoveTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionUpdateOne {
	ttsuo.mutation.RemoveTeammateTaskIDs(ids...)
	return ttsuo
}

// RemoveTeammateTasks removes "teammate_tasks" edges to TeammateTask entities.
func (ttsuo *TeammateTaskSectionUpdateOne) RemoveTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsuo.RemoveTeammateTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttsuo *TeammateTaskSectionUpdateOne) Select(field string, fields ...string) *TeammateTaskSectionUpdateOne {
	ttsuo.fields = append([]string{field}, fields...)
	return ttsuo
}

// Save executes the query and returns the updated TeammateTaskSection entity.
func (ttsuo *TeammateTaskSectionUpdateOne) Save(ctx context.Context) (*TeammateTaskSection, error) {
	var (
		err  error
		node *TeammateTaskSection
	)
	if len(ttsuo.hooks) == 0 {
		if err = ttsuo.check(); err != nil {
			return nil, err
		}
		node, err = ttsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttsuo.check(); err != nil {
				return nil, err
			}
			ttsuo.mutation = mutation
			node, err = ttsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttsuo.hooks) - 1; i >= 0; i-- {
			if ttsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttsuo *TeammateTaskSectionUpdateOne) SaveX(ctx context.Context) *TeammateTaskSection {
	node, err := ttsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttsuo *TeammateTaskSectionUpdateOne) Exec(ctx context.Context) error {
	_, err := ttsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttsuo *TeammateTaskSectionUpdateOne) ExecX(ctx context.Context) {
	if err := ttsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttsuo *TeammateTaskSectionUpdateOne) check() error {
	if v, ok := ttsuo.mutation.Name(); ok {
		if err := teammatetasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ttsuo.mutation.TeammateID(); ttsuo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	if _, ok := ttsuo.mutation.WorkspaceID(); ttsuo.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	return nil
}

func (ttsuo *TeammateTaskSectionUpdateOne) sqlSave(ctx context.Context) (_node *TeammateTaskSection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teammatetasksection.Table,
			Columns: teammatetasksection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasksection.FieldID,
			},
		},
	}
	id, ok := ttsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TeammateTaskSection.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teammatetasksection.FieldID)
		for _, f := range fields {
			if !teammatetasksection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teammatetasksection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammatetasksection.FieldName,
		})
	}
	if value, ok := ttsuo.mutation.Assigned(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetasksection.FieldAssigned,
		})
	}
	if ttsuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if nodes := ttsuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if ttsuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if nodes := ttsuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if ttsuo.mutation.TeammateTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttsuo.mutation.RemovedTeammateTasksIDs(); len(nodes) > 0 && !ttsuo.mutation.TeammateTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttsuo.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TeammateTaskSection{config: ttsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teammatetasksection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}