// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tag"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetaskcolumn"
	"project-management-demo-backend/ent/teammatetaskliststatus"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/teammatetasktabstatus"
	"project-management-demo-backend/ent/workspace"
	"project-management-demo-backend/ent/workspaceteammate"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceCreate is the builder for creating a Workspace entity.
type WorkspaceCreate struct {
	config
	mutation *WorkspaceMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (wc *WorkspaceCreate) SetCreatedBy(u ulid.ID) *WorkspaceCreate {
	wc.mutation.SetCreatedBy(u)
	return wc
}

// SetName sets the "name" field.
func (wc *WorkspaceCreate) SetName(s string) *WorkspaceCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkspaceCreate) SetDescription(e editor.Description) *WorkspaceCreate {
	wc.mutation.SetDescription(e)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkspaceCreate) SetCreatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableCreatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkspaceCreate) SetUpdatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableUpdatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkspaceCreate) SetID(u ulid.ID) *WorkspaceCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableID(u *ulid.ID) *WorkspaceCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (wc *WorkspaceCreate) SetTeammateID(id ulid.ID) *WorkspaceCreate {
	wc.mutation.SetTeammateID(id)
	return wc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wc *WorkspaceCreate) SetTeammate(t *Teammate) *WorkspaceCreate {
	return wc.SetTeammateID(t.ID)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (wc *WorkspaceCreate) AddProjectIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddProjectIDs(ids...)
	return wc
}

// AddProjects adds the "projects" edges to the Project entity.
func (wc *WorkspaceCreate) AddProjects(p ...*Project) *WorkspaceCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wc.AddProjectIDs(ids...)
}

// AddWorkspaceTeammateIDs adds the "workspaceTeammates" edge to the WorkspaceTeammate entity by IDs.
func (wc *WorkspaceCreate) AddWorkspaceTeammateIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddWorkspaceTeammateIDs(ids...)
	return wc
}

// AddWorkspaceTeammates adds the "workspaceTeammates" edges to the WorkspaceTeammate entity.
func (wc *WorkspaceCreate) AddWorkspaceTeammates(w ...*WorkspaceTeammate) *WorkspaceCreate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkspaceTeammateIDs(ids...)
}

// AddFavoriteWorkspaceIDs adds the "favoriteWorkspaces" edge to the FavoriteWorkspace entity by IDs.
func (wc *WorkspaceCreate) AddFavoriteWorkspaceIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddFavoriteWorkspaceIDs(ids...)
	return wc
}

// AddFavoriteWorkspaces adds the "favoriteWorkspaces" edges to the FavoriteWorkspace entity.
func (wc *WorkspaceCreate) AddFavoriteWorkspaces(f ...*FavoriteWorkspace) *WorkspaceCreate {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return wc.AddFavoriteWorkspaceIDs(ids...)
}

// AddTeammateTaskTabStatuseIDs adds the "teammateTaskTabStatuses" edge to the TeammateTaskTabStatus entity by IDs.
func (wc *WorkspaceCreate) AddTeammateTaskTabStatuseIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTeammateTaskTabStatuseIDs(ids...)
	return wc
}

// AddTeammateTaskTabStatuses adds the "teammateTaskTabStatuses" edges to the TeammateTaskTabStatus entity.
func (wc *WorkspaceCreate) AddTeammateTaskTabStatuses(t ...*TeammateTaskTabStatus) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTeammateTaskTabStatuseIDs(ids...)
}

// AddTeammateTaskListStatuseIDs adds the "teammateTaskListStatuses" edge to the TeammateTaskListStatus entity by IDs.
func (wc *WorkspaceCreate) AddTeammateTaskListStatuseIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTeammateTaskListStatuseIDs(ids...)
	return wc
}

// AddTeammateTaskListStatuses adds the "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (wc *WorkspaceCreate) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTeammateTaskListStatuseIDs(ids...)
}

// AddTeammateTaskSectionIDs adds the "teammateTaskSections" edge to the TeammateTaskSection entity by IDs.
func (wc *WorkspaceCreate) AddTeammateTaskSectionIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTeammateTaskSectionIDs(ids...)
	return wc
}

// AddTeammateTaskSections adds the "teammateTaskSections" edges to the TeammateTaskSection entity.
func (wc *WorkspaceCreate) AddTeammateTaskSections(t ...*TeammateTaskSection) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTeammateTaskSectionIDs(ids...)
}

// AddTaskLikeIDs adds the "taskLikes" edge to the TaskLike entity by IDs.
func (wc *WorkspaceCreate) AddTaskLikeIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTaskLikeIDs(ids...)
	return wc
}

// AddTaskLikes adds the "taskLikes" edges to the TaskLike entity.
func (wc *WorkspaceCreate) AddTaskLikes(t ...*TaskLike) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTaskLikeIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (wc *WorkspaceCreate) AddTagIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTagIDs(ids...)
	return wc
}

// AddTags adds the "tags" edges to the Tag entity.
func (wc *WorkspaceCreate) AddTags(t ...*Tag) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTagIDs(ids...)
}

// AddTeammateTaskColumnIDs adds the "teammateTaskColumns" edge to the TeammateTaskColumn entity by IDs.
func (wc *WorkspaceCreate) AddTeammateTaskColumnIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTeammateTaskColumnIDs(ids...)
	return wc
}

// AddTeammateTaskColumns adds the "teammateTaskColumns" edges to the TeammateTaskColumn entity.
func (wc *WorkspaceCreate) AddTeammateTaskColumns(t ...*TeammateTaskColumn) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTeammateTaskColumnIDs(ids...)
}

// AddTeammateTaskIDs adds the "teammateTasks" edge to the TeammateTask entity by IDs.
func (wc *WorkspaceCreate) AddTeammateTaskIDs(ids ...ulid.ID) *WorkspaceCreate {
	wc.mutation.AddTeammateTaskIDs(ids...)
	return wc
}

// AddTeammateTasks adds the "teammateTasks" edges to the TeammateTask entity.
func (wc *WorkspaceCreate) AddTeammateTasks(t ...*TeammateTask) *WorkspaceCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddTeammateTaskIDs(ids...)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wc *WorkspaceCreate) Mutation() *WorkspaceMutation {
	return wc.mutation
}

// Save creates the Workspace in the database.
func (wc *WorkspaceCreate) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkspaceCreate) SaveX(ctx context.Context) *Workspace {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkspaceCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkspaceCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkspaceCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workspace.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workspace.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workspace.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkspaceCreate) check() error {
	if _, ok := wc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "created_by"`)}
	}
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := wc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	return nil
}

func (wc *WorkspaceCreate) sqlSave(ctx context.Context) (*Workspace, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(ulid.ID)
	}
	return _node, nil
}

func (wc *WorkspaceCreate) createSpec() (*Workspace, *sqlgraph.CreateSpec) {
	var (
		_node = &Workspace{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workspace.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspace.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
		_node.Name = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wc.mutation.TeammateIDs(); len(nodes) > 0 {
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
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkspaceTeammatesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.FavoriteWorkspacesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TeammateTaskTabStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TeammateTaskTabStatusesTable,
			Columns: []string{workspace.TeammateTaskTabStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetasktabstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TeammateTaskListStatusesTable,
			Columns: []string{workspace.TeammateTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TeammateTaskSectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TeammateTaskSectionsTable,
			Columns: []string{workspace.TeammateTaskSectionsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TaskLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TaskLikesTable,
			Columns: []string{workspace.TaskLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklike.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TagsTable,
			Columns: []string{workspace.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TeammateTaskColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TeammateTaskColumnsTable,
			Columns: []string{workspace.TeammateTaskColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetaskcolumn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workspace.TeammateTasksTable,
			Columns: []string{workspace.TeammateTasksColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkspaceCreateBulk is the builder for creating many Workspace entities in bulk.
type WorkspaceCreateBulk struct {
	config
	builders []*WorkspaceCreate
}

// Save creates the Workspace entities in the database.
func (wcb *WorkspaceCreateBulk) Save(ctx context.Context) ([]*Workspace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workspace, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) SaveX(ctx context.Context) []*Workspace {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkspaceCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
