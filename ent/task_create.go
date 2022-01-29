// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/ent/taskpriority"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetTaskParentID sets the "task_parent_id" field.
func (tc *TaskCreate) SetTaskParentID(u ulid.ID) *TaskCreate {
	tc.mutation.SetTaskParentID(u)
	return tc
}

// SetNillableTaskParentID sets the "task_parent_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableTaskParentID(u *ulid.ID) *TaskCreate {
	if u != nil {
		tc.SetTaskParentID(*u)
	}
	return tc
}

// SetTaskPriorityID sets the "task_priority_id" field.
func (tc *TaskCreate) SetTaskPriorityID(u ulid.ID) *TaskCreate {
	tc.mutation.SetTaskPriorityID(u)
	return tc
}

// SetAssigneeID sets the "assignee_id" field.
func (tc *TaskCreate) SetAssigneeID(u ulid.ID) *TaskCreate {
	tc.mutation.SetAssigneeID(u)
	return tc
}

// SetNillableAssigneeID sets the "assignee_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableAssigneeID(u *ulid.ID) *TaskCreate {
	if u != nil {
		tc.SetAssigneeID(*u)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TaskCreate) SetCreatedBy(u ulid.ID) *TaskCreate {
	tc.mutation.SetCreatedBy(u)
	return tc
}

// SetCompleted sets the "completed" field.
func (tc *TaskCreate) SetCompleted(b bool) *TaskCreate {
	tc.mutation.SetCompleted(b)
	return tc
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCompleted(b *bool) *TaskCreate {
	if b != nil {
		tc.SetCompleted(*b)
	}
	return tc
}

// SetCompletedAt sets the "completed_at" field.
func (tc *TaskCreate) SetCompletedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCompletedAt(t)
	return tc
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCompletedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCompletedAt(*t)
	}
	return tc
}

// SetIsNew sets the "is_new" field.
func (tc *TaskCreate) SetIsNew(b bool) *TaskCreate {
	tc.mutation.SetIsNew(b)
	return tc
}

// SetNillableIsNew sets the "is_new" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsNew(b *bool) *TaskCreate {
	if b != nil {
		tc.SetIsNew(*b)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDueDate sets the "due_date" field.
func (tc *TaskCreate) SetDueDate(t time.Time) *TaskCreate {
	tc.mutation.SetDueDate(t)
	return tc
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDueDate(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetDueDate(*t)
	}
	return tc
}

// SetDueTime sets the "due_time" field.
func (tc *TaskCreate) SetDueTime(t time.Time) *TaskCreate {
	tc.mutation.SetDueTime(t)
	return tc
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDueTime(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetDueTime(*t)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TaskCreate) SetUpdatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUpdatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(u ulid.ID) *TaskCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableID(u *ulid.ID) *TaskCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (tc *TaskCreate) SetTeammateID(id ulid.ID) *TaskCreate {
	tc.mutation.SetTeammateID(id)
	return tc
}

// SetNillableTeammateID sets the "teammate" edge to the Teammate entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableTeammateID(id *ulid.ID) *TaskCreate {
	if id != nil {
		tc = tc.SetTeammateID(*id)
	}
	return tc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tc *TaskCreate) SetTeammate(t *Teammate) *TaskCreate {
	return tc.SetTeammateID(t.ID)
}

// SetTaskPriority sets the "task_priority" edge to the TaskPriority entity.
func (tc *TaskCreate) SetTaskPriority(t *TaskPriority) *TaskCreate {
	return tc.SetTaskPriorityID(t.ID)
}

// SetParentID sets the "parent" edge to the Task entity by ID.
func (tc *TaskCreate) SetParentID(id ulid.ID) *TaskCreate {
	tc.mutation.SetParentID(id)
	return tc
}

// SetNillableParentID sets the "parent" edge to the Task entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableParentID(id *ulid.ID) *TaskCreate {
	if id != nil {
		tc = tc.SetParentID(*id)
	}
	return tc
}

// SetParent sets the "parent" edge to the Task entity.
func (tc *TaskCreate) SetParent(t *Task) *TaskCreate {
	return tc.SetParentID(t.ID)
}

// AddSubTaskIDs adds the "sub_tasks" edge to the Task entity by IDs.
func (tc *TaskCreate) AddSubTaskIDs(ids ...ulid.ID) *TaskCreate {
	tc.mutation.AddSubTaskIDs(ids...)
	return tc
}

// AddSubTasks adds the "sub_tasks" edges to the Task entity.
func (tc *TaskCreate) AddSubTasks(t ...*Task) *TaskCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddSubTaskIDs(ids...)
}

// AddTeammateTaskIDs adds the "teammate_tasks" edge to the TeammateTask entity by IDs.
func (tc *TaskCreate) AddTeammateTaskIDs(ids ...ulid.ID) *TaskCreate {
	tc.mutation.AddTeammateTaskIDs(ids...)
	return tc
}

// AddTeammateTasks adds the "teammate_tasks" edges to the TeammateTask entity.
func (tc *TaskCreate) AddTeammateTasks(t ...*TeammateTask) *TaskCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTeammateTaskIDs(ids...)
}

// AddProjectTaskIDs adds the "project_tasks" edge to the ProjectTask entity by IDs.
func (tc *TaskCreate) AddProjectTaskIDs(ids ...ulid.ID) *TaskCreate {
	tc.mutation.AddProjectTaskIDs(ids...)
	return tc
}

// AddProjectTasks adds the "project_tasks" edges to the ProjectTask entity.
func (tc *TaskCreate) AddProjectTasks(p ...*ProjectTask) *TaskCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddProjectTaskIDs(ids...)
}

// AddTaskLikeIDs adds the "task_likes" edge to the TaskLike entity by IDs.
func (tc *TaskCreate) AddTaskLikeIDs(ids ...ulid.ID) *TaskCreate {
	tc.mutation.AddTaskLikeIDs(ids...)
	return tc
}

// AddTaskLikes adds the "task_likes" edges to the TaskLike entity.
func (tc *TaskCreate) AddTaskLikes(t ...*TaskLike) *TaskCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTaskLikeIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.Completed(); !ok {
		v := task.DefaultCompleted
		tc.mutation.SetCompleted(v)
	}
	if _, ok := tc.mutation.IsNew(); !ok {
		v := task.DefaultIsNew
		tc.mutation.SetIsNew(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := task.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := task.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.TaskPriorityID(); !ok {
		return &ValidationError{Name: "task_priority_id", err: errors.New(`ent: missing required field "task_priority_id"`)}
	}
	if _, ok := tc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "created_by"`)}
	}
	if _, ok := tc.mutation.Completed(); !ok {
		return &ValidationError{Name: "completed", err: errors.New(`ent: missing required field "completed"`)}
	}
	if _, ok := tc.mutation.IsNew(); !ok {
		return &ValidationError{Name: "is_new", err: errors.New(`ent: missing required field "is_new"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tc.mutation.TaskPriorityID(); !ok {
		return &ValidationError{Name: "task_priority", err: errors.New("ent: missing required edge \"task_priority\"")}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: task.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.Completed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldCompleted,
		})
		_node.Completed = value
	}
	if value, ok := tc.mutation.CompletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCompletedAt,
		})
		_node.CompletedAt = &value
	}
	if value, ok := tc.mutation.IsNew(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldIsNew,
		})
		_node.IsNew = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.DueDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueDate,
		})
		_node.DueDate = &value
	}
	if value, ok := tc.mutation.DueTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDueTime,
		})
		_node.DueTime = &value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.TeammateIDs(); len(nodes) > 0 {
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
		_node.AssigneeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TaskPriorityIDs(); len(nodes) > 0 {
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
		_node.TaskPriorityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.TaskParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SubTasksIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TeammateTasksTable,
			Columns: []string{task.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ProjectTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ProjectTasksTable,
			Columns: []string{task.ProjectTasksColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TaskLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TaskLikesTable,
			Columns: []string{task.TaskLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tasklike.FieldID,
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

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
