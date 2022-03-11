// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/tasklistsortstatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskListStatusUpdate is the builder for updating ProjectTaskListStatus entities.
type ProjectTaskListStatusUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTaskListStatusMutation
}

// Where appends a list predicates to the ProjectTaskListStatusUpdate builder.
func (ptlsu *ProjectTaskListStatusUpdate) Where(ps ...predicate.ProjectTaskListStatus) *ProjectTaskListStatusUpdate {
	ptlsu.mutation.Where(ps...)
	return ptlsu
}

// SetProjectID sets the "project_id" field.
func (ptlsu *ProjectTaskListStatusUpdate) SetProjectID(u ulid.ID) *ProjectTaskListStatusUpdate {
	ptlsu.mutation.SetProjectID(u)
	return ptlsu
}

// SetTaskListCompletedStatusID sets the "task_list_completed_status_id" field.
func (ptlsu *ProjectTaskListStatusUpdate) SetTaskListCompletedStatusID(u ulid.ID) *ProjectTaskListStatusUpdate {
	ptlsu.mutation.SetTaskListCompletedStatusID(u)
	return ptlsu
}

// SetTaskListSortStatusID sets the "task_list_sort_status_id" field.
func (ptlsu *ProjectTaskListStatusUpdate) SetTaskListSortStatusID(u ulid.ID) *ProjectTaskListStatusUpdate {
	ptlsu.mutation.SetTaskListSortStatusID(u)
	return ptlsu
}

// SetProject sets the "project" edge to the Project entity.
func (ptlsu *ProjectTaskListStatusUpdate) SetProject(p *Project) *ProjectTaskListStatusUpdate {
	return ptlsu.SetProjectID(p.ID)
}

// SetTaskListCompletedStatus sets the "taskListCompletedStatus" edge to the TaskListCompletedStatus entity.
func (ptlsu *ProjectTaskListStatusUpdate) SetTaskListCompletedStatus(t *TaskListCompletedStatus) *ProjectTaskListStatusUpdate {
	return ptlsu.SetTaskListCompletedStatusID(t.ID)
}

// SetTaskListSortStatus sets the "taskListSortStatus" edge to the TaskListSortStatus entity.
func (ptlsu *ProjectTaskListStatusUpdate) SetTaskListSortStatus(t *TaskListSortStatus) *ProjectTaskListStatusUpdate {
	return ptlsu.SetTaskListSortStatusID(t.ID)
}

// Mutation returns the ProjectTaskListStatusMutation object of the builder.
func (ptlsu *ProjectTaskListStatusUpdate) Mutation() *ProjectTaskListStatusMutation {
	return ptlsu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptlsu *ProjectTaskListStatusUpdate) ClearProject() *ProjectTaskListStatusUpdate {
	ptlsu.mutation.ClearProject()
	return ptlsu
}

// ClearTaskListCompletedStatus clears the "taskListCompletedStatus" edge to the TaskListCompletedStatus entity.
func (ptlsu *ProjectTaskListStatusUpdate) ClearTaskListCompletedStatus() *ProjectTaskListStatusUpdate {
	ptlsu.mutation.ClearTaskListCompletedStatus()
	return ptlsu
}

// ClearTaskListSortStatus clears the "taskListSortStatus" edge to the TaskListSortStatus entity.
func (ptlsu *ProjectTaskListStatusUpdate) ClearTaskListSortStatus() *ProjectTaskListStatusUpdate {
	ptlsu.mutation.ClearTaskListSortStatus()
	return ptlsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptlsu *ProjectTaskListStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptlsu.hooks) == 0 {
		if err = ptlsu.check(); err != nil {
			return 0, err
		}
		affected, err = ptlsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskListStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptlsu.check(); err != nil {
				return 0, err
			}
			ptlsu.mutation = mutation
			affected, err = ptlsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptlsu.hooks) - 1; i >= 0; i-- {
			if ptlsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptlsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptlsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptlsu *ProjectTaskListStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := ptlsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptlsu *ProjectTaskListStatusUpdate) Exec(ctx context.Context) error {
	_, err := ptlsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptlsu *ProjectTaskListStatusUpdate) ExecX(ctx context.Context) {
	if err := ptlsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptlsu *ProjectTaskListStatusUpdate) check() error {
	if _, ok := ptlsu.mutation.ProjectID(); ptlsu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.project"`)
	}
	if _, ok := ptlsu.mutation.TaskListCompletedStatusID(); ptlsu.mutation.TaskListCompletedStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.taskListCompletedStatus"`)
	}
	if _, ok := ptlsu.mutation.TaskListSortStatusID(); ptlsu.mutation.TaskListSortStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.taskListSortStatus"`)
	}
	return nil
}

func (ptlsu *ProjectTaskListStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskliststatus.Table,
			Columns: projecttaskliststatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskliststatus.FieldID,
			},
		},
	}
	if ps := ptlsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ptlsu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.ProjectTable,
			Columns: []string{projecttaskliststatus.ProjectColumn},
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
	if nodes := ptlsu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.ProjectTable,
			Columns: []string{projecttaskliststatus.ProjectColumn},
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
	if ptlsu.mutation.TaskListCompletedStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListCompletedStatusTable,
			Columns: []string{projecttaskliststatus.TaskListCompletedStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistcompletedstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptlsu.mutation.TaskListCompletedStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListCompletedStatusTable,
			Columns: []string{projecttaskliststatus.TaskListCompletedStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistcompletedstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptlsu.mutation.TaskListSortStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListSortStatusTable,
			Columns: []string{projecttaskliststatus.TaskListSortStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistsortstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptlsu.mutation.TaskListSortStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListSortStatusTable,
			Columns: []string{projecttaskliststatus.TaskListSortStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistsortstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptlsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttaskliststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectTaskListStatusUpdateOne is the builder for updating a single ProjectTaskListStatus entity.
type ProjectTaskListStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectTaskListStatusMutation
}

// SetProjectID sets the "project_id" field.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetProjectID(u ulid.ID) *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.SetProjectID(u)
	return ptlsuo
}

// SetTaskListCompletedStatusID sets the "task_list_completed_status_id" field.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetTaskListCompletedStatusID(u ulid.ID) *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.SetTaskListCompletedStatusID(u)
	return ptlsuo
}

// SetTaskListSortStatusID sets the "task_list_sort_status_id" field.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetTaskListSortStatusID(u ulid.ID) *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.SetTaskListSortStatusID(u)
	return ptlsuo
}

// SetProject sets the "project" edge to the Project entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetProject(p *Project) *ProjectTaskListStatusUpdateOne {
	return ptlsuo.SetProjectID(p.ID)
}

// SetTaskListCompletedStatus sets the "taskListCompletedStatus" edge to the TaskListCompletedStatus entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetTaskListCompletedStatus(t *TaskListCompletedStatus) *ProjectTaskListStatusUpdateOne {
	return ptlsuo.SetTaskListCompletedStatusID(t.ID)
}

// SetTaskListSortStatus sets the "taskListSortStatus" edge to the TaskListSortStatus entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SetTaskListSortStatus(t *TaskListSortStatus) *ProjectTaskListStatusUpdateOne {
	return ptlsuo.SetTaskListSortStatusID(t.ID)
}

// Mutation returns the ProjectTaskListStatusMutation object of the builder.
func (ptlsuo *ProjectTaskListStatusUpdateOne) Mutation() *ProjectTaskListStatusMutation {
	return ptlsuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) ClearProject() *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.ClearProject()
	return ptlsuo
}

// ClearTaskListCompletedStatus clears the "taskListCompletedStatus" edge to the TaskListCompletedStatus entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) ClearTaskListCompletedStatus() *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.ClearTaskListCompletedStatus()
	return ptlsuo
}

// ClearTaskListSortStatus clears the "taskListSortStatus" edge to the TaskListSortStatus entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) ClearTaskListSortStatus() *ProjectTaskListStatusUpdateOne {
	ptlsuo.mutation.ClearTaskListSortStatus()
	return ptlsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptlsuo *ProjectTaskListStatusUpdateOne) Select(field string, fields ...string) *ProjectTaskListStatusUpdateOne {
	ptlsuo.fields = append([]string{field}, fields...)
	return ptlsuo
}

// Save executes the query and returns the updated ProjectTaskListStatus entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) Save(ctx context.Context) (*ProjectTaskListStatus, error) {
	var (
		err  error
		node *ProjectTaskListStatus
	)
	if len(ptlsuo.hooks) == 0 {
		if err = ptlsuo.check(); err != nil {
			return nil, err
		}
		node, err = ptlsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskListStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptlsuo.check(); err != nil {
				return nil, err
			}
			ptlsuo.mutation = mutation
			node, err = ptlsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptlsuo.hooks) - 1; i >= 0; i-- {
			if ptlsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptlsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptlsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptlsuo *ProjectTaskListStatusUpdateOne) SaveX(ctx context.Context) *ProjectTaskListStatus {
	node, err := ptlsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptlsuo *ProjectTaskListStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := ptlsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptlsuo *ProjectTaskListStatusUpdateOne) ExecX(ctx context.Context) {
	if err := ptlsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptlsuo *ProjectTaskListStatusUpdateOne) check() error {
	if _, ok := ptlsuo.mutation.ProjectID(); ptlsuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.project"`)
	}
	if _, ok := ptlsuo.mutation.TaskListCompletedStatusID(); ptlsuo.mutation.TaskListCompletedStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.taskListCompletedStatus"`)
	}
	if _, ok := ptlsuo.mutation.TaskListSortStatusID(); ptlsuo.mutation.TaskListSortStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTaskListStatus.taskListSortStatus"`)
	}
	return nil
}

func (ptlsuo *ProjectTaskListStatusUpdateOne) sqlSave(ctx context.Context) (_node *ProjectTaskListStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttaskliststatus.Table,
			Columns: projecttaskliststatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskliststatus.FieldID,
			},
		},
	}
	id, ok := ptlsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectTaskListStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptlsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttaskliststatus.FieldID)
		for _, f := range fields {
			if !projecttaskliststatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projecttaskliststatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptlsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ptlsuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.ProjectTable,
			Columns: []string{projecttaskliststatus.ProjectColumn},
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
	if nodes := ptlsuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.ProjectTable,
			Columns: []string{projecttaskliststatus.ProjectColumn},
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
	if ptlsuo.mutation.TaskListCompletedStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListCompletedStatusTable,
			Columns: []string{projecttaskliststatus.TaskListCompletedStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistcompletedstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptlsuo.mutation.TaskListCompletedStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListCompletedStatusTable,
			Columns: []string{projecttaskliststatus.TaskListCompletedStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistcompletedstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptlsuo.mutation.TaskListSortStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListSortStatusTable,
			Columns: []string{projecttaskliststatus.TaskListSortStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistsortstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptlsuo.mutation.TaskListSortStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttaskliststatus.TaskListSortStatusTable,
			Columns: []string{projecttaskliststatus.TaskListSortStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklistsortstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectTaskListStatus{config: ptlsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptlsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttaskliststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
