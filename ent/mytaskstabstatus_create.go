// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/mytaskstabstatus"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MyTasksTabStatusCreate is the builder for creating a MyTasksTabStatus entity.
type MyTasksTabStatusCreate struct {
	config
	mutation *MyTasksTabStatusMutation
	hooks    []Hook
}

// SetWorkspaceID sets the "workspace_id" field.
func (mttsc *MyTasksTabStatusCreate) SetWorkspaceID(u ulid.ID) *MyTasksTabStatusCreate {
	mttsc.mutation.SetWorkspaceID(u)
	return mttsc
}

// SetTeammateID sets the "teammate_id" field.
func (mttsc *MyTasksTabStatusCreate) SetTeammateID(u ulid.ID) *MyTasksTabStatusCreate {
	mttsc.mutation.SetTeammateID(u)
	return mttsc
}

// SetStatus sets the "status" field.
func (mttsc *MyTasksTabStatusCreate) SetStatus(m mytaskstabstatus.Status) *MyTasksTabStatusCreate {
	mttsc.mutation.SetStatus(m)
	return mttsc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mttsc *MyTasksTabStatusCreate) SetNillableStatus(m *mytaskstabstatus.Status) *MyTasksTabStatusCreate {
	if m != nil {
		mttsc.SetStatus(*m)
	}
	return mttsc
}

// SetCreatedAt sets the "created_at" field.
func (mttsc *MyTasksTabStatusCreate) SetCreatedAt(t time.Time) *MyTasksTabStatusCreate {
	mttsc.mutation.SetCreatedAt(t)
	return mttsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mttsc *MyTasksTabStatusCreate) SetNillableCreatedAt(t *time.Time) *MyTasksTabStatusCreate {
	if t != nil {
		mttsc.SetCreatedAt(*t)
	}
	return mttsc
}

// SetUpdatedAt sets the "updated_at" field.
func (mttsc *MyTasksTabStatusCreate) SetUpdatedAt(t time.Time) *MyTasksTabStatusCreate {
	mttsc.mutation.SetUpdatedAt(t)
	return mttsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mttsc *MyTasksTabStatusCreate) SetNillableUpdatedAt(t *time.Time) *MyTasksTabStatusCreate {
	if t != nil {
		mttsc.SetUpdatedAt(*t)
	}
	return mttsc
}

// SetID sets the "id" field.
func (mttsc *MyTasksTabStatusCreate) SetID(u ulid.ID) *MyTasksTabStatusCreate {
	mttsc.mutation.SetID(u)
	return mttsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mttsc *MyTasksTabStatusCreate) SetNillableID(u *ulid.ID) *MyTasksTabStatusCreate {
	if u != nil {
		mttsc.SetID(*u)
	}
	return mttsc
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (mttsc *MyTasksTabStatusCreate) SetWorkspace(w *Workspace) *MyTasksTabStatusCreate {
	return mttsc.SetWorkspaceID(w.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (mttsc *MyTasksTabStatusCreate) SetTeammate(t *Teammate) *MyTasksTabStatusCreate {
	return mttsc.SetTeammateID(t.ID)
}

// Mutation returns the MyTasksTabStatusMutation object of the builder.
func (mttsc *MyTasksTabStatusCreate) Mutation() *MyTasksTabStatusMutation {
	return mttsc.mutation
}

// Save creates the MyTasksTabStatus in the database.
func (mttsc *MyTasksTabStatusCreate) Save(ctx context.Context) (*MyTasksTabStatus, error) {
	var (
		err  error
		node *MyTasksTabStatus
	)
	mttsc.defaults()
	if len(mttsc.hooks) == 0 {
		if err = mttsc.check(); err != nil {
			return nil, err
		}
		node, err = mttsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MyTasksTabStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mttsc.check(); err != nil {
				return nil, err
			}
			mttsc.mutation = mutation
			if node, err = mttsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mttsc.hooks) - 1; i >= 0; i-- {
			if mttsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mttsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mttsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mttsc *MyTasksTabStatusCreate) SaveX(ctx context.Context) *MyTasksTabStatus {
	v, err := mttsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mttsc *MyTasksTabStatusCreate) Exec(ctx context.Context) error {
	_, err := mttsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mttsc *MyTasksTabStatusCreate) ExecX(ctx context.Context) {
	if err := mttsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mttsc *MyTasksTabStatusCreate) defaults() {
	if _, ok := mttsc.mutation.Status(); !ok {
		v := mytaskstabstatus.DefaultStatus
		mttsc.mutation.SetStatus(v)
	}
	if _, ok := mttsc.mutation.CreatedAt(); !ok {
		v := mytaskstabstatus.DefaultCreatedAt()
		mttsc.mutation.SetCreatedAt(v)
	}
	if _, ok := mttsc.mutation.UpdatedAt(); !ok {
		v := mytaskstabstatus.DefaultUpdatedAt()
		mttsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mttsc.mutation.ID(); !ok {
		v := mytaskstabstatus.DefaultID()
		mttsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mttsc *MyTasksTabStatusCreate) check() error {
	if _, ok := mttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "workspace_id"`)}
	}
	if _, ok := mttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "teammate_id"`)}
	}
	if _, ok := mttsc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := mttsc.mutation.Status(); ok {
		if err := mytaskstabstatus.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := mttsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mttsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New("ent: missing required edge \"workspace\"")}
	}
	if _, ok := mttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	return nil
}

func (mttsc *MyTasksTabStatusCreate) sqlSave(ctx context.Context) (*MyTasksTabStatus, error) {
	_node, _spec := mttsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mttsc.driver, _spec); err != nil {
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

func (mttsc *MyTasksTabStatusCreate) createSpec() (*MyTasksTabStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &MyTasksTabStatus{config: mttsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mytaskstabstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mytaskstabstatus.FieldID,
			},
		}
	)
	if id, ok := mttsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mttsc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mytaskstabstatus.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := mttsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mytaskstabstatus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mttsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mytaskstabstatus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := mttsc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mytaskstabstatus.WorkspaceTable,
			Columns: []string{mytaskstabstatus.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkspaceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mttsc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mytaskstabstatus.TeammateTable,
			Columns: []string{mytaskstabstatus.TeammateColumn},
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
	return _node, _spec
}

// MyTasksTabStatusCreateBulk is the builder for creating many MyTasksTabStatus entities in bulk.
type MyTasksTabStatusCreateBulk struct {
	config
	builders []*MyTasksTabStatusCreate
}

// Save creates the MyTasksTabStatus entities in the database.
func (mttscb *MyTasksTabStatusCreateBulk) Save(ctx context.Context) ([]*MyTasksTabStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mttscb.builders))
	nodes := make([]*MyTasksTabStatus, len(mttscb.builders))
	mutators := make([]Mutator, len(mttscb.builders))
	for i := range mttscb.builders {
		func(i int, root context.Context) {
			builder := mttscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MyTasksTabStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, mttscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mttscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mttscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mttscb *MyTasksTabStatusCreateBulk) SaveX(ctx context.Context) []*MyTasksTabStatus {
	v, err := mttscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mttscb *MyTasksTabStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := mttscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mttscb *MyTasksTabStatusCreateBulk) ExecX(ctx context.Context) {
	if err := mttscb.Exec(ctx); err != nil {
		panic(err)
	}
}