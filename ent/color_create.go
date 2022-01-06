// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ColorCreate is the builder for creating a Color entity.
type ColorCreate struct {
	config
	mutation *ColorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ColorCreate) SetName(s string) *ColorCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetColor sets the "color" field.
func (cc *ColorCreate) SetColor(s string) *ColorCreate {
	cc.mutation.SetColor(s)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ColorCreate) SetCreatedAt(t time.Time) *ColorCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ColorCreate) SetNillableCreatedAt(t *time.Time) *ColorCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ColorCreate) SetUpdatedAt(t time.Time) *ColorCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ColorCreate) SetNillableUpdatedAt(t *time.Time) *ColorCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ColorCreate) SetID(u ulid.ID) *ColorCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ColorCreate) SetNillableID(u *ulid.ID) *ColorCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the ColorMutation object of the builder.
func (cc *ColorCreate) Mutation() *ColorMutation {
	return cc.mutation
}

// Save creates the Color in the database.
func (cc *ColorCreate) Save(ctx context.Context) (*Color, error) {
	var (
		err  error
		node *Color
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ColorCreate) SaveX(ctx context.Context) *Color {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ColorCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ColorCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ColorCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := color.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := color.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := color.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ColorCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := color.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "color"`)}
	}
	if v, ok := cc.mutation.Color(); ok {
		if err := color.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "color": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (cc *ColorCreate) sqlSave(ctx context.Context) (*Color, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ColorCreate) createSpec() (*Color, *sqlgraph.CreateSpec) {
	var (
		_node = &Color{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: color.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: color.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: color.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: color.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: color.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: color.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// ColorCreateBulk is the builder for creating many Color entities in bulk.
type ColorCreateBulk struct {
	config
	builders []*ColorCreate
}

// Save creates the Color entities in the database.
func (ccb *ColorCreateBulk) Save(ctx context.Context) ([]*Color, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Color, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ColorMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ColorCreateBulk) SaveX(ctx context.Context) []*Color {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ColorCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ColorCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
