// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectbasecolor"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/projectteammate"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetWorkspaceID sets the "workspace_id" field.
func (pu *ProjectUpdate) SetWorkspaceID(u ulid.ID) *ProjectUpdate {
	pu.mutation.SetWorkspaceID(u)
	return pu
}

// SetProjectBaseColorID sets the "project_base_color_id" field.
func (pu *ProjectUpdate) SetProjectBaseColorID(u ulid.ID) *ProjectUpdate {
	pu.mutation.SetProjectBaseColorID(u)
	return pu
}

// SetProjectLightColorID sets the "project_light_color_id" field.
func (pu *ProjectUpdate) SetProjectLightColorID(u ulid.ID) *ProjectUpdate {
	pu.mutation.SetProjectLightColorID(u)
	return pu
}

// SetIconID sets the "icon_id" field.
func (pu *ProjectUpdate) SetIconID(u ulid.ID) *ProjectUpdate {
	pu.mutation.SetIconID(u)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *ProjectUpdate) SetCreatedBy(u ulid.ID) *ProjectUpdate {
	pu.mutation.SetCreatedBy(u)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(e editor.Description) *ProjectUpdate {
	pu.mutation.SetDescription(e)
	return pu
}

// SetDescriptionTitle sets the "description_title" field.
func (pu *ProjectUpdate) SetDescriptionTitle(s string) *ProjectUpdate {
	pu.mutation.SetDescriptionTitle(s)
	return pu
}

// SetDueDate sets the "due_date" field.
func (pu *ProjectUpdate) SetDueDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetDueDate(t)
	return pu
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDueDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetDueDate(*t)
	}
	return pu
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (pu *ProjectUpdate) SetWorkspace(w *Workspace) *ProjectUpdate {
	return pu.SetWorkspaceID(w.ID)
}

// SetProjectBaseColor sets the "project_base_color" edge to the ProjectBaseColor entity.
func (pu *ProjectUpdate) SetProjectBaseColor(p *ProjectBaseColor) *ProjectUpdate {
	return pu.SetProjectBaseColorID(p.ID)
}

// SetProjectLightColor sets the "project_light_color" edge to the ProjectLightColor entity.
func (pu *ProjectUpdate) SetProjectLightColor(p *ProjectLightColor) *ProjectUpdate {
	return pu.SetProjectLightColorID(p.ID)
}

// SetIcon sets the "icon" edge to the Icon entity.
func (pu *ProjectUpdate) SetIcon(i *Icon) *ProjectUpdate {
	return pu.SetIconID(i.ID)
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (pu *ProjectUpdate) SetTeammateID(id ulid.ID) *ProjectUpdate {
	pu.mutation.SetTeammateID(id)
	return pu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (pu *ProjectUpdate) SetTeammate(t *Teammate) *ProjectUpdate {
	return pu.SetTeammateID(t.ID)
}

// AddProjectTeammateIDs adds the "project_teammates" edge to the ProjectTeammate entity by IDs.
func (pu *ProjectUpdate) AddProjectTeammateIDs(ids ...ulid.ID) *ProjectUpdate {
	pu.mutation.AddProjectTeammateIDs(ids...)
	return pu
}

// AddProjectTeammates adds the "project_teammates" edges to the ProjectTeammate entity.
func (pu *ProjectUpdate) AddProjectTeammates(p ...*ProjectTeammate) *ProjectUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProjectTeammateIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (pu *ProjectUpdate) ClearWorkspace() *ProjectUpdate {
	pu.mutation.ClearWorkspace()
	return pu
}

// ClearProjectBaseColor clears the "project_base_color" edge to the ProjectBaseColor entity.
func (pu *ProjectUpdate) ClearProjectBaseColor() *ProjectUpdate {
	pu.mutation.ClearProjectBaseColor()
	return pu
}

// ClearProjectLightColor clears the "project_light_color" edge to the ProjectLightColor entity.
func (pu *ProjectUpdate) ClearProjectLightColor() *ProjectUpdate {
	pu.mutation.ClearProjectLightColor()
	return pu
}

// ClearIcon clears the "icon" edge to the Icon entity.
func (pu *ProjectUpdate) ClearIcon() *ProjectUpdate {
	pu.mutation.ClearIcon()
	return pu
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (pu *ProjectUpdate) ClearTeammate() *ProjectUpdate {
	pu.mutation.ClearTeammate()
	return pu
}

// ClearProjectTeammates clears all "project_teammates" edges to the ProjectTeammate entity.
func (pu *ProjectUpdate) ClearProjectTeammates() *ProjectUpdate {
	pu.mutation.ClearProjectTeammates()
	return pu
}

// RemoveProjectTeammateIDs removes the "project_teammates" edge to ProjectTeammate entities by IDs.
func (pu *ProjectUpdate) RemoveProjectTeammateIDs(ids ...ulid.ID) *ProjectUpdate {
	pu.mutation.RemoveProjectTeammateIDs(ids...)
	return pu
}

// RemoveProjectTeammates removes "project_teammates" edges to ProjectTeammate entities.
func (pu *ProjectUpdate) RemoveProjectTeammates(p ...*ProjectTeammate) *ProjectUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProjectTeammateIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.DescriptionTitle(); ok {
		if err := project.DescriptionTitleValidator(v); err != nil {
			return &ValidationError{Name: "description_title", err: fmt.Errorf("ent: validator failed for field \"description_title\": %w", err)}
		}
	}
	if _, ok := pu.mutation.WorkspaceID(); pu.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	if _, ok := pu.mutation.ProjectBaseColorID(); pu.mutation.ProjectBaseColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_base_color\"")
	}
	if _, ok := pu.mutation.ProjectLightColorID(); pu.mutation.ProjectLightColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_light_color\"")
	}
	if _, ok := pu.mutation.IconID(); pu.mutation.IconCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"icon\"")
	}
	if _, ok := pu.mutation.TeammateID(); pu.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: project.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: project.FieldDescription,
		})
	}
	if value, ok := pu.mutation.DescriptionTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescriptionTitle,
		})
	}
	if value, ok := pu.mutation.DueDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldDueDate,
		})
	}
	if pu.mutation.WorkspaceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.WorkspaceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProjectBaseColorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProjectBaseColorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProjectLightColorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProjectLightColorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.IconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.IconTable,
			Columns: []string{project.IconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: icon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.IconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.IconTable,
			Columns: []string{project.IconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: icon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TeammateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TeammateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProjectTeammatesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedProjectTeammatesIDs(); len(nodes) > 0 && !pu.mutation.ProjectTeammatesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProjectTeammatesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetWorkspaceID sets the "workspace_id" field.
func (puo *ProjectUpdateOne) SetWorkspaceID(u ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetWorkspaceID(u)
	return puo
}

// SetProjectBaseColorID sets the "project_base_color_id" field.
func (puo *ProjectUpdateOne) SetProjectBaseColorID(u ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetProjectBaseColorID(u)
	return puo
}

// SetProjectLightColorID sets the "project_light_color_id" field.
func (puo *ProjectUpdateOne) SetProjectLightColorID(u ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetProjectLightColorID(u)
	return puo
}

// SetIconID sets the "icon_id" field.
func (puo *ProjectUpdateOne) SetIconID(u ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetIconID(u)
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *ProjectUpdateOne) SetCreatedBy(u ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetCreatedBy(u)
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(e editor.Description) *ProjectUpdateOne {
	puo.mutation.SetDescription(e)
	return puo
}

// SetDescriptionTitle sets the "description_title" field.
func (puo *ProjectUpdateOne) SetDescriptionTitle(s string) *ProjectUpdateOne {
	puo.mutation.SetDescriptionTitle(s)
	return puo
}

// SetDueDate sets the "due_date" field.
func (puo *ProjectUpdateOne) SetDueDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetDueDate(t)
	return puo
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDueDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetDueDate(*t)
	}
	return puo
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (puo *ProjectUpdateOne) SetWorkspace(w *Workspace) *ProjectUpdateOne {
	return puo.SetWorkspaceID(w.ID)
}

// SetProjectBaseColor sets the "project_base_color" edge to the ProjectBaseColor entity.
func (puo *ProjectUpdateOne) SetProjectBaseColor(p *ProjectBaseColor) *ProjectUpdateOne {
	return puo.SetProjectBaseColorID(p.ID)
}

// SetProjectLightColor sets the "project_light_color" edge to the ProjectLightColor entity.
func (puo *ProjectUpdateOne) SetProjectLightColor(p *ProjectLightColor) *ProjectUpdateOne {
	return puo.SetProjectLightColorID(p.ID)
}

// SetIcon sets the "icon" edge to the Icon entity.
func (puo *ProjectUpdateOne) SetIcon(i *Icon) *ProjectUpdateOne {
	return puo.SetIconID(i.ID)
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (puo *ProjectUpdateOne) SetTeammateID(id ulid.ID) *ProjectUpdateOne {
	puo.mutation.SetTeammateID(id)
	return puo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (puo *ProjectUpdateOne) SetTeammate(t *Teammate) *ProjectUpdateOne {
	return puo.SetTeammateID(t.ID)
}

// AddProjectTeammateIDs adds the "project_teammates" edge to the ProjectTeammate entity by IDs.
func (puo *ProjectUpdateOne) AddProjectTeammateIDs(ids ...ulid.ID) *ProjectUpdateOne {
	puo.mutation.AddProjectTeammateIDs(ids...)
	return puo
}

// AddProjectTeammates adds the "project_teammates" edges to the ProjectTeammate entity.
func (puo *ProjectUpdateOne) AddProjectTeammates(p ...*ProjectTeammate) *ProjectUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProjectTeammateIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (puo *ProjectUpdateOne) ClearWorkspace() *ProjectUpdateOne {
	puo.mutation.ClearWorkspace()
	return puo
}

// ClearProjectBaseColor clears the "project_base_color" edge to the ProjectBaseColor entity.
func (puo *ProjectUpdateOne) ClearProjectBaseColor() *ProjectUpdateOne {
	puo.mutation.ClearProjectBaseColor()
	return puo
}

// ClearProjectLightColor clears the "project_light_color" edge to the ProjectLightColor entity.
func (puo *ProjectUpdateOne) ClearProjectLightColor() *ProjectUpdateOne {
	puo.mutation.ClearProjectLightColor()
	return puo
}

// ClearIcon clears the "icon" edge to the Icon entity.
func (puo *ProjectUpdateOne) ClearIcon() *ProjectUpdateOne {
	puo.mutation.ClearIcon()
	return puo
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (puo *ProjectUpdateOne) ClearTeammate() *ProjectUpdateOne {
	puo.mutation.ClearTeammate()
	return puo
}

// ClearProjectTeammates clears all "project_teammates" edges to the ProjectTeammate entity.
func (puo *ProjectUpdateOne) ClearProjectTeammates() *ProjectUpdateOne {
	puo.mutation.ClearProjectTeammates()
	return puo
}

// RemoveProjectTeammateIDs removes the "project_teammates" edge to ProjectTeammate entities by IDs.
func (puo *ProjectUpdateOne) RemoveProjectTeammateIDs(ids ...ulid.ID) *ProjectUpdateOne {
	puo.mutation.RemoveProjectTeammateIDs(ids...)
	return puo
}

// RemoveProjectTeammates removes "project_teammates" edges to ProjectTeammate entities.
func (puo *ProjectUpdateOne) RemoveProjectTeammates(p ...*ProjectTeammate) *ProjectUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProjectTeammateIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.DescriptionTitle(); ok {
		if err := project.DescriptionTitleValidator(v); err != nil {
			return &ValidationError{Name: "description_title", err: fmt.Errorf("ent: validator failed for field \"description_title\": %w", err)}
		}
	}
	if _, ok := puo.mutation.WorkspaceID(); puo.mutation.WorkspaceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workspace\"")
	}
	if _, ok := puo.mutation.ProjectBaseColorID(); puo.mutation.ProjectBaseColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_base_color\"")
	}
	if _, ok := puo.mutation.ProjectLightColorID(); puo.mutation.ProjectLightColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_light_color\"")
	}
	if _, ok := puo.mutation.IconID(); puo.mutation.IconCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"icon\"")
	}
	if _, ok := puo.mutation.TeammateID(); puo.mutation.TeammateCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teammate\"")
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: project.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Project.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: project.FieldDescription,
		})
	}
	if value, ok := puo.mutation.DescriptionTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescriptionTitle,
		})
	}
	if value, ok := puo.mutation.DueDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldDueDate,
		})
	}
	if puo.mutation.WorkspaceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.WorkspaceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProjectBaseColorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProjectBaseColorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProjectLightColorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProjectLightColorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.IconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.IconTable,
			Columns: []string{project.IconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: icon.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.IconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.IconTable,
			Columns: []string{project.IconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: icon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TeammateCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TeammateIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProjectTeammatesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedProjectTeammatesIDs(); len(nodes) > 0 && !puo.mutation.ProjectTeammatesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProjectTeammatesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
