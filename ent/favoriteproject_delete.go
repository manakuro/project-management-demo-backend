// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteProjectDelete is the builder for deleting a FavoriteProject entity.
type FavoriteProjectDelete struct {
	config
	hooks    []Hook
	mutation *FavoriteProjectMutation
}

// Where appends a list predicates to the FavoriteProjectDelete builder.
func (fpd *FavoriteProjectDelete) Where(ps ...predicate.FavoriteProject) *FavoriteProjectDelete {
	fpd.mutation.Where(ps...)
	return fpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fpd *FavoriteProjectDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fpd.hooks) == 0 {
		affected, err = fpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fpd.mutation = mutation
			affected, err = fpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fpd.hooks) - 1; i >= 0; i-- {
			if fpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpd *FavoriteProjectDelete) ExecX(ctx context.Context) int {
	n, err := fpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fpd *FavoriteProjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: favoriteproject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteproject.FieldID,
			},
		},
	}
	if ps := fpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fpd.driver, _spec)
}

// FavoriteProjectDeleteOne is the builder for deleting a single FavoriteProject entity.
type FavoriteProjectDeleteOne struct {
	fpd *FavoriteProjectDelete
}

// Exec executes the deletion query.
func (fpdo *FavoriteProjectDeleteOne) Exec(ctx context.Context) error {
	n, err := fpdo.fpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{favoriteproject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fpdo *FavoriteProjectDeleteOne) ExecX(ctx context.Context) {
	fpdo.fpd.ExecX(ctx)
}
