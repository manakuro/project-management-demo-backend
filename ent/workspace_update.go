// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/mytaskstabstatus"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceUpdate is the builder for updating Workspace entities.
type WorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where appends a list predicates to the WorkspaceUpdate builder.
func (wu *WorkspaceUpdate) Where(ps ...predicate.Workspace) *WorkspaceUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetCreatedBy sets the "created_by" field.
func (wu *WorkspaceUpdate) SetCreatedBy(u ulid.ID) *WorkspaceUpdate {
	wu.mutation.SetCreatedBy(u)
	return wu
}

// SetName sets the "name" field.
func (wu *WorkspaceUpdate) SetName(s string) *WorkspaceUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetDescription sets the "description" field.
func (wu *WorkspaceUpdate) SetDescription(e editor.Description) *WorkspaceUpdate {
	wu.mutation.SetDescription(e)
	return wu
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (wu *WorkspaceUpdate) SetTeammateID(id ulid.ID) *WorkspaceUpdate {
	wu.mutation.SetTeammateID(id)
	return wu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wu *WorkspaceUpdate) SetTeammate(t *Teammate) *WorkspaceUpdate {
	return wu.SetTeammateID(t.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (wu *WorkspaceUpdate) AddProjectIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.AddProjectIDs(ids...)
	return wu
}

// AddProjects adds the "projects" edges to the Project entity.
func (wu *WorkspaceUpdate) AddProjects(p ...*Project) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.AddProjectIDs(ids...)
}

// AddWorkspaceTeammateIDs adds the "workspace_teammates" edge to the WorkspaceTeammate entity by IDs.
func (wu *WorkspaceUpdate) AddWorkspaceTeammateIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.AddWorkspaceTeammateIDs(ids...)
	return wu
}

// AddWorkspaceTeammates adds the "workspace_teammates" edges to the WorkspaceTeammate entity.
func (wu *WorkspaceUpdate) AddWorkspaceTeammates(w ...*WorkspaceTeammate) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWorkspaceTeammateIDs(ids...)
}

// AddFavoriteWorkspaceIDs adds the "favorite_workspaces" edge to the FavoriteWorkspace entity by IDs.
func (wu *WorkspaceUpdate) AddFavoriteWorkspaceIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.AddFavoriteWorkspaceIDs(ids...)
	return wu
}

// AddFavoriteWorkspaces adds the "favorite_workspaces" edges to the FavoriteWorkspace entity.
func (wu *WorkspaceUpdate) AddFavoriteWorkspaces(f ...*FavoriteWorkspace) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return wu.AddFavoriteWorkspaceIDs(ids...)
}

// AddMyTasksTabStatusIDs adds the "my_tasks_tab_statuses" edge to the MyTasksTabStatus entity by IDs.
func (wu *WorkspaceUpdate) AddMyTasksTabStatusIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.AddMyTasksTabStatusIDs(ids...)
	return wu
}

// AddMyTasksTabStatuses adds the "my_tasks_tab_statuses" edges to the MyTasksTabStatus entity.
func (wu *WorkspaceUpdate) AddMyTasksTabStatuses(m ...*MyTasksTabStatus) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wu.AddMyTasksTabStatusIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wu *WorkspaceUpdate) Mutation() *WorkspaceMutation {
	return wu.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (wu *WorkspaceUpdate) ClearTeammate() *WorkspaceUpdate {
	wu.mutation.ClearTeammate()
	return wu
}

// ClearProjects clears all "projects" edges to the Project entity.
func (wu *WorkspaceUpdate) ClearProjects() *WorkspaceUpdate {
	wu.mutation.ClearProjects()
	return wu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (wu *WorkspaceUpdate) RemoveProjectIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.RemoveProjectIDs(ids...)
	return wu
}

// RemoveProjects removes "projects" edges to Project entities.
func (wu *WorkspaceUpdate) RemoveProjects(p ...*Project) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.RemoveProjectIDs(ids...)
}

// ClearWorkspaceTeammates clears all "workspace_teammates" edges to the WorkspaceTeammate entity.
func (wu *WorkspaceUpdate) ClearWorkspaceTeammates() *WorkspaceUpdate {
	wu.mutation.ClearWorkspaceTeammates()
	return wu
}

// RemoveWorkspaceTeammateIDs removes the "workspace_teammates" edge to WorkspaceTeammate entities by IDs.
func (wu *WorkspaceUpdate) RemoveWorkspaceTeammateIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.RemoveWorkspaceTeammateIDs(ids...)
	return wu
}

// RemoveWorkspaceTeammates removes "workspace_teammates" edges to WorkspaceTeammate entities.
func (wu *WorkspaceUpdate) RemoveWorkspaceTeammates(w ...*WorkspaceTeammate) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWorkspaceTeammateIDs(ids...)
}

// ClearFavoriteWorkspaces clears all "favorite_workspaces" edges to the FavoriteWorkspace entity.
func (wu *WorkspaceUpdate) ClearFavoriteWorkspaces() *WorkspaceUpdate {
	wu.mutation.ClearFavoriteWorkspaces()
	return wu
}

// RemoveFavoriteWorkspaceIDs removes the "favorite_workspaces" edge to FavoriteWorkspace entities by IDs.
func (wu *WorkspaceUpdate) RemoveFavoriteWorkspaceIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.RemoveFavoriteWorkspaceIDs(ids...)
	return wu
}

// RemoveFavoriteWorkspaces removes "favorite_workspaces" edges to FavoriteWorkspace entities.
func (wu *WorkspaceUpdate) RemoveFavoriteWorkspaces(f ...*FavoriteWorkspace) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return wu.RemoveFavoriteWorkspaceIDs(ids...)
}

// ClearMyTasksTabStatuses clears all "my_tasks_tab_statuses" edges to the MyTasksTabStatus entity.
func (wu *WorkspaceUpdate) ClearMyTasksTabStatuses() *WorkspaceUpdate {
	wu.mutation.ClearMyTasksTabStatuses()
	return wu
}

// RemoveMyTasksTabStatusIDs removes the "my_tasks_tab_statuses" edge to MyTasksTabStatus entities by IDs.
func (wu *WorkspaceUpdate) RemoveMyTasksTabStatusIDs(ids ...ulid.ID) *WorkspaceUpdate {
	wu.mutation.RemoveMyTasksTabStatusIDs(ids...)
	return wu
}

// RemoveMyTasksTabStatuses removes "my_tasks_tab_statuses" edges to MyTasksTabStatus entities.
func (wu *WorkspaceUpdate) RemoveMyTasksTabStatuses(m ...*MyTasksTabStatus) *WorkspaceUpdate {
	ids := make([]ulid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wu.RemoveMyTasksTabStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkspaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkspaceUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkspaceUpdate) check() error {
	if v, ok := wu.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wu.mutation.TeammateID(); wu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (wu *WorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspace.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldDescription,
		})
	}
	if wu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspace.TeammateTable,
			Columns: []string{workspace.TeammateColumn},
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
	if nodes := wu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspace.TeammateTable,
			Columns: []string{workspace.TeammateColumn},
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
	if wu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
	if nodes := wu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !wu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
	if wu.mutation.WorkspaceTeammatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWorkspaceTeammatesIDs(); len(nodes) > 0 && !wu.mutation.WorkspaceTeammatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WorkspaceTeammatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.FavoriteWorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedFavoriteWorkspacesIDs(); len(nodes) > 0 && !wu.mutation.FavoriteWorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.FavoriteWorkspacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.MyTasksTabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedMyTasksTabStatusesIDs(); len(nodes) > 0 && !wu.mutation.MyTasksTabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.MyTasksTabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceUpdateOne is the builder for updating a single Workspace entity.
type WorkspaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkspaceMutation
}

// SetCreatedBy sets the "created_by" field.
func (wuo *WorkspaceUpdateOne) SetCreatedBy(u ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.SetCreatedBy(u)
	return wuo
}

// SetName sets the "name" field.
func (wuo *WorkspaceUpdateOne) SetName(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetDescription sets the "description" field.
func (wuo *WorkspaceUpdateOne) SetDescription(e editor.Description) *WorkspaceUpdateOne {
	wuo.mutation.SetDescription(e)
	return wuo
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (wuo *WorkspaceUpdateOne) SetTeammateID(id ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.SetTeammateID(id)
	return wuo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wuo *WorkspaceUpdateOne) SetTeammate(t *Teammate) *WorkspaceUpdateOne {
	return wuo.SetTeammateID(t.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (wuo *WorkspaceUpdateOne) AddProjectIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.AddProjectIDs(ids...)
	return wuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (wuo *WorkspaceUpdateOne) AddProjects(p ...*Project) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.AddProjectIDs(ids...)
}

// AddWorkspaceTeammateIDs adds the "workspace_teammates" edge to the WorkspaceTeammate entity by IDs.
func (wuo *WorkspaceUpdateOne) AddWorkspaceTeammateIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.AddWorkspaceTeammateIDs(ids...)
	return wuo
}

// AddWorkspaceTeammates adds the "workspace_teammates" edges to the WorkspaceTeammate entity.
func (wuo *WorkspaceUpdateOne) AddWorkspaceTeammates(w ...*WorkspaceTeammate) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWorkspaceTeammateIDs(ids...)
}

// AddFavoriteWorkspaceIDs adds the "favorite_workspaces" edge to the FavoriteWorkspace entity by IDs.
func (wuo *WorkspaceUpdateOne) AddFavoriteWorkspaceIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.AddFavoriteWorkspaceIDs(ids...)
	return wuo
}

// AddFavoriteWorkspaces adds the "favorite_workspaces" edges to the FavoriteWorkspace entity.
func (wuo *WorkspaceUpdateOne) AddFavoriteWorkspaces(f ...*FavoriteWorkspace) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return wuo.AddFavoriteWorkspaceIDs(ids...)
}

// AddMyTasksTabStatusIDs adds the "my_tasks_tab_statuses" edge to the MyTasksTabStatus entity by IDs.
func (wuo *WorkspaceUpdateOne) AddMyTasksTabStatusIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.AddMyTasksTabStatusIDs(ids...)
	return wuo
}

// AddMyTasksTabStatuses adds the "my_tasks_tab_statuses" edges to the MyTasksTabStatus entity.
func (wuo *WorkspaceUpdateOne) AddMyTasksTabStatuses(m ...*MyTasksTabStatus) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wuo.AddMyTasksTabStatusIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wuo *WorkspaceUpdateOne) Mutation() *WorkspaceMutation {
	return wuo.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (wuo *WorkspaceUpdateOne) ClearTeammate() *WorkspaceUpdateOne {
	wuo.mutation.ClearTeammate()
	return wuo
}

// ClearProjects clears all "projects" edges to the Project entity.
func (wuo *WorkspaceUpdateOne) ClearProjects() *WorkspaceUpdateOne {
	wuo.mutation.ClearProjects()
	return wuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (wuo *WorkspaceUpdateOne) RemoveProjectIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveProjectIDs(ids...)
	return wuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (wuo *WorkspaceUpdateOne) RemoveProjects(p ...*Project) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.RemoveProjectIDs(ids...)
}

// ClearWorkspaceTeammates clears all "workspace_teammates" edges to the WorkspaceTeammate entity.
func (wuo *WorkspaceUpdateOne) ClearWorkspaceTeammates() *WorkspaceUpdateOne {
	wuo.mutation.ClearWorkspaceTeammates()
	return wuo
}

// RemoveWorkspaceTeammateIDs removes the "workspace_teammates" edge to WorkspaceTeammate entities by IDs.
func (wuo *WorkspaceUpdateOne) RemoveWorkspaceTeammateIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveWorkspaceTeammateIDs(ids...)
	return wuo
}

// RemoveWorkspaceTeammates removes "workspace_teammates" edges to WorkspaceTeammate entities.
func (wuo *WorkspaceUpdateOne) RemoveWorkspaceTeammates(w ...*WorkspaceTeammate) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWorkspaceTeammateIDs(ids...)
}

// ClearFavoriteWorkspaces clears all "favorite_workspaces" edges to the FavoriteWorkspace entity.
func (wuo *WorkspaceUpdateOne) ClearFavoriteWorkspaces() *WorkspaceUpdateOne {
	wuo.mutation.ClearFavoriteWorkspaces()
	return wuo
}

// RemoveFavoriteWorkspaceIDs removes the "favorite_workspaces" edge to FavoriteWorkspace entities by IDs.
func (wuo *WorkspaceUpdateOne) RemoveFavoriteWorkspaceIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveFavoriteWorkspaceIDs(ids...)
	return wuo
}

// RemoveFavoriteWorkspaces removes "favorite_workspaces" edges to FavoriteWorkspace entities.
func (wuo *WorkspaceUpdateOne) RemoveFavoriteWorkspaces(f ...*FavoriteWorkspace) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return wuo.RemoveFavoriteWorkspaceIDs(ids...)
}

// ClearMyTasksTabStatuses clears all "my_tasks_tab_statuses" edges to the MyTasksTabStatus entity.
func (wuo *WorkspaceUpdateOne) ClearMyTasksTabStatuses() *WorkspaceUpdateOne {
	wuo.mutation.ClearMyTasksTabStatuses()
	return wuo
}

// RemoveMyTasksTabStatusIDs removes the "my_tasks_tab_statuses" edge to MyTasksTabStatus entities by IDs.
func (wuo *WorkspaceUpdateOne) RemoveMyTasksTabStatusIDs(ids ...ulid.ID) *WorkspaceUpdateOne {
	wuo.mutation.RemoveMyTasksTabStatusIDs(ids...)
	return wuo
}

// RemoveMyTasksTabStatuses removes "my_tasks_tab_statuses" edges to MyTasksTabStatus entities.
func (wuo *WorkspaceUpdateOne) RemoveMyTasksTabStatuses(m ...*MyTasksTabStatus) *WorkspaceUpdateOne {
	ids := make([]ulid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wuo.RemoveMyTasksTabStatusIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkspaceUpdateOne) Select(field string, fields ...string) *WorkspaceUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workspace entity.
func (wuo *WorkspaceUpdateOne) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) SaveX(ctx context.Context) *Workspace {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkspaceUpdateOne) check() error {
	if v, ok := wuo.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wuo.mutation.TeammateID(); wuo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (wuo *WorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *Workspace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspace.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Workspace.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspace.FieldID)
		for _, f := range fields {
			if !workspace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldDescription,
		})
	}
	if wuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspace.TeammateTable,
			Columns: []string{workspace.TeammateColumn},
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
	if nodes := wuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspace.TeammateTable,
			Columns: []string{workspace.TeammateColumn},
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
	if wuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
	if nodes := wuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !wuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.ProjectsTable,
			Columns: []string{workspace.ProjectsColumn},
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
	if wuo.mutation.WorkspaceTeammatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWorkspaceTeammatesIDs(); len(nodes) > 0 && !wuo.mutation.WorkspaceTeammatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WorkspaceTeammatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.WorkspaceTeammatesTable,
			Columns: []string{workspace.WorkspaceTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceteammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.FavoriteWorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedFavoriteWorkspacesIDs(); len(nodes) > 0 && !wuo.mutation.FavoriteWorkspacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.FavoriteWorkspacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.FavoriteWorkspacesTable,
			Columns: []string{workspace.FavoriteWorkspacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteworkspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.MyTasksTabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedMyTasksTabStatusesIDs(); len(nodes) > 0 && !wuo.mutation.MyTasksTabStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.MyTasksTabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.MyTasksTabStatusesTable,
			Columns: []string{workspace.MyTasksTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mytaskstabstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Workspace{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
