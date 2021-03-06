// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projecttaskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskColumnDelete is the builder for deleting a ProjectTaskColumn entity.
type ProjectTaskColumnDelete struct {
	config
	hooks    []Hook
	mutation *ProjectTaskColumnMutation
}

// Where appends a list predicates to the ProjectTaskColumnDelete builder.
func (ptcd *ProjectTaskColumnDelete) Where(ps ...predicate.ProjectTaskColumn) *ProjectTaskColumnDelete {
	ptcd.mutation.Where(ps...)
	return ptcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptcd *ProjectTaskColumnDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptcd.hooks) == 0 {
		affected, err = ptcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptcd.mutation = mutation
			affected, err = ptcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptcd.hooks) - 1; i >= 0; i-- {
			if ptcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcd *ProjectTaskColumnDelete) ExecX(ctx context.Context) int {
	n, err := ptcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptcd *ProjectTaskColumnDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: projecttaskcolumn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttaskcolumn.FieldID,
			},
		},
	}
	if ps := ptcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ptcd.driver, _spec)
}

// ProjectTaskColumnDeleteOne is the builder for deleting a single ProjectTaskColumn entity.
type ProjectTaskColumnDeleteOne struct {
	ptcd *ProjectTaskColumnDelete
}

// Exec executes the deletion query.
func (ptcdo *ProjectTaskColumnDeleteOne) Exec(ctx context.Context) error {
	n, err := ptcdo.ptcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projecttaskcolumn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcdo *ProjectTaskColumnDeleteOne) ExecX(ctx context.Context) {
	ptcdo.ptcd.ExecX(ctx)
}
