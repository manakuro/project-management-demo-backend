// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskcolumn"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskColumnDelete is the builder for deleting a TaskColumn entity.
type TaskColumnDelete struct {
	config
	hooks    []Hook
	mutation *TaskColumnMutation
}

// Where appends a list predicates to the TaskColumnDelete builder.
func (tcd *TaskColumnDelete) Where(ps ...predicate.TaskColumn) *TaskColumnDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TaskColumnDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tcd.hooks) == 0 {
		affected, err = tcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskColumnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcd.mutation = mutation
			affected, err = tcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcd.hooks) - 1; i >= 0; i-- {
			if tcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TaskColumnDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TaskColumnDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskcolumn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskcolumn.FieldID,
			},
		},
	}
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
}

// TaskColumnDeleteOne is the builder for deleting a single TaskColumn entity.
type TaskColumnDeleteOne struct {
	tcd *TaskColumnDelete
}

// Exec executes the deletion query.
func (tcdo *TaskColumnDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskcolumn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TaskColumnDeleteOne) ExecX(ctx context.Context) {
	tcdo.tcd.ExecX(ctx)
}
