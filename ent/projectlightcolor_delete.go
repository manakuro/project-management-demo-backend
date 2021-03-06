// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projectlightcolor"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectLightColorDelete is the builder for deleting a ProjectLightColor entity.
type ProjectLightColorDelete struct {
	config
	hooks    []Hook
	mutation *ProjectLightColorMutation
}

// Where appends a list predicates to the ProjectLightColorDelete builder.
func (plcd *ProjectLightColorDelete) Where(ps ...predicate.ProjectLightColor) *ProjectLightColorDelete {
	plcd.mutation.Where(ps...)
	return plcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (plcd *ProjectLightColorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(plcd.hooks) == 0 {
		affected, err = plcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLightColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			plcd.mutation = mutation
			affected, err = plcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(plcd.hooks) - 1; i >= 0; i-- {
			if plcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (plcd *ProjectLightColorDelete) ExecX(ctx context.Context) int {
	n, err := plcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (plcd *ProjectLightColorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: projectlightcolor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projectlightcolor.FieldID,
			},
		},
	}
	if ps := plcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, plcd.driver, _spec)
}

// ProjectLightColorDeleteOne is the builder for deleting a single ProjectLightColor entity.
type ProjectLightColorDeleteOne struct {
	plcd *ProjectLightColorDelete
}

// Exec executes the deletion query.
func (plcdo *ProjectLightColorDeleteOne) Exec(ctx context.Context) error {
	n, err := plcdo.plcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projectlightcolor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (plcdo *ProjectLightColorDeleteOne) ExecX(ctx context.Context) {
	plcdo.plcd.ExecX(ctx)
}
