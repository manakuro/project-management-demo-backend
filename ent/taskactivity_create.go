// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskactivity"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskActivityCreate is the builder for creating a TaskActivity entity.
type TaskActivityCreate struct {
	config
	mutation *TaskActivityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetActivityID sets the "activity_id" field.
func (tac *TaskActivityCreate) SetActivityID(u ulid.ID) *TaskActivityCreate {
	tac.mutation.SetActivityID(u)
	return tac
}

// SetTeammateID sets the "teammate_id" field.
func (tac *TaskActivityCreate) SetTeammateID(u ulid.ID) *TaskActivityCreate {
	tac.mutation.SetTeammateID(u)
	return tac
}

// SetCreatedAt sets the "created_at" field.
func (tac *TaskActivityCreate) SetCreatedAt(t time.Time) *TaskActivityCreate {
	tac.mutation.SetCreatedAt(t)
	return tac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tac *TaskActivityCreate) SetNillableCreatedAt(t *time.Time) *TaskActivityCreate {
	if t != nil {
		tac.SetCreatedAt(*t)
	}
	return tac
}

// SetUpdatedAt sets the "updated_at" field.
func (tac *TaskActivityCreate) SetUpdatedAt(t time.Time) *TaskActivityCreate {
	tac.mutation.SetUpdatedAt(t)
	return tac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tac *TaskActivityCreate) SetNillableUpdatedAt(t *time.Time) *TaskActivityCreate {
	if t != nil {
		tac.SetUpdatedAt(*t)
	}
	return tac
}

// SetID sets the "id" field.
func (tac *TaskActivityCreate) SetID(u ulid.ID) *TaskActivityCreate {
	tac.mutation.SetID(u)
	return tac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tac *TaskActivityCreate) SetNillableID(u *ulid.ID) *TaskActivityCreate {
	if u != nil {
		tac.SetID(*u)
	}
	return tac
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tac *TaskActivityCreate) SetTeammate(t *Teammate) *TaskActivityCreate {
	return tac.SetTeammateID(t.ID)
}

// SetActivityTypeID sets the "activityType" edge to the ActivityType entity by ID.
func (tac *TaskActivityCreate) SetActivityTypeID(id ulid.ID) *TaskActivityCreate {
	tac.mutation.SetActivityTypeID(id)
	return tac
}

// SetActivityType sets the "activityType" edge to the ActivityType entity.
func (tac *TaskActivityCreate) SetActivityType(a *ActivityType) *TaskActivityCreate {
	return tac.SetActivityTypeID(a.ID)
}

// Mutation returns the TaskActivityMutation object of the builder.
func (tac *TaskActivityCreate) Mutation() *TaskActivityMutation {
	return tac.mutation
}

// Save creates the TaskActivity in the database.
func (tac *TaskActivityCreate) Save(ctx context.Context) (*TaskActivity, error) {
	var (
		err  error
		node *TaskActivity
	)
	tac.defaults()
	if len(tac.hooks) == 0 {
		if err = tac.check(); err != nil {
			return nil, err
		}
		node, err = tac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tac.check(); err != nil {
				return nil, err
			}
			tac.mutation = mutation
			if node, err = tac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tac.hooks) - 1; i >= 0; i-- {
			if tac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tac *TaskActivityCreate) SaveX(ctx context.Context) *TaskActivity {
	v, err := tac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tac *TaskActivityCreate) Exec(ctx context.Context) error {
	_, err := tac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tac *TaskActivityCreate) ExecX(ctx context.Context) {
	if err := tac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tac *TaskActivityCreate) defaults() {
	if _, ok := tac.mutation.CreatedAt(); !ok {
		v := taskactivity.DefaultCreatedAt()
		tac.mutation.SetCreatedAt(v)
	}
	if _, ok := tac.mutation.UpdatedAt(); !ok {
		v := taskactivity.DefaultUpdatedAt()
		tac.mutation.SetUpdatedAt(v)
	}
	if _, ok := tac.mutation.ID(); !ok {
		v := taskactivity.DefaultID()
		tac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tac *TaskActivityCreate) check() error {
	if _, ok := tac.mutation.ActivityID(); !ok {
		return &ValidationError{Name: "activity_id", err: errors.New(`ent: missing required field "TaskActivity.activity_id"`)}
	}
	if _, ok := tac.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "TaskActivity.teammate_id"`)}
	}
	if _, ok := tac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TaskActivity.created_at"`)}
	}
	if _, ok := tac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TaskActivity.updated_at"`)}
	}
	if _, ok := tac.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New(`ent: missing required edge "TaskActivity.teammate"`)}
	}
	if _, ok := tac.mutation.ActivityTypeID(); !ok {
		return &ValidationError{Name: "activityType", err: errors.New(`ent: missing required edge "TaskActivity.activityType"`)}
	}
	return nil
}

func (tac *TaskActivityCreate) sqlSave(ctx context.Context) (*TaskActivity, error) {
	_node, _spec := tac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tac.driver, _spec); err != nil {
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

func (tac *TaskActivityCreate) createSpec() (*TaskActivity, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskActivity{config: tac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: taskactivity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskactivity.FieldID,
			},
		}
	)
	_spec.OnConflict = tac.conflict
	if id, ok := tac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskactivity.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskactivity.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tac.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskactivity.TeammateTable,
			Columns: []string{taskactivity.TeammateColumn},
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
	if nodes := tac.mutation.ActivityTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskactivity.ActivityTypeTable,
			Columns: []string{taskactivity.ActivityTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActivityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskActivity.Create().
//		SetActivityID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskActivityUpsert) {
//			SetActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (tac *TaskActivityCreate) OnConflict(opts ...sql.ConflictOption) *TaskActivityUpsertOne {
	tac.conflict = opts
	return &TaskActivityUpsertOne{
		create: tac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tac *TaskActivityCreate) OnConflictColumns(columns ...string) *TaskActivityUpsertOne {
	tac.conflict = append(tac.conflict, sql.ConflictColumns(columns...))
	return &TaskActivityUpsertOne{
		create: tac,
	}
}

type (
	// TaskActivityUpsertOne is the builder for "upsert"-ing
	//  one TaskActivity node.
	TaskActivityUpsertOne struct {
		create *TaskActivityCreate
	}

	// TaskActivityUpsert is the "OnConflict" setter.
	TaskActivityUpsert struct {
		*sql.UpdateSet
	}
)

// SetActivityID sets the "activity_id" field.
func (u *TaskActivityUpsert) SetActivityID(v ulid.ID) *TaskActivityUpsert {
	u.Set(taskactivity.FieldActivityID, v)
	return u
}

// UpdateActivityID sets the "activity_id" field to the value that was provided on create.
func (u *TaskActivityUpsert) UpdateActivityID() *TaskActivityUpsert {
	u.SetExcluded(taskactivity.FieldActivityID)
	return u
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskActivityUpsert) SetTeammateID(v ulid.ID) *TaskActivityUpsert {
	u.Set(taskactivity.FieldTeammateID, v)
	return u
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskActivityUpsert) UpdateTeammateID() *TaskActivityUpsert {
	u.SetExcluded(taskactivity.FieldTeammateID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskActivityUpsert) SetCreatedAt(v time.Time) *TaskActivityUpsert {
	u.Set(taskactivity.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskActivityUpsert) UpdateCreatedAt() *TaskActivityUpsert {
	u.SetExcluded(taskactivity.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskActivityUpsert) SetUpdatedAt(v time.Time) *TaskActivityUpsert {
	u.Set(taskactivity.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskActivityUpsert) UpdateUpdatedAt() *TaskActivityUpsert {
	u.SetExcluded(taskactivity.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TaskActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskActivityUpsertOne) UpdateNewValues() *TaskActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(taskactivity.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(taskactivity.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(taskactivity.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TaskActivity.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TaskActivityUpsertOne) Ignore() *TaskActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskActivityUpsertOne) DoNothing() *TaskActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskActivityCreate.OnConflict
// documentation for more info.
func (u *TaskActivityUpsertOne) Update(set func(*TaskActivityUpsert)) *TaskActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetActivityID sets the "activity_id" field.
func (u *TaskActivityUpsertOne) SetActivityID(v ulid.ID) *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetActivityID(v)
	})
}

// UpdateActivityID sets the "activity_id" field to the value that was provided on create.
func (u *TaskActivityUpsertOne) UpdateActivityID() *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateActivityID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskActivityUpsertOne) SetTeammateID(v ulid.ID) *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskActivityUpsertOne) UpdateTeammateID() *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskActivityUpsertOne) SetCreatedAt(v time.Time) *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskActivityUpsertOne) UpdateCreatedAt() *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskActivityUpsertOne) SetUpdatedAt(v time.Time) *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskActivityUpsertOne) UpdateUpdatedAt() *TaskActivityUpsertOne {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskActivityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskActivityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskActivityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskActivityUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TaskActivityUpsertOne.ID is not supported by MySQL driver. Use TaskActivityUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskActivityUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskActivityCreateBulk is the builder for creating many TaskActivity entities in bulk.
type TaskActivityCreateBulk struct {
	config
	builders []*TaskActivityCreate
	conflict []sql.ConflictOption
}

// Save creates the TaskActivity entities in the database.
func (tacb *TaskActivityCreateBulk) Save(ctx context.Context) ([]*TaskActivity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tacb.builders))
	nodes := make([]*TaskActivity, len(tacb.builders))
	mutators := make([]Mutator, len(tacb.builders))
	for i := range tacb.builders {
		func(i int, root context.Context) {
			builder := tacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskActivityMutation)
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
					_, err = mutators[i+1].Mutate(root, tacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tacb *TaskActivityCreateBulk) SaveX(ctx context.Context) []*TaskActivity {
	v, err := tacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tacb *TaskActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := tacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tacb *TaskActivityCreateBulk) ExecX(ctx context.Context) {
	if err := tacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TaskActivity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskActivityUpsert) {
//			SetActivityID(v+v).
//		}).
//		Exec(ctx)
//
func (tacb *TaskActivityCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskActivityUpsertBulk {
	tacb.conflict = opts
	return &TaskActivityUpsertBulk{
		create: tacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TaskActivity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tacb *TaskActivityCreateBulk) OnConflictColumns(columns ...string) *TaskActivityUpsertBulk {
	tacb.conflict = append(tacb.conflict, sql.ConflictColumns(columns...))
	return &TaskActivityUpsertBulk{
		create: tacb,
	}
}

// TaskActivityUpsertBulk is the builder for "upsert"-ing
// a bulk of TaskActivity nodes.
type TaskActivityUpsertBulk struct {
	create *TaskActivityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TaskActivity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taskactivity.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TaskActivityUpsertBulk) UpdateNewValues() *TaskActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(taskactivity.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(taskactivity.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(taskactivity.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TaskActivity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TaskActivityUpsertBulk) Ignore() *TaskActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskActivityUpsertBulk) DoNothing() *TaskActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskActivityCreateBulk.OnConflict
// documentation for more info.
func (u *TaskActivityUpsertBulk) Update(set func(*TaskActivityUpsert)) *TaskActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetActivityID sets the "activity_id" field.
func (u *TaskActivityUpsertBulk) SetActivityID(v ulid.ID) *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetActivityID(v)
	})
}

// UpdateActivityID sets the "activity_id" field to the value that was provided on create.
func (u *TaskActivityUpsertBulk) UpdateActivityID() *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateActivityID()
	})
}

// SetTeammateID sets the "teammate_id" field.
func (u *TaskActivityUpsertBulk) SetTeammateID(v ulid.ID) *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetTeammateID(v)
	})
}

// UpdateTeammateID sets the "teammate_id" field to the value that was provided on create.
func (u *TaskActivityUpsertBulk) UpdateTeammateID() *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateTeammateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskActivityUpsertBulk) SetCreatedAt(v time.Time) *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskActivityUpsertBulk) UpdateCreatedAt() *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TaskActivityUpsertBulk) SetUpdatedAt(v time.Time) *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TaskActivityUpsertBulk) UpdateUpdatedAt() *TaskActivityUpsertBulk {
	return u.Update(func(s *TaskActivityUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TaskActivityUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskActivityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskActivityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskActivityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
