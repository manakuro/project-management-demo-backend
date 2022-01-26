// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/teammatetaskliststatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskListStatusDelete is the builder for deleting a TeammateTaskListStatus entity.
type TeammateTaskListStatusDelete struct {
	config
	hooks    []Hook
	mutation *TeammateTaskListStatusMutation
}

// Where appends a list predicates to the TeammateTaskListStatusDelete builder.
func (ttlsd *TeammateTaskListStatusDelete) Where(ps ...predicate.TeammateTaskListStatus) *TeammateTaskListStatusDelete {
	ttlsd.mutation.Where(ps...)
	return ttlsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttlsd *TeammateTaskListStatusDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttlsd.hooks) == 0 {
		affected, err = ttlsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskListStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttlsd.mutation = mutation
			affected, err = ttlsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttlsd.hooks) - 1; i >= 0; i-- {
			if ttlsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttlsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttlsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttlsd *TeammateTaskListStatusDelete) ExecX(ctx context.Context) int {
	n, err := ttlsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttlsd *TeammateTaskListStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: teammatetaskliststatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetaskliststatus.FieldID,
			},
		},
	}
	if ps := ttlsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ttlsd.driver, _spec)
}

// TeammateTaskListStatusDeleteOne is the builder for deleting a single TeammateTaskListStatus entity.
type TeammateTaskListStatusDeleteOne struct {
	ttlsd *TeammateTaskListStatusDelete
}

// Exec executes the deletion query.
func (ttlsdo *TeammateTaskListStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := ttlsdo.ttlsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teammatetaskliststatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttlsdo *TeammateTaskListStatusDeleteOne) ExecX(ctx context.Context) {
	ttlsdo.ttlsd.ExecX(ctx)
}