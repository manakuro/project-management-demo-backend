// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projecticon"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetWorkspaceID sets the "workspace_id" field.
func (pc *ProjectCreate) SetWorkspaceID(u ulid.ID) *ProjectCreate {
	pc.mutation.SetWorkspaceID(u)
	return pc
}

// SetProjectBaseColorID sets the "project_base_color_id" field.
func (pc *ProjectCreate) SetProjectBaseColorID(u ulid.ID) *ProjectCreate {
	pc.mutation.SetProjectBaseColorID(u)
	return pc
}

// SetProjectLightColorID sets the "project_light_color_id" field.
func (pc *ProjectCreate) SetProjectLightColorID(u ulid.ID) *ProjectCreate {
	pc.mutation.SetProjectLightColorID(u)
	return pc
}

// SetProjectIconID sets the "project_icon_id" field.
func (pc *ProjectCreate) SetProjectIconID(u ulid.ID) *ProjectCreate {
	pc.mutation.SetProjectIconID(u)
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *ProjectCreate) SetCreatedBy(u ulid.ID) *ProjectCreate {
	pc.mutation.SetCreatedBy(u)
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(e editor.Description) *ProjectCreate {
	pc.mutation.SetDescription(e)
	return pc
}

// SetDescriptionTitle sets the "description_title" field.
func (pc *ProjectCreate) SetDescriptionTitle(s string) *ProjectCreate {
	pc.mutation.SetDescriptionTitle(s)
	return pc
}

// SetDueDate sets the "due_date" field.
func (pc *ProjectCreate) SetDueDate(t time.Time) *ProjectCreate {
	pc.mutation.SetDueDate(t)
	return pc
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDueDate(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetDueDate(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(u ulid.ID) *ProjectCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableID(u *ulid.ID) *ProjectCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (pc *ProjectCreate) SetWorkspace(w *Workspace) *ProjectCreate {
	return pc.SetWorkspaceID(w.ID)
}

// SetProjectBaseColor sets the "project_base_color" edge to the ProjectBaseColor entity.
func (pc *ProjectCreate) SetProjectBaseColor(p *ProjectBaseColor) *ProjectCreate {
	return pc.SetProjectBaseColorID(p.ID)
}

// SetProjectLightColor sets the "project_light_color" edge to the ProjectLightColor entity.
func (pc *ProjectCreate) SetProjectLightColor(p *ProjectLightColor) *ProjectCreate {
	return pc.SetProjectLightColorID(p.ID)
}

// SetProjectIcon sets the "project_icon" edge to the ProjectIcon entity.
func (pc *ProjectCreate) SetProjectIcon(p *ProjectIcon) *ProjectCreate {
	return pc.SetProjectIconID(p.ID)
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (pc *ProjectCreate) SetTeammateID(id ulid.ID) *ProjectCreate {
	pc.mutation.SetTeammateID(id)
	return pc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (pc *ProjectCreate) SetTeammate(t *Teammate) *ProjectCreate {
	return pc.SetTeammateID(t.ID)
}

// AddProjectTeammateIDs adds the "project_teammates" edge to the ProjectTeammate entity by IDs.
func (pc *ProjectCreate) AddProjectTeammateIDs(ids ...ulid.ID) *ProjectCreate {
	pc.mutation.AddProjectTeammateIDs(ids...)
	return pc
}

// AddProjectTeammates adds the "project_teammates" edges to the ProjectTeammate entity.
func (pc *ProjectCreate) AddProjectTeammates(p ...*ProjectTeammate) *ProjectCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProjectTeammateIDs(ids...)
}

// AddFavoriteProjectIDs adds the "favorite_projects" edge to the FavoriteProject entity by IDs.
func (pc *ProjectCreate) AddFavoriteProjectIDs(ids ...ulid.ID) *ProjectCreate {
	pc.mutation.AddFavoriteProjectIDs(ids...)
	return pc
}

// AddFavoriteProjects adds the "favorite_projects" edges to the FavoriteProject entity.
func (pc *ProjectCreate) AddFavoriteProjects(f ...*FavoriteProject) *ProjectCreate {
	ids := make([]ulid.ID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddFavoriteProjectIDs(ids...)
}

// AddProjectTaskColumnIDs adds the "project_task_columns" edge to the ProjectTaskColumn entity by IDs.
func (pc *ProjectCreate) AddProjectTaskColumnIDs(ids ...ulid.ID) *ProjectCreate {
	pc.mutation.AddProjectTaskColumnIDs(ids...)
	return pc
}

// AddProjectTaskColumns adds the "project_task_columns" edges to the ProjectTaskColumn entity.
func (pc *ProjectCreate) AddProjectTaskColumns(p ...*ProjectTaskColumn) *ProjectCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProjectTaskColumnIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.DueDate(); !ok {
		v := project.DefaultDueDate()
		pc.mutation.SetDueDate(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := project.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := project.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "workspace_id"`)}
	}
	if _, ok := pc.mutation.ProjectBaseColorID(); !ok {
		return &ValidationError{Name: "project_base_color_id", err: errors.New(`ent: missing required field "project_base_color_id"`)}
	}
	if _, ok := pc.mutation.ProjectLightColorID(); !ok {
		return &ValidationError{Name: "project_light_color_id", err: errors.New(`ent: missing required field "project_light_color_id"`)}
	}
	if _, ok := pc.mutation.ProjectIconID(); !ok {
		return &ValidationError{Name: "project_icon_id", err: errors.New(`ent: missing required field "project_icon_id"`)}
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "created_by"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := pc.mutation.DescriptionTitle(); !ok {
		return &ValidationError{Name: "description_title", err: errors.New(`ent: missing required field "description_title"`)}
	}
	if v, ok := pc.mutation.DescriptionTitle(); ok {
		if err := project.DescriptionTitleValidator(v); err != nil {
			return &ValidationError{Name: "description_title", err: fmt.Errorf(`ent: validator failed for field "description_title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.DueDate(); !ok {
		return &ValidationError{Name: "due_date", err: errors.New(`ent: missing required field "due_date"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := pc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New("ent: missing required edge \"workspace\"")}
	}
	if _, ok := pc.mutation.ProjectBaseColorID(); !ok {
		return &ValidationError{Name: "project_base_color", err: errors.New("ent: missing required edge \"project_base_color\"")}
	}
	if _, ok := pc.mutation.ProjectLightColorID(); !ok {
		return &ValidationError{Name: "project_light_color", err: errors.New("ent: missing required edge \"project_light_color\"")}
	}
	if _, ok := pc.mutation.ProjectIconID(); !ok {
		return &ValidationError{Name: "project_icon", err: errors.New("ent: missing required edge \"project_icon\"")}
	}
	if _, ok := pc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: project.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: project.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: project.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pc.mutation.DescriptionTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescriptionTitle,
		})
		_node.DescriptionTitle = value
	}
	if value, ok := pc.mutation.DueDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldDueDate,
		})
		_node.DueDate = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.WorkspaceTable,
			Columns: []string{project.WorkspaceColumn},
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
		_node.WorkspaceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectBaseColorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectBaseColorTable,
			Columns: []string{project.ProjectBaseColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projectbasecolor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectBaseColorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectLightColorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectLightColorTable,
			Columns: []string{project.ProjectLightColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projectlightcolor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectLightColorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectIconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ProjectIconTable,
			Columns: []string{project.ProjectIconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecticon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectIconID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.TeammateTable,
			Columns: []string{project.TeammateColumn},
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
	if nodes := pc.mutation.ProjectTeammatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ProjectTeammatesTable,
			Columns: []string{project.ProjectTeammatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projectteammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FavoriteProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.FavoriteProjectsTable,
			Columns: []string{project.FavoriteProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: favoriteproject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectTaskColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ProjectTaskColumnsTable,
			Columns: []string{project.ProjectTaskColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskcolumn.FieldID,
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

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
