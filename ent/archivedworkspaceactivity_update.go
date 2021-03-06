// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/archivedworkspaceactivity"
	"project-management-demo-backend/ent/archivedworkspaceactivitytask"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedWorkspaceActivityUpdate is the builder for updating ArchivedWorkspaceActivity entities.
type ArchivedWorkspaceActivityUpdate struct {
	config
	hooks    []Hook
	mutation *ArchivedWorkspaceActivityMutation
}

// Where appends a list predicates to the ArchivedWorkspaceActivityUpdate builder.
func (awau *ArchivedWorkspaceActivityUpdate) Where(ps ...predicate.ArchivedWorkspaceActivity) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.Where(ps...)
	return awau
}

// SetActivityTypeID sets the "activity_type_id" field.
func (awau *ArchivedWorkspaceActivityUpdate) SetActivityTypeID(u ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.SetActivityTypeID(u)
	return awau
}

// SetWorkspaceID sets the "workspace_id" field.
func (awau *ArchivedWorkspaceActivityUpdate) SetWorkspaceID(u ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.SetWorkspaceID(u)
	return awau
}

// SetProjectID sets the "project_id" field.
func (awau *ArchivedWorkspaceActivityUpdate) SetProjectID(u ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.SetProjectID(u)
	return awau
}

// SetTeammateID sets the "teammate_id" field.
func (awau *ArchivedWorkspaceActivityUpdate) SetTeammateID(u ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.SetTeammateID(u)
	return awau
}

// SetActivityType sets the "activityType" edge to the ActivityType entity.
func (awau *ArchivedWorkspaceActivityUpdate) SetActivityType(a *ActivityType) *ArchivedWorkspaceActivityUpdate {
	return awau.SetActivityTypeID(a.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (awau *ArchivedWorkspaceActivityUpdate) SetWorkspace(w *Workspace) *ArchivedWorkspaceActivityUpdate {
	return awau.SetWorkspaceID(w.ID)
}

// SetProject sets the "project" edge to the Project entity.
func (awau *ArchivedWorkspaceActivityUpdate) SetProject(p *Project) *ArchivedWorkspaceActivityUpdate {
	return awau.SetProjectID(p.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (awau *ArchivedWorkspaceActivityUpdate) SetTeammate(t *Teammate) *ArchivedWorkspaceActivityUpdate {
	return awau.SetTeammateID(t.ID)
}

// AddArchivedWorkspaceActivityTaskIDs adds the "archivedWorkspaceActivityTasks" edge to the ArchivedWorkspaceActivityTask entity by IDs.
func (awau *ArchivedWorkspaceActivityUpdate) AddArchivedWorkspaceActivityTaskIDs(ids ...ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.AddArchivedWorkspaceActivityTaskIDs(ids...)
	return awau
}

// AddArchivedWorkspaceActivityTasks adds the "archivedWorkspaceActivityTasks" edges to the ArchivedWorkspaceActivityTask entity.
func (awau *ArchivedWorkspaceActivityUpdate) AddArchivedWorkspaceActivityTasks(a ...*ArchivedWorkspaceActivityTask) *ArchivedWorkspaceActivityUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return awau.AddArchivedWorkspaceActivityTaskIDs(ids...)
}

// Mutation returns the ArchivedWorkspaceActivityMutation object of the builder.
func (awau *ArchivedWorkspaceActivityUpdate) Mutation() *ArchivedWorkspaceActivityMutation {
	return awau.mutation
}

// ClearActivityType clears the "activityType" edge to the ActivityType entity.
func (awau *ArchivedWorkspaceActivityUpdate) ClearActivityType() *ArchivedWorkspaceActivityUpdate {
	awau.mutation.ClearActivityType()
	return awau
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (awau *ArchivedWorkspaceActivityUpdate) ClearWorkspace() *ArchivedWorkspaceActivityUpdate {
	awau.mutation.ClearWorkspace()
	return awau
}

// ClearProject clears the "project" edge to the Project entity.
func (awau *ArchivedWorkspaceActivityUpdate) ClearProject() *ArchivedWorkspaceActivityUpdate {
	awau.mutation.ClearProject()
	return awau
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (awau *ArchivedWorkspaceActivityUpdate) ClearTeammate() *ArchivedWorkspaceActivityUpdate {
	awau.mutation.ClearTeammate()
	return awau
}

// ClearArchivedWorkspaceActivityTasks clears all "archivedWorkspaceActivityTasks" edges to the ArchivedWorkspaceActivityTask entity.
func (awau *ArchivedWorkspaceActivityUpdate) ClearArchivedWorkspaceActivityTasks() *ArchivedWorkspaceActivityUpdate {
	awau.mutation.ClearArchivedWorkspaceActivityTasks()
	return awau
}

// RemoveArchivedWorkspaceActivityTaskIDs removes the "archivedWorkspaceActivityTasks" edge to ArchivedWorkspaceActivityTask entities by IDs.
func (awau *ArchivedWorkspaceActivityUpdate) RemoveArchivedWorkspaceActivityTaskIDs(ids ...ulid.ID) *ArchivedWorkspaceActivityUpdate {
	awau.mutation.RemoveArchivedWorkspaceActivityTaskIDs(ids...)
	return awau
}

// RemoveArchivedWorkspaceActivityTasks removes "archivedWorkspaceActivityTasks" edges to ArchivedWorkspaceActivityTask entities.
func (awau *ArchivedWorkspaceActivityUpdate) RemoveArchivedWorkspaceActivityTasks(a ...*ArchivedWorkspaceActivityTask) *ArchivedWorkspaceActivityUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return awau.RemoveArchivedWorkspaceActivityTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (awau *ArchivedWorkspaceActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(awau.hooks) == 0 {
		if err = awau.check(); err != nil {
			return 0, err
		}
		affected, err = awau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedWorkspaceActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awau.check(); err != nil {
				return 0, err
			}
			awau.mutation = mutation
			affected, err = awau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(awau.hooks) - 1; i >= 0; i-- {
			if awau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = awau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (awau *ArchivedWorkspaceActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := awau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (awau *ArchivedWorkspaceActivityUpdate) Exec(ctx context.Context) error {
	_, err := awau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awau *ArchivedWorkspaceActivityUpdate) ExecX(ctx context.Context) {
	if err := awau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awau *ArchivedWorkspaceActivityUpdate) check() error {
	if _, ok := awau.mutation.ActivityTypeID(); awau.mutation.ActivityTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.activityType"`)
	}
	if _, ok := awau.mutation.WorkspaceID(); awau.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.workspace"`)
	}
	if _, ok := awau.mutation.ProjectID(); awau.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.project"`)
	}
	if _, ok := awau.mutation.TeammateID(); awau.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.teammate"`)
	}
	return nil
}

func (awau *ArchivedWorkspaceActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedworkspaceactivity.Table,
			Columns: archivedworkspaceactivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedworkspaceactivity.FieldID,
			},
		},
	}
	if ps := awau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if awau.mutation.ActivityTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ActivityTypeTable,
			Columns: []string{archivedworkspaceactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awau.mutation.ActivityTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ActivityTypeTable,
			Columns: []string{archivedworkspaceactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if awau.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.WorkspaceTable,
			Columns: []string{archivedworkspaceactivity.WorkspaceColumn},
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
	if nodes := awau.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.WorkspaceTable,
			Columns: []string{archivedworkspaceactivity.WorkspaceColumn},
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
	if awau.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ProjectTable,
			Columns: []string{archivedworkspaceactivity.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awau.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ProjectTable,
			Columns: []string{archivedworkspaceactivity.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if awau.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.TeammateTable,
			Columns: []string{archivedworkspaceactivity.TeammateColumn},
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
	if nodes := awau.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.TeammateTable,
			Columns: []string{archivedworkspaceactivity.TeammateColumn},
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
	if awau.mutation.ArchivedWorkspaceActivityTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awau.mutation.RemovedArchivedWorkspaceActivityTasksIDs(); len(nodes) > 0 && !awau.mutation.ArchivedWorkspaceActivityTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awau.mutation.ArchivedWorkspaceActivityTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, awau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{archivedworkspaceactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ArchivedWorkspaceActivityUpdateOne is the builder for updating a single ArchivedWorkspaceActivity entity.
type ArchivedWorkspaceActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArchivedWorkspaceActivityMutation
}

// SetActivityTypeID sets the "activity_type_id" field.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetActivityTypeID(u ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.SetActivityTypeID(u)
	return awauo
}

// SetWorkspaceID sets the "workspace_id" field.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetWorkspaceID(u ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.SetWorkspaceID(u)
	return awauo
}

// SetProjectID sets the "project_id" field.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetProjectID(u ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.SetProjectID(u)
	return awauo
}

// SetTeammateID sets the "teammate_id" field.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetTeammateID(u ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.SetTeammateID(u)
	return awauo
}

// SetActivityType sets the "activityType" edge to the ActivityType entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetActivityType(a *ActivityType) *ArchivedWorkspaceActivityUpdateOne {
	return awauo.SetActivityTypeID(a.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetWorkspace(w *Workspace) *ArchivedWorkspaceActivityUpdateOne {
	return awauo.SetWorkspaceID(w.ID)
}

// SetProject sets the "project" edge to the Project entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetProject(p *Project) *ArchivedWorkspaceActivityUpdateOne {
	return awauo.SetProjectID(p.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SetTeammate(t *Teammate) *ArchivedWorkspaceActivityUpdateOne {
	return awauo.SetTeammateID(t.ID)
}

// AddArchivedWorkspaceActivityTaskIDs adds the "archivedWorkspaceActivityTasks" edge to the ArchivedWorkspaceActivityTask entity by IDs.
func (awauo *ArchivedWorkspaceActivityUpdateOne) AddArchivedWorkspaceActivityTaskIDs(ids ...ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.AddArchivedWorkspaceActivityTaskIDs(ids...)
	return awauo
}

// AddArchivedWorkspaceActivityTasks adds the "archivedWorkspaceActivityTasks" edges to the ArchivedWorkspaceActivityTask entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) AddArchivedWorkspaceActivityTasks(a ...*ArchivedWorkspaceActivityTask) *ArchivedWorkspaceActivityUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return awauo.AddArchivedWorkspaceActivityTaskIDs(ids...)
}

// Mutation returns the ArchivedWorkspaceActivityMutation object of the builder.
func (awauo *ArchivedWorkspaceActivityUpdateOne) Mutation() *ArchivedWorkspaceActivityMutation {
	return awauo.mutation
}

// ClearActivityType clears the "activityType" edge to the ActivityType entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ClearActivityType() *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.ClearActivityType()
	return awauo
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ClearWorkspace() *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.ClearWorkspace()
	return awauo
}

// ClearProject clears the "project" edge to the Project entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ClearProject() *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.ClearProject()
	return awauo
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ClearTeammate() *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.ClearTeammate()
	return awauo
}

// ClearArchivedWorkspaceActivityTasks clears all "archivedWorkspaceActivityTasks" edges to the ArchivedWorkspaceActivityTask entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ClearArchivedWorkspaceActivityTasks() *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.ClearArchivedWorkspaceActivityTasks()
	return awauo
}

// RemoveArchivedWorkspaceActivityTaskIDs removes the "archivedWorkspaceActivityTasks" edge to ArchivedWorkspaceActivityTask entities by IDs.
func (awauo *ArchivedWorkspaceActivityUpdateOne) RemoveArchivedWorkspaceActivityTaskIDs(ids ...ulid.ID) *ArchivedWorkspaceActivityUpdateOne {
	awauo.mutation.RemoveArchivedWorkspaceActivityTaskIDs(ids...)
	return awauo
}

// RemoveArchivedWorkspaceActivityTasks removes "archivedWorkspaceActivityTasks" edges to ArchivedWorkspaceActivityTask entities.
func (awauo *ArchivedWorkspaceActivityUpdateOne) RemoveArchivedWorkspaceActivityTasks(a ...*ArchivedWorkspaceActivityTask) *ArchivedWorkspaceActivityUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return awauo.RemoveArchivedWorkspaceActivityTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (awauo *ArchivedWorkspaceActivityUpdateOne) Select(field string, fields ...string) *ArchivedWorkspaceActivityUpdateOne {
	awauo.fields = append([]string{field}, fields...)
	return awauo
}

// Save executes the query and returns the updated ArchivedWorkspaceActivity entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) Save(ctx context.Context) (*ArchivedWorkspaceActivity, error) {
	var (
		err  error
		node *ArchivedWorkspaceActivity
	)
	if len(awauo.hooks) == 0 {
		if err = awauo.check(); err != nil {
			return nil, err
		}
		node, err = awauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedWorkspaceActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awauo.check(); err != nil {
				return nil, err
			}
			awauo.mutation = mutation
			node, err = awauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(awauo.hooks) - 1; i >= 0; i-- {
			if awauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = awauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (awauo *ArchivedWorkspaceActivityUpdateOne) SaveX(ctx context.Context) *ArchivedWorkspaceActivity {
	node, err := awauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (awauo *ArchivedWorkspaceActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := awauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awauo *ArchivedWorkspaceActivityUpdateOne) ExecX(ctx context.Context) {
	if err := awauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awauo *ArchivedWorkspaceActivityUpdateOne) check() error {
	if _, ok := awauo.mutation.ActivityTypeID(); awauo.mutation.ActivityTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.activityType"`)
	}
	if _, ok := awauo.mutation.WorkspaceID(); awauo.mutation.WorkspaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.workspace"`)
	}
	if _, ok := awauo.mutation.ProjectID(); awauo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.project"`)
	}
	if _, ok := awauo.mutation.TeammateID(); awauo.mutation.TeammateCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArchivedWorkspaceActivity.teammate"`)
	}
	return nil
}

func (awauo *ArchivedWorkspaceActivityUpdateOne) sqlSave(ctx context.Context) (_node *ArchivedWorkspaceActivity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   archivedworkspaceactivity.Table,
			Columns: archivedworkspaceactivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedworkspaceactivity.FieldID,
			},
		},
	}
	id, ok := awauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArchivedWorkspaceActivity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := awauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, archivedworkspaceactivity.FieldID)
		for _, f := range fields {
			if !archivedworkspaceactivity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != archivedworkspaceactivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := awauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if awauo.mutation.ActivityTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ActivityTypeTable,
			Columns: []string{archivedworkspaceactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awauo.mutation.ActivityTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ActivityTypeTable,
			Columns: []string{archivedworkspaceactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if awauo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.WorkspaceTable,
			Columns: []string{archivedworkspaceactivity.WorkspaceColumn},
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
	if nodes := awauo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.WorkspaceTable,
			Columns: []string{archivedworkspaceactivity.WorkspaceColumn},
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
	if awauo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ProjectTable,
			Columns: []string{archivedworkspaceactivity.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awauo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.ProjectTable,
			Columns: []string{archivedworkspaceactivity.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if awauo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.TeammateTable,
			Columns: []string{archivedworkspaceactivity.TeammateColumn},
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
	if nodes := awauo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivity.TeammateTable,
			Columns: []string{archivedworkspaceactivity.TeammateColumn},
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
	if awauo.mutation.ArchivedWorkspaceActivityTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awauo.mutation.RemovedArchivedWorkspaceActivityTasksIDs(); len(nodes) > 0 && !awauo.mutation.ArchivedWorkspaceActivityTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := awauo.mutation.ArchivedWorkspaceActivityTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   archivedworkspaceactivity.ArchivedWorkspaceActivityTasksTable,
			Columns: []string{archivedworkspaceactivity.ArchivedWorkspaceActivityTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivitytask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ArchivedWorkspaceActivity{config: awauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, awauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{archivedworkspaceactivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
