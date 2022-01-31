// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfeedlike"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskFeedLikeCreate is the builder for creating a TaskFeedLike entity.
type TaskFeedLikeCreate struct {
	config
	mutation *TaskFeedLikeMutation
	hooks    []Hook
}

// SetTaskID sets the "task_id" field.
func (tflc *TaskFeedLikeCreate) SetTaskID(u ulid.ID) *TaskFeedLikeCreate {
	tflc.mutation.SetTaskID(u)
	return tflc
}

// SetTeammateID sets the "teammate_id" field.
func (tflc *TaskFeedLikeCreate) SetTeammateID(u ulid.ID) *TaskFeedLikeCreate {
	tflc.mutation.SetTeammateID(u)
	return tflc
}

// SetTaskFeedID sets the "task_feed_id" field.
func (tflc *TaskFeedLikeCreate) SetTaskFeedID(u ulid.ID) *TaskFeedLikeCreate {
	tflc.mutation.SetTaskFeedID(u)
	return tflc
}

// SetCreatedAt sets the "created_at" field.
func (tflc *TaskFeedLikeCreate) SetCreatedAt(t time.Time) *TaskFeedLikeCreate {
	tflc.mutation.SetCreatedAt(t)
	return tflc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tflc *TaskFeedLikeCreate) SetNillableCreatedAt(t *time.Time) *TaskFeedLikeCreate {
	if t != nil {
		tflc.SetCreatedAt(*t)
	}
	return tflc
}

// SetUpdatedAt sets the "updated_at" field.
func (tflc *TaskFeedLikeCreate) SetUpdatedAt(t time.Time) *TaskFeedLikeCreate {
	tflc.mutation.SetUpdatedAt(t)
	return tflc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tflc *TaskFeedLikeCreate) SetNillableUpdatedAt(t *time.Time) *TaskFeedLikeCreate {
	if t != nil {
		tflc.SetUpdatedAt(*t)
	}
	return tflc
}

// SetID sets the "id" field.
func (tflc *TaskFeedLikeCreate) SetID(u ulid.ID) *TaskFeedLikeCreate {
	tflc.mutation.SetID(u)
	return tflc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tflc *TaskFeedLikeCreate) SetNillableID(u *ulid.ID) *TaskFeedLikeCreate {
	if u != nil {
		tflc.SetID(*u)
	}
	return tflc
}

// SetTask sets the "task" edge to the Task entity.
func (tflc *TaskFeedLikeCreate) SetTask(t *Task) *TaskFeedLikeCreate {
	return tflc.SetTaskID(t.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (tflc *TaskFeedLikeCreate) SetTeammate(t *Teammate) *TaskFeedLikeCreate {
	return tflc.SetTeammateID(t.ID)
}

// SetFeedID sets the "feed" edge to the TaskFeed entity by ID.
func (tflc *TaskFeedLikeCreate) SetFeedID(id ulid.ID) *TaskFeedLikeCreate {
	tflc.mutation.SetFeedID(id)
	return tflc
}

// SetFeed sets the "feed" edge to the TaskFeed entity.
func (tflc *TaskFeedLikeCreate) SetFeed(t *TaskFeed) *TaskFeedLikeCreate {
	return tflc.SetFeedID(t.ID)
}

// Mutation returns the TaskFeedLikeMutation object of the builder.
func (tflc *TaskFeedLikeCreate) Mutation() *TaskFeedLikeMutation {
	return tflc.mutation
}

// Save creates the TaskFeedLike in the database.
func (tflc *TaskFeedLikeCreate) Save(ctx context.Context) (*TaskFeedLike, error) {
	var (
		err  error
		node *TaskFeedLike
	)
	tflc.defaults()
	if len(tflc.hooks) == 0 {
		if err = tflc.check(); err != nil {
			return nil, err
		}
		node, err = tflc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskFeedLikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tflc.check(); err != nil {
				return nil, err
			}
			tflc.mutation = mutation
			if node, err = tflc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tflc.hooks) - 1; i >= 0; i-- {
			if tflc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tflc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tflc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tflc *TaskFeedLikeCreate) SaveX(ctx context.Context) *TaskFeedLike {
	v, err := tflc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tflc *TaskFeedLikeCreate) Exec(ctx context.Context) error {
	_, err := tflc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tflc *TaskFeedLikeCreate) ExecX(ctx context.Context) {
	if err := tflc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tflc *TaskFeedLikeCreate) defaults() {
	if _, ok := tflc.mutation.CreatedAt(); !ok {
		v := taskfeedlike.DefaultCreatedAt()
		tflc.mutation.SetCreatedAt(v)
	}
	if _, ok := tflc.mutation.UpdatedAt(); !ok {
		v := taskfeedlike.DefaultUpdatedAt()
		tflc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tflc.mutation.ID(); !ok {
		v := taskfeedlike.DefaultID()
		tflc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tflc *TaskFeedLikeCreate) check() error {
	if _, ok := tflc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "task_id"`)}
	}
	if _, ok := tflc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "teammate_id"`)}
	}
	if _, ok := tflc.mutation.TaskFeedID(); !ok {
		return &ValidationError{Name: "task_feed_id", err: errors.New(`ent: missing required field "task_feed_id"`)}
	}
	if _, ok := tflc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tflc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tflc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required edge \"task\"")}
	}
	if _, ok := tflc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	if _, ok := tflc.mutation.FeedID(); !ok {
		return &ValidationError{Name: "feed", err: errors.New("ent: missing required edge \"feed\"")}
	}
	return nil
}

func (tflc *TaskFeedLikeCreate) sqlSave(ctx context.Context) (*TaskFeedLike, error) {
	_node, _spec := tflc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tflc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(ulid.ID)
	}
	return _node, nil
}

func (tflc *TaskFeedLikeCreate) createSpec() (*TaskFeedLike, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskFeedLike{config: tflc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: taskfeedlike.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: taskfeedlike.FieldID,
			},
		}
	)
	if id, ok := tflc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tflc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskfeedlike.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tflc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: taskfeedlike.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := tflc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfeedlike.TaskTable,
			Columns: []string{taskfeedlike.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tflc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfeedlike.TeammateTable,
			Columns: []string{taskfeedlike.TeammateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeammateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tflc.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taskfeedlike.FeedTable,
			Columns: []string{taskfeedlike.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: taskfeed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskFeedID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskFeedLikeCreateBulk is the builder for creating many TaskFeedLike entities in bulk.
type TaskFeedLikeCreateBulk struct {
	config
	builders []*TaskFeedLikeCreate
}

// Save creates the TaskFeedLike entities in the database.
func (tflcb *TaskFeedLikeCreateBulk) Save(ctx context.Context) ([]*TaskFeedLike, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tflcb.builders))
	nodes := make([]*TaskFeedLike, len(tflcb.builders))
	mutators := make([]Mutator, len(tflcb.builders))
	for i := range tflcb.builders {
		func(i int, root context.Context) {
			builder := tflcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskFeedLikeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tflcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tflcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tflcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tflcb *TaskFeedLikeCreateBulk) SaveX(ctx context.Context) []*TaskFeedLike {
	v, err := tflcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tflcb *TaskFeedLikeCreateBulk) Exec(ctx context.Context) error {
	_, err := tflcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tflcb *TaskFeedLikeCreateBulk) ExecX(ctx context.Context) {
	if err := tflcb.Exec(ctx); err != nil {
		panic(err)
	}
}
