// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tag"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/tasktag"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskTagCreate is the builder for creating a TaskTag entity.
type TaskTagCreate struct {
	config
	mutation *TaskTagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTaskID sets the "task_id" field.
func (ttc *TaskTagCreate) SetTaskID(u ulid.ID) *TaskTagCreate {
	ttc.mutation.SetTaskID(u)
	return ttc
}

// SetTagID sets the "tag_id" field.
func (ttc *TaskTagCreate) SetTagID(u ulid.ID) *TaskTagCreate {
	ttc.mutation.SetTagID(u)
	return ttc
}

// SetCreatedAt sets the "created_at" field.
func (ttc *TaskTagCreate) SetCreatedAt(t time.Time) *TaskTagCreate {
	ttc.mutation.SetCreatedAt(t)
	return ttc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttc *TaskTagCreate) SetNillableCreatedAt(t *time.Time) *TaskTagCreate {
	if t != nil {
		ttc.SetCreatedAt(*t)
	}
	return ttc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttc *TaskTagCreate) SetUpdatedAt(t time.Time) *TaskTagCreate {
	ttc.mutation.SetUpdatedAt(t)
	return ttc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttc *TaskTagCreate) SetNillableUpdatedAt(t *time.Time) *TaskTagCreate {
	if t != nil {
		ttc.SetUpdatedAt(*t)
	}
	return ttc
}

// SetID sets the "id" field.
func (ttc *TaskTagCreate) SetID(u ulid.ID) *TaskTagCreate {
	ttc.mutation.SetID(u)
	return ttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttc *TaskTagCreate) SetNillableID(u *ulid.ID) *TaskTagCreate {
	if u != nil {
		ttc.SetID(*u)
	}
	return ttc
}

// SetTask sets the "task" edge to the Task entity.
func (ttc *TaskTagCreate) SetTask(t *Task) *TaskTagCreate {
	return ttc.SetTaskID(t.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttc *TaskTagCreate) SetTag(t *Tag) *TaskTagCreate {
	return ttc.SetTagID(t.ID)
}

// Mutation returns the TaskTagMutation object of the builder.
func (ttc *TaskTagCreate) Mutation() *TaskTagMutation {
	return ttc.mutation
}

// Save creates the TaskTag in the database.
func (ttc *TaskTagCreate) Save(ctx context.Context) (*TaskTag, error) {
	var (
		err  error
		node *TaskTag
	)
	ttc.defaults()
	if len(ttc.hooks) == 0 {
		if err = ttc.check(); err != nil {
			return nil, err
		}
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttc.check(); err != nil {
				return nil, err
			}
			ttc.mutation = mutation
			if node, err = ttc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttc.hooks) - 1; i >= 0; i-- {
			if ttc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TaskTagCreate) SaveX(ctx context.Context) *TaskTag {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TaskTagCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TaskTagCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TaskTagCreate) defaults() {
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		v := tasktag.DefaultCreatedAt()
		ttc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		v := tasktag.DefaultUpdatedAt()
		ttc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttc.mutation.ID(); !ok {
		v := tasktag.DefaultID()
		ttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TaskTagCreate) check() error {
	if _, ok := ttc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "task_id"`)}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "tag_id"`)}
	}
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ttc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required edge \"task\"")}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag", err: errors.New("ent: missing required edge \"tag\"")}
	}
	return nil
}

func (ttc *TaskTagCreate) sqlSave(ctx context.Context) (*TaskTag, error) {
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
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

func (ttc *TaskTagCreate) createSpec() (*TaskTag, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskTag{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tasktag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasktag.FieldID,
			},
		}
	)
	_spec.OnConflict = ttc.conflict
	if id, ok := ttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ttc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tasktag.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ttc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tasktag.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ttc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasktag.TaskTable,
			Columns: []string{tasktag.TaskColumn},
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
	if nodes := ttc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tasktag.TagTable,
			Columns: []string{tasktag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskTag.Create().
//		SetTaskID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskTagUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (ttc *TaskTagCreate) OnConflict(opts ...sql.ConflictOption) *TaskTagUpsertOne {
	ttc.conflict = opts
	return &TaskTagUpsertOne{
		create: ttc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttc *TaskTagCreate) OnConflictColumns(columns ...string) *TaskTagUpsertOne {
	ttc.conflict = append(ttc.conflict, sql.ConflictColumns(columns...))
	return &TaskTagUpsertOne{
		create: ttc,
	}
}

type (
	// TaskTagUpsertOne is the builder for "upsert"-ing
	//  one TaskTag node.
	TaskTagUpsertOne struct {
		create *TaskTagCreate
	}

	// TaskTagUpsert is the "OnConflict" setter.
	TaskTagUpsert struct {
		*sql.UpdateSet
	}
)

// SetTaskID sets the "task_id" field.
func (u *TaskTagUpsert) SetTaskID(v ulid.ID) *TaskTagUpsert {
	u.Set(tasktag.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskTagUpsert) UpdateTaskID() *TaskTagUpsert {
	u.SetExcluded(tasktag.FieldTaskID)
	return u
}

// SetTagID sets the "tag_id" field.
func (u *TaskTagUpsert) SetTagID(v ulid.ID) *TaskTagUpsert {
	u.Set(tasktag.FieldTagID, v)
	return u
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TaskTagUpsert) UpdateTagID() *TaskTagUpsert {
	u.SetExcluded(tasktag.FieldTagID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskTagUpsert) SetCreatedAt(v time.Time) *TaskTagUpsert {
	u.Set(tasktag.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskTagUpsert) UpdateCreatedAt() *TaskTagUpsert {
	u.SetExcluded(tasktag.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskTagUpsert) SetUpdatedAt(v time.Time) *TaskTagUpsert {
	u.Set(tasktag.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskTagUpsert) UpdateUpdatedAt() *TaskTagUpsert {
	u.SetExcluded(tasktag.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TaskTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tasktag.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskTagUpsertOne) UpdateNewValues() *TaskTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tasktag.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TaskTag.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TaskTagUpsertOne) Ignore() *TaskTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskTagUpsertOne) DoNothing() *TaskTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskTagCreate.OnConflict
// documentation for more info.
func (u *TaskTagUpsertOne) Update(set func(*TaskTagUpsert)) *TaskTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *TaskTagUpsertOne) SetTaskID(v ulid.ID) *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskTagUpsertOne) UpdateTaskID() *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateTaskID()
	})
}

// SetTagID sets the "tag_id" field.
func (u *TaskTagUpsertOne) SetTagID(v ulid.ID) *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TaskTagUpsertOne) UpdateTagID() *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateTagID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskTagUpsertOne) SetCreatedAt(v time.Time) *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskTagUpsertOne) UpdateCreatedAt() *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskTagUpsertOne) SetUpdatedAt(v time.Time) *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskTagUpsertOne) UpdateUpdatedAt() *TaskTagUpsertOne {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskTagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskTagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskTagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskTagUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TaskTagUpsertOne.ID is not supported by MySQL driver. Use TaskTagUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskTagUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskTagCreateBulk is the builder for creating many TaskTag entities in bulk.
type TaskTagCreateBulk struct {
	config
	builders []*TaskTagCreate
	conflict []sql.ConflictOption
}

// Save creates the TaskTag entities in the database.
func (ttcb *TaskTagCreateBulk) Save(ctx context.Context) ([]*TaskTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TaskTag, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskTagMutation)
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
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ttcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TaskTagCreateBulk) SaveX(ctx context.Context) []*TaskTag {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TaskTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TaskTagCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskTag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskTagUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (ttcb *TaskTagCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskTagUpsertBulk {
	ttcb.conflict = opts
	return &TaskTagUpsertBulk{
		create: ttcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ttcb *TaskTagCreateBulk) OnConflictColumns(columns ...string) *TaskTagUpsertBulk {
	ttcb.conflict = append(ttcb.conflict, sql.ConflictColumns(columns...))
	return &TaskTagUpsertBulk{
		create: ttcb,
	}
}

// TaskTagUpsertBulk is the builder for "upsert"-ing
// a bulk of TaskTag nodes.
type TaskTagUpsertBulk struct {
	create *TaskTagCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TaskTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tasktag.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskTagUpsertBulk) UpdateNewValues() *TaskTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tasktag.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TaskTag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TaskTagUpsertBulk) Ignore() *TaskTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskTagUpsertBulk) DoNothing() *TaskTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskTagCreateBulk.OnConflict
// documentation for more info.
func (u *TaskTagUpsertBulk) Update(set func(*TaskTagUpsert)) *TaskTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *TaskTagUpsertBulk) SetTaskID(v ulid.ID) *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskTagUpsertBulk) UpdateTaskID() *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateTaskID()
	})
}

// SetTagID sets the "tag_id" field.
func (u *TaskTagUpsertBulk) SetTagID(v ulid.ID) *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TaskTagUpsertBulk) UpdateTagID() *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateTagID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskTagUpsertBulk) SetCreatedAt(v time.Time) *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskTagUpsertBulk) UpdateCreatedAt() *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskTagUpsertBulk) SetUpdatedAt(v time.Time) *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskTagUpsertBulk) UpdateUpdatedAt() *TaskTagUpsertBulk {
	return u.Update(func(s *TaskTagUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskTagUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskTagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskTagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskTagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
