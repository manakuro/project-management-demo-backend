// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/pulid"
	"project-management-demo-backend/ent/testtodo"
	"project-management-demo-backend/ent/testuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestUserCreate is the builder for creating a TestUser entity.
type TestUserCreate struct {
	config
	mutation *TestUserMutation
	hooks    []Hook
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
func (tuc *TestUserCreate) SetID(pu pulid.ID) *TestUserCreate {
	tuc.mutation.SetID(pu)
	return tuc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tuc *TestUserCreate) SetNillableID(pu *pulid.ID) *TestUserCreate {
	if pu != nil {
		tuc.SetID(*pu)
	}
	return tuc
}

// AddTestTodoIDs adds the "test_todos" edge to the TestTodo entity by IDs.
func (tuc *TestUserCreate) AddTestTodoIDs(ids ...pulid.ID) *TestUserCreate {
	tuc.mutation.AddTestTodoIDs(ids...)
	return tuc
}

// AddTestTodos adds the "test_todos" edges to the TestTodo entity.
func (tuc *TestUserCreate) AddTestTodos(t ...*TestTodo) *TestUserCreate {
	ids := make([]pulid.ID, len(t))
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
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tuc.mutation.Name(); ok {
		if err := testuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tuc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "age"`)}
	}
	if _, ok := tuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
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
		_node.ID = _spec.ID.Value.(pulid.ID)
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
	if id, ok := tuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
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

// TestUserCreateBulk is the builder for creating many TestUser entities in bulk.
type TestUserCreateBulk struct {
	config
	builders []*TestUserCreate
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
