// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttaskcolumn"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskColumnUpdate is the builder for updating ProjectTaskColumn entities.
type ProjectTaskColumnUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTaskColumnMutation
}

// Where appends a list predicates to the ProjectTaskColumnUpdate builder.
func (ptcu *ProjectTaskColumnUpdate) Where(ps ...predicate.ProjectTaskColumn) *ProjectTaskColumnUpdate {
	ptcu.mutation.Where(ps...)
	return ptcu
}

// SetProjectID sets the "project_id" field.
func (ptcu *ProjectTaskColumnUpdate) SetProjectID(u ulid.ID) *ProjectTaskColumnUpdate {
	ptcu.mutation.SetProjectID(u)
	return ptcu
}

// SetTaskColumnID sets the "task_column_id" field.
func (ptcu *ProjectTaskColumnUpdate) SetTaskColumnID(u ulid.ID) *ProjectTaskColumnUpdate {
	ptcu.mutation.SetTaskColumnID(u)
	return ptcu
}

// SetWidth sets the "width" field.
func (ptcu *ProjectTaskColumnUpdate) SetWidth(s string) *ProjectTaskColumnUpdate {
	ptcu.mutation.SetWidth(s)
	return ptcu
}

// SetDisabled sets the "disabled" field.
func (ptcu *ProjectTaskColumnUpdate) SetDisabled(b bool) *ProjectTaskColumnUpdate {
	ptcu.mutation.SetDisabled(b)
	return ptcu
}

// SetCustomizable sets the "customizable" field.
func (ptcu *ProjectTaskColumnUpdate) SetCustomizable(b bool) *ProjectTaskColumnUpdate {
	ptcu.mutation.SetCustomizable(b)
	return ptcu
}

// SetOrder sets the "order" field.
func (ptcu *ProjectTaskColumnUpdate) SetOrder(i int) *ProjectTaskColumnUpdate {
	ptcu.mutation.ResetOrder()
	ptcu.mutation.SetOrder(i)
	return ptcu
}

// AddOrder adds i to the "order" field.
func (ptcu *ProjectTaskColumnUpdate) AddOrder(i int) *ProjectTaskColumnUpdate {
	ptcu.mutation.AddOrder(i)
	return ptcu
}

// SetProject sets the "project" edge to the Project entity.
func (ptcu *ProjectTaskColumnUpdate) SetProject(p *Project) *ProjectTaskColumnUpdate {
	return ptcu.SetProjectID(p.ID)
}

// SetTaskColumn sets the "taskColumn" edge to the TaskColumn entity.
func (ptcu *ProjectTaskColumnUpdate) SetTaskColumn(t *TaskColumn) *ProjectTaskColumnUpdate {
	return ptcu.SetTaskColumnID(t.ID)
}

// Mutation returns the ProjectTaskColumnMutation object of the builder.
func (ptcu *ProjectTaskColumnUpdate) Mutation() *ProjectTaskColumnMutation {
	return ptcu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptcu *ProjectTaskColumnUpdate) ClearProject() *ProjectTaskColumnUpdate {
	ptcu.mutation.ClearProject()
	return ptcu
}

// ClearTaskColumn clears the "taskColumn" edge to the TaskColumn entity.
func (ptcu *ProjectTaskColumnUpdate) ClearTaskColumn() *ProjectTaskColumnUpdate {
	ptcu.mutation.ClearTaskColumn()
	return ptcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptcu *ProjectTaskColumnUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptcu.hooks) == 0 {
		if err = ptcu.check(); err != nil {
			return 0, err
		}
		affected, err = ptcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptcu.check(); err != nil {
				return 0, err
			}
			ptcu.mutation = mutation
			affected, err = ptcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptcu.hooks) - 1; i >= 0; i-- {
			if ptcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptcu *ProjectTaskColumnUpdate) SaveX(ctx context.Context) int {
	affected, err := ptcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptcu *ProjectTaskColumnUpdate) Exec(ctx context.Context) error {
	_, err := ptcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcu *ProjectTaskColumnUpdate) ExecX(ctx context.Context) {
	if err := ptcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptcu *ProjectTaskColumnUpdate) check() error {
	if v, ok := ptcu.mutation.Width(); ok {
		if err := projecttaskcolumn.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "ProjectTaskColumn.width": %w`, err)}
		}
	}
	if _, ok := ptcu.mutation.ProjectID(); ptcu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskColumn.project"`)
	}
	if _, ok := ptcu.mutation.TaskColumnID(); ptcu.mutation.TaskColumnCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskColumn.taskColumn"`)
	}
	return nil
}

func (ptcu *ProjectTaskColumnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskcolumn.Table,
			Columns: projecttaskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskcolumn.FieldID,
			},
		},
	}
	if ps := ptcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptcu.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttaskcolumn.FieldWidth,
		})
	}
	if value, ok := ptcu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: projecttaskcolumn.FieldDisabled,
		})
	}
	if value, ok := ptcu.mutation.Customizable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: projecttaskcolumn.FieldCustomizable,
		})
	}
	if value, ok := ptcu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: projecttaskcolumn.FieldOrder,
		})
	}
	if value, ok := ptcu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: projecttaskcolumn.FieldOrder,
		})
	}
	if ptcu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.ProjectTable,
			Columns: []string{projecttaskcolumn.ProjectColumn},
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
	if nodes := ptcu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.ProjectTable,
			Columns: []string{projecttaskcolumn.ProjectColumn},
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
	if ptcu.mutation.TaskColumnCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.TaskColumnTable,
			Columns: []string{projecttaskcolumn.TaskColumnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskcolumn.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptcu.mutation.TaskColumnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.TaskColumnTable,
			Columns: []string{projecttaskcolumn.TaskColumnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskcolumn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttaskcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectTaskColumnUpdateOne is the builder for updating a single ProjectTaskColumn entity.
type ProjectTaskColumnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectTaskColumnMutation
}

// SetProjectID sets the "project_id" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetProjectID(u ulid.ID) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.SetProjectID(u)
	return ptcuo
}

// SetTaskColumnID sets the "task_column_id" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetTaskColumnID(u ulid.ID) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.SetTaskColumnID(u)
	return ptcuo
}

// SetWidth sets the "width" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetWidth(s string) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.SetWidth(s)
	return ptcuo
}

// SetDisabled sets the "disabled" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetDisabled(b bool) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.SetDisabled(b)
	return ptcuo
}

// SetCustomizable sets the "customizable" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetCustomizable(b bool) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.SetCustomizable(b)
	return ptcuo
}

// SetOrder sets the "order" field.
func (ptcuo *ProjectTaskColumnUpdateOne) SetOrder(i int) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.ResetOrder()
	ptcuo.mutation.SetOrder(i)
	return ptcuo
}

// AddOrder adds i to the "order" field.
func (ptcuo *ProjectTaskColumnUpdateOne) AddOrder(i int) *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.AddOrder(i)
	return ptcuo
}

// SetProject sets the "project" edge to the Project entity.
func (ptcuo *ProjectTaskColumnUpdateOne) SetProject(p *Project) *ProjectTaskColumnUpdateOne {
	return ptcuo.SetProjectID(p.ID)
}

// SetTaskColumn sets the "taskColumn" edge to the TaskColumn entity.
func (ptcuo *ProjectTaskColumnUpdateOne) SetTaskColumn(t *TaskColumn) *ProjectTaskColumnUpdateOne {
	return ptcuo.SetTaskColumnID(t.ID)
}

// Mutation returns the ProjectTaskColumnMutation object of the builder.
func (ptcuo *ProjectTaskColumnUpdateOne) Mutation() *ProjectTaskColumnMutation {
	return ptcuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptcuo *ProjectTaskColumnUpdateOne) ClearProject() *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.ClearProject()
	return ptcuo
}

// ClearTaskColumn clears the "taskColumn" edge to the TaskColumn entity.
func (ptcuo *ProjectTaskColumnUpdateOne) ClearTaskColumn() *ProjectTaskColumnUpdateOne {
	ptcuo.mutation.ClearTaskColumn()
	return ptcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptcuo *ProjectTaskColumnUpdateOne) Select(field string, fields ...string) *ProjectTaskColumnUpdateOne {
	ptcuo.fields = append([]string{field}, fields...)
	return ptcuo
}

// Save executes the query and returns the updated ProjectTaskColumn entity.
func (ptcuo *ProjectTaskColumnUpdateOne) Save(ctx context.Context) (*ProjectTaskColumn, error) {
	var (
		err  error
		node *ProjectTaskColumn
	)
	if len(ptcuo.hooks) == 0 {
		if err = ptcuo.check(); err != nil {
			return nil, err
		}
		node, err = ptcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptcuo.check(); err != nil {
				return nil, err
			}
			ptcuo.mutation = mutation
			node, err = ptcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptcuo.hooks) - 1; i >= 0; i-- {
			if ptcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptcuo *ProjectTaskColumnUpdateOne) SaveX(ctx context.Context) *ProjectTaskColumn {
	node, err := ptcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptcuo *ProjectTaskColumnUpdateOne) Exec(ctx context.Context) error {
	_, err := ptcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcuo *ProjectTaskColumnUpdateOne) ExecX(ctx context.Context) {
	if err := ptcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptcuo *ProjectTaskColumnUpdateOne) check() error {
	if v, ok := ptcuo.mutation.Width(); ok {
		if err := projecttaskcolumn.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "ProjectTaskColumn.width": %w`, err)}
		}
	}
	if _, ok := ptcuo.mutation.ProjectID(); ptcuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskColumn.project"`)
	}
	if _, ok := ptcuo.mutation.TaskColumnID(); ptcuo.mutation.TaskColumnCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskColumn.taskColumn"`)
	}
	return nil
}

func (ptcuo *ProjectTaskColumnUpdateOne) sqlSave(ctx context.Context) (_node *ProjectTaskColumn, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskcolumn.Table,
			Columns: projecttaskcolumn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskcolumn.FieldID,
			},
		},
	}
	id, ok := ptcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectTaskColumn.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttaskcolumn.FieldID)
		for _, f := range fields {
			if !projecttaskcolumn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projecttaskcolumn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptcuo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttaskcolumn.FieldWidth,
		})
	}
	if value, ok := ptcuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: projecttaskcolumn.FieldDisabled,
		})
	}
	if value, ok := ptcuo.mutation.Customizable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: projecttaskcolumn.FieldCustomizable,
		})
	}
	if value, ok := ptcuo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: projecttaskcolumn.FieldOrder,
		})
	}
	if value, ok := ptcuo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: projecttaskcolumn.FieldOrder,
		})
	}
	if ptcuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.ProjectTable,
			Columns: []string{projecttaskcolumn.ProjectColumn},
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
	if nodes := ptcuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.ProjectTable,
			Columns: []string{projecttaskcolumn.ProjectColumn},
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
	if ptcuo.mutation.TaskColumnCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.TaskColumnTable,
			Columns: []string{projecttaskcolumn.TaskColumnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskcolumn.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptcuo.mutation.TaskColumnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskcolumn.TaskColumnTable,
			Columns: []string{projecttaskcolumn.TaskColumnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskcolumn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectTaskColumn{config: ptcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttaskcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
