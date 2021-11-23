// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestUserUpdate is the builder for updating TestUser entities.
type TestUserUpdate struct {
	config
	hooks    []Hook
	mutation *TestUserMutation
}

// Where appends a list predicates to the TestUserUpdate builder.
func (tuu *TestUserUpdate) Where(ps ...predicate.TestUser) *TestUserUpdate {
	tuu.mutation.Where(ps...)
	return tuu
}

// SetName sets the "name" field.
func (tuu *TestUserUpdate) SetName(s string) *TestUserUpdate {
	tuu.mutation.SetName(s)
	return tuu
}

// SetAge sets the "age" field.
func (tuu *TestUserUpdate) SetAge(i int) *TestUserUpdate {
	tuu.mutation.ResetAge()
	tuu.mutation.SetAge(i)
	return tuu
}

// AddAge adds i to the "age" field.
func (tuu *TestUserUpdate) AddAge(i int) *TestUserUpdate {
	tuu.mutation.AddAge(i)
	return tuu
}

// SetCreatedAt sets the "created_at" field.
func (tuu *TestUserUpdate) SetCreatedAt(t time.Time) *TestUserUpdate {
	tuu.mutation.SetCreatedAt(t)
	return tuu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuu *TestUserUpdate) SetNillableCreatedAt(t *time.Time) *TestUserUpdate {
	if t != nil {
		tuu.SetCreatedAt(*t)
	}
	return tuu
}

// SetUpdatedAt sets the "updated_at" field.
func (tuu *TestUserUpdate) SetUpdatedAt(t time.Time) *TestUserUpdate {
	tuu.mutation.SetUpdatedAt(t)
	return tuu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuu *TestUserUpdate) SetNillableUpdatedAt(t *time.Time) *TestUserUpdate {
	if t != nil {
		tuu.SetUpdatedAt(*t)
	}
	return tuu
}

// AddTestTodoIDs adds the "test_todos" edge to the TestTodo entity by IDs.
func (tuu *TestUserUpdate) AddTestTodoIDs(ids ...int) *TestUserUpdate {
	tuu.mutation.AddTestTodoIDs(ids...)
	return tuu
}

// AddTestTodos adds the "test_todos" edges to the TestTodo entity.
func (tuu *TestUserUpdate) AddTestTodos(t ...*TestTodo) *TestUserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuu.AddTestTodoIDs(ids...)
}

// Mutation returns the TestUserMutation object of the builder.
func (tuu *TestUserUpdate) Mutation() *TestUserMutation {
	return tuu.mutation
}

// ClearTestTodos clears all "test_todos" edges to the TestTodo entity.
func (tuu *TestUserUpdate) ClearTestTodos() *TestUserUpdate {
	tuu.mutation.ClearTestTodos()
	return tuu
}

// RemoveTestTodoIDs removes the "test_todos" edge to TestTodo entities by IDs.
func (tuu *TestUserUpdate) RemoveTestTodoIDs(ids ...int) *TestUserUpdate {
	tuu.mutation.RemoveTestTodoIDs(ids...)
	return tuu
}

// RemoveTestTodos removes "test_todos" edges to TestTodo entities.
func (tuu *TestUserUpdate) RemoveTestTodos(t ...*TestTodo) *TestUserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuu.RemoveTestTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tuu *TestUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tuu.hooks) == 0 {
		if err = tuu.check(); err != nil {
			return 0, err
		}
		affected, err = tuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuu.check(); err != nil {
				return 0, err
			}
			tuu.mutation = mutation
			affected, err = tuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tuu.hooks) - 1; i >= 0; i-- {
			if tuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuu *TestUserUpdate) SaveX(ctx context.Context) int {
	affected, err := tuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tuu *TestUserUpdate) Exec(ctx context.Context) error {
	_, err := tuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuu *TestUserUpdate) ExecX(ctx context.Context) {
	if err := tuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuu *TestUserUpdate) check() error {
	if v, ok := tuu.mutation.Name(); ok {
		if err := testuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (tuu *TestUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testuser.Table,
			Columns: testuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: testuser.FieldID,
			},
		},
	}
	if ps := tuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testuser.FieldName,
		})
	}
	if value, ok := tuu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testuser.FieldAge,
		})
	}
	if value, ok := tuu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testuser.FieldAge,
		})
	}
	if value, ok := tuu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldCreatedAt,
		})
	}
	if value, ok := tuu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldUpdatedAt,
		})
	}
	if tuu.mutation.TestTodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuu.mutation.RemovedTestTodosIDs(); len(nodes) > 0 && !tuu.mutation.TestTodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuu.mutation.TestTodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TestUserUpdateOne is the builder for updating a single TestUser entity.
type TestUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestUserMutation
}

// SetName sets the "name" field.
func (tuuo *TestUserUpdateOne) SetName(s string) *TestUserUpdateOne {
	tuuo.mutation.SetName(s)
	return tuuo
}

// SetAge sets the "age" field.
func (tuuo *TestUserUpdateOne) SetAge(i int) *TestUserUpdateOne {
	tuuo.mutation.ResetAge()
	tuuo.mutation.SetAge(i)
	return tuuo
}

// AddAge adds i to the "age" field.
func (tuuo *TestUserUpdateOne) AddAge(i int) *TestUserUpdateOne {
	tuuo.mutation.AddAge(i)
	return tuuo
}

// SetCreatedAt sets the "created_at" field.
func (tuuo *TestUserUpdateOne) SetCreatedAt(t time.Time) *TestUserUpdateOne {
	tuuo.mutation.SetCreatedAt(t)
	return tuuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuuo *TestUserUpdateOne) SetNillableCreatedAt(t *time.Time) *TestUserUpdateOne {
	if t != nil {
		tuuo.SetCreatedAt(*t)
	}
	return tuuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuuo *TestUserUpdateOne) SetUpdatedAt(t time.Time) *TestUserUpdateOne {
	tuuo.mutation.SetUpdatedAt(t)
	return tuuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuuo *TestUserUpdateOne) SetNillableUpdatedAt(t *time.Time) *TestUserUpdateOne {
	if t != nil {
		tuuo.SetUpdatedAt(*t)
	}
	return tuuo
}

// AddTestTodoIDs adds the "test_todos" edge to the TestTodo entity by IDs.
func (tuuo *TestUserUpdateOne) AddTestTodoIDs(ids ...int) *TestUserUpdateOne {
	tuuo.mutation.AddTestTodoIDs(ids...)
	return tuuo
}

// AddTestTodos adds the "test_todos" edges to the TestTodo entity.
func (tuuo *TestUserUpdateOne) AddTestTodos(t ...*TestTodo) *TestUserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuuo.AddTestTodoIDs(ids...)
}

// Mutation returns the TestUserMutation object of the builder.
func (tuuo *TestUserUpdateOne) Mutation() *TestUserMutation {
	return tuuo.mutation
}

// ClearTestTodos clears all "test_todos" edges to the TestTodo entity.
func (tuuo *TestUserUpdateOne) ClearTestTodos() *TestUserUpdateOne {
	tuuo.mutation.ClearTestTodos()
	return tuuo
}

// RemoveTestTodoIDs removes the "test_todos" edge to TestTodo entities by IDs.
func (tuuo *TestUserUpdateOne) RemoveTestTodoIDs(ids ...int) *TestUserUpdateOne {
	tuuo.mutation.RemoveTestTodoIDs(ids...)
	return tuuo
}

// RemoveTestTodos removes "test_todos" edges to TestTodo entities.
func (tuuo *TestUserUpdateOne) RemoveTestTodos(t ...*TestTodo) *TestUserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuuo.RemoveTestTodoIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuuo *TestUserUpdateOne) Select(field string, fields ...string) *TestUserUpdateOne {
	tuuo.fields = append([]string{field}, fields...)
	return tuuo
}

// Save executes the query and returns the updated TestUser entity.
func (tuuo *TestUserUpdateOne) Save(ctx context.Context) (*TestUser, error) {
	var (
		err  error
		node *TestUser
	)
	if len(tuuo.hooks) == 0 {
		if err = tuuo.check(); err != nil {
			return nil, err
		}
		node, err = tuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuuo.check(); err != nil {
				return nil, err
			}
			tuuo.mutation = mutation
			node, err = tuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuuo.hooks) - 1; i >= 0; i-- {
			if tuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuuo *TestUserUpdateOne) SaveX(ctx context.Context) *TestUser {
	node, err := tuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuuo *TestUserUpdateOne) Exec(ctx context.Context) error {
	_, err := tuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuuo *TestUserUpdateOne) ExecX(ctx context.Context) {
	if err := tuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuuo *TestUserUpdateOne) check() error {
	if v, ok := tuuo.mutation.Name(); ok {
		if err := testuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (tuuo *TestUserUpdateOne) sqlSave(ctx context.Context) (_node *TestUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testuser.Table,
			Columns: testuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: testuser.FieldID,
			},
		},
	}
	id, ok := tuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TestUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testuser.FieldID)
		for _, f := range fields {
			if !testuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testuser.FieldName,
		})
	}
	if value, ok := tuuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testuser.FieldAge,
		})
	}
	if value, ok := tuuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testuser.FieldAge,
		})
	}
	if value, ok := tuuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldCreatedAt,
		})
	}
	if value, ok := tuuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testuser.FieldUpdatedAt,
		})
	}
	if tuuo.mutation.TestTodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuuo.mutation.RemovedTestTodosIDs(); len(nodes) > 0 && !tuuo.mutation.TestTodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuuo.mutation.TestTodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testuser.TestTodosTable,
			Columns: []string{testuser.TestTodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: testtodo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TestUser{config: tuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
