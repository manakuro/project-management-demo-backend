// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/schema/pulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestTodoUpdate is the builder for updating TestTodo entities.
type TestTodoUpdate struct {
	config
	hooks    []Hook
	mutation *TestTodoMutation
}

// Where appends a list predicates to the TestTodoUpdate builder.
func (ttu *TestTodoUpdate) Where(ps ...predicate.TestTodo) *TestTodoUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetTestUserID sets the "test_user_id" field.
func (ttu *TestTodoUpdate) SetTestUserID(pu pulid.ID) *TestTodoUpdate {
	ttu.mutation.SetTestUserID(pu)
	return ttu
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (ttu *TestTodoUpdate) SetNillableTestUserID(pu *pulid.ID) *TestTodoUpdate {
	if pu != nil {
		ttu.SetTestUserID(*pu)
	}
	return ttu
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (ttu *TestTodoUpdate) ClearTestUserID() *TestTodoUpdate {
	ttu.mutation.ClearTestUserID()
	return ttu
}

// SetName sets the "name" field.
func (ttu *TestTodoUpdate) SetName(s string) *TestTodoUpdate {
	ttu.mutation.SetName(s)
	return ttu
}

// SetStatus sets the "status" field.
func (ttu *TestTodoUpdate) SetStatus(t testtodo.Status) *TestTodoUpdate {
	ttu.mutation.SetStatus(t)
	return ttu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ttu *TestTodoUpdate) SetNillableStatus(t *testtodo.Status) *TestTodoUpdate {
	if t != nil {
		ttu.SetStatus(*t)
	}
	return ttu
}

// SetPriority sets the "priority" field.
func (ttu *TestTodoUpdate) SetPriority(i int) *TestTodoUpdate {
	ttu.mutation.ResetPriority()
	ttu.mutation.SetPriority(i)
	return ttu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ttu *TestTodoUpdate) SetNillablePriority(i *int) *TestTodoUpdate {
	if i != nil {
		ttu.SetPriority(*i)
	}
	return ttu
}

// AddPriority adds i to the "priority" field.
func (ttu *TestTodoUpdate) AddPriority(i int) *TestTodoUpdate {
	ttu.mutation.AddPriority(i)
	return ttu
}

// SetTestUser sets the "test_user" edge to the TestUser entity.
func (ttu *TestTodoUpdate) SetTestUser(t *TestUser) *TestTodoUpdate {
	return ttu.SetTestUserID(t.ID)
}

// Mutation returns the TestTodoMutation object of the builder.
func (ttu *TestTodoUpdate) Mutation() *TestTodoMutation {
	return ttu.mutation
}

// ClearTestUser clears the "test_user" edge to the TestUser entity.
func (ttu *TestTodoUpdate) ClearTestUser() *TestTodoUpdate {
	ttu.mutation.ClearTestUser()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TestTodoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttu.hooks) == 0 {
		if err = ttu.check(); err != nil {
			return 0, err
		}
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestTodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttu.check(); err != nil {
				return 0, err
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			if ttu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TestTodoUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TestTodoUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TestTodoUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TestTodoUpdate) check() error {
	if v, ok := ttu.mutation.Name(); ok {
		if err := testtodo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ttu.mutation.Status(); ok {
		if err := testtodo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ttu *TestTodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testtodo.Table,
			Columns: testtodo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testtodo.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testtodo.FieldName,
		})
	}
	if value, ok := ttu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: testtodo.FieldStatus,
		})
	}
	if value, ok := ttu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testtodo.FieldPriority,
		})
	}
	if value, ok := ttu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testtodo.FieldPriority,
		})
	}
	if ttu.mutation.TestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testtodo.TestUserTable,
			Columns: []string{testtodo.TestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: testuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TestUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testtodo.TestUserTable,
			Columns: []string{testtodo.TestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: testuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testtodo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TestTodoUpdateOne is the builder for updating a single TestTodo entity.
type TestTodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestTodoMutation
}

// SetTestUserID sets the "test_user_id" field.
func (ttuo *TestTodoUpdateOne) SetTestUserID(pu pulid.ID) *TestTodoUpdateOne {
	ttuo.mutation.SetTestUserID(pu)
	return ttuo
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (ttuo *TestTodoUpdateOne) SetNillableTestUserID(pu *pulid.ID) *TestTodoUpdateOne {
	if pu != nil {
		ttuo.SetTestUserID(*pu)
	}
	return ttuo
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (ttuo *TestTodoUpdateOne) ClearTestUserID() *TestTodoUpdateOne {
	ttuo.mutation.ClearTestUserID()
	return ttuo
}

// SetName sets the "name" field.
func (ttuo *TestTodoUpdateOne) SetName(s string) *TestTodoUpdateOne {
	ttuo.mutation.SetName(s)
	return ttuo
}

// SetStatus sets the "status" field.
func (ttuo *TestTodoUpdateOne) SetStatus(t testtodo.Status) *TestTodoUpdateOne {
	ttuo.mutation.SetStatus(t)
	return ttuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ttuo *TestTodoUpdateOne) SetNillableStatus(t *testtodo.Status) *TestTodoUpdateOne {
	if t != nil {
		ttuo.SetStatus(*t)
	}
	return ttuo
}

// SetPriority sets the "priority" field.
func (ttuo *TestTodoUpdateOne) SetPriority(i int) *TestTodoUpdateOne {
	ttuo.mutation.ResetPriority()
	ttuo.mutation.SetPriority(i)
	return ttuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ttuo *TestTodoUpdateOne) SetNillablePriority(i *int) *TestTodoUpdateOne {
	if i != nil {
		ttuo.SetPriority(*i)
	}
	return ttuo
}

// AddPriority adds i to the "priority" field.
func (ttuo *TestTodoUpdateOne) AddPriority(i int) *TestTodoUpdateOne {
	ttuo.mutation.AddPriority(i)
	return ttuo
}

// SetTestUser sets the "test_user" edge to the TestUser entity.
func (ttuo *TestTodoUpdateOne) SetTestUser(t *TestUser) *TestTodoUpdateOne {
	return ttuo.SetTestUserID(t.ID)
}

// Mutation returns the TestTodoMutation object of the builder.
func (ttuo *TestTodoUpdateOne) Mutation() *TestTodoMutation {
	return ttuo.mutation
}

// ClearTestUser clears the "test_user" edge to the TestUser entity.
func (ttuo *TestTodoUpdateOne) ClearTestUser() *TestTodoUpdateOne {
	ttuo.mutation.ClearTestUser()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TestTodoUpdateOne) Select(field string, fields ...string) *TestTodoUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TestTodo entity.
func (ttuo *TestTodoUpdateOne) Save(ctx context.Context) (*TestTodo, error) {
	var (
		err  error
		node *TestTodo
	)
	if len(ttuo.hooks) == 0 {
		if err = ttuo.check(); err != nil {
			return nil, err
		}
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestTodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttuo.check(); err != nil {
				return nil, err
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			if ttuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TestTodoUpdateOne) SaveX(ctx context.Context) *TestTodo {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TestTodoUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TestTodoUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TestTodoUpdateOne) check() error {
	if v, ok := ttuo.mutation.Name(); ok {
		if err := testtodo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ttuo.mutation.Status(); ok {
		if err := testtodo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ttuo *TestTodoUpdateOne) sqlSave(ctx context.Context) (_node *TestTodo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testtodo.Table,
			Columns: testtodo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testtodo.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TestTodo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testtodo.FieldID)
		for _, f := range fields {
			if !testtodo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testtodo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testtodo.FieldName,
		})
	}
	if value, ok := ttuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: testtodo.FieldStatus,
		})
	}
	if value, ok := ttuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testtodo.FieldPriority,
		})
	}
	if value, ok := ttuo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testtodo.FieldPriority,
		})
	}
	if ttuo.mutation.TestUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testtodo.TestUserTable,
			Columns: []string{testtodo.TestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: testuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TestUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testtodo.TestUserTable,
			Columns: []string{testtodo.TestUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: testuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TestTodo{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testtodo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
