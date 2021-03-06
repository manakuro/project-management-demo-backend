// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projecttaskliststatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistsortstatus"
	"project-management-demo-backend/ent/teammatetaskliststatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskListSortStatusUpdate is the builder for updating TaskListSortStatus entities.
type TaskListSortStatusUpdate struct {
	config
	hooks    []Hook
	mutation *TaskListSortStatusMutation
}

// Where appends a list predicates to the TaskListSortStatusUpdate builder.
func (tlssu *TaskListSortStatusUpdate) Where(ps ...predicate.TaskListSortStatus) *TaskListSortStatusUpdate {
	tlssu.mutation.Where(ps...)
	return tlssu
}

// SetName sets the "name" field.
func (tlssu *TaskListSortStatusUpdate) SetName(s string) *TaskListSortStatusUpdate {
	tlssu.mutation.SetName(s)
	return tlssu
}

// SetStatusCode sets the "status_code" field.
func (tlssu *TaskListSortStatusUpdate) SetStatusCode(tc tasklistsortstatus.StatusCode) *TaskListSortStatusUpdate {
	tlssu.mutation.SetStatusCode(tc)
	return tlssu
}

// AddTeammateTaskListStatuseIDs adds the "teammateTaskListStatuses" edge to the TeammateTaskListStatus entity by IDs.
func (tlssu *TaskListSortStatusUpdate) AddTeammateTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdate {
	tlssu.mutation.AddTeammateTaskListStatuseIDs(ids...)
	return tlssu
}

// AddTeammateTaskListStatuses adds the "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (tlssu *TaskListSortStatusUpdate) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListSortStatusUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlssu.AddTeammateTaskListStatuseIDs(ids...)
}

// AddProjectTaskListStatuseIDs adds the "projectTaskListStatuses" edge to the ProjectTaskListStatus entity by IDs.
func (tlssu *TaskListSortStatusUpdate) AddProjectTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdate {
	tlssu.mutation.AddProjectTaskListStatuseIDs(ids...)
	return tlssu
}

// AddProjectTaskListStatuses adds the "projectTaskListStatuses" edges to the ProjectTaskListStatus entity.
func (tlssu *TaskListSortStatusUpdate) AddProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListSortStatusUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlssu.AddProjectTaskListStatuseIDs(ids...)
}

// Mutation returns the TaskListSortStatusMutation object of the builder.
func (tlssu *TaskListSortStatusUpdate) Mutation() *TaskListSortStatusMutation {
	return tlssu.mutation
}

// ClearTeammateTaskListStatuses clears all "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (tlssu *TaskListSortStatusUpdate) ClearTeammateTaskListStatuses() *TaskListSortStatusUpdate {
	tlssu.mutation.ClearTeammateTaskListStatuses()
	return tlssu
}

// RemoveTeammateTaskListStatuseIDs removes the "teammateTaskListStatuses" edge to TeammateTaskListStatus entities by IDs.
func (tlssu *TaskListSortStatusUpdate) RemoveTeammateTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdate {
	tlssu.mutation.RemoveTeammateTaskListStatuseIDs(ids...)
	return tlssu
}

// RemoveTeammateTaskListStatuses removes "teammateTaskListStatuses" edges to TeammateTaskListStatus entities.
func (tlssu *TaskListSortStatusUpdate) RemoveTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListSortStatusUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlssu.RemoveTeammateTaskListStatuseIDs(ids...)
}

// ClearProjectTaskListStatuses clears all "projectTaskListStatuses" edges to the ProjectTaskListStatus entity.
func (tlssu *TaskListSortStatusUpdate) ClearProjectTaskListStatuses() *TaskListSortStatusUpdate {
	tlssu.mutation.ClearProjectTaskListStatuses()
	return tlssu
}

// RemoveProjectTaskListStatuseIDs removes the "projectTaskListStatuses" edge to ProjectTaskListStatus entities by IDs.
func (tlssu *TaskListSortStatusUpdate) RemoveProjectTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdate {
	tlssu.mutation.RemoveProjectTaskListStatuseIDs(ids...)
	return tlssu
}

// RemoveProjectTaskListStatuses removes "projectTaskListStatuses" edges to ProjectTaskListStatus entities.
func (tlssu *TaskListSortStatusUpdate) RemoveProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListSortStatusUpdate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlssu.RemoveProjectTaskListStatuseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlssu *TaskListSortStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlssu.hooks) == 0 {
		if err = tlssu.check(); err != nil {
			return 0, err
		}
		affected, err = tlssu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskListSortStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlssu.check(); err != nil {
				return 0, err
			}
			tlssu.mutation = mutation
			affected, err = tlssu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlssu.hooks) - 1; i >= 0; i-- {
			if tlssu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlssu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlssu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlssu *TaskListSortStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := tlssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlssu *TaskListSortStatusUpdate) Exec(ctx context.Context) error {
	_, err := tlssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlssu *TaskListSortStatusUpdate) ExecX(ctx context.Context) {
	if err := tlssu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlssu *TaskListSortStatusUpdate) check() error {
	if v, ok := tlssu.mutation.Name(); ok {
		if err := tasklistsortstatus.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TaskListSortStatus.name": %w`, err)}
		}
	}
	if v, ok := tlssu.mutation.StatusCode(); ok {
		if err := tasklistsortstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf(`ent: validator failed for field "TaskListSortStatus.status_code": %w`, err)}
		}
	}
	return nil
}

func (tlssu *TaskListSortStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklistsortstatus.Table,
			Columns: tasklistsortstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistsortstatus.FieldID,
			},
		},
	}
	if ps := tlssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlssu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklistsortstatus.FieldName,
		})
	}
	if value, ok := tlssu.mutation.StatusCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tasklistsortstatus.FieldStatusCode,
		})
	}
	if tlssu.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if nodes := tlssu.mutation.RemovedTeammateTaskListStatusesIDs(); len(nodes) > 0 && !tlssu.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if nodes := tlssu.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if tlssu.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	if nodes := tlssu.mutation.RemovedProjectTaskListStatusesIDs(); len(nodes) > 0 && !tlssu.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	if nodes := tlssu.mutation.ProjectTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tlssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklistsortstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskListSortStatusUpdateOne is the builder for updating a single TaskListSortStatus entity.
type TaskListSortStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskListSortStatusMutation
}

// SetName sets the "name" field.
func (tlssuo *TaskListSortStatusUpdateOne) SetName(s string) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.SetName(s)
	return tlssuo
}

// SetStatusCode sets the "status_code" field.
func (tlssuo *TaskListSortStatusUpdateOne) SetStatusCode(tc tasklistsortstatus.StatusCode) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.SetStatusCode(tc)
	return tlssuo
}

// AddTeammateTaskListStatuseIDs adds the "teammateTaskListStatuses" edge to the TeammateTaskListStatus entity by IDs.
func (tlssuo *TaskListSortStatusUpdateOne) AddTeammateTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.AddTeammateTaskListStatuseIDs(ids...)
	return tlssuo
}

// AddTeammateTaskListStatuses adds the "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (tlssuo *TaskListSortStatusUpdateOne) AddTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListSortStatusUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlssuo.AddTeammateTaskListStatuseIDs(ids...)
}

// AddProjectTaskListStatuseIDs adds the "projectTaskListStatuses" edge to the ProjectTaskListStatus entity by IDs.
func (tlssuo *TaskListSortStatusUpdateOne) AddProjectTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.AddProjectTaskListStatuseIDs(ids...)
	return tlssuo
}

// AddProjectTaskListStatuses adds the "projectTaskListStatuses" edges to the ProjectTaskListStatus entity.
func (tlssuo *TaskListSortStatusUpdateOne) AddProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListSortStatusUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlssuo.AddProjectTaskListStatuseIDs(ids...)
}

// Mutation returns the TaskListSortStatusMutation object of the builder.
func (tlssuo *TaskListSortStatusUpdateOne) Mutation() *TaskListSortStatusMutation {
	return tlssuo.mutation
}

// ClearTeammateTaskListStatuses clears all "teammateTaskListStatuses" edges to the TeammateTaskListStatus entity.
func (tlssuo *TaskListSortStatusUpdateOne) ClearTeammateTaskListStatuses() *TaskListSortStatusUpdateOne {
	tlssuo.mutation.ClearTeammateTaskListStatuses()
	return tlssuo
}

// RemoveTeammateTaskListStatuseIDs removes the "teammateTaskListStatuses" edge to TeammateTaskListStatus entities by IDs.
func (tlssuo *TaskListSortStatusUpdateOne) RemoveTeammateTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.RemoveTeammateTaskListStatuseIDs(ids...)
	return tlssuo
}

// RemoveTeammateTaskListStatuses removes "teammateTaskListStatuses" edges to TeammateTaskListStatus entities.
func (tlssuo *TaskListSortStatusUpdateOne) RemoveTeammateTaskListStatuses(t ...*TeammateTaskListStatus) *TaskListSortStatusUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlssuo.RemoveTeammateTaskListStatuseIDs(ids...)
}

// ClearProjectTaskListStatuses clears all "projectTaskListStatuses" edges to the ProjectTaskListStatus entity.
func (tlssuo *TaskListSortStatusUpdateOne) ClearProjectTaskListStatuses() *TaskListSortStatusUpdateOne {
	tlssuo.mutation.ClearProjectTaskListStatuses()
	return tlssuo
}

// RemoveProjectTaskListStatuseIDs removes the "projectTaskListStatuses" edge to ProjectTaskListStatus entities by IDs.
func (tlssuo *TaskListSortStatusUpdateOne) RemoveProjectTaskListStatuseIDs(ids ...ulid.ID) *TaskListSortStatusUpdateOne {
	tlssuo.mutation.RemoveProjectTaskListStatuseIDs(ids...)
	return tlssuo
}

// RemoveProjectTaskListStatuses removes "projectTaskListStatuses" edges to ProjectTaskListStatus entities.
func (tlssuo *TaskListSortStatusUpdateOne) RemoveProjectTaskListStatuses(p ...*ProjectTaskListStatus) *TaskListSortStatusUpdateOne {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tlssuo.RemoveProjectTaskListStatuseIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tlssuo *TaskListSortStatusUpdateOne) Select(field string, fields ...string) *TaskListSortStatusUpdateOne {
	tlssuo.fields = append([]string{field}, fields...)
	return tlssuo
}

// Save executes the query and returns the updated TaskListSortStatus entity.
func (tlssuo *TaskListSortStatusUpdateOne) Save(ctx context.Context) (*TaskListSortStatus, error) {
	var (
		err  error
		node *TaskListSortStatus
	)
	if len(tlssuo.hooks) == 0 {
		if err = tlssuo.check(); err != nil {
			return nil, err
		}
		node, err = tlssuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskListSortStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlssuo.check(); err != nil {
				return nil, err
			}
			tlssuo.mutation = mutation
			node, err = tlssuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tlssuo.hooks) - 1; i >= 0; i-- {
			if tlssuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlssuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlssuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlssuo *TaskListSortStatusUpdateOne) SaveX(ctx context.Context) *TaskListSortStatus {
	node, err := tlssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tlssuo *TaskListSortStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := tlssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlssuo *TaskListSortStatusUpdateOne) ExecX(ctx context.Context) {
	if err := tlssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlssuo *TaskListSortStatusUpdateOne) check() error {
	if v, ok := tlssuo.mutation.Name(); ok {
		if err := tasklistsortstatus.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TaskListSortStatus.name": %w`, err)}
		}
	}
	if v, ok := tlssuo.mutation.StatusCode(); ok {
		if err := tasklistsortstatus.StatusCodeValidator(v); err != nil {
			return &ValidationError{Name: "status_code", err: fmt.Errorf(`ent: validator failed for field "TaskListSortStatus.status_code": %w`, err)}
		}
	}
	return nil
}

func (tlssuo *TaskListSortStatusUpdateOne) sqlSave(ctx context.Context) (_node *TaskListSortStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklistsortstatus.Table,
			Columns: tasklistsortstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklistsortstatus.FieldID,
			},
		},
	}
	id, ok := tlssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskListSortStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tlssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklistsortstatus.FieldID)
		for _, f := range fields {
			if !tasklistsortstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tasklistsortstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tlssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlssuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklistsortstatus.FieldName,
		})
	}
	if value, ok := tlssuo.mutation.StatusCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: tasklistsortstatus.FieldStatusCode,
		})
	}
	if tlssuo.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if nodes := tlssuo.mutation.RemovedTeammateTaskListStatusesIDs(); len(nodes) > 0 && !tlssuo.mutation.TeammateTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if nodes := tlssuo.mutation.TeammateTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.TeammateTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.TeammateTaskListStatusesColumn},
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
	if tlssuo.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	if nodes := tlssuo.mutation.RemovedProjectTaskListStatusesIDs(); len(nodes) > 0 && !tlssuo.mutation.ProjectTaskListStatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	if nodes := tlssuo.mutation.ProjectTaskListStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklistsortstatus.ProjectTaskListStatusesTable,
			Columns: []string{tasklistsortstatus.ProjectTaskListStatusesColumn},
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
	_node = &TaskListSortStatus{config: tlssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tlssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklistsortstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
