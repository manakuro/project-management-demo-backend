// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/taskpriority"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskPriorityDelete is the builder for deleting a TaskPriority entity.
type TaskPriorityDelete struct {
	config
	hooks    []Hook
	mutation *TaskPriorityMutation
}

// Where appends a list predicates to the TaskPriorityDelete builder.
func (tpd *TaskPriorityDelete) Where(ps ...predicate.TaskPriority) *TaskPriorityDelete {
	tpd.mutation.Where(ps...)
	return tpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpd *TaskPriorityDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tpd.hooks) == 0 {
		affected, err = tpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskPriorityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tpd.mutation = mutation
			affected, err = tpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tpd.hooks) - 1; i >= 0; i-- {
			if tpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpd *TaskPriorityDelete) ExecX(ctx context.Context) int {
	n, err := tpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpd *TaskPriorityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskpriority.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskpriority.FieldID,
			},
		},
	}
	if ps := tpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tpd.driver, _spec)
}

// TaskPriorityDeleteOne is the builder for deleting a single TaskPriority entity.
type TaskPriorityDeleteOne struct {
	tpd *TaskPriorityDelete
}

// Exec executes the deletion query.
func (tpdo *TaskPriorityDeleteOne) Exec(ctx context.Context) error {
	n, err := tpdo.tpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskpriority.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpdo *TaskPriorityDeleteOne) ExecX(ctx context.Context) {
	tpdo.tpd.ExecX(ctx)
}