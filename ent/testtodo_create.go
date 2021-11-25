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

// TestTodoCreate is the builder for creating a TestTodo entity.
type TestTodoCreate struct {
	config
	mutation *TestTodoMutation
	hooks    []Hook
}

// SetTestUserID sets the "test_user_id" field.
func (ttc *TestTodoCreate) SetTestUserID(pu pulid.ID) *TestTodoCreate {
	ttc.mutation.SetTestUserID(pu)
	return ttc
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillableTestUserID(pu *pulid.ID) *TestTodoCreate {
	if pu != nil {
		ttc.SetTestUserID(*pu)
	}
	return ttc
}

// SetName sets the "name" field.
func (ttc *TestTodoCreate) SetName(s string) *TestTodoCreate {
	ttc.mutation.SetName(s)
	return ttc
}

// SetStatus sets the "status" field.
func (ttc *TestTodoCreate) SetStatus(t testtodo.Status) *TestTodoCreate {
	ttc.mutation.SetStatus(t)
	return ttc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillableStatus(t *testtodo.Status) *TestTodoCreate {
	if t != nil {
		ttc.SetStatus(*t)
	}
	return ttc
}

// SetPriority sets the "priority" field.
func (ttc *TestTodoCreate) SetPriority(i int) *TestTodoCreate {
	ttc.mutation.SetPriority(i)
	return ttc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillablePriority(i *int) *TestTodoCreate {
	if i != nil {
		ttc.SetPriority(*i)
	}
	return ttc
}

// SetCreatedAt sets the "created_at" field.
func (ttc *TestTodoCreate) SetCreatedAt(t time.Time) *TestTodoCreate {
	ttc.mutation.SetCreatedAt(t)
	return ttc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillableCreatedAt(t *time.Time) *TestTodoCreate {
	if t != nil {
		ttc.SetCreatedAt(*t)
	}
	return ttc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttc *TestTodoCreate) SetUpdatedAt(t time.Time) *TestTodoCreate {
	ttc.mutation.SetUpdatedAt(t)
	return ttc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillableUpdatedAt(t *time.Time) *TestTodoCreate {
	if t != nil {
		ttc.SetUpdatedAt(*t)
	}
	return ttc
}

// SetID sets the "id" field.
func (ttc *TestTodoCreate) SetID(pu pulid.ID) *TestTodoCreate {
	ttc.mutation.SetID(pu)
	return ttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttc *TestTodoCreate) SetNillableID(pu *pulid.ID) *TestTodoCreate {
	if pu != nil {
		ttc.SetID(*pu)
	}
	return ttc
}

// SetTestUser sets the "test_user" edge to the TestUser entity.
func (ttc *TestTodoCreate) SetTestUser(t *TestUser) *TestTodoCreate {
	return ttc.SetTestUserID(t.ID)
}

// Mutation returns the TestTodoMutation object of the builder.
func (ttc *TestTodoCreate) Mutation() *TestTodoMutation {
	return ttc.mutation
}

// Save creates the TestTodo in the database.
func (ttc *TestTodoCreate) Save(ctx context.Context) (*TestTodo, error) {
	var (
		err  error
		node *TestTodo
	)
	ttc.defaults()
	if len(ttc.hooks) == 0 {
		if err = ttc.check(); err != nil {
			return nil, err
		}
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestTodoMutation)
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
func (ttc *TestTodoCreate) SaveX(ctx context.Context) *TestTodo {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TestTodoCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TestTodoCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TestTodoCreate) defaults() {
	if _, ok := ttc.mutation.Status(); !ok {
		v := testtodo.DefaultStatus
		ttc.mutation.SetStatus(v)
	}
	if _, ok := ttc.mutation.Priority(); !ok {
		v := testtodo.DefaultPriority
		ttc.mutation.SetPriority(v)
	}
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		v := testtodo.DefaultCreatedAt()
		ttc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		v := testtodo.DefaultUpdatedAt()
		ttc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttc.mutation.ID(); !ok {
		v := testtodo.DefaultID()
		ttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TestTodoCreate) check() error {
	if _, ok := ttc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := ttc.mutation.Name(); ok {
		if err := testtodo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := ttc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := ttc.mutation.Status(); ok {
		if err := testtodo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := ttc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "priority"`)}
	}
	if _, ok := ttc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ttc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (ttc *TestTodoCreate) sqlSave(ctx context.Context) (*TestTodo, error) {
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
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

func (ttc *TestTodoCreate) createSpec() (*TestTodo, *sqlgraph.CreateSpec) {
	var (
		_node = &TestTodo{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: testtodo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testtodo.FieldID,
			},
		}
	)
	if id, ok := ttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ttc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: testtodo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ttc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: testtodo.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ttc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: testtodo.FieldPriority,
		})
		_node.Priority = value
	}
	if value, ok := ttc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testtodo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ttc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: testtodo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ttc.mutation.TestUserIDs(); len(nodes) > 0 {
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
		_node.TestUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TestTodoCreateBulk is the builder for creating many TestTodo entities in bulk.
type TestTodoCreateBulk struct {
	config
	builders []*TestTodoCreate
}

// Save creates the TestTodo entities in the database.
func (ttcb *TestTodoCreateBulk) Save(ctx context.Context) ([]*TestTodo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TestTodo, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestTodoMutation)
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
func (ttcb *TestTodoCreateBulk) SaveX(ctx context.Context) []*TestTodo {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TestTodoCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TestTodoCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}
