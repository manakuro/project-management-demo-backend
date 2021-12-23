// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateCreate is the builder for creating a Teammate entity.
type TeammateCreate struct {
	config
	mutation *TeammateMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TeammateCreate) SetName(s string) *TeammateCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetImage sets the "image" field.
func (tc *TeammateCreate) SetImage(s string) *TeammateCreate {
	tc.mutation.SetImage(s)
	return tc
}

// SetEmail sets the "email" field.
func (tc *TeammateCreate) SetEmail(s string) *TeammateCreate {
	tc.mutation.SetEmail(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TeammateCreate) SetCreatedAt(t time.Time) *TeammateCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TeammateCreate) SetNillableCreatedAt(t *time.Time) *TeammateCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TeammateCreate) SetUpdatedAt(t time.Time) *TeammateCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TeammateCreate) SetNillableUpdatedAt(t *time.Time) *TeammateCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TeammateCreate) SetID(u ulid.ID) *TeammateCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TeammateCreate) SetNillableID(u *ulid.ID) *TeammateCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TeammateMutation object of the builder.
func (tc *TeammateCreate) Mutation() *TeammateMutation {
	return tc.mutation
}

// Save creates the Teammate in the database.
func (tc *TeammateCreate) Save(ctx context.Context) (*Teammate, error) {
	var (
		err  error
		node *Teammate
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeammateCreate) SaveX(ctx context.Context) *Teammate {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeammateCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeammateCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TeammateCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := teammate.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := teammate.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := teammate.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeammateCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := teammate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "image"`)}
	}
	if v, ok := tc.mutation.Image(); ok {
		if err := teammate.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "image": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if v, ok := tc.mutation.Email(); ok {
		if err := teammate.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "email": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (tc *TeammateCreate) sqlSave(ctx context.Context) (*Teammate, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TeammateCreate) createSpec() (*Teammate, *sqlgraph.CreateSpec) {
	var (
		_node = &Teammate{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teammate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammate.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammate.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammate.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := tc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammate.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammate.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TeammateCreateBulk is the builder for creating many Teammate entities in bulk.
type TeammateCreateBulk struct {
	config
	builders []*TeammateCreate
}

// Save creates the Teammate entities in the database.
func (tcb *TeammateCreateBulk) Save(ctx context.Context) ([]*Teammate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Teammate, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeammateMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeammateCreateBulk) SaveX(ctx context.Context) []*Teammate {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeammateCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeammateCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}