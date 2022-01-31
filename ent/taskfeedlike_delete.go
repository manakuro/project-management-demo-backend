// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskfeedlike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFeedLikeDelete is the builder for deleting a TaskFeedLike entity.
type TaskFeedLikeDelete struct {
	config
	hooks    []Hook
	mutation *TaskFeedLikeMutation
}

// Where appends a list predicates to the TaskFeedLikeDelete builder.
func (tfld *TaskFeedLikeDelete) Where(ps ...predicate.TaskFeedLike) *TaskFeedLikeDelete {
	tfld.mutation.Where(ps...)
	return tfld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tfld *TaskFeedLikeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tfld.hooks) == 0 {
		affected, err = tfld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFeedLikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tfld.mutation = mutation
			affected, err = tfld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tfld.hooks) - 1; i >= 0; i-- {
			if tfld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tfld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfld *TaskFeedLikeDelete) ExecX(ctx context.Context) int {
	n, err := tfld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tfld *TaskFeedLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskfeedlike.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfeedlike.FieldID,
			},
		},
	}
	if ps := tfld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tfld.driver, _spec)
}

// TaskFeedLikeDeleteOne is the builder for deleting a single TaskFeedLike entity.
type TaskFeedLikeDeleteOne struct {
	tfld *TaskFeedLikeDelete
}

// Exec executes the deletion query.
func (tfldo *TaskFeedLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := tfldo.tfld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskfeedlike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tfldo *TaskFeedLikeDeleteOne) ExecX(ctx context.Context) {
	tfldo.tfld.ExecX(ctx)
}