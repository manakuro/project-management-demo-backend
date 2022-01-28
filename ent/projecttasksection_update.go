// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/ent/schema/ulid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskSectionUpdate is the builder for updating ProjectTaskSection entities.
type ProjectTaskSectionUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTaskSectionMutation
}

// Where appends a list predicates to the ProjectTaskSectionUpdate builder.
func (ptsu *ProjectTaskSectionUpdate) Where(ps ...predicate.ProjectTaskSection) *ProjectTaskSectionUpdate {
	ptsu.mutation.Where(ps...)
	return ptsu
}

// SetProjectID sets the "project_id" field.
func (ptsu *ProjectTaskSectionUpdate) SetProjectID(u ulid.ID) *ProjectTaskSectionUpdate {
	ptsu.mutation.SetProjectID(u)
	return ptsu
}

// SetName sets the "name" field.
func (ptsu *ProjectTaskSectionUpdate) SetName(s string) *ProjectTaskSectionUpdate {
	ptsu.mutation.SetName(s)
	return ptsu
}

// SetProject sets the "project" edge to the Project entity.
func (ptsu *ProjectTaskSectionUpdate) SetProject(p *Project) *ProjectTaskSectionUpdate {
	return ptsu.SetProjectID(p.ID)
}

// AddProjectTaskIDs adds the "project_tasks" edge to the ProjectTask entity by IDs.
func (ptsu *ProjectTaskSectionUpdate) AddProjectTaskIDs(ids ...ulid.ID) *ProjectTaskSectionUpdate {
	ptsu.mutation.AddProjectTaskIDs(ids...)
	return ptsu
}

// AddProjectTasks adds the "project_tasks" edges to the ProjectTask entity.
func (ptsu *ProjectTaskSectionUpdate) AddProjectTasks(p ...*ProjectTask) *ProjectTaskSectionUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptsu.AddProjectTaskIDs(ids...)
}

// Mutation returns the ProjectTaskSectionMutation object of the builder.
func (ptsu *ProjectTaskSectionUpdate) Mutation() *ProjectTaskSectionMutation {
	return ptsu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptsu *ProjectTaskSectionUpdate) ClearProject() *ProjectTaskSectionUpdate {
	ptsu.mutation.ClearProject()
	return ptsu
}

// ClearProjectTasks clears all "project_tasks" edges to the ProjectTask entity.
func (ptsu *ProjectTaskSectionUpdate) ClearProjectTasks() *ProjectTaskSectionUpdate {
	ptsu.mutation.ClearProjectTasks()
	return ptsu
}

// RemoveProjectTaskIDs removes the "project_tasks" edge to ProjectTask entities by IDs.
func (ptsu *ProjectTaskSectionUpdate) RemoveProjectTaskIDs(ids ...ulid.ID) *ProjectTaskSectionUpdate {
	ptsu.mutation.RemoveProjectTaskIDs(ids...)
	return ptsu
}

// RemoveProjectTasks removes "project_tasks" edges to ProjectTask entities.
func (ptsu *ProjectTaskSectionUpdate) RemoveProjectTasks(p ...*ProjectTask) *ProjectTaskSectionUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptsu.RemoveProjectTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptsu *ProjectTaskSectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptsu.hooks) == 0 {
		if err = ptsu.check(); err != nil {
			return 0, err
		}
		affected, err = ptsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptsu.check(); err != nil {
				return 0, err
			}
			ptsu.mutation = mutation
			affected, err = ptsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptsu.hooks) - 1; i >= 0; i-- {
			if ptsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptsu *ProjectTaskSectionUpdate) SaveX(ctx context.Context) int {
	affected, err := ptsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptsu *ProjectTaskSectionUpdate) Exec(ctx context.Context) error {
	_, err := ptsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsu *ProjectTaskSectionUpdate) ExecX(ctx context.Context) {
	if err := ptsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptsu *ProjectTaskSectionUpdate) check() error {
	if v, ok := ptsu.mutation.Name(); ok {
		if err := projecttasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ptsu.mutation.ProjectID(); ptsu.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (ptsu *ProjectTaskSectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttasksection.Table,
			Columns: projecttasksection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttasksection.FieldID,
			},
		},
	}
	if ps := ptsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttasksection.FieldName,
		})
	}
	if ptsu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttasksection.ProjectTable,
			Columns: []string{projecttasksection.ProjectColumn},
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
	if nodes := ptsu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttasksection.ProjectTable,
			Columns: []string{projecttasksection.ProjectColumn},
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
	if ptsu.mutation.ProjectTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptsu.mutation.RemovedProjectTasksIDs(); len(nodes) > 0 && !ptsu.mutation.ProjectTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptsu.mutation.ProjectTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttasksection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectTaskSectionUpdateOne is the builder for updating a single ProjectTaskSection entity.
type ProjectTaskSectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectTaskSectionMutation
}

// SetProjectID sets the "project_id" field.
func (ptsuo *ProjectTaskSectionUpdateOne) SetProjectID(u ulid.ID) *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.SetProjectID(u)
	return ptsuo
}

// SetName sets the "name" field.
func (ptsuo *ProjectTaskSectionUpdateOne) SetName(s string) *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.SetName(s)
	return ptsuo
}

// SetProject sets the "project" edge to the Project entity.
func (ptsuo *ProjectTaskSectionUpdateOne) SetProject(p *Project) *ProjectTaskSectionUpdateOne {
	return ptsuo.SetProjectID(p.ID)
}

// AddProjectTaskIDs adds the "project_tasks" edge to the ProjectTask entity by IDs.
func (ptsuo *ProjectTaskSectionUpdateOne) AddProjectTaskIDs(ids ...ulid.ID) *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.AddProjectTaskIDs(ids...)
	return ptsuo
}

// AddProjectTasks adds the "project_tasks" edges to the ProjectTask entity.
func (ptsuo *ProjectTaskSectionUpdateOne) AddProjectTasks(p ...*ProjectTask) *ProjectTaskSectionUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptsuo.AddProjectTaskIDs(ids...)
}

// Mutation returns the ProjectTaskSectionMutation object of the builder.
func (ptsuo *ProjectTaskSectionUpdateOne) Mutation() *ProjectTaskSectionMutation {
	return ptsuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptsuo *ProjectTaskSectionUpdateOne) ClearProject() *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.ClearProject()
	return ptsuo
}

// ClearProjectTasks clears all "project_tasks" edges to the ProjectTask entity.
func (ptsuo *ProjectTaskSectionUpdateOne) ClearProjectTasks() *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.ClearProjectTasks()
	return ptsuo
}

// RemoveProjectTaskIDs removes the "project_tasks" edge to ProjectTask entities by IDs.
func (ptsuo *ProjectTaskSectionUpdateOne) RemoveProjectTaskIDs(ids ...ulid.ID) *ProjectTaskSectionUpdateOne {
	ptsuo.mutation.RemoveProjectTaskIDs(ids...)
	return ptsuo
}

// RemoveProjectTasks removes "project_tasks" edges to ProjectTask entities.
func (ptsuo *ProjectTaskSectionUpdateOne) RemoveProjectTasks(p ...*ProjectTask) *ProjectTaskSectionUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptsuo.RemoveProjectTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptsuo *ProjectTaskSectionUpdateOne) Select(field string, fields ...string) *ProjectTaskSectionUpdateOne {
	ptsuo.fields = append([]string{field}, fields...)
	return ptsuo
}

// Save executes the query and returns the updated ProjectTaskSection entity.
func (ptsuo *ProjectTaskSectionUpdateOne) Save(ctx context.Context) (*ProjectTaskSection, error) {
	var (
		err  error
		node *ProjectTaskSection
	)
	if len(ptsuo.hooks) == 0 {
		if err = ptsuo.check(); err != nil {
			return nil, err
		}
		node, err = ptsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptsuo.check(); err != nil {
				return nil, err
			}
			ptsuo.mutation = mutation
			node, err = ptsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptsuo.hooks) - 1; i >= 0; i-- {
			if ptsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptsuo *ProjectTaskSectionUpdateOne) SaveX(ctx context.Context) *ProjectTaskSection {
	node, err := ptsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptsuo *ProjectTaskSectionUpdateOne) Exec(ctx context.Context) error {
	_, err := ptsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsuo *ProjectTaskSectionUpdateOne) ExecX(ctx context.Context) {
	if err := ptsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptsuo *ProjectTaskSectionUpdateOne) check() error {
	if v, ok := ptsuo.mutation.Name(); ok {
		if err := projecttasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ptsuo.mutation.ProjectID(); ptsuo.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (ptsuo *ProjectTaskSectionUpdateOne) sqlSave(ctx context.Context) (_node *ProjectTaskSection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttasksection.Table,
			Columns: projecttasksection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttasksection.FieldID,
			},
		},
	}
	id, ok := ptsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectTaskSection.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ptsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttasksection.FieldID)
		for _, f := range fields {
			if !projecttasksection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projecttasksection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttasksection.FieldName,
		})
	}
	if ptsuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttasksection.ProjectTable,
			Columns: []string{projecttasksection.ProjectColumn},
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
	if nodes := ptsuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttasksection.ProjectTable,
			Columns: []string{projecttasksection.ProjectColumn},
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
	if ptsuo.mutation.ProjectTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptsuo.mutation.RemovedProjectTasksIDs(); len(nodes) > 0 && !ptsuo.mutation.ProjectTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptsuo.mutation.ProjectTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttasksection.ProjectTasksTable,
			Columns: []string{projecttasksection.ProjectTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectTaskSection{config: ptsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttasksection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
