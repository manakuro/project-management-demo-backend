// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskfile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFileDelete is the builder for deleting a TaskFile entity.
type TaskFileDelete struct {
	config
	hooks    []Hook
	mutation *TaskFileMutation
}

// Where appends a list predicates to the TaskFileDelete builder.
func (tfd *TaskFileDelete) Where(ps ...predicate.TaskFile) *TaskFileDelete {
	tfd.mutation.Where(ps...)
	return tfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tfd *TaskFileDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tfd.hooks) == 0 {
		affected, err = tfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tfd.mutation = mutation
			affected, err = tfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tfd.hooks) - 1; i >= 0; i-- {
			if tfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfd *TaskFileDelete) ExecX(ctx context.Context) int {
	n, err := tfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tfd *TaskFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskfile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfile.FieldID,
			},
		},
	}
	if ps := tfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tfd.driver, _spec)
}

// TaskFileDeleteOne is the builder for deleting a single TaskFile entity.
type TaskFileDeleteOne struct {
	tfd *TaskFileDelete
}

// Exec executes the deletion query.
func (tfdo *TaskFileDeleteOne) Exec(ctx context.Context) error {
	n, err := tfdo.tfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskfile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tfdo *TaskFileDeleteOne) ExecX(ctx context.Context) {
	tfdo.tfd.ExecX(ctx)
}
