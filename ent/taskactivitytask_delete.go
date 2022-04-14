// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskactivitytask"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskActivityTaskDelete is the builder for deleting a TaskActivityTask entity.
type TaskActivityTaskDelete struct {
	config
	hooks    []Hook
	mutation *TaskActivityTaskMutation
}

// Where appends a list predicates to the TaskActivityTaskDelete builder.
func (tatd *TaskActivityTaskDelete) Where(ps ...predicate.TaskActivityTask) *TaskActivityTaskDelete {
	tatd.mutation.Where(ps...)
	return tatd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tatd *TaskActivityTaskDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tatd.hooks) == 0 {
		affected, err = tatd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskActivityTaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tatd.mutation = mutation
			affected, err = tatd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tatd.hooks) - 1; i >= 0; i-- {
			if tatd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tatd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tatd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tatd *TaskActivityTaskDelete) ExecX(ctx context.Context) int {
	n, err := tatd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tatd *TaskActivityTaskDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskactivitytask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskactivitytask.FieldID,
			},
		},
	}
	if ps := tatd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tatd.driver, _spec)
}

// TaskActivityTaskDeleteOne is the builder for deleting a single TaskActivityTask entity.
type TaskActivityTaskDeleteOne struct {
	tatd *TaskActivityTaskDelete
}

// Exec executes the deletion query.
func (tatdo *TaskActivityTaskDeleteOne) Exec(ctx context.Context) error {
	n, err := tatdo.tatd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskactivitytask.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tatdo *TaskActivityTaskDeleteOne) ExecX(ctx context.Context) {
	tatdo.tatd.ExecX(ctx)
}
