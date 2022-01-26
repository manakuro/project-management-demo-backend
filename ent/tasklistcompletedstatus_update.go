// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"project-management-demo-backend/ent/teammatetaskliststatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskListCompletedStatusUpdate is the builder for updating TaskListCompletedStatus entities.
type TaskListCompletedStatusUpdate struct {
	config
	hooks    []Hook
	mutation *TaskListCompletedStatusMutation
}

// Where appends a list predicates to the TaskListCompletedStatusUpdate builder.
func (tlcsu *TaskListCompletedStatusUpdate) Where(ps ...predicate.TaskListCompletedStatus) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.Where(ps...)
	return tlcsu
}

// SetName sets the "name" field.
func (tlcsu *TaskListCompletedStatusUpdate) SetName(s string) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.SetName(s)
	return tlcsu
}

// SetStatusCode sets the "status_code" field.
func (tlcsu *TaskListCompletedStatusUpdate) SetStatusCode(tc tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.SetStatusCode(tc)
	return tlcsu
}

// AddTeammateTaskListStatusIDs adds the "teammate_task_list_statuses" edge to the TeammateTaskListStatus entity by IDs.
func (tlcsu *TaskListCompletedStatusUpdate) AddTeammateTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.AddTeammateTaskListStatusIDs(ids...)
	return tlcsu
}

// AddTeammateTaskListStatuses adds the "teammate_task_list_statuses" edges to the TeammateTaskListStatus entity.
func (tlcsu *TaskListCompletedStatusUpdate) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListCompletedStatusUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlcsu.AddTeammateTaskListStatusIDs(ids...)
}

// AddProjectTaskListStatusIDs adds the "project_task_list_statuses" edge to the ProjectTaskListStatus entity by IDs.
func (tlcsu *TaskListCompletedStatusUpdate) AddProjectTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.AddProjectTaskListStatusIDs(ids...)
	return tlcsu
}

// AddProjectTaskListStatuses adds the "project_task_list_statuses" edges to the ProjectTaskListStatus entity.
func (tlcsu *TaskListCompletedStatusUpdate) AddProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListCompletedStatusUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlcsu.AddProjectTaskListStatusIDs(ids...)
}

// Mutation returns the TaskListCompletedStatusMutation object of the builder.
func (tlcsu *TaskListCompletedStatusUpdate) Mutation() *TaskListCompletedStatusMutation {
	return tlcsu.mutation
}

// ClearTeammateTaskListStatuses clears all "teammate_task_list_statuses" edges to the TeammateTaskListStatus entity.
func (tlcsu *TaskListCompletedStatusUpdate) ClearTeammateTaskListStatuses() *TaskListCompletedStatusUpdate {
	tlcsu.mutation.ClearTeammateTaskListStatuses()
	return tlcsu
}

// RemoveTeammateTaskListStatusIDs removes the "teammate_task_list_statuses" edge to TeammateTaskListStatus entities by IDs.
func (tlcsu *TaskListCompletedStatusUpdate) RemoveTeammateTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.RemoveTeammateTaskListStatusIDs(ids...)
	return tlcsu
}

// RemoveTeammateTaskListStatuses removes "teammate_task_list_statuses" edges to TeammateTaskListStatus entities.
func (tlcsu *TaskListCompletedStatusUpdate) RemoveTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListCompletedStatusUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlcsu.RemoveTeammateTaskListStatusIDs(ids...)
}

// ClearProjectTaskListStatuses clears all "project_task_list_statuses" edges to the ProjectTaskListStatus entity.
func (tlcsu *TaskListCompletedStatusUpdate) ClearProjectTaskListStatuses() *TaskListCompletedStatusUpdate {
	tlcsu.mutation.ClearProjectTaskListStatuses()
	return tlcsu
}

// RemoveProjectTaskListStatusIDs removes the "project_task_list_statuses" edge to ProjectTaskListStatus entities by IDs.
func (tlcsu *TaskListCompletedStatusUpdate) RemoveProjectTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdate {
	tlcsu.mutation.RemoveProjectTaskListStatusIDs(ids...)
	return tlcsu
}

// RemoveProjectTaskListStatuses removes "project_task_list_statuses" edges to ProjectTaskListStatus entities.
func (tlcsu *TaskListCompletedStatusUpdate) RemoveProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListCompletedStatusUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlcsu.RemoveProjectTaskListStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlcsu *TaskListCompletedStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlcsu.hooks) == 0 {
		if err = tlcsu.check(); err != nil {
			return 0, err
		}
		affected, err = tlcsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskListCompletedStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlcsu.check(); err != nil {
				return 0, err
			}
			tlcsu.mutation = mutation
			affected, err = tlcsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlcsu.hooks) - 1; i >= 0; i-- {
			if tlcsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlcsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlcsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlcsu *TaskListCompletedStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := tlcsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlcsu *TaskListCompletedStatusUpdate) Exec(ctx context.Context) error {
	_, err := tlcsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcsu *TaskListCompletedStatusUpdate) ExecX(ctx context.Context) {
	if err := tlcsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlcsu *TaskListCompletedStatusUpdate) check() error {
	if v, ok := tlcsu.mutation.Name(); ok {
		if err := tasklistcompletedstatus.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tlcsu.mutation.StatusCode(); ok {
		if err := tasklistcompletedstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf("ent: validator failed for field \"status_code\": %w", err)}
		}
	}
	return nil
}

func (tlcsu *TaskListCompletedStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklistcompletedstatus.Table,
			Columns: tasklistcompletedstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistcompletedstatus.FieldID,
			},
		},
	}
	if ps := tlcsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlcsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklistcompletedstatus.FieldName,
		})
	}
	if value, ok := tlcsu.mutation.StatusCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tasklistcompletedstatus.FieldStatusCode,
		})
	}
	if tlcsu.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetaskliststatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsu.mutation.RemovedTeammateTaskListStatusesIDs(); len(nodes) > 0 && !tlcsu.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsu.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tlcsu.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsu.mutation.RemovedProjectTaskListStatusesIDs(); len(nodes) > 0 && !tlcsu.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsu.mutation.ProjectTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tlcsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklistcompletedstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskListCompletedStatusUpdateOne is the builder for updating a single TaskListCompletedStatus entity.
type TaskListCompletedStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskListCompletedStatusMutation
}

// SetName sets the "name" field.
func (tlcsuo *TaskListCompletedStatusUpdateOne) SetName(s string) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.SetName(s)
	return tlcsuo
}

// SetStatusCode sets the "status_code" field.
func (tlcsuo *TaskListCompletedStatusUpdateOne) SetStatusCode(tc tasklistcompletedstatus.StatusCode) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.SetStatusCode(tc)
	return tlcsuo
}

// AddTeammateTaskListStatusIDs adds the "teammate_task_list_statuses" edge to the TeammateTaskListStatus entity by IDs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) AddTeammateTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.AddTeammateTaskListStatusIDs(ids...)
	return tlcsuo
}

// AddTeammateTaskListStatuses adds the "teammate_task_list_statuses" edges to the TeammateTaskListStatus entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListCompletedStatusUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlcsuo.AddTeammateTaskListStatusIDs(ids...)
}

// AddProjectTaskListStatusIDs adds the "project_task_list_statuses" edge to the ProjectTaskListStatus entity by IDs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) AddProjectTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.AddProjectTaskListStatusIDs(ids...)
	return tlcsuo
}

// AddProjectTaskListStatuses adds the "project_task_list_statuses" edges to the ProjectTaskListStatus entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) AddProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListCompletedStatusUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlcsuo.AddProjectTaskListStatusIDs(ids...)
}

// Mutation returns the TaskListCompletedStatusMutation object of the builder.
func (tlcsuo *TaskListCompletedStatusUpdateOne) Mutation() *TaskListCompletedStatusMutation {
	return tlcsuo.mutation
}

// ClearTeammateTaskListStatuses clears all "teammate_task_list_statuses" edges to the TeammateTaskListStatus entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) ClearTeammateTaskListStatuses() *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.ClearTeammateTaskListStatuses()
	return tlcsuo
}

// RemoveTeammateTaskListStatusIDs removes the "teammate_task_list_statuses" edge to TeammateTaskListStatus entities by IDs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) RemoveTeammateTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.RemoveTeammateTaskListStatusIDs(ids...)
	return tlcsuo
}

// RemoveTeammateTaskListStatuses removes "teammate_task_list_statuses" edges to TeammateTaskListStatus entities.
func (tlcsuo *TaskListCompletedStatusUpdateOne) RemoveTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListCompletedStatusUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlcsuo.RemoveTeammateTaskListStatusIDs(ids...)
}

// ClearProjectTaskListStatuses clears all "project_task_list_statuses" edges to the ProjectTaskListStatus entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) ClearProjectTaskListStatuses() *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.ClearProjectTaskListStatuses()
	return tlcsuo
}

// RemoveProjectTaskListStatusIDs removes the "project_task_list_statuses" edge to ProjectTaskListStatus entities by IDs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) RemoveProjectTaskListStatusIDs(ids ...ulid.ID) *TaskListCompletedStatusUpdateOne {
	tlcsuo.mutation.RemoveProjectTaskListStatusIDs(ids...)
	return tlcsuo
}

// RemoveProjectTaskListStatuses removes "project_task_list_statuses" edges to ProjectTaskListStatus entities.
func (tlcsuo *TaskListCompletedStatusUpdateOne) RemoveProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListCompletedStatusUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlcsuo.RemoveProjectTaskListStatusIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tlcsuo *TaskListCompletedStatusUpdateOne) Select(field string, fields ...string) *TaskListCompletedStatusUpdateOne {
	tlcsuo.fields = append([]string{field}, fields...)
	return tlcsuo
}

// Save executes the query and returns the updated TaskListCompletedStatus entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) Save(ctx context.Context) (*TaskListCompletedStatus, error) {
	var (
		err  error
		node *TaskListCompletedStatus
	)
	if len(tlcsuo.hooks) == 0 {
		if err = tlcsuo.check(); err != nil {
			return nil, err
		}
		node, err = tlcsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskListCompletedStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlcsuo.check(); err != nil {
				return nil, err
			}
			tlcsuo.mutation = mutation
			node, err = tlcsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tlcsuo.hooks) - 1; i >= 0; i-- {
			if tlcsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlcsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlcsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) SaveX(ctx context.Context) *TaskListCompletedStatus {
	node, err := tlcsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tlcsuo *TaskListCompletedStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := tlcsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcsuo *TaskListCompletedStatusUpdateOne) ExecX(ctx context.Context) {
	if err := tlcsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlcsuo *TaskListCompletedStatusUpdateOne) check() error {
	if v, ok := tlcsuo.mutation.Name(); ok {
		if err := tasklistcompletedstatus.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tlcsuo.mutation.StatusCode(); ok {
		if err := tasklistcompletedstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf("ent: validator failed for field \"status_code\": %w", err)}
		}
	}
	return nil
}

func (tlcsuo *TaskListCompletedStatusUpdateOne) sqlSave(ctx context.Context) (_node *TaskListCompletedStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklistcompletedstatus.Table,
			Columns: tasklistcompletedstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistcompletedstatus.FieldID,
			},
		},
	}
	id, ok := tlcsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TaskListCompletedStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tlcsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklistcompletedstatus.FieldID)
		for _, f := range fields {
			if !tasklistcompletedstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tasklistcompletedstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tlcsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlcsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklistcompletedstatus.FieldName,
		})
	}
	if value, ok := tlcsuo.mutation.StatusCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tasklistcompletedstatus.FieldStatusCode,
		})
	}
	if tlcsuo.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetaskliststatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsuo.mutation.RemovedTeammateTaskListStatusesIDs(); len(nodes) > 0 && !tlcsuo.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsuo.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.TeammateTaskListStatusesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tlcsuo.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsuo.mutation.RemovedProjectTaskListStatusesIDs(); len(nodes) > 0 && !tlcsuo.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlcsuo.mutation.ProjectTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistcompletedstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistcompletedstatus.ProjectTaskListStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: projecttaskliststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TaskListCompletedStatus{config: tlcsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tlcsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklistcompletedstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
