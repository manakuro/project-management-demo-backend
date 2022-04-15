// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/archivedworkspaceactivity"
	"project-management-demo-backend/ent/archivedworkspaceactivitytask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArchivedWorkspaceActivityTaskCreate is the builder for creating a ArchivedWorkspaceActivityTask entity.
type ArchivedWorkspaceActivityTaskCreate struct {
	config
	mutation *ArchivedWorkspaceActivityTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetArchivedWorkspaceActivityID(u ulid.ID) *ArchivedWorkspaceActivityTaskCreate {
	awatc.mutation.SetArchivedWorkspaceActivityID(u)
	return awatc
}

// SetTaskID sets the "task_id" field.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetTaskID(u ulid.ID) *ArchivedWorkspaceActivityTaskCreate {
	awatc.mutation.SetTaskID(u)
	return awatc
}

// SetCreatedAt sets the "created_at" field.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetCreatedAt(t time.Time) *ArchivedWorkspaceActivityTaskCreate {
	awatc.mutation.SetCreatedAt(t)
	return awatc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetNillableCreatedAt(t *time.Time) *ArchivedWorkspaceActivityTaskCreate {
	if t != nil {
		awatc.SetCreatedAt(*t)
	}
	return awatc
}

// SetUpdatedAt sets the "updated_at" field.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetUpdatedAt(t time.Time) *ArchivedWorkspaceActivityTaskCreate {
	awatc.mutation.SetUpdatedAt(t)
	return awatc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetNillableUpdatedAt(t *time.Time) *ArchivedWorkspaceActivityTaskCreate {
	if t != nil {
		awatc.SetUpdatedAt(*t)
	}
	return awatc
}

// SetID sets the "id" field.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetID(u ulid.ID) *ArchivedWorkspaceActivityTaskCreate {
	awatc.mutation.SetID(u)
	return awatc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetNillableID(u *ulid.ID) *ArchivedWorkspaceActivityTaskCreate {
	if u != nil {
		awatc.SetID(*u)
	}
	return awatc
}

// SetTask sets the "task" edge to the Task entity.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetTask(t *Task) *ArchivedWorkspaceActivityTaskCreate {
	return awatc.SetTaskID(t.ID)
}

// SetArchivedWorkspaceActivity sets the "archivedWorkspaceActivity" edge to the ArchivedWorkspaceActivity entity.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SetArchivedWorkspaceActivity(a *ArchivedWorkspaceActivity) *ArchivedWorkspaceActivityTaskCreate {
	return awatc.SetArchivedWorkspaceActivityID(a.ID)
}

// Mutation returns the ArchivedWorkspaceActivityTaskMutation object of the builder.
func (awatc *ArchivedWorkspaceActivityTaskCreate) Mutation() *ArchivedWorkspaceActivityTaskMutation {
	return awatc.mutation
}

// Save creates the ArchivedWorkspaceActivityTask in the database.
func (awatc *ArchivedWorkspaceActivityTaskCreate) Save(ctx context.Context) (*ArchivedWorkspaceActivityTask, error) {
	var (
		err  error
		node *ArchivedWorkspaceActivityTask
	)
	awatc.defaults()
	if len(awatc.hooks) == 0 {
		if err = awatc.check(); err != nil {
			return nil, err
		}
		node, err = awatc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArchivedWorkspaceActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = awatc.check(); err != nil {
				return nil, err
			}
			awatc.mutation = mutation
			if node, err = awatc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(awatc.hooks) - 1; i >= 0; i-- {
			if awatc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = awatc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, awatc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (awatc *ArchivedWorkspaceActivityTaskCreate) SaveX(ctx context.Context) *ArchivedWorkspaceActivityTask {
	v, err := awatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (awatc *ArchivedWorkspaceActivityTaskCreate) Exec(ctx context.Context) error {
	_, err := awatc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awatc *ArchivedWorkspaceActivityTaskCreate) ExecX(ctx context.Context) {
	if err := awatc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (awatc *ArchivedWorkspaceActivityTaskCreate) defaults() {
	if _, ok := awatc.mutation.CreatedAt(); !ok {
		v := archivedworkspaceactivitytask.DefaultCreatedAt()
		awatc.mutation.SetCreatedAt(v)
	}
	if _, ok := awatc.mutation.UpdatedAt(); !ok {
		v := archivedworkspaceactivitytask.DefaultUpdatedAt()
		awatc.mutation.SetUpdatedAt(v)
	}
	if _, ok := awatc.mutation.ID(); !ok {
		v := archivedworkspaceactivitytask.DefaultID()
		awatc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awatc *ArchivedWorkspaceActivityTaskCreate) check() error {
	if _, ok := awatc.mutation.ArchivedWorkspaceActivityID(); !ok {
		return &ValidationError{Name: "archived_workspace_activity_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivityTask.archived_workspace_activity_id"`)}
	}
	if _, ok := awatc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivityTask.task_id"`)}
	}
	if _, ok := awatc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivityTask.created_at"`)}
	}
	if _, ok := awatc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ArchivedWorkspaceActivityTask.updated_at"`)}
	}
	if _, ok := awatc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivityTask.task"`)}
	}
	if _, ok := awatc.mutation.ArchivedWorkspaceActivityID(); !ok {
		return &ValidationError{Name: "archivedWorkspaceActivity", err: errors.New(`ent: missing required edge "ArchivedWorkspaceActivityTask.archivedWorkspaceActivity"`)}
	}
	return nil
}

func (awatc *ArchivedWorkspaceActivityTaskCreate) sqlSave(ctx context.Context) (*ArchivedWorkspaceActivityTask, error) {
	_node, _spec := awatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, awatc.driver, _spec); err != nil {
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

func (awatc *ArchivedWorkspaceActivityTaskCreate) createSpec() (*ArchivedWorkspaceActivityTask, *sqlgraph.CreateSpec) {
	var (
		_node = &ArchivedWorkspaceActivityTask{config: awatc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: archivedworkspaceactivitytask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: archivedworkspaceactivitytask.FieldID,
			},
		}
	)
	_spec.OnConflict = awatc.conflict
	if id, ok := awatc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := awatc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedworkspaceactivitytask.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := awatc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: archivedworkspaceactivitytask.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := awatc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivitytask.TaskTable,
			Columns: []string{archivedworkspaceactivitytask.TaskColumn},
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
	if nodes := awatc.mutation.ArchivedWorkspaceActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   archivedworkspaceactivitytask.ArchivedWorkspaceActivityTable,
			Columns: []string{archivedworkspaceactivitytask.ArchivedWorkspaceActivityColumn},
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
		_node.ArchivedWorkspaceActivityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		SetArchivedWorkspaceActivityID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedWorkspaceActivityTaskUpsert) {
//			SetArchivedWorkspaceActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (awatc *ArchivedWorkspaceActivityTaskCreate) OnConflict(opts ...sql.ConflictOption) *ArchivedWorkspaceActivityTaskUpsertOne {
	awatc.conflict = opts
	return &ArchivedWorkspaceActivityTaskUpsertOne{
		create: awatc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (awatc *ArchivedWorkspaceActivityTaskCreate) OnConflictColumns(columns ...string) *ArchivedWorkspaceActivityTaskUpsertOne {
	awatc.conflict = append(awatc.conflict, sql.ConflictColumns(columns...))
	return &ArchivedWorkspaceActivityTaskUpsertOne{
		create: awatc,
	}
}

type (
	// ArchivedWorkspaceActivityTaskUpsertOne is the builder for "upsert"-ing
	//  one ArchivedWorkspaceActivityTask node.
	ArchivedWorkspaceActivityTaskUpsertOne struct {
		create *ArchivedWorkspaceActivityTaskCreate
	}

	// ArchivedWorkspaceActivityTaskUpsert is the "OnConflict" setter.
	ArchivedWorkspaceActivityTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsert) SetArchivedWorkspaceActivityID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsert {
	u.Set(archivedworkspaceactivitytask.FieldArchivedWorkspaceActivityID, v)
	return u
}

// UpdateArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsert) UpdateArchivedWorkspaceActivityID() *ArchivedWorkspaceActivityTaskUpsert {
	u.SetExcluded(archivedworkspaceactivitytask.FieldArchivedWorkspaceActivityID)
	return u
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsert) SetTaskID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsert {
	u.Set(archivedworkspaceactivitytask.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsert) UpdateTaskID() *ArchivedWorkspaceActivityTaskUpsert {
	u.SetExcluded(archivedworkspaceactivitytask.FieldTaskID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsert) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsert {
	u.Set(archivedworkspaceactivitytask.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsert) UpdateCreatedAt() *ArchivedWorkspaceActivityTaskUpsert {
	u.SetExcluded(archivedworkspaceactivitytask.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsert) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsert {
	u.Set(archivedworkspaceactivitytask.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsert) UpdateUpdatedAt() *ArchivedWorkspaceActivityTaskUpsert {
	u.SetExcluded(archivedworkspaceactivitytask.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedworkspaceactivitytask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityTaskUpsertOne) UpdateNewValues() *ArchivedWorkspaceActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(archivedworkspaceactivitytask.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(archivedworkspaceactivitytask.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(archivedworkspaceactivitytask.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ArchivedWorkspaceActivityTask.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ArchivedWorkspaceActivityTaskUpsertOne) Ignore() *ArchivedWorkspaceActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) DoNothing() *ArchivedWorkspaceActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedWorkspaceActivityTaskCreate.OnConflict
// documentation for more info.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) Update(set func(*ArchivedWorkspaceActivityTaskUpsert)) *ArchivedWorkspaceActivityTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedWorkspaceActivityTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) SetArchivedWorkspaceActivityID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetArchivedWorkspaceActivityID(v)
	})
}

// UpdateArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) UpdateArchivedWorkspaceActivityID() *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateArchivedWorkspaceActivityID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) SetTaskID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) UpdateTaskID() *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) UpdateCreatedAt() *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) UpdateUpdatedAt() *ArchivedWorkspaceActivityTaskUpsertOne {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedWorkspaceActivityTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArchivedWorkspaceActivityTaskUpsertOne.ID is not supported by MySQL driver. Use ArchivedWorkspaceActivityTaskUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityTaskUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArchivedWorkspaceActivityTaskCreateBulk is the builder for creating many ArchivedWorkspaceActivityTask entities in bulk.
type ArchivedWorkspaceActivityTaskCreateBulk struct {
	config
	builders []*ArchivedWorkspaceActivityTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the ArchivedWorkspaceActivityTask entities in the database.
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) Save(ctx context.Context) ([]*ArchivedWorkspaceActivityTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(awatcb.builders))
	nodes := make([]*ArchivedWorkspaceActivityTask, len(awatcb.builders))
	mutators := make([]Mutator, len(awatcb.builders))
	for i := range awatcb.builders {
		func(i int, root context.Context) {
			builder := awatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArchivedWorkspaceActivityTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, awatcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = awatcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, awatcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, awatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) SaveX(ctx context.Context) []*ArchivedWorkspaceActivityTask {
	v, err := awatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := awatcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) ExecX(ctx context.Context) {
	if err := awatcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArchivedWorkspaceActivityTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArchivedWorkspaceActivityTaskUpsert) {
//			SetArchivedWorkspaceActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArchivedWorkspaceActivityTaskUpsertBulk {
	awatcb.conflict = opts
	return &ArchivedWorkspaceActivityTaskUpsertBulk{
		create: awatcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (awatcb *ArchivedWorkspaceActivityTaskCreateBulk) OnConflictColumns(columns ...string) *ArchivedWorkspaceActivityTaskUpsertBulk {
	awatcb.conflict = append(awatcb.conflict, sql.ConflictColumns(columns...))
	return &ArchivedWorkspaceActivityTaskUpsertBulk{
		create: awatcb,
	}
}

// ArchivedWorkspaceActivityTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of ArchivedWorkspaceActivityTask nodes.
type ArchivedWorkspaceActivityTaskUpsertBulk struct {
	create *ArchivedWorkspaceActivityTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(archivedworkspaceactivitytask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) UpdateNewValues() *ArchivedWorkspaceActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(archivedworkspaceactivitytask.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(archivedworkspaceactivitytask.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(archivedworkspaceactivitytask.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArchivedWorkspaceActivityTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) Ignore() *ArchivedWorkspaceActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) DoNothing() *ArchivedWorkspaceActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArchivedWorkspaceActivityTaskCreateBulk.OnConflict
// documentation for more info.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) Update(set func(*ArchivedWorkspaceActivityTaskUpsert)) *ArchivedWorkspaceActivityTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArchivedWorkspaceActivityTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) SetArchivedWorkspaceActivityID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetArchivedWorkspaceActivityID(v)
	})
}

// UpdateArchivedWorkspaceActivityID sets the "archived_workspace_activity_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) UpdateArchivedWorkspaceActivityID() *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateArchivedWorkspaceActivityID()
	})
}

// SetTaskID sets the "task_id" field.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) SetTaskID(v ulid.ID) *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) UpdateTaskID() *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) SetCreatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) UpdateCreatedAt() *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) SetUpdatedAt(v time.Time) *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) UpdateUpdatedAt() *ArchivedWorkspaceActivityTaskUpsertBulk {
	return u.Update(func(s *ArchivedWorkspaceActivityTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArchivedWorkspaceActivityTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArchivedWorkspaceActivityTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArchivedWorkspaceActivityTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
