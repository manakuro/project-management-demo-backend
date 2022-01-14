// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projectlightcolor"
	"project-management-demo-backend/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectLightColorUpdate is the builder for updating ProjectLightColor entities.
type ProjectLightColorUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectLightColorMutation
}

// Where appends a list predicates to the ProjectLightColorUpdate builder.
func (plcu *ProjectLightColorUpdate) Where(ps ...predicate.ProjectLightColor) *ProjectLightColorUpdate {
	plcu.mutation.Where(ps...)
	return plcu
}

// SetColorID sets the "color_id" field.
func (plcu *ProjectLightColorUpdate) SetColorID(u ulid.ID) *ProjectLightColorUpdate {
	plcu.mutation.SetColorID(u)
	return plcu
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (plcu *ProjectLightColorUpdate) AddProjectIDs(ids ...ulid.ID) *ProjectLightColorUpdate {
	plcu.mutation.AddProjectIDs(ids...)
	return plcu
}

// AddProjects adds the "projects" edges to the Project entity.
func (plcu *ProjectLightColorUpdate) AddProjects(p ...*Project) *ProjectLightColorUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plcu.AddProjectIDs(ids...)
}

// SetColor sets the "color" edge to the Color entity.
func (plcu *ProjectLightColorUpdate) SetColor(c *Color) *ProjectLightColorUpdate {
	return plcu.SetColorID(c.ID)
}

// Mutation returns the ProjectLightColorMutation object of the builder.
func (plcu *ProjectLightColorUpdate) Mutation() *ProjectLightColorMutation {
	return plcu.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (plcu *ProjectLightColorUpdate) ClearProjects() *ProjectLightColorUpdate {
	plcu.mutation.ClearProjects()
	return plcu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (plcu *ProjectLightColorUpdate) RemoveProjectIDs(ids ...ulid.ID) *ProjectLightColorUpdate {
	plcu.mutation.RemoveProjectIDs(ids...)
	return plcu
}

// RemoveProjects removes "projects" edges to Project entities.
func (plcu *ProjectLightColorUpdate) RemoveProjects(p ...*Project) *ProjectLightColorUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plcu.RemoveProjectIDs(ids...)
}

// ClearColor clears the "color" edge to the Color entity.
func (plcu *ProjectLightColorUpdate) ClearColor() *ProjectLightColorUpdate {
	plcu.mutation.ClearColor()
	return plcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plcu *ProjectLightColorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(plcu.hooks) == 0 {
		if err = plcu.check(); err != nil {
			return 0, err
		}
		affected, err = plcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLightColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plcu.check(); err != nil {
				return 0, err
			}
			plcu.mutation = mutation
			affected, err = plcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(plcu.hooks) - 1; i >= 0; i-- {
			if plcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (plcu *ProjectLightColorUpdate) SaveX(ctx context.Context) int {
	affected, err := plcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plcu *ProjectLightColorUpdate) Exec(ctx context.Context) error {
	_, err := plcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plcu *ProjectLightColorUpdate) ExecX(ctx context.Context) {
	if err := plcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plcu *ProjectLightColorUpdate) check() error {
	if _, ok := plcu.mutation.ColorID(); plcu.mutation.ColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"color\"")
	}
	return nil
}

func (plcu *ProjectLightColorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlightcolor.Table,
			Columns: projectlightcolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projectlightcolor.FieldID,
			},
		},
	}
	if ps := plcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if plcu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if nodes := plcu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !plcu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if nodes := plcu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if plcu.mutation.ColorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlightcolor.ColorTable,
			Columns: []string{projectlightcolor.ColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: color.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plcu.mutation.ColorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlightcolor.ColorTable,
			Columns: []string{projectlightcolor.ColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: color.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlightcolor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectLightColorUpdateOne is the builder for updating a single ProjectLightColor entity.
type ProjectLightColorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectLightColorMutation
}

// SetColorID sets the "color_id" field.
func (plcuo *ProjectLightColorUpdateOne) SetColorID(u ulid.ID) *ProjectLightColorUpdateOne {
	plcuo.mutation.SetColorID(u)
	return plcuo
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (plcuo *ProjectLightColorUpdateOne) AddProjectIDs(ids ...ulid.ID) *ProjectLightColorUpdateOne {
	plcuo.mutation.AddProjectIDs(ids...)
	return plcuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (plcuo *ProjectLightColorUpdateOne) AddProjects(p ...*Project) *ProjectLightColorUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plcuo.AddProjectIDs(ids...)
}

// SetColor sets the "color" edge to the Color entity.
func (plcuo *ProjectLightColorUpdateOne) SetColor(c *Color) *ProjectLightColorUpdateOne {
	return plcuo.SetColorID(c.ID)
}

// Mutation returns the ProjectLightColorMutation object of the builder.
func (plcuo *ProjectLightColorUpdateOne) Mutation() *ProjectLightColorMutation {
	return plcuo.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (plcuo *ProjectLightColorUpdateOne) ClearProjects() *ProjectLightColorUpdateOne {
	plcuo.mutation.ClearProjects()
	return plcuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (plcuo *ProjectLightColorUpdateOne) RemoveProjectIDs(ids ...ulid.ID) *ProjectLightColorUpdateOne {
	plcuo.mutation.RemoveProjectIDs(ids...)
	return plcuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (plcuo *ProjectLightColorUpdateOne) RemoveProjects(p ...*Project) *ProjectLightColorUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plcuo.RemoveProjectIDs(ids...)
}

// ClearColor clears the "color" edge to the Color entity.
func (plcuo *ProjectLightColorUpdateOne) ClearColor() *ProjectLightColorUpdateOne {
	plcuo.mutation.ClearColor()
	return plcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (plcuo *ProjectLightColorUpdateOne) Select(field string, fields ...string) *ProjectLightColorUpdateOne {
	plcuo.fields = append([]string{field}, fields...)
	return plcuo
}

// Save executes the query and returns the updated ProjectLightColor entity.
func (plcuo *ProjectLightColorUpdateOne) Save(ctx context.Context) (*ProjectLightColor, error) {
	var (
		err  error
		node *ProjectLightColor
	)
	if len(plcuo.hooks) == 0 {
		if err = plcuo.check(); err != nil {
			return nil, err
		}
		node, err = plcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLightColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plcuo.check(); err != nil {
				return nil, err
			}
			plcuo.mutation = mutation
			node, err = plcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(plcuo.hooks) - 1; i >= 0; i-- {
			if plcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (plcuo *ProjectLightColorUpdateOne) SaveX(ctx context.Context) *ProjectLightColor {
	node, err := plcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (plcuo *ProjectLightColorUpdateOne) Exec(ctx context.Context) error {
	_, err := plcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plcuo *ProjectLightColorUpdateOne) ExecX(ctx context.Context) {
	if err := plcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plcuo *ProjectLightColorUpdateOne) check() error {
	if _, ok := plcuo.mutation.ColorID(); plcuo.mutation.ColorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"color\"")
	}
	return nil
}

func (plcuo *ProjectLightColorUpdateOne) sqlSave(ctx context.Context) (_node *ProjectLightColor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlightcolor.Table,
			Columns: projectlightcolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projectlightcolor.FieldID,
			},
		},
	}
	id, ok := plcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectLightColor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := plcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectlightcolor.FieldID)
		for _, f := range fields {
			if !projectlightcolor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectlightcolor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := plcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if plcuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if nodes := plcuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !plcuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if nodes := plcuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlightcolor.ProjectsTable,
			Columns: []string{projectlightcolor.ProjectsColumn},
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
	if plcuo.mutation.ColorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlightcolor.ColorTable,
			Columns: []string{projectlightcolor.ColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: color.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plcuo.mutation.ColorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlightcolor.ColorTable,
			Columns: []string{projectlightcolor.ColorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: color.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectLightColor{config: plcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, plcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlightcolor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
