// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/archivedtaskactivity"
	"project-management-demo-backend/ent/archivedworkspaceactivity"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskactivity"
	"project-management-demo-backend/ent/workspaceactivity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivityTypeUpdate is the builder for updating ActivityType entities.
type ActivityTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityTypeMutation
}

// Where appends a list predicates to the ActivityTypeUpdate builder.
func (atu *ActivityTypeUpdate) Where(ps ...predicate.ActivityType) *ActivityTypeUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetName sets the "name" field.
func (atu *ActivityTypeUpdate) SetName(s string) *ActivityTypeUpdate {
	atu.mutation.SetName(s)
	return atu
}

// SetTypeCode sets the "type_code" field.
func (atu *ActivityTypeUpdate) SetTypeCode(ac activitytype.TypeCode) *ActivityTypeUpdate {
	atu.mutation.SetTypeCode(ac)
	return atu
}

// AddTaskActivityIDs adds the "taskActivities" edge to the TaskActivity entity by IDs.
func (atu *ActivityTypeUpdate) AddTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.AddTaskActivityIDs(ids...)
	return atu
}

// AddTaskActivities adds the "taskActivities" edges to the TaskActivity entity.
func (atu *ActivityTypeUpdate) AddTaskActivities(t ...*TaskActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return atu.AddTaskActivityIDs(ids...)
}

// AddWorkspaceActivityIDs adds the "workspaceActivities" edge to the WorkspaceActivity entity by IDs.
func (atu *ActivityTypeUpdate) AddWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.AddWorkspaceActivityIDs(ids...)
	return atu
}

// AddWorkspaceActivities adds the "workspaceActivities" edges to the WorkspaceActivity entity.
func (atu *ActivityTypeUpdate) AddWorkspaceActivities(w ...*WorkspaceActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return atu.AddWorkspaceActivityIDs(ids...)
}

// AddArchivedTaskActivityIDs adds the "archivedTaskActivities" edge to the ArchivedTaskActivity entity by IDs.
func (atu *ActivityTypeUpdate) AddArchivedTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.AddArchivedTaskActivityIDs(ids...)
	return atu
}

// AddArchivedTaskActivities adds the "archivedTaskActivities" edges to the ArchivedTaskActivity entity.
func (atu *ActivityTypeUpdate) AddArchivedTaskActivities(a ...*ArchivedTaskActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.AddArchivedTaskActivityIDs(ids...)
}

// AddArchivedWorkspaceActivityIDs adds the "archivedWorkspaceActivities" edge to the ArchivedWorkspaceActivity entity by IDs.
func (atu *ActivityTypeUpdate) AddArchivedWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.AddArchivedWorkspaceActivityIDs(ids...)
	return atu
}

// AddArchivedWorkspaceActivities adds the "archivedWorkspaceActivities" edges to the ArchivedWorkspaceActivity entity.
func (atu *ActivityTypeUpdate) AddArchivedWorkspaceActivities(a ...*ArchivedWorkspaceActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.AddArchivedWorkspaceActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atu *ActivityTypeUpdate) Mutation() *ActivityTypeMutation {
	return atu.mutation
}

// ClearTaskActivities clears all "taskActivities" edges to the TaskActivity entity.
func (atu *ActivityTypeUpdate) ClearTaskActivities() *ActivityTypeUpdate {
	atu.mutation.ClearTaskActivities()
	return atu
}

// RemoveTaskActivityIDs removes the "taskActivities" edge to TaskActivity entities by IDs.
func (atu *ActivityTypeUpdate) RemoveTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.RemoveTaskActivityIDs(ids...)
	return atu
}

// RemoveTaskActivities removes "taskActivities" edges to TaskActivity entities.
func (atu *ActivityTypeUpdate) RemoveTaskActivities(t ...*TaskActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return atu.RemoveTaskActivityIDs(ids...)
}

// ClearWorkspaceActivities clears all "workspaceActivities" edges to the WorkspaceActivity entity.
func (atu *ActivityTypeUpdate) ClearWorkspaceActivities() *ActivityTypeUpdate {
	atu.mutation.ClearWorkspaceActivities()
	return atu
}

// RemoveWorkspaceActivityIDs removes the "workspaceActivities" edge to WorkspaceActivity entities by IDs.
func (atu *ActivityTypeUpdate) RemoveWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.RemoveWorkspaceActivityIDs(ids...)
	return atu
}

// RemoveWorkspaceActivities removes "workspaceActivities" edges to WorkspaceActivity entities.
func (atu *ActivityTypeUpdate) RemoveWorkspaceActivities(w ...*WorkspaceActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return atu.RemoveWorkspaceActivityIDs(ids...)
}

// ClearArchivedTaskActivities clears all "archivedTaskActivities" edges to the ArchivedTaskActivity entity.
func (atu *ActivityTypeUpdate) ClearArchivedTaskActivities() *ActivityTypeUpdate {
	atu.mutation.ClearArchivedTaskActivities()
	return atu
}

// RemoveArchivedTaskActivityIDs removes the "archivedTaskActivities" edge to ArchivedTaskActivity entities by IDs.
func (atu *ActivityTypeUpdate) RemoveArchivedTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.RemoveArchivedTaskActivityIDs(ids...)
	return atu
}

// RemoveArchivedTaskActivities removes "archivedTaskActivities" edges to ArchivedTaskActivity entities.
func (atu *ActivityTypeUpdate) RemoveArchivedTaskActivities(a ...*ArchivedTaskActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.RemoveArchivedTaskActivityIDs(ids...)
}

// ClearArchivedWorkspaceActivities clears all "archivedWorkspaceActivities" edges to the ArchivedWorkspaceActivity entity.
func (atu *ActivityTypeUpdate) ClearArchivedWorkspaceActivities() *ActivityTypeUpdate {
	atu.mutation.ClearArchivedWorkspaceActivities()
	return atu
}

// RemoveArchivedWorkspaceActivityIDs removes the "archivedWorkspaceActivities" edge to ArchivedWorkspaceActivity entities by IDs.
func (atu *ActivityTypeUpdate) RemoveArchivedWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdate {
	atu.mutation.RemoveArchivedWorkspaceActivityIDs(ids...)
	return atu
}

// RemoveArchivedWorkspaceActivities removes "archivedWorkspaceActivities" edges to ArchivedWorkspaceActivity entities.
func (atu *ActivityTypeUpdate) RemoveArchivedWorkspaceActivities(a ...*ArchivedWorkspaceActivity) *ActivityTypeUpdate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.RemoveArchivedWorkspaceActivityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ActivityTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ActivityTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ActivityTypeUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ActivityTypeUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ActivityTypeUpdate) check() error {
	if v, ok := atu.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ActivityType.name": %w`, err)}
		}
	}
	if v, ok := atu.mutation.TypeCode(); ok {
		if err := activitytype.TypeCodeValidator(v); err != nil {
			return &ValidationError{Name: "type_code", err: fmt.Errorf(`ent: validator failed for field "ActivityType.type_code": %w`, err)}
		}
	}
	return nil
}

func (atu *ActivityTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activitytype.Table,
			Columns: activitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activitytype.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
	}
	if value, ok := atu.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: activitytype.FieldTypeCode,
		})
	}
	if atu.mutation.TaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedTaskActivitiesIDs(); len(nodes) > 0 && !atu.mutation.TaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.TaskActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.WorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedWorkspaceActivitiesIDs(); len(nodes) > 0 && !atu.mutation.WorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.WorkspaceActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.ArchivedTaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedArchivedTaskActivitiesIDs(); len(nodes) > 0 && !atu.mutation.ArchivedTaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ArchivedTaskActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.ArchivedWorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedArchivedWorkspaceActivitiesIDs(); len(nodes) > 0 && !atu.mutation.ArchivedWorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ArchivedWorkspaceActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ActivityTypeUpdateOne is the builder for updating a single ActivityType entity.
type ActivityTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityTypeMutation
}

// SetName sets the "name" field.
func (atuo *ActivityTypeUpdateOne) SetName(s string) *ActivityTypeUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// SetTypeCode sets the "type_code" field.
func (atuo *ActivityTypeUpdateOne) SetTypeCode(ac activitytype.TypeCode) *ActivityTypeUpdateOne {
	atuo.mutation.SetTypeCode(ac)
	return atuo
}

// AddTaskActivityIDs adds the "taskActivities" edge to the TaskActivity entity by IDs.
func (atuo *ActivityTypeUpdateOne) AddTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.AddTaskActivityIDs(ids...)
	return atuo
}

// AddTaskActivities adds the "taskActivities" edges to the TaskActivity entity.
func (atuo *ActivityTypeUpdateOne) AddTaskActivities(t ...*TaskActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return atuo.AddTaskActivityIDs(ids...)
}

// AddWorkspaceActivityIDs adds the "workspaceActivities" edge to the WorkspaceActivity entity by IDs.
func (atuo *ActivityTypeUpdateOne) AddWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.AddWorkspaceActivityIDs(ids...)
	return atuo
}

// AddWorkspaceActivities adds the "workspaceActivities" edges to the WorkspaceActivity entity.
func (atuo *ActivityTypeUpdateOne) AddWorkspaceActivities(w ...*WorkspaceActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return atuo.AddWorkspaceActivityIDs(ids...)
}

// AddArchivedTaskActivityIDs adds the "archivedTaskActivities" edge to the ArchivedTaskActivity entity by IDs.
func (atuo *ActivityTypeUpdateOne) AddArchivedTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.AddArchivedTaskActivityIDs(ids...)
	return atuo
}

// AddArchivedTaskActivities adds the "archivedTaskActivities" edges to the ArchivedTaskActivity entity.
func (atuo *ActivityTypeUpdateOne) AddArchivedTaskActivities(a ...*ArchivedTaskActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.AddArchivedTaskActivityIDs(ids...)
}

// AddArchivedWorkspaceActivityIDs adds the "archivedWorkspaceActivities" edge to the ArchivedWorkspaceActivity entity by IDs.
func (atuo *ActivityTypeUpdateOne) AddArchivedWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.AddArchivedWorkspaceActivityIDs(ids...)
	return atuo
}

// AddArchivedWorkspaceActivities adds the "archivedWorkspaceActivities" edges to the ArchivedWorkspaceActivity entity.
func (atuo *ActivityTypeUpdateOne) AddArchivedWorkspaceActivities(a ...*ArchivedWorkspaceActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.AddArchivedWorkspaceActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atuo *ActivityTypeUpdateOne) Mutation() *ActivityTypeMutation {
	return atuo.mutation
}

// ClearTaskActivities clears all "taskActivities" edges to the TaskActivity entity.
func (atuo *ActivityTypeUpdateOne) ClearTaskActivities() *ActivityTypeUpdateOne {
	atuo.mutation.ClearTaskActivities()
	return atuo
}

// RemoveTaskActivityIDs removes the "taskActivities" edge to TaskActivity entities by IDs.
func (atuo *ActivityTypeUpdateOne) RemoveTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.RemoveTaskActivityIDs(ids...)
	return atuo
}

// RemoveTaskActivities removes "taskActivities" edges to TaskActivity entities.
func (atuo *ActivityTypeUpdateOne) RemoveTaskActivities(t ...*TaskActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return atuo.RemoveTaskActivityIDs(ids...)
}

// ClearWorkspaceActivities clears all "workspaceActivities" edges to the WorkspaceActivity entity.
func (atuo *ActivityTypeUpdateOne) ClearWorkspaceActivities() *ActivityTypeUpdateOne {
	atuo.mutation.ClearWorkspaceActivities()
	return atuo
}

// RemoveWorkspaceActivityIDs removes the "workspaceActivities" edge to WorkspaceActivity entities by IDs.
func (atuo *ActivityTypeUpdateOne) RemoveWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.RemoveWorkspaceActivityIDs(ids...)
	return atuo
}

// RemoveWorkspaceActivities removes "workspaceActivities" edges to WorkspaceActivity entities.
func (atuo *ActivityTypeUpdateOne) RemoveWorkspaceActivities(w ...*WorkspaceActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return atuo.RemoveWorkspaceActivityIDs(ids...)
}

// ClearArchivedTaskActivities clears all "archivedTaskActivities" edges to the ArchivedTaskActivity entity.
func (atuo *ActivityTypeUpdateOne) ClearArchivedTaskActivities() *ActivityTypeUpdateOne {
	atuo.mutation.ClearArchivedTaskActivities()
	return atuo
}

// RemoveArchivedTaskActivityIDs removes the "archivedTaskActivities" edge to ArchivedTaskActivity entities by IDs.
func (atuo *ActivityTypeUpdateOne) RemoveArchivedTaskActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.RemoveArchivedTaskActivityIDs(ids...)
	return atuo
}

// RemoveArchivedTaskActivities removes "archivedTaskActivities" edges to ArchivedTaskActivity entities.
func (atuo *ActivityTypeUpdateOne) RemoveArchivedTaskActivities(a ...*ArchivedTaskActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.RemoveArchivedTaskActivityIDs(ids...)
}

// ClearArchivedWorkspaceActivities clears all "archivedWorkspaceActivities" edges to the ArchivedWorkspaceActivity entity.
func (atuo *ActivityTypeUpdateOne) ClearArchivedWorkspaceActivities() *ActivityTypeUpdateOne {
	atuo.mutation.ClearArchivedWorkspaceActivities()
	return atuo
}

// RemoveArchivedWorkspaceActivityIDs removes the "archivedWorkspaceActivities" edge to ArchivedWorkspaceActivity entities by IDs.
func (atuo *ActivityTypeUpdateOne) RemoveArchivedWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeUpdateOne {
	atuo.mutation.RemoveArchivedWorkspaceActivityIDs(ids...)
	return atuo
}

// RemoveArchivedWorkspaceActivities removes "archivedWorkspaceActivities" edges to ArchivedWorkspaceActivity entities.
func (atuo *ActivityTypeUpdateOne) RemoveArchivedWorkspaceActivities(a ...*ArchivedWorkspaceActivity) *ActivityTypeUpdateOne {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.RemoveArchivedWorkspaceActivityIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ActivityTypeUpdateOne) Select(field string, fields ...string) *ActivityTypeUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ActivityType entity.
func (atuo *ActivityTypeUpdateOne) Save(ctx context.Context) (*ActivityType, error) {
	var (
		err  error
		node *ActivityType
	)
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ActivityTypeUpdateOne) SaveX(ctx context.Context) *ActivityType {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ActivityTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ActivityTypeUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ActivityTypeUpdateOne) check() error {
	if v, ok := atuo.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ActivityType.name": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.TypeCode(); ok {
		if err := activitytype.TypeCodeValidator(v); err != nil {
			return &ValidationError{Name: "type_code", err: fmt.Errorf(`ent: validator failed for field "ActivityType.type_code": %w`, err)}
		}
	}
	return nil
}

func (atuo *ActivityTypeUpdateOne) sqlSave(ctx context.Context) (_node *ActivityType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activitytype.Table,
			Columns: activitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activitytype.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ActivityType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activitytype.FieldID)
		for _, f := range fields {
			if !activitytype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activitytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
	}
	if value, ok := atuo.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: activitytype.FieldTypeCode,
		})
	}
	if atuo.mutation.TaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedTaskActivitiesIDs(); len(nodes) > 0 && !atuo.mutation.TaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.TaskActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.TaskActivitiesTable,
			Columns: []string{activitytype.TaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.WorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedWorkspaceActivitiesIDs(); len(nodes) > 0 && !atuo.mutation.WorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.WorkspaceActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.WorkspaceActivitiesTable,
			Columns: []string{activitytype.WorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.ArchivedTaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedArchivedTaskActivitiesIDs(); len(nodes) > 0 && !atuo.mutation.ArchivedTaskActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ArchivedTaskActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedTaskActivitiesTable,
			Columns: []string{activitytype.ArchivedTaskActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedtaskactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.ArchivedWorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedArchivedWorkspaceActivitiesIDs(); len(nodes) > 0 && !atuo.mutation.ArchivedWorkspaceActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ArchivedWorkspaceActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ArchivedWorkspaceActivitiesTable,
			Columns: []string{activitytype.ArchivedWorkspaceActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: archivedworkspaceactivity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ActivityType{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
