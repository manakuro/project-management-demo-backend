// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskcollaborator"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCollaboratorCreate is the builder for creating a TaskCollaborator entity.
type TaskCollaboratorCreate struct {
	config
	mutation *TaskCollaboratorMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTaskID sets the "task_id" field.
func (tcc *TaskCollaboratorCreate) SetTaskID(u ulid.ID) *TaskCollaboratorCreate {
	tcc.mutation.SetTaskID(u)
	return tcc
}

// SetTeammateID sets the "teammate_id" field.
func (tcc *TaskCollaboratorCreate) SetTeammateID(u ulid.ID) *TaskCollaboratorCreate {
	tcc.mutation.SetTeammateID(u)
	return tcc
}

// SetCreatedAt sets the "created_at" field.
func (tcc *TaskCollaboratorCreate) SetCreatedAt(t time.Time) *TaskCollaboratorCreate {
	tcc.mutation.SetCreatedAt(t)
	return tcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcc *TaskCollaboratorCreate) SetNillableCreatedAt(t *time.Time) *TaskCollaboratorCreate {
	if t != nil {
		tcc.SetCreatedAt(*t)
	}
	return tcc
}

// SetUpdatedAt sets the "updated_at" field.
func (tcc *TaskCollaboratorCreate) SetUpdatedAt(t time.Time) *TaskCollaboratorCreate {
	tcc.mutation.SetUpdatedAt(t)
	return tcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tcc *TaskCollaboratorCreate) SetNillableUpdatedAt(t *time.Time) *TaskCollaboratorCreate {
	if t != nil {
		tcc.SetUpdatedAt(*t)
	}
	return tcc
}

// SetID sets the "id" field.
func (tcc *TaskCollaboratorCreate) SetID(u ulid.ID) *TaskCollaboratorCreate {
	tcc.mutation.SetID(u)
	return tcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tcc *TaskCollaboratorCreate) SetNillableID(u *ulid.ID) *TaskCollaboratorCreate {
	if u != nil {
		tcc.SetID(*u)
	}
	return tcc
}

// SetTask sets the "task" edge to the Task entity.
func (tcc *TaskCollaboratorCreate) SetTask(t *Task) *TaskCollaboratorCreate {
	return tcc.SetTaskID(t.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tcc *TaskCollaboratorCreate) SetTeammate(t *Teammate) *TaskCollaboratorCreate {
	return tcc.SetTeammateID(t.ID)
}

// Mutation returns the TaskCollaboratorMutation object of the builder.
func (tcc *TaskCollaboratorCreate) Mutation() *TaskCollaboratorMutation {
	return tcc.mutation
}

// Save creates the TaskCollaborator in the database.
func (tcc *TaskCollaboratorCreate) Save(ctx context.Context) (*TaskCollaborator, error) {
	var (
		err  error
		node *TaskCollaborator
	)
	tcc.defaults()
	if len(tcc.hooks) == 0 {
		if err = tcc.check(); err != nil {
			return nil, err
		}
		node, err = tcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskCollaboratorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcc.check(); err != nil {
				return nil, err
			}
			tcc.mutation = mutation
			if node, err = tcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tcc.hooks) - 1; i >= 0; i-- {
			if tcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TaskCollaboratorCreate) SaveX(ctx context.Context) *TaskCollaborator {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *TaskCollaboratorCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *TaskCollaboratorCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcc *TaskCollaboratorCreate) defaults() {
	if _, ok := tcc.mutation.CreatedAt(); !ok {
		v := taskcollaborator.DefaultCreatedAt()
		tcc.mutation.SetCreatedAt(v)
	}
	if _, ok := tcc.mutation.UpdatedAt(); !ok {
		v := taskcollaborator.DefaultUpdatedAt()
		tcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tcc.mutation.ID(); !ok {
		v := taskcollaborator.DefaultID()
		tcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TaskCollaboratorCreate) check() error {
	if _, ok := tcc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "TaskCollaborator.task_id"`)}
	}
	if _, ok := tcc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "TaskCollaborator.teammate_id"`)}
	}
	if _, ok := tcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TaskCollaborator.created_at"`)}
	}
	if _, ok := tcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TaskCollaborator.updated_at"`)}
	}
	if _, ok := tcc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "TaskCollaborator.task"`)}
	}
	if _, ok := tcc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "TaskCollaborator.teammate"`)}
	}
	return nil
}

func (tcc *TaskCollaboratorCreate) sqlSave(ctx context.Context) (*TaskCollaborator, error) {
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
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

func (tcc *TaskCollaboratorCreate) createSpec() (*TaskCollaborator, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskCollaborator{config: tcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: taskcollaborator.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcollaborator.FieldID,
			},
		}
	)
	_spec.OnConflict = tcc.conflict
	if id, ok := tcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskcollaborator.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskcollaborator.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tcc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskcollaborator.TaskTable,
			Columns: []string{taskcollaborator.TaskColumn},
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
	if nodes := tcc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskcollaborator.TeammateTable,
			Columns: []string{taskcollaborator.TeammateColumn},
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
		_node.TeammateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskCollaborator.Create().
//		SetTaskID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskCollaboratorUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (tcc *TaskCollaboratorCreate) OnConflict(opts ...sql.ConflictOption) *TaskCollaboratorUpsertOne {
	tcc.conflict = opts
	return &TaskCollaboratorUpsertOne{
		create: tcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskCollaborator.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcc *TaskCollaboratorCreate) OnConflictColumns(columns ...string) *TaskCollaboratorUpsertOne {
	tcc.conflict = append(tcc.conflict, sql.ConflictColumns(columns...))
	return &TaskCollaboratorUpsertOne{
		create: tcc,
	}
}

type (
	// TaskCollaboratorUpsertOne is the builder for "upsert"-ing
	//  one TaskCollaborator node.
	TaskCollaboratorUpsertOne struct {
		create *TaskCollaboratorCreate
	}

	// TaskCollaboratorUpsert is the "OnConflict" setter.
	TaskCollaboratorUpsert struct {
		*sql.UpdateSet
	}
)

// SetTaskID sets the "task_id" field.
func (u *TaskCollaboratorUpsert) SetTaskID(v ulid.ID) *TaskCollaboratorUpsert {
	u.Set(taskcollaborator.FieldTaskID, v)
	return u
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsert) UpdateTaskID() *TaskCollaboratorUpsert {
	u.SetExcluded(taskcollaborator.FieldTaskID)
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskCollaboratorUpsert) SetTeammateID(v ulid.ID) *TaskCollaboratorUpsert {
	u.Set(taskcollaborator.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsert) UpdateTeammateID() *TaskCollaboratorUpsert {
	u.SetExcluded(taskcollaborator.FieldTeammateID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskCollaboratorUpsert) SetCreatedAt(v time.Time) *TaskCollaboratorUpsert {
	u.Set(taskcollaborator.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsert) UpdateCreatedAt() *TaskCollaboratorUpsert {
	u.SetExcluded(taskcollaborator.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskCollaboratorUpsert) SetUpdatedAt(v time.Time) *TaskCollaboratorUpsert {
	u.Set(taskcollaborator.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsert) UpdateUpdatedAt() *TaskCollaboratorUpsert {
	u.SetExcluded(taskcollaborator.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TaskCollaborator.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskcollaborator.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskCollaboratorUpsertOne) UpdateNewValues() *TaskCollaboratorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(taskcollaborator.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(taskcollaborator.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(taskcollaborator.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TaskCollaborator.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TaskCollaboratorUpsertOne) Ignore() *TaskCollaboratorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskCollaboratorUpsertOne) DoNothing() *TaskCollaboratorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCollaboratorCreate.OnConflict
// documentation for more info.
func (u *TaskCollaboratorUpsertOne) Update(set func(*TaskCollaboratorUpsert)) *TaskCollaboratorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskCollaboratorUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *TaskCollaboratorUpsertOne) SetTaskID(v ulid.ID) *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertOne) UpdateTaskID() *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateTaskID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskCollaboratorUpsertOne) SetTeammateID(v ulid.ID) *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertOne) UpdateTeammateID() *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskCollaboratorUpsertOne) SetCreatedAt(v time.Time) *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertOne) UpdateCreatedAt() *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskCollaboratorUpsertOne) SetUpdatedAt(v time.Time) *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertOne) UpdateUpdatedAt() *TaskCollaboratorUpsertOne {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskCollaboratorUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCollaboratorCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskCollaboratorUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskCollaboratorUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TaskCollaboratorUpsertOne.ID is not supported by MySQL driver. Use TaskCollaboratorUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskCollaboratorUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCollaboratorCreateBulk is the builder for creating many TaskCollaborator entities in bulk.
type TaskCollaboratorCreateBulk struct {
	config
	builders []*TaskCollaboratorCreate
	conflict []sql.ConflictOption
}

// Save creates the TaskCollaborator entities in the database.
func (tccb *TaskCollaboratorCreateBulk) Save(ctx context.Context) ([]*TaskCollaborator, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TaskCollaborator, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskCollaboratorMutation)
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
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TaskCollaboratorCreateBulk) SaveX(ctx context.Context) []*TaskCollaborator {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *TaskCollaboratorCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *TaskCollaboratorCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskCollaborator.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskCollaboratorUpsert) {
//			SetTaskID(v+v).
//		}).
//		Exec(ctx)
//
func (tccb *TaskCollaboratorCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskCollaboratorUpsertBulk {
	tccb.conflict = opts
	return &TaskCollaboratorUpsertBulk{
		create: tccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskCollaborator.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tccb *TaskCollaboratorCreateBulk) OnConflictColumns(columns ...string) *TaskCollaboratorUpsertBulk {
	tccb.conflict = append(tccb.conflict, sql.ConflictColumns(columns...))
	return &TaskCollaboratorUpsertBulk{
		create: tccb,
	}
}

// TaskCollaboratorUpsertBulk is the builder for "upsert"-ing
// a bulk of TaskCollaborator nodes.
type TaskCollaboratorUpsertBulk struct {
	create *TaskCollaboratorCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TaskCollaborator.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskcollaborator.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskCollaboratorUpsertBulk) UpdateNewValues() *TaskCollaboratorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(taskcollaborator.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(taskcollaborator.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(taskcollaborator.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TaskCollaborator.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TaskCollaboratorUpsertBulk) Ignore() *TaskCollaboratorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskCollaboratorUpsertBulk) DoNothing() *TaskCollaboratorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCollaboratorCreateBulk.OnConflict
// documentation for more info.
func (u *TaskCollaboratorUpsertBulk) Update(set func(*TaskCollaboratorUpsert)) *TaskCollaboratorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskCollaboratorUpsert{UpdateSet: update})
	}))
	return u
}

// SetTaskID sets the "task_id" field.
func (u *TaskCollaboratorUpsertBulk) SetTaskID(v ulid.ID) *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetTaskID(v)
	})
}

// UpdateTaskID sets the "task_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertBulk) UpdateTaskID() *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateTaskID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskCollaboratorUpsertBulk) SetTeammateID(v ulid.ID) *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertBulk) UpdateTeammateID() *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskCollaboratorUpsertBulk) SetCreatedAt(v time.Time) *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertBulk) UpdateCreatedAt() *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskCollaboratorUpsertBulk) SetUpdatedAt(v time.Time) *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskCollaboratorUpsertBulk) UpdateUpdatedAt() *TaskCollaboratorUpsertBulk {
	return u.Update(func(s *TaskCollaboratorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskCollaboratorUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCollaboratorCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCollaboratorCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskCollaboratorUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
