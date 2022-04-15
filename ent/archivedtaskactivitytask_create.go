// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/archivedtaskactivitytask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskactivity"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedTaskActivityTaskCreate is the builder for creating a ArchivedTaskActivityTask entity.
type ArchivedTaskActivityTaskCreate struct {
	config
	mutation *ArchivedTaskActivityTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTaskActivityID sets the "task_activity_id" field.
func (atatc *ArchivedTaskActivityTaskCreate) SetTaskActivityID(u ulid.ID) *ArchivedTaskActivityTaskCreate {
	atatc.mutation.SetTaskActivityID(u)
	return atatc
}

// SetTaskID sets the "task_id" field.
func (atatc *ArchivedTaskActivityTaskCreate) SetTaskID(u ulid.ID) *ArchivedTaskActivityTaskCreate {
	atatc.mutation.SetTaskID(u)
	return atatc
}

// SetCreatedAt sets the "created_at" field.
func (atatc *ArchivedTaskActivityTaskCreate) SetCreatedAt(t time.Time) *ArchivedTaskActivityTaskCreate {
	atatc.mutation.SetCreatedAt(t)
	return atatc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atatc *ArchivedTaskActivityTaskCreate) SetNillableCreatedAt(t *time.Time) *ArchivedTaskActivityTaskCreate {
	if t != nil {
		atatc.SetCreatedAt(*t)
	}
	return atatc
}

// SetUpdatedAt sets the "updated_at" field.
func (atatc *ArchivedTaskActivityTaskCreate) SetUpdatedAt(t time.Time) *ArchivedTaskActivityTaskCreate {
	atatc.mutation.SetUpdatedAt(t)
	return atatc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atatc *ArchivedTaskActivityTaskCreate) SetNillableUpdatedAt(t *time.Time) *ArchivedTaskActivityTaskCreate {
	if t != nil {
		atatc.SetUpdatedAt(*t)
	}
	return atatc
}

// SetID sets the "id" field.
func (atatc *ArchivedTaskActivityTaskCreate) SetID(u ulid.ID) *ArchivedTaskActivityTaskCreate {
	atatc.mutation.SetID(u)
	return atatc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atatc *ArchivedTaskActivityTaskCreate) SetNillableID(u *ulid.ID) *ArchivedTaskActivityTaskCreate {
	if u != nil {
		atatc.SetID(*u)
	}
	return atatc
}

// SetTask sets the "task" edge to the Task entity.
func (atatc *ArchivedTaskActivityTaskCreate) SetTask(t *Task) *ArchivedTaskActivityTaskCreate {
	return atatc.SetTaskID(t.ID)
}

// SetTaskActivity sets the "taskActivity" edge to the TaskActivity entity.
func (atatc *ArchivedTaskActivityTaskCreate) SetTaskActivity(t *TaskActivity) *ArchivedTaskActivityTaskCreate {
	return atatc.SetTaskActivityID(t.ID)
}

// Mutation returns the ArchivedTaskActivityTaskMutation object of the builder.
func (atatc *ArchivedTaskActivityTaskCreate) Mutation() *ArchivedTaskActivityTaskMutation {
	return atatc.mutation
}

// Save creates the ArchivedTaskActivityTask in the database.
func (atatc *ArchivedTaskActivityTaskCreate) Save(ctx context.Context) (*ArchivedTaskActivityTask, error) {
	var (
		err  error
		node *ArchivedTaskActivityTask
	)
	atatc.defaults()
	if len(atatc.hooks) == 0 {
		if err = atatc.check(); err != nil {
			return nil, err
		}
		node, err = atatc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedTaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atatc.check(); err != nil {
				return nil, err
			}
			atatc.mutation = mutation
			if node, err = atatc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(atatc.hooks) - 1; i >= 0; i-- {
			if atatc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atatc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atatc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atatc *ArchivedTaskActivityTaskCreate) SaveX(ctx context.Context) *ArchivedTaskActivityTask {
	v, err := atatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atatc *ArchivedTaskActivityTaskCreate) Exec(ctx context.Context) error {
	_, err := atatc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atatc *ArchivedTaskActivityTaskCreate) ExecX(ctx context.Context) {
	if err := atatc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atatc *ArchivedTaskActivityTaskCreate) defaults() {
	if _, ok := atatc.mutation.CreatedAt(); !ok {
		v := archivedtaskactivitytask.DefaultCreatedAt()
		atatc.mutation.SetCreatedAt(v)
	}
	if _, ok := atatc.mutation.UpdatedAt(); !ok {
		v := archivedtaskactivitytask.DefaultUpdatedAt()
		atatc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atatc.mutation.ID(); !ok {
		v := archivedtaskactivitytask.DefaultID()
		atatc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atatc *ArchivedTaskActivityTaskCreate) check() error {
	if _, ok := atatc.mutation.TaskActivityID(); !ok {
		return &ValidationError{Name: "task_activity_id", err: errors.New(`ent: missing required field "ArchivedTaskActivityTask.task_activity_id"`)}
	}
	if _, ok := atatc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "ArchivedTaskActivityTask.task_id"`)}
	}
	if _, ok := atatc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ArchivedTaskActivityTask.created_at"`)}
	}
	if _, ok := atatc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ArchivedTaskActivityTask.updated_at"`)}
	}
	if _, ok := atatc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "ArchivedTaskActivityTask.task"`)}
	}
	if _, ok := atatc.mutation.TaskActivityID(); !ok {
		return &ValidationError{Name: "taskActivity", err: errors.New(`ent: missing required edge "ArchivedTaskActivityTask.taskActivity"`)}
	}
	return nil
}

func (atatc *ArchivedTaskActivityTaskCreate) sqlSave(ctx context.Context) (*ArchivedTaskActivityTask, error) {
	_node, _spec := atatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atatc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*ulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (atatc *ArchivedTaskActivityTaskCreate) createSpec() (*ArchivedTaskActivityTask, *sqlgraph.CreateSpec) {
	var (
		_node = &ArchivedTaskActivityTask{config: atatc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: archivedtaskactivitytask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedtaskactivitytask.FieldID,
			},
		}
	)
	_spec.OnConflict = atatc.conflict
	if id, ok := atatc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atatc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedtaskactivitytask.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := atatc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedtaskactivitytask.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := atatc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskTable,
			Columns: []string{archivedtaskactivitytask.TaskColumn},
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
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atatc.mutation.TaskActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedtaskactivitytask.TaskActivityTable,
			Columns: []string{archivedtaskactivitytask.TaskActivityColumn},
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
		_node.TaskActivityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArchivedTaskActivityTask.Create().
//		SetTaskActivityID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedTaskActivityTaskUpsert) {
//			SetTaskActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (atatc *ArchivedTaskActivityTaskCreate) OnConflict(opts ...sql.ConflictOption) *ArchivedTaskActivityTaskUpsertOne {
	atatc.conflict = opts
	return &ArchivedTaskActivityTaskUpsertOne{
		create: atatc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedTaskActivityTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (atatc *ArchivedTaskActivityTaskCreate) OnConflictColumns(columns ...string) *ArchivedTaskActivityTaskUpsertOne {
	atatc.conflict = append(atatc.conflict, sql.ConflictColumns(columns...))
	return &ArchivedTaskActivityTaskUpsertOne{
		create: atatc,
	}
}

type (
	// ArchivedTaskActivityTaskUpsertOne is the builder for "upsert"-ing
	//  one ArchivedTaskActivityTask node.
	ArchivedTaskActivityTaskUpsertOne struct {
		create *ArchivedTaskActivityTaskCreate
	}

	// ArchivedTaskActivityTaskUpsert is the "OnConflict" setter.
	ArchivedTaskActivityTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetTaskActivityID sets the "task_activity_id" field.
func (u *ArchivedTaskActivityTaskUpsert) SetTaskActivityID(v ulid.ID) *ArchivedTaskActivityTaskUpsert {
	u.Set(archivedtaskactivitytask.FieldTaskActivityID, v)
	return u
}

// UpdateTaskActivityID sets the "task_activity_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsert) UpdateTaskActivityID() *ArchivedTaskActivityTaskUpsert {
	u.SetExcluded(archivedtaskactivitytask.FieldTaskActivityID)
	return u
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedTaskActivityTaskUpsert) SetTaskID(v ulid.ID) *ArchivedTaskActivityTaskUpsert {
	u.Set(archivedtaskactivitytask.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsert) UpdateTaskID() *ArchivedTaskActivityTaskUpsert {
	u.SetExcluded(archivedtaskactivitytask.FieldTaskID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedTaskActivityTaskUpsert) SetCreatedAt(v time.Time) *ArchivedTaskActivityTaskUpsert {
	u.Set(archivedtaskactivitytask.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsert) UpdateCreatedAt() *ArchivedTaskActivityTaskUpsert {
	u.SetExcluded(archivedtaskactivitytask.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedTaskActivityTaskUpsert) SetUpdatedAt(v time.Time) *ArchivedTaskActivityTaskUpsert {
	u.Set(archivedtaskactivitytask.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsert) UpdateUpdatedAt() *ArchivedTaskActivityTaskUpsert {
	u.SetExcluded(archivedtaskactivitytask.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ArchivedTaskActivityTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedtaskactivitytask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedTaskActivityTaskUpsertOne) UpdateNewValues() *ArchivedTaskActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(archivedtaskactivitytask.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(archivedtaskactivitytask.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(archivedtaskactivitytask.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ArchivedTaskActivityTask.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ArchivedTaskActivityTaskUpsertOne) Ignore() *ArchivedTaskActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedTaskActivityTaskUpsertOne) DoNothing() *ArchivedTaskActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedTaskActivityTaskCreate.OnConflict
// documentation for more info.
func (u *ArchivedTaskActivityTaskUpsertOne) Update(set func(*ArchivedTaskActivityTaskUpsert)) *ArchivedTaskActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedTaskActivityTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskActivityID sets the "task_activity_id" field.
func (u *ArchivedTaskActivityTaskUpsertOne) SetTaskActivityID(v ulid.ID) *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetTaskActivityID(v)
	})
}

// UpdateTaskActivityID sets the "task_activity_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertOne) UpdateTaskActivityID() *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateTaskActivityID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedTaskActivityTaskUpsertOne) SetTaskID(v ulid.ID) *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertOne) UpdateTaskID() *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedTaskActivityTaskUpsertOne) SetCreatedAt(v time.Time) *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertOne) UpdateCreatedAt() *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedTaskActivityTaskUpsertOne) SetUpdatedAt(v time.Time) *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertOne) UpdateUpdatedAt() *ArchivedTaskActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedTaskActivityTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedTaskActivityTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedTaskActivityTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArchivedTaskActivityTaskUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArchivedTaskActivityTaskUpsertOne.ID is not supported by MySQL driver. Use ArchivedTaskActivityTaskUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArchivedTaskActivityTaskUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArchivedTaskActivityTaskCreateBulk is the builder for creating many ArchivedTaskActivityTask entities in bulk.
type ArchivedTaskActivityTaskCreateBulk struct {
	config
	builders []*ArchivedTaskActivityTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the ArchivedTaskActivityTask entities in the database.
func (atatcb *ArchivedTaskActivityTaskCreateBulk) Save(ctx context.Context) ([]*ArchivedTaskActivityTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atatcb.builders))
	nodes := make([]*ArchivedTaskActivityTask, len(atatcb.builders))
	mutators := make([]Mutator, len(atatcb.builders))
	for i := range atatcb.builders {
		func(i int, root context.Context) {
			builder := atatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArchivedTaskActivityTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, atatcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atatcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atatcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atatcb *ArchivedTaskActivityTaskCreateBulk) SaveX(ctx context.Context) []*ArchivedTaskActivityTask {
	v, err := atatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atatcb *ArchivedTaskActivityTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := atatcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atatcb *ArchivedTaskActivityTaskCreateBulk) ExecX(ctx context.Context) {
	if err := atatcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArchivedTaskActivityTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedTaskActivityTaskUpsert) {
//			SetTaskActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (atatcb *ArchivedTaskActivityTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArchivedTaskActivityTaskUpsertBulk {
	atatcb.conflict = opts
	return &ArchivedTaskActivityTaskUpsertBulk{
		create: atatcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedTaskActivityTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (atatcb *ArchivedTaskActivityTaskCreateBulk) OnConflictColumns(columns ...string) *ArchivedTaskActivityTaskUpsertBulk {
	atatcb.conflict = append(atatcb.conflict, sql.ConflictColumns(columns...))
	return &ArchivedTaskActivityTaskUpsertBulk{
		create: atatcb,
	}
}

// ArchivedTaskActivityTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of ArchivedTaskActivityTask nodes.
type ArchivedTaskActivityTaskUpsertBulk struct {
	create *ArchivedTaskActivityTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ArchivedTaskActivityTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedtaskactivitytask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedTaskActivityTaskUpsertBulk) UpdateNewValues() *ArchivedTaskActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(archivedtaskactivitytask.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(archivedtaskactivitytask.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(archivedtaskactivitytask.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArchivedTaskActivityTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ArchivedTaskActivityTaskUpsertBulk) Ignore() *ArchivedTaskActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedTaskActivityTaskUpsertBulk) DoNothing() *ArchivedTaskActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedTaskActivityTaskCreateBulk.OnConflict
// documentation for more info.
func (u *ArchivedTaskActivityTaskUpsertBulk) Update(set func(*ArchivedTaskActivityTaskUpsert)) *ArchivedTaskActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedTaskActivityTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskActivityID sets the "task_activity_id" field.
func (u *ArchivedTaskActivityTaskUpsertBulk) SetTaskActivityID(v ulid.ID) *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetTaskActivityID(v)
	})
}

// UpdateTaskActivityID sets the "task_activity_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertBulk) UpdateTaskActivityID() *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateTaskActivityID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedTaskActivityTaskUpsertBulk) SetTaskID(v ulid.ID) *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertBulk) UpdateTaskID() *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedTaskActivityTaskUpsertBulk) SetCreatedAt(v time.Time) *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertBulk) UpdateCreatedAt() *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedTaskActivityTaskUpsertBulk) SetUpdatedAt(v time.Time) *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedTaskActivityTaskUpsertBulk) UpdateUpdatedAt() *ArchivedTaskActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedTaskActivityTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedTaskActivityTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArchivedTaskActivityTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedTaskActivityTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedTaskActivityTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
