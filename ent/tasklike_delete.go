// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/tasklike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskLikeDelete is the builder for deleting a TaskLike entity.
type TaskLikeDelete struct {
	config
	hooks    []Hook
	mutation *TaskLikeMutation
}

// Where appends a list predicates to the TaskLikeDelete builder.
func (tld *TaskLikeDelete) Where(ps ...predicate.TaskLike) *TaskLikeDelete {
	tld.mutation.Where(ps...)
	return tld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tld *TaskLikeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tld.hooks) == 0 {
		affected, err = tld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tld.mutation = mutation
			affected, err = tld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tld.hooks) - 1; i >= 0; i-- {
			if tld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tld *TaskLikeDelete) ExecX(ctx context.Context) int {
	n, err := tld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tld *TaskLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tasklike.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tasklike.FieldID,
			},
		},
	}
	if ps := tld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tld.driver, _spec)
}

// TaskLikeDeleteOne is the builder for deleting a single TaskLike entity.
type TaskLikeDeleteOne struct {
	tld *TaskLikeDelete
}

// Exec executes the deletion query.
func (tldo *TaskLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := tldo.tld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tasklike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tldo *TaskLikeDeleteOne) ExecX(ctx context.Context) {
	tldo.tld.ExecX(ctx)
}
