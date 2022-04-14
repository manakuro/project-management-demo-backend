// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/workspaceactivity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceActivityDelete is the builder for deleting a WorkspaceActivity entity.
type WorkspaceActivityDelete struct {
	config
	hooks    []Hook
	mutation *WorkspaceActivityMutation
}

// Where appends a list predicates to the WorkspaceActivityDelete builder.
func (wad *WorkspaceActivityDelete) Where(ps ...predicate.WorkspaceActivity) *WorkspaceActivityDelete {
	wad.mutation.Where(ps...)
	return wad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wad *WorkspaceActivityDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wad.hooks) == 0 {
		affected, err = wad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wad.mutation = mutation
			affected, err = wad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wad.hooks) - 1; i >= 0; i-- {
			if wad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wad *WorkspaceActivityDelete) ExecX(ctx context.Context) int {
	n, err := wad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wad *WorkspaceActivityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workspaceactivity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspaceactivity.FieldID,
			},
		},
	}
	if ps := wad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wad.driver, _spec)
}

// WorkspaceActivityDeleteOne is the builder for deleting a single WorkspaceActivity entity.
type WorkspaceActivityDeleteOne struct {
	wad *WorkspaceActivityDelete
}

// Exec executes the deletion query.
func (wado *WorkspaceActivityDeleteOne) Exec(ctx context.Context) error {
	n, err := wado.wad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workspaceactivity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wado *WorkspaceActivityDeleteOne) ExecX(ctx context.Context) {
	wado.wad.ExecX(ctx)
}
