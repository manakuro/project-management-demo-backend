// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent/predicate"
	"project-management-demo-backend/ent/projecttasksection"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectTaskSectionDelete is the builder for deleting a ProjectTaskSection entity.
type ProjectTaskSectionDelete struct {
	config
	hooks    []Hook
	mutation *ProjectTaskSectionMutation
}

// Where appends a list predicates to the ProjectTaskSectionDelete builder.
func (ptsd *ProjectTaskSectionDelete) Where(ps ...predicate.ProjectTaskSection) *ProjectTaskSectionDelete {
	ptsd.mutation.Where(ps...)
	return ptsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptsd *ProjectTaskSectionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptsd.hooks) == 0 {
		affected, err = ptsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptsd.mutation = mutation
			affected, err = ptsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptsd.hooks) - 1; i >= 0; i-- {
			if ptsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsd *ProjectTaskSectionDelete) ExecX(ctx context.Context) int {
	n, err := ptsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptsd *ProjectTaskSectionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: projecttasksection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecttasksection.FieldID,
			},
		},
	}
	if ps := ptsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ptsd.driver, _spec)
}

// ProjectTaskSectionDeleteOne is the builder for deleting a single ProjectTaskSection entity.
type ProjectTaskSectionDeleteOne struct {
	ptsd *ProjectTaskSectionDelete
}

// Exec executes the deletion query.
func (ptsdo *ProjectTaskSectionDeleteOne) Exec(ctx context.Context) error {
	n, err := ptsdo.ptsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projecttasksection.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptsdo *ProjectTaskSectionDeleteOne) ExecX(ctx context.Context) {
	ptsdo.ptsd.ExecX(ctx)
}
