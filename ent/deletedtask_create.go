// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeletedTaskCreate is the builder for creating a DeletedTask entity.
type DeletedTaskCreate struct {
	config
	mutation *DeletedTaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTaskID sets the "task_id" field.
func (dtc *DeletedTaskCreate) SetTaskID(u ulid.ID) *DeletedTaskCreate {
	dtc.mutation.SetTaskID(u)
	return dtc
}

// SetWorkspaceID sets the "workspace_id" field.
func (dtc *DeletedTaskCreate) SetWorkspaceID(u ulid.ID) *DeletedTaskCreate {
	dtc.mutation.SetWorkspaceID(u)
	return dtc
}

// SetCreatedAt sets the "created_at" field.
func (dtc *DeletedTaskCreate) SetCreatedAt(t time.Time) *DeletedTaskCreate {
	dtc.mutation.SetCreatedAt(t)
	return dtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dtc *DeletedTaskCreate) SetNillableCreatedAt(t *time.Time) *DeletedTaskCreate {
	if t != nil {
		dtc.SetCreatedAt(*t)
	}
	return dtc
}

// SetUpdatedAt sets the "updated_at" field.
func (dtc *DeletedTaskCreate) SetUpdatedAt(t time.Time) *DeletedTaskCreate {
	dtc.mutation.SetUpdatedAt(t)
	return dtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dtc *DeletedTaskCreate) SetNillableUpdatedAt(t *time.Time) *DeletedTaskCreate {
	if t != nil {
		dtc.SetUpdatedAt(*t)
	}
	return dtc
}

// SetID sets the "id" field.
func (dtc *DeletedTaskCreate) SetID(u ulid.ID) *DeletedTaskCreate {
	dtc.mutation.SetID(u)
	return dtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dtc *DeletedTaskCreate) SetNillableID(u *ulid.ID) *DeletedTaskCreate {
	if u != nil {
		dtc.SetID(*u)
	}
	return dtc
}

// SetTask sets the "task" edge to the Task entity.
func (dtc *DeletedTaskCreate) SetTask(t *Task) *DeletedTaskCreate {
	return dtc.SetTaskID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (dtc *DeletedTaskCreate) SetWorkspace(w *Workspace) *DeletedTaskCreate {
	return dtc.SetWorkspaceID(w.ID)
}

// Mutation returns the DeletedTaskMutation object of the builder.
func (dtc *DeletedTaskCreate) Mutation() *DeletedTaskMutation {
	return dtc.mutation
}

// Save creates the DeletedTask in the database.
func (dtc *DeletedTaskCreate) Save(ctx context.Context) (*DeletedTask, error) {
	var (
		err  error
		node *DeletedTask
	)
	dtc.defaults()
	if len(dtc.hooks) == 0 {
		if err = dtc.check(); err != nil {
			return nil, err
		}
		node, err = dtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeletedTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dtc.check(); err != nil {
				return nil, err
			}
			dtc.mutation = mutation
			if node, err = dtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dtc.hooks) - 1; i >= 0; i-- {
			if dtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DeletedTaskCreate) SaveX(ctx context.Context) *DeletedTask {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DeletedTaskCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DeletedTaskCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dtc *DeletedTaskCreate) defaults() {
	if _, ok := dtc.mutation.CreatedAt(); !ok {
		v := deletedtask.DefaultCreatedAt()
		dtc.mutation.SetCreatedAt(v)
	}
	if _, ok := dtc.mutation.UpdatedAt(); !ok {
		v := deletedtask.DefaultUpdatedAt()
		dtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dtc.mutation.ID(); !ok {
		v := deletedtask.DefaultID()
		dtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DeletedTaskCreate) check() error {
	if _, ok := dtc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "DeletedTask.task_id"`)}
	}
	if _, ok := dtc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "DeletedTask.workspace_id"`)}
	}
	if _, ok := dtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeletedTask.created_at"`)}
	}
	if _, ok := dtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeletedTask.updated_at"`)}
	}
	if _, ok := dtc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "DeletedTask.task"`)}
	}
	if _, ok := dtc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New(`ent: missing required edge "DeletedTask.workspace"`)}
	}
	return nil
}

func (dtc *DeletedTaskCreate) sqlSave(ctx context.Context) (*DeletedTask, error) {
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
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

func (dtc *DeletedTaskCreate) createSpec() (*DeletedTask, *sqlgraph.CreateSpec) {
	var (
		_node = &DeletedTask{config: dtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deletedtask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: deletedtask.FieldID,
			},
		}
	)
	_spec.OnConflict = dtc.conflict
	if id, ok := dtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtask.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dtc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deletedtask.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := dtc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtask.TaskTable,
			Columns: []string{deletedtask.TaskColumn},
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
	if nodes := dtc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deletedtask.WorkspaceTable,
			Columns: []string{deletedtask.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkspaceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeletedTask.Create().
//		SetTaskID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeletedTaskUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (dtc *DeletedTaskCreate) OnConflict(opts ...sql.ConflictOption) *DeletedTaskUpsertOne {
	dtc.conflict = opts
	return &DeletedTaskUpsertOne{
		create: dtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeletedTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dtc *DeletedTaskCreate) OnConflictColumns(columns ...string) *DeletedTaskUpsertOne {
	dtc.conflict = append(dtc.conflict, sql.ConflictColumns(columns...))
	return &DeletedTaskUpsertOne{
		create: dtc,
	}
}

type (
	// DeletedTaskUpsertOne is the builder for "upsert"-ing
	//  one DeletedTask node.
	DeletedTaskUpsertOne struct {
		create *DeletedTaskCreate
	}

	// DeletedTaskUpsert is the "OnConflict" setter.
	DeletedTaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetTaskID sets the "task_id" field.
func (u *DeletedTaskUpsert) SetTaskID(v ulid.ID) *DeletedTaskUpsert {
	u.Set(deletedtask.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTaskUpsert) UpdateTaskID() *DeletedTaskUpsert {
	u.SetExcluded(deletedtask.FieldTaskID)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTaskUpsert) SetWorkspaceID(v ulid.ID) *DeletedTaskUpsert {
	u.Set(deletedtask.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTaskUpsert) UpdateWorkspaceID() *DeletedTaskUpsert {
	u.SetExcluded(deletedtask.FieldWorkspaceID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTaskUpsert) SetCreatedAt(v time.Time) *DeletedTaskUpsert {
	u.Set(deletedtask.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTaskUpsert) UpdateCreatedAt() *DeletedTaskUpsert {
	u.SetExcluded(deletedtask.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTaskUpsert) SetUpdatedAt(v time.Time) *DeletedTaskUpsert {
	u.Set(deletedtask.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTaskUpsert) UpdateUpdatedAt() *DeletedTaskUpsert {
	u.SetExcluded(deletedtask.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeletedTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deletedtask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeletedTaskUpsertOne) UpdateNewValues() *DeletedTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deletedtask.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(deletedtask.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(deletedtask.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DeletedTask.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DeletedTaskUpsertOne) Ignore() *DeletedTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeletedTaskUpsertOne) DoNothing() *DeletedTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeletedTaskCreate.OnConflict
// documentation for more info.
func (u *DeletedTaskUpsertOne) Update(set func(*DeletedTaskUpsert)) *DeletedTaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeletedTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *DeletedTaskUpsertOne) SetTaskID(v ulid.ID) *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTaskUpsertOne) UpdateTaskID() *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTaskUpsertOne) SetWorkspaceID(v ulid.ID) *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTaskUpsertOne) UpdateWorkspaceID() *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTaskUpsertOne) SetCreatedAt(v time.Time) *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTaskUpsertOne) UpdateCreatedAt() *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTaskUpsertOne) SetUpdatedAt(v time.Time) *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTaskUpsertOne) UpdateUpdatedAt() *DeletedTaskUpsertOne {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeletedTaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeletedTaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeletedTaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeletedTaskUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DeletedTaskUpsertOne.ID is not supported by MySQL driver. Use DeletedTaskUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeletedTaskUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeletedTaskCreateBulk is the builder for creating many DeletedTask entities in bulk.
type DeletedTaskCreateBulk struct {
	config
	builders []*DeletedTaskCreate
	conflict []sql.ConflictOption
}

// Save creates the DeletedTask entities in the database.
func (dtcb *DeletedTaskCreateBulk) Save(ctx context.Context) ([]*DeletedTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DeletedTask, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeletedTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DeletedTaskCreateBulk) SaveX(ctx context.Context) []*DeletedTask {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DeletedTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DeletedTaskCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeletedTask.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeletedTaskUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (dtcb *DeletedTaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeletedTaskUpsertBulk {
	dtcb.conflict = opts
	return &DeletedTaskUpsertBulk{
		create: dtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeletedTask.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dtcb *DeletedTaskCreateBulk) OnConflictColumns(columns ...string) *DeletedTaskUpsertBulk {
	dtcb.conflict = append(dtcb.conflict, sql.ConflictColumns(columns...))
	return &DeletedTaskUpsertBulk{
		create: dtcb,
	}
}

// DeletedTaskUpsertBulk is the builder for "upsert"-ing
// a bulk of DeletedTask nodes.
type DeletedTaskUpsertBulk struct {
	create *DeletedTaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeletedTask.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deletedtask.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DeletedTaskUpsertBulk) UpdateNewValues() *DeletedTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deletedtask.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(deletedtask.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(deletedtask.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeletedTask.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DeletedTaskUpsertBulk) Ignore() *DeletedTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeletedTaskUpsertBulk) DoNothing() *DeletedTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeletedTaskCreateBulk.OnConflict
// documentation for more info.
func (u *DeletedTaskUpsertBulk) Update(set func(*DeletedTaskUpsert)) *DeletedTaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeletedTaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *DeletedTaskUpsertBulk) SetTaskID(v ulid.ID) *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *DeletedTaskUpsertBulk) UpdateTaskID() *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateTaskID()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *DeletedTaskUpsertBulk) SetWorkspaceID(v ulid.ID) *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *DeletedTaskUpsertBulk) UpdateWorkspaceID() *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateWorkspaceID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeletedTaskUpsertBulk) SetCreatedAt(v time.Time) *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeletedTaskUpsertBulk) UpdateCreatedAt() *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeletedTaskUpsertBulk) SetUpdatedAt(v time.Time) *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeletedTaskUpsertBulk) UpdateUpdatedAt() *DeletedTaskUpsertBulk {
	return u.Update(func(s *DeletedTaskUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeletedTaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeletedTaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeletedTaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeletedTaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
