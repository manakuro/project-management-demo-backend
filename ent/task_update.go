// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskpriority"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTaskParentID sets the "task_parent_id" field.
func (tu *TaskUpdate) SetTaskParentID(u ulid.ID) *TaskUpdate {
	tu.mutation.SetTaskParentID(u)
	return tu
}

// SetNillableTaskParentID sets the "task_parent_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTaskParentID(u *ulid.ID) *TaskUpdate {
	if u != nil {
		tu.SetTaskParentID(*u)
	}
	return tu
}

// ClearTaskParentID clears the value of the "task_parent_id" field.
func (tu *TaskUpdate) ClearTaskParentID() *TaskUpdate {
	tu.mutation.ClearTaskParentID()
	return tu
}

// SetTaskPriorityID sets the "task_priority_id" field.
func (tu *TaskUpdate) SetTaskPriorityID(u ulid.ID) *TaskUpdate {
	tu.mutation.SetTaskPriorityID(u)
	return tu
}

// SetAssigneeID sets the "assignee_id" field.
func (tu *TaskUpdate) SetAssigneeID(u ulid.ID) *TaskUpdate {
	tu.mutation.SetAssigneeID(u)
	return tu
}

// SetNillableAssigneeID sets the "assignee_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableAssigneeID(u *ulid.ID) *TaskUpdate {
	if u != nil {
		tu.SetAssigneeID(*u)
	}
	return tu
}

// ClearAssigneeID clears the value of the "assignee_id" field.
func (tu *TaskUpdate) ClearAssigneeID() *TaskUpdate {
	tu.mutation.ClearAssigneeID()
	return tu
}

// SetCreatedBy sets the "created_by" field.
func (tu *TaskUpdate) SetCreatedBy(u ulid.ID) *TaskUpdate {
	tu.mutation.SetCreatedBy(u)
	return tu
}

// SetCompleted sets the "completed" field.
func (tu *TaskUpdate) SetCompleted(b bool) *TaskUpdate {
	tu.mutation.SetCompleted(b)
	return tu
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompleted(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetCompleted(*b)
	}
	return tu
}

// SetCompletedAt sets the "completed_at" field.
func (tu *TaskUpdate) SetCompletedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCompletedAt(t)
	return tu
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompletedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCompletedAt(*t)
	}
	return tu
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tu *TaskUpdate) ClearCompletedAt() *TaskUpdate {
	tu.mutation.ClearCompletedAt()
	return tu
}

// SetIsNew sets the "is_new" field.
func (tu *TaskUpdate) SetIsNew(b bool) *TaskUpdate {
	tu.mutation.SetIsNew(b)
	return tu
}

// SetNillableIsNew sets the "is_new" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableIsNew(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetIsNew(*b)
	}
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDueDate sets the "due_date" field.
func (tu *TaskUpdate) SetDueDate(t time.Time) *TaskUpdate {
	tu.mutation.SetDueDate(t)
	return tu
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDueDate(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDueDate(*t)
	}
	return tu
}

// ClearDueDate clears the value of the "due_date" field.
func (tu *TaskUpdate) ClearDueDate() *TaskUpdate {
	tu.mutation.ClearDueDate()
	return tu
}

// SetDueTime sets the "due_time" field.
func (tu *TaskUpdate) SetDueTime(t time.Time) *TaskUpdate {
	tu.mutation.SetDueTime(t)
	return tu
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDueTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDueTime(*t)
	}
	return tu
}

// ClearDueTime clears the value of the "due_time" field.
func (tu *TaskUpdate) ClearDueTime() *TaskUpdate {
	tu.mutation.ClearDueTime()
	return tu
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (tu *TaskUpdate) SetTeammateID(id ulid.ID) *TaskUpdate {
	tu.mutation.SetTeammateID(id)
	return tu
}

// SetNillableTeammateID sets the "teammate" edge to the Teammate entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableTeammateID(id *ulid.ID) *TaskUpdate {
	if id != nil {
		tu = tu.SetTeammateID(*id)
	}
	return tu
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tu *TaskUpdate) SetTeammate(t *Teammate) *TaskUpdate {
	return tu.SetTeammateID(t.ID)
}

// SetTaskPriority sets the "task_priority" edge to the TaskPriority entity.
func (tu *TaskUpdate) SetTaskPriority(t *TaskPriority) *TaskUpdate {
	return tu.SetTaskPriorityID(t.ID)
}

// SetParentID sets the "parent" edge to the Task entity by ID.
func (tu *TaskUpdate) SetParentID(id ulid.ID) *TaskUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Task entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableParentID(id *ulid.ID) *TaskUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Task entity.
func (tu *TaskUpdate) SetParent(t *Task) *TaskUpdate {
	return tu.SetParentID(t.ID)
}

// AddSubTaskIDs adds the "sub_tasks" edge to the Task entity by IDs.
func (tu *TaskUpdate) AddSubTaskIDs(ids ...ulid.ID) *TaskUpdate {
	tu.mutation.AddSubTaskIDs(ids...)
	return tu
}

// AddSubTasks adds the "sub_tasks" edges to the Task entity.
func (tu *TaskUpdate) AddSubTasks(t ...*Task) *TaskUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddSubTaskIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (tu *TaskUpdate) ClearTeammate() *TaskUpdate {
	tu.mutation.ClearTeammate()
	return tu
}

// ClearTaskPriority clears the "task_priority" edge to the TaskPriority entity.
func (tu *TaskUpdate) ClearTaskPriority() *TaskUpdate {
	tu.mutation.ClearTaskPriority()
	return tu
}

// ClearParent clears the "parent" edge to the Task entity.
func (tu *TaskUpdate) ClearParent() *TaskUpdate {
	tu.mutation.ClearParent()
	return tu
}

// ClearSubTasks clears all "sub_tasks" edges to the Task entity.
func (tu *TaskUpdate) ClearSubTasks() *TaskUpdate {
	tu.mutation.ClearSubTasks()
	return tu
}

// RemoveSubTaskIDs removes the "sub_tasks" edge to Task entities by IDs.
func (tu *TaskUpdate) RemoveSubTaskIDs(ids ...ulid.ID) *TaskUpdate {
	tu.mutation.RemoveSubTaskIDs(ids...)
	return tu
}

// RemoveSubTasks removes "sub_tasks" edges to Task entities.
func (tu *TaskUpdate) RemoveSubTasks(t ...*Task) *TaskUpdate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveSubTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := tu.mutation.TaskPriorityID(); tu.mutation.TaskPriorityCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task_priority\"")
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldCreatedBy,
		})
	}
	if value, ok := tu.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldCompleted,
		})
	}
	if value, ok := tu.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCompletedAt,
		})
	}
	if tu.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCompletedAt,
		})
	}
	if value, ok := tu.mutation.IsNew(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldIsNew,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldName,
		})
	}
	if value, ok := tu.mutation.DueDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueDate,
		})
	}
	if tu.mutation.DueDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDueDate,
		})
	}
	if value, ok := tu.mutation.DueTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueTime,
		})
	}
	if tu.mutation.DueTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDueTime,
		})
	}
	if tu.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TeammateTable,
			Columns: []string{task.TeammateColumn},
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
	if nodes := tu.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TeammateTable,
			Columns: []string{task.TeammateColumn},
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
	if tu.mutation.TaskPriorityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TaskPriorityTable,
			Columns: []string{task.TaskPriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskpriority.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TaskPriorityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TaskPriorityTable,
			Columns: []string{task.TaskPriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskpriority.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ParentTable,
			Columns: []string{task.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ParentTable,
			Columns: []string{task.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SubTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSubTasksIDs(); len(nodes) > 0 && !tu.mutation.SubTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SubTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetTaskParentID sets the "task_parent_id" field.
func (tuo *TaskUpdateOne) SetTaskParentID(u ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetTaskParentID(u)
	return tuo
}

// SetNillableTaskParentID sets the "task_parent_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTaskParentID(u *ulid.ID) *TaskUpdateOne {
	if u != nil {
		tuo.SetTaskParentID(*u)
	}
	return tuo
}

// ClearTaskParentID clears the value of the "task_parent_id" field.
func (tuo *TaskUpdateOne) ClearTaskParentID() *TaskUpdateOne {
	tuo.mutation.ClearTaskParentID()
	return tuo
}

// SetTaskPriorityID sets the "task_priority_id" field.
func (tuo *TaskUpdateOne) SetTaskPriorityID(u ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetTaskPriorityID(u)
	return tuo
}

// SetAssigneeID sets the "assignee_id" field.
func (tuo *TaskUpdateOne) SetAssigneeID(u ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetAssigneeID(u)
	return tuo
}

// SetNillableAssigneeID sets the "assignee_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableAssigneeID(u *ulid.ID) *TaskUpdateOne {
	if u != nil {
		tuo.SetAssigneeID(*u)
	}
	return tuo
}

// ClearAssigneeID clears the value of the "assignee_id" field.
func (tuo *TaskUpdateOne) ClearAssigneeID() *TaskUpdateOne {
	tuo.mutation.ClearAssigneeID()
	return tuo
}

// SetCreatedBy sets the "created_by" field.
func (tuo *TaskUpdateOne) SetCreatedBy(u ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetCreatedBy(u)
	return tuo
}

// SetCompleted sets the "completed" field.
func (tuo *TaskUpdateOne) SetCompleted(b bool) *TaskUpdateOne {
	tuo.mutation.SetCompleted(b)
	return tuo
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompleted(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetCompleted(*b)
	}
	return tuo
}

// SetCompletedAt sets the "completed_at" field.
func (tuo *TaskUpdateOne) SetCompletedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCompletedAt(t)
	return tuo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompletedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCompletedAt(*t)
	}
	return tuo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tuo *TaskUpdateOne) ClearCompletedAt() *TaskUpdateOne {
	tuo.mutation.ClearCompletedAt()
	return tuo
}

// SetIsNew sets the "is_new" field.
func (tuo *TaskUpdateOne) SetIsNew(b bool) *TaskUpdateOne {
	tuo.mutation.SetIsNew(b)
	return tuo
}

// SetNillableIsNew sets the "is_new" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableIsNew(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetIsNew(*b)
	}
	return tuo
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDueDate sets the "due_date" field.
func (tuo *TaskUpdateOne) SetDueDate(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDueDate(t)
	return tuo
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDueDate(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDueDate(*t)
	}
	return tuo
}

// ClearDueDate clears the value of the "due_date" field.
func (tuo *TaskUpdateOne) ClearDueDate() *TaskUpdateOne {
	tuo.mutation.ClearDueDate()
	return tuo
}

// SetDueTime sets the "due_time" field.
func (tuo *TaskUpdateOne) SetDueTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDueTime(t)
	return tuo
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDueTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDueTime(*t)
	}
	return tuo
}

// ClearDueTime clears the value of the "due_time" field.
func (tuo *TaskUpdateOne) ClearDueTime() *TaskUpdateOne {
	tuo.mutation.ClearDueTime()
	return tuo
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (tuo *TaskUpdateOne) SetTeammateID(id ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetTeammateID(id)
	return tuo
}

// SetNillableTeammateID sets the "teammate" edge to the Teammate entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTeammateID(id *ulid.ID) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetTeammateID(*id)
	}
	return tuo
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tuo *TaskUpdateOne) SetTeammate(t *Teammate) *TaskUpdateOne {
	return tuo.SetTeammateID(t.ID)
}

// SetTaskPriority sets the "task_priority" edge to the TaskPriority entity.
func (tuo *TaskUpdateOne) SetTaskPriority(t *TaskPriority) *TaskUpdateOne {
	return tuo.SetTaskPriorityID(t.ID)
}

// SetParentID sets the "parent" edge to the Task entity by ID.
func (tuo *TaskUpdateOne) SetParentID(id ulid.ID) *TaskUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Task entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableParentID(id *ulid.ID) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Task entity.
func (tuo *TaskUpdateOne) SetParent(t *Task) *TaskUpdateOne {
	return tuo.SetParentID(t.ID)
}

// AddSubTaskIDs adds the "sub_tasks" edge to the Task entity by IDs.
func (tuo *TaskUpdateOne) AddSubTaskIDs(ids ...ulid.ID) *TaskUpdateOne {
	tuo.mutation.AddSubTaskIDs(ids...)
	return tuo
}

// AddSubTasks adds the "sub_tasks" edges to the Task entity.
func (tuo *TaskUpdateOne) AddSubTasks(t ...*Task) *TaskUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddSubTaskIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearTeammate clears the "teammate" edge to the Teammate entity.
func (tuo *TaskUpdateOne) ClearTeammate() *TaskUpdateOne {
	tuo.mutation.ClearTeammate()
	return tuo
}

// ClearTaskPriority clears the "task_priority" edge to the TaskPriority entity.
func (tuo *TaskUpdateOne) ClearTaskPriority() *TaskUpdateOne {
	tuo.mutation.ClearTaskPriority()
	return tuo
}

// ClearParent clears the "parent" edge to the Task entity.
func (tuo *TaskUpdateOne) ClearParent() *TaskUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// ClearSubTasks clears all "sub_tasks" edges to the Task entity.
func (tuo *TaskUpdateOne) ClearSubTasks() *TaskUpdateOne {
	tuo.mutation.ClearSubTasks()
	return tuo
}

// RemoveSubTaskIDs removes the "sub_tasks" edge to Task entities by IDs.
func (tuo *TaskUpdateOne) RemoveSubTaskIDs(ids ...ulid.ID) *TaskUpdateOne {
	tuo.mutation.RemoveSubTaskIDs(ids...)
	return tuo
}

// RemoveSubTasks removes "sub_tasks" edges to Task entities.
func (tuo *TaskUpdateOne) RemoveSubTasks(t ...*Task) *TaskUpdateOne {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveSubTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := tuo.mutation.TaskPriorityID(); tuo.mutation.TaskPriorityCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"task_priority\"")
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldCreatedBy,
		})
	}
	if value, ok := tuo.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldCompleted,
		})
	}
	if value, ok := tuo.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCompletedAt,
		})
	}
	if tuo.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldCompletedAt,
		})
	}
	if value, ok := tuo.mutation.IsNew(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldIsNew,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldName,
		})
	}
	if value, ok := tuo.mutation.DueDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueDate,
		})
	}
	if tuo.mutation.DueDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDueDate,
		})
	}
	if value, ok := tuo.mutation.DueTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueTime,
		})
	}
	if tuo.mutation.DueTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDueTime,
		})
	}
	if tuo.mutation.TeammateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TeammateTable,
			Columns: []string{task.TeammateColumn},
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
	if nodes := tuo.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TeammateTable,
			Columns: []string{task.TeammateColumn},
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
	if tuo.mutation.TaskPriorityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TaskPriorityTable,
			Columns: []string{task.TaskPriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskpriority.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TaskPriorityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TaskPriorityTable,
			Columns: []string{task.TaskPriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskpriority.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ParentTable,
			Columns: []string{task.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ParentTable,
			Columns: []string{task.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SubTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSubTasksIDs(); len(nodes) > 0 && !tuo.mutation.SubTasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SubTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.SubTasksTable,
			Columns: []string{task.SubTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
