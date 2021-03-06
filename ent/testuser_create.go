// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/testuserprofile"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestUserCreate is the builder for creating a TestUser entity.
type TestUserCreate struct {
	config
	mutation *TestUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tuc *TestUserCreate) SetName(s string) *TestUserCreate {
	tuc.mutation.SetName(s)
	return tuc
}

// SetAge sets the "age" field.
func (tuc *TestUserCreate) SetAge(i int) *TestUserCreate {
	tuc.mutation.SetAge(i)
	return tuc
}

// SetProfile sets the "profile" field.
func (tuc *TestUserCreate) SetProfile(tup testuserprofile.TestUserProfile) *TestUserCreate {
	tuc.mutation.SetProfile(tup)
	return tuc
}

// SetDescription sets the "description" field.
func (tuc *TestUserCreate) SetDescription(m map[string]interface{}) *TestUserCreate {
	tuc.mutation.SetDescription(m)
	return tuc
}

// SetCreatedAt sets the "created_at" field.
func (tuc *TestUserCreate) SetCreatedAt(t time.Time) *TestUserCreate {
	tuc.mutation.SetCreatedAt(t)
	return tuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuc *TestUserCreate) SetNillableCreatedAt(t *time.Time) *TestUserCreate {
	if t != nil {
		tuc.SetCreatedAt(*t)
	}
	return tuc
}

// SetUpdatedAt sets the "updated_at" field.
func (tuc *TestUserCreate) SetUpdatedAt(t time.Time) *TestUserCreate {
	tuc.mutation.SetUpdatedAt(t)
	return tuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuc *TestUserCreate) SetNillableUpdatedAt(t *time.Time) *TestUserCreate {
	if t != nil {
		tuc.SetUpdatedAt(*t)
	}
	return tuc
}

// SetID sets the "id" field.
func (tuc *TestUserCreate) SetID(u ulid.ID) *TestUserCreate {
	tuc.mutation.SetID(u)
	return tuc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tuc *TestUserCreate) SetNillableID(u *ulid.ID) *TestUserCreate {
	if u != nil {
		tuc.SetID(*u)
	}
	return tuc
}

// AddTestTodoIDs adds the "testTodos" edge to the TestTodo entity by IDs.
func (tuc *TestUserCreate) AddTestTodoIDs(ids ...ulid.ID) *TestUserCreate {
	tuc.mutation.AddTestTodoIDs(ids...)
	return tuc
}

// AddTestTodos adds the "testTodos" edges to the TestTodo entity.
func (tuc *TestUserCreate) AddTestTodos(t ...*TestTodo) *TestUserCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuc.AddTestTodoIDs(ids...)
}

// Mutation returns the TestUserMutation object of the builder.
func (tuc *TestUserCreate) Mutation() *TestUserMutation {
	return tuc.mutation
}

// Save creates the TestUser in the database.
func (tuc *TestUserCreate) Save(ctx context.Context) (*TestUser, error) {
	var (
		err  error
		node *TestUser
	)
	tuc.defaults()
	if len(tuc.hooks) == 0 {
		if err = tuc.check(); err != nil {
			return nil, err
		}
		node, err = tuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuc.check(); err != nil {
				return nil, err
			}
			tuc.mutation = mutation
			if node, err = tuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tuc.hooks) - 1; i >= 0; i-- {
			if tuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tuc *TestUserCreate) SaveX(ctx context.Context) *TestUser {
	v, err := tuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tuc *TestUserCreate) Exec(ctx context.Context) error {
	_, err := tuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuc *TestUserCreate) ExecX(ctx context.Context) {
	if err := tuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuc *TestUserCreate) defaults() {
	if _, ok := tuc.mutation.CreatedAt(); !ok {
		v := testuser.DefaultCreatedAt()
		tuc.mutation.SetCreatedAt(v)
	}
	if _, ok := tuc.mutation.UpdatedAt(); !ok {
		v := testuser.DefaultUpdatedAt()
		tuc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tuc.mutation.ID(); !ok {
		v := testuser.DefaultID()
		tuc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuc *TestUserCreate) check() error {
	if _, ok := tuc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TestUser.name"`)}
	}
	if v, ok := tuc.mutation.Name(); ok {
		if err := testuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestUser.name": %w`, err)}
		}
	}
	if _, ok := tuc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "TestUser.age"`)}
	}
	if _, ok := tuc.mutation.Profile(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required field "TestUser.profile"`)}
	}
	if _, ok := tuc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "TestUser.description"`)}
	}
	if _, ok := tuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TestUser.created_at"`)}
	}
	if _, ok := tuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TestUser.updated_at"`)}
	}
	return nil
}

func (tuc *TestUserCreate) sqlSave(ctx context.Context) (*TestUser, error) {
	_node, _spec := tuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tuc.driver, _spec); err != nil {
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

func (tuc *TestUserCreate) createSpec() (*TestUser, *sqlgraph.CreateSpec) {
	var (
		_node = &TestUser{config: tuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: testuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testuser.FieldID,
			},
		}
	)
	_spec.OnConflict = tuc.conflict
	if id, ok := tuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tuc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testuser.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tuc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testuser.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := tuc.mutation.Profile(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: testuser.FieldProfile,
		})
		_node.Profile = value
	}
	if value, ok := tuc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: testuser.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tuc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tuc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tuc.mutation.TestTodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: testtodo.FieldID,
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
//	client.TestUser.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestUserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tuc *TestUserCreate) OnConflict(opts ...sql.ConflictOption) *TestUserUpsertOne {
	tuc.conflict = opts
	return &TestUserUpsertOne{
		create: tuc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tuc *TestUserCreate) OnConflictColumns(columns ...string) *TestUserUpsertOne {
	tuc.conflict = append(tuc.conflict, sql.ConflictColumns(columns...))
	return &TestUserUpsertOne{
		create: tuc,
	}
}

type (
	// TestUserUpsertOne is the builder for "upsert"-ing
	//  one TestUser node.
	TestUserUpsertOne struct {
		create *TestUserCreate
	}

	// TestUserUpsert is the "OnConflict" setter.
	TestUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TestUserUpsert) SetName(v string) *TestUserUpsert {
	u.Set(testuser.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateName() *TestUserUpsert {
	u.SetExcluded(testuser.FieldName)
	return u
}

// SetAge sets the "age" field.
func (u *TestUserUpsert) SetAge(v int) *TestUserUpsert {
	u.Set(testuser.FieldAge, v)
	return u
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateAge() *TestUserUpsert {
	u.SetExcluded(testuser.FieldAge)
	return u
}

// AddAge adds v to the "age" field.
func (u *TestUserUpsert) AddAge(v int) *TestUserUpsert {
	u.Add(testuser.FieldAge, v)
	return u
}

// SetProfile sets the "profile" field.
func (u *TestUserUpsert) SetProfile(v testuserprofile.TestUserProfile) *TestUserUpsert {
	u.Set(testuser.FieldProfile, v)
	return u
}

// UpdateProfile sets the "profile" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateProfile() *TestUserUpsert {
	u.SetExcluded(testuser.FieldProfile)
	return u
}

// SetDescription sets the "description" field.
func (u *TestUserUpsert) SetDescription(v map[string]interface{}) *TestUserUpsert {
	u.Set(testuser.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateDescription() *TestUserUpsert {
	u.SetExcluded(testuser.FieldDescription)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TestUserUpsert) SetCreatedAt(v time.Time) *TestUserUpsert {
	u.Set(testuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateCreatedAt() *TestUserUpsert {
	u.SetExcluded(testuser.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestUserUpsert) SetUpdatedAt(v time.Time) *TestUserUpsert {
	u.Set(testuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestUserUpsert) UpdateUpdatedAt() *TestUserUpsert {
	u.SetExcluded(testuser.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TestUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TestUserUpsertOne) UpdateNewValues() *TestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(testuser.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(testuser.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.UpdatedAt(); exists {
			s.SetIgnore(testuser.FieldUpdatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TestUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TestUserUpsertOne) Ignore() *TestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestUserUpsertOne) DoNothing() *TestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestUserCreate.OnConflict
// documentation for more info.
func (u *TestUserUpsertOne) Update(set func(*TestUserUpsert)) *TestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestUserUpsertOne) SetName(v string) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateName() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateName()
	})
}

// SetAge sets the "age" field.
func (u *TestUserUpsertOne) SetAge(v int) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *TestUserUpsertOne) AddAge(v int) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateAge() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateAge()
	})
}

// SetProfile sets the "profile" field.
func (u *TestUserUpsertOne) SetProfile(v testuserprofile.TestUserProfile) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetProfile(v)
	})
}

// UpdateProfile sets the "profile" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateProfile() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateProfile()
	})
}

// SetDescription sets the "description" field.
func (u *TestUserUpsertOne) SetDescription(v map[string]interface{}) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateDescription() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateDescription()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestUserUpsertOne) SetCreatedAt(v time.Time) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateCreatedAt() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestUserUpsertOne) SetUpdatedAt(v time.Time) *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestUserUpsertOne) UpdateUpdatedAt() *TestUserUpsertOne {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TestUserUpsertOne) ID(ctx context.Context) (id ulid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TestUserUpsertOne.ID is not supported by MySQL driver. Use TestUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TestUserUpsertOne) IDX(ctx context.Context) ulid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TestUserCreateBulk is the builder for creating many TestUser entities in bulk.
type TestUserCreateBulk struct {
	config
	builders []*TestUserCreate
	conflict []sql.ConflictOption
}

// Save creates the TestUser entities in the database.
func (tucb *TestUserCreateBulk) Save(ctx context.Context) ([]*TestUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tucb.builders))
	nodes := make([]*TestUser, len(tucb.builders))
	mutators := make([]Mutator, len(tucb.builders))
	for i := range tucb.builders {
		func(i int, root context.Context) {
			builder := tucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestUserMutation)
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
					_, err = mutators[i+1].Mutate(root, tucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tucb *TestUserCreateBulk) SaveX(ctx context.Context) []*TestUser {
	v, err := tucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tucb *TestUserCreateBulk) Exec(ctx context.Context) error {
	_, err := tucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tucb *TestUserCreateBulk) ExecX(ctx context.Context) {
	if err := tucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TestUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestUserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (tucb *TestUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *TestUserUpsertBulk {
	tucb.conflict = opts
	return &TestUserUpsertBulk{
		create: tucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tucb *TestUserCreateBulk) OnConflictColumns(columns ...string) *TestUserUpsertBulk {
	tucb.conflict = append(tucb.conflict, sql.ConflictColumns(columns...))
	return &TestUserUpsertBulk{
		create: tucb,
	}
}

// TestUserUpsertBulk is the builder for "upsert"-ing
// a bulk of TestUser nodes.
type TestUserUpsertBulk struct {
	create *TestUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TestUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TestUserUpsertBulk) UpdateNewValues() *TestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(testuser.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(testuser.FieldCreatedAt)
			}
			if _, exists := b.mutation.UpdatedAt(); exists {
				s.SetIgnore(testuser.FieldUpdatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TestUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TestUserUpsertBulk) Ignore() *TestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestUserUpsertBulk) DoNothing() *TestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestUserCreateBulk.OnConflict
// documentation for more info.
func (u *TestUserUpsertBulk) Update(set func(*TestUserUpsert)) *TestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TestUserUpsertBulk) SetName(v string) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateName() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateName()
	})
}

// SetAge sets the "age" field.
func (u *TestUserUpsertBulk) SetAge(v int) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *TestUserUpsertBulk) AddAge(v int) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateAge() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateAge()
	})
}

// SetProfile sets the "profile" field.
func (u *TestUserUpsertBulk) SetProfile(v testuserprofile.TestUserProfile) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetProfile(v)
	})
}

// UpdateProfile sets the "profile" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateProfile() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateProfile()
	})
}

// SetDescription sets the "description" field.
func (u *TestUserUpsertBulk) SetDescription(v map[string]interface{}) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateDescription() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateDescription()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestUserUpsertBulk) SetCreatedAt(v time.Time) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateCreatedAt() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestUserUpsertBulk) SetUpdatedAt(v time.Time) *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestUserUpsertBulk) UpdateUpdatedAt() *TestUserUpsertBulk {
	return u.Update(func(s *TestUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TestUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
