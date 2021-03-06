// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/archivedtaskactivity"
	"project-management-demo-backend/ent/archivedworkspaceactivity"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskactivity"
	"project-management-demo-backend/ent/workspaceactivity"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivityTypeCreate is the builder for creating a ActivityType entity.
type ActivityTypeCreate struct {
	config
	mutation *ActivityTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (atc *ActivityTypeCreate) SetName(s string) *ActivityTypeCreate {
	atc.mutation.SetName(s)
	return atc
}

// SetTypeCode sets the "type_code" field.
func (atc *ActivityTypeCreate) SetTypeCode(ac activitytype.TypeCode) *ActivityTypeCreate {
	atc.mutation.SetTypeCode(ac)
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *ActivityTypeCreate) SetCreatedAt(t time.Time) *ActivityTypeCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *ActivityTypeCreate) SetNillableCreatedAt(t *time.Time) *ActivityTypeCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *ActivityTypeCreate) SetUpdatedAt(t time.Time) *ActivityTypeCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *ActivityTypeCreate) SetNillableUpdatedAt(t *time.Time) *ActivityTypeCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *ActivityTypeCreate) SetID(u ulid.ID) *ActivityTypeCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ActivityTypeCreate) SetNillableID(u *ulid.ID) *ActivityTypeCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// AddTaskActivityIDs adds the "taskActivities" edge to the TaskActivity entity by IDs.
func (atc *ActivityTypeCreate) AddTaskActivityIDs(ids ...ulid.ID) *ActivityTypeCreate {
	atc.mutation.AddTaskActivityIDs(ids...)
	return atc
}

// AddTaskActivities adds the "taskActivities" edges to the TaskActivity entity.
func (atc *ActivityTypeCreate) AddTaskActivities(t ...*TaskActivity) *ActivityTypeCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return atc.AddTaskActivityIDs(ids...)
}

// AddWorkspaceActivityIDs adds the "workspaceActivities" edge to the WorkspaceActivity entity by IDs.
func (atc *ActivityTypeCreate) AddWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeCreate {
	atc.mutation.AddWorkspaceActivityIDs(ids...)
	return atc
}

// AddWorkspaceActivities adds the "workspaceActivities" edges to the WorkspaceActivity entity.
func (atc *ActivityTypeCreate) AddWorkspaceActivities(w ...*WorkspaceActivity) *ActivityTypeCreate {
	ids := make([]ulid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return atc.AddWorkspaceActivityIDs(ids...)
}

// AddArchivedTaskActivityIDs adds the "archivedTaskActivities" edge to the ArchivedTaskActivity entity by IDs.
func (atc *ActivityTypeCreate) AddArchivedTaskActivityIDs(ids ...ulid.ID) *ActivityTypeCreate {
	atc.mutation.AddArchivedTaskActivityIDs(ids...)
	return atc
}

// AddArchivedTaskActivities adds the "archivedTaskActivities" edges to the ArchivedTaskActivity entity.
func (atc *ActivityTypeCreate) AddArchivedTaskActivities(a ...*ArchivedTaskActivity) *ActivityTypeCreate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atc.AddArchivedTaskActivityIDs(ids...)
}

// AddArchivedWorkspaceActivityIDs adds the "archivedWorkspaceActivities" edge to the ArchivedWorkspaceActivity entity by IDs.
func (atc *ActivityTypeCreate) AddArchivedWorkspaceActivityIDs(ids ...ulid.ID) *ActivityTypeCreate {
	atc.mutation.AddArchivedWorkspaceActivityIDs(ids...)
	return atc
}

// AddArchivedWorkspaceActivities adds the "archivedWorkspaceActivities" edges to the ArchivedWorkspaceActivity entity.
func (atc *ActivityTypeCreate) AddArchivedWorkspaceActivities(a ...*ArchivedWorkspaceActivity) *ActivityTypeCreate {
	ids := make([]ulid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atc.AddArchivedWorkspaceActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atc *ActivityTypeCreate) Mutation() *ActivityTypeMutation {
	return atc.mutation
}

// Save creates the ActivityType in the database.
func (atc *ActivityTypeCreate) Save(ctx context.Context) (*ActivityType, error) {
	var (
		err  error
		node *ActivityType
	)
	atc.defaults()
	if len(atc.hooks) == 0 {
		if err = atc.check(); err != nil {
			return nil, err
		}
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atc.check(); err != nil {
				return nil, err
			}
			atc.mutation = mutation
			if node, err = atc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			if atc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ActivityTypeCreate) SaveX(ctx context.Context) *ActivityType {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ActivityTypeCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ActivityTypeCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ActivityTypeCreate) defaults() {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := activitytype.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := activitytype.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := activitytype.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ActivityTypeCreate) check() error {
	if _, ok := atc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ActivityType.name"`)}
	}
	if v, ok := atc.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ActivityType.name": %w`, err)}
		}
	}
	if _, ok := atc.mutation.TypeCode(); !ok {
		return &ValidationError{Name: "type_code", err: errors.New(`ent: missing required field "ActivityType.type_code"`)}
	}
	if v, ok := atc.mutation.TypeCode(); ok {
		if err := activitytype.TypeCodeValidator(v); err != nil {
			return &ValidationError{Name: "type_code", err: fmt.Errorf(`ent: validator failed for field "ActivityType.type_code": %w`, err)}
		}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ActivityType.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ActivityType.updated_at"`)}
	}
	return nil
}

func (atc *ActivityTypeCreate) sqlSave(ctx context.Context) (*ActivityType, error) {
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
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

func (atc *ActivityTypeCreate) createSpec() (*ActivityType, *sqlgraph.CreateSpec) {
	var (
		_node = &ActivityType{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: activitytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activitytype.FieldID,
			},
		}
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
		_node.Name = value
	}
	if value, ok := atc.mutation.TypeCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: activitytype.FieldTypeCode,
		})
		_node.TypeCode = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activitytype.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activitytype.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := atc.mutation.TaskActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.WorkspaceActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.ArchivedTaskActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.ArchivedWorkspaceActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ActivityType.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActivityTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (atc *ActivityTypeCreate) OnConflict(opts ...sql.ConflictOption) *ActivityTypeUpsertOne {
	atc.conflict = opts
	return &ActivityTypeUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ActivityType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (atc *ActivityTypeCreate) OnConflictColumns(columns ...string) *ActivityTypeUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &ActivityTypeUpsertOne{
		create: atc,
	}
}

type (
	// ActivityTypeUpsertOne is the builder for "upsert"-ing
	//  one ActivityType node.
	ActivityTypeUpsertOne struct {
		create *ActivityTypeCreate
	}

	// ActivityTypeUpsert is the "OnConflict" setter.
	ActivityTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ActivityTypeUpsert) SetName(v string) *ActivityTypeUpsert {
	u.Set(activitytype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ActivityTypeUpsert) UpdateName() *ActivityTypeUpsert {
	u.SetExcluded(activitytype.FieldName)
	return u
}

// SetTypeCode sets the "type_code" field.
func (u *ActivityTypeUpsert) SetTypeCode(v activitytype.TypeCode) *ActivityTypeUpsert {
	u.Set(activitytype.FieldTypeCode, v)
	return u
}

// UpdateTypeCode sets the "type_code" field to the value that was provided on create.
func (u *ActivityTypeUpsert) UpdateTypeCode() *ActivityTypeUpsert {
	u.SetExcluded(activitytype.FieldTypeCode)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ActivityTypeUpsert) SetCreatedAt(v time.Time) *ActivityTypeUpsert {
	u.Set(activitytype.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ActivityTypeUpsert) UpdateCreatedAt() *ActivityTypeUpsert {
	u.SetExcluded(activitytype.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ActivityTypeUpsert) SetUpdatedAt(v time.Time) *ActivityTypeUpsert {
	u.Set(activitytype.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActivityTypeUpsert) UpdateUpdatedAt() *ActivityTypeUpsert {
	u.SetExcluded(activitytype.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ActivityType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(activitytype.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ActivityTypeUpsertOne) UpdateNewValues() *ActivityTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(activitytype.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(activitytype.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(activitytype.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ActivityType.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ActivityTypeUpsertOne) Ignore() *ActivityTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActivityTypeUpsertOne) DoNothing() *ActivityTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActivityTypeCreate.OnConflict
// documentation for more info.
func (u *ActivityTypeUpsertOne) Update(set func(*ActivityTypeUpsert)) *ActivityTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActivityTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ActivityTypeUpsertOne) SetName(v string) *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ActivityTypeUpsertOne) UpdateName() *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateName()
	})
}

// SetTypeCode sets the "type_code" field.
func (u *ActivityTypeUpsertOne) SetTypeCode(v activitytype.TypeCode) *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetTypeCode(v)
	})
}

// UpdateTypeCode sets the "type_code" field to the value that was provided on create.
func (u *ActivityTypeUpsertOne) UpdateTypeCode() *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateTypeCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ActivityTypeUpsertOne) SetCreatedAt(v time.Time) *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ActivityTypeUpsertOne) UpdateCreatedAt() *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ActivityTypeUpsertOne) SetUpdatedAt(v time.Time) *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActivityTypeUpsertOne) UpdateUpdatedAt() *ActivityTypeUpsertOne {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ActivityTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActivityTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActivityTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ActivityTypeUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ActivityTypeUpsertOne.ID is not supported by MySQL driver. Use ActivityTypeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ActivityTypeUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ActivityTypeCreateBulk is the builder for creating many ActivityType entities in bulk.
type ActivityTypeCreateBulk struct {
	config
	builders []*ActivityTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the ActivityType entities in the database.
func (atcb *ActivityTypeCreateBulk) Save(ctx context.Context) ([]*ActivityType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ActivityType, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ActivityTypeCreateBulk) SaveX(ctx context.Context) []*ActivityType {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ActivityTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ActivityTypeCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ActivityType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActivityTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (atcb *ActivityTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *ActivityTypeUpsertBulk {
	atcb.conflict = opts
	return &ActivityTypeUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ActivityType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (atcb *ActivityTypeCreateBulk) OnConflictColumns(columns ...string) *ActivityTypeUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &ActivityTypeUpsertBulk{
		create: atcb,
	}
}

// ActivityTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of ActivityType nodes.
type ActivityTypeUpsertBulk struct {
	create *ActivityTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ActivityType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(activitytype.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ActivityTypeUpsertBulk) UpdateNewValues() *ActivityTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(activitytype.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(activitytype.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(activitytype.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ActivityType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ActivityTypeUpsertBulk) Ignore() *ActivityTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActivityTypeUpsertBulk) DoNothing() *ActivityTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActivityTypeCreateBulk.OnConflict
// documentation for more info.
func (u *ActivityTypeUpsertBulk) Update(set func(*ActivityTypeUpsert)) *ActivityTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActivityTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ActivityTypeUpsertBulk) SetName(v string) *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ActivityTypeUpsertBulk) UpdateName() *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateName()
	})
}

// SetTypeCode sets the "type_code" field.
func (u *ActivityTypeUpsertBulk) SetTypeCode(v activitytype.TypeCode) *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetTypeCode(v)
	})
}

// UpdateTypeCode sets the "type_code" field to the value that was provided on create.
func (u *ActivityTypeUpsertBulk) UpdateTypeCode() *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateTypeCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ActivityTypeUpsertBulk) SetCreatedAt(v time.Time) *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ActivityTypeUpsertBulk) UpdateCreatedAt() *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ActivityTypeUpsertBulk) SetUpdatedAt(v time.Time) *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActivityTypeUpsertBulk) UpdateUpdatedAt() *ActivityTypeUpsertBulk {
	return u.Update(func(s *ActivityTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ActivityTypeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ActivityTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActivityTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActivityTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
