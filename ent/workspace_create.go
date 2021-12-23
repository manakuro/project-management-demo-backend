// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceCreate is the builder for creating a Workspace entity.
type WorkspaceCreate struct {
	config
	mutation *WorkspaceMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (wc *WorkspaceCreate) SetCreatedBy(u ulid.ID) *WorkspaceCreate {
	wc.mutation.SetCreatedBy(u)
	return wc
}

// SetName sets the "name" field.
func (wc *WorkspaceCreate) SetName(s string) *WorkspaceCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkspaceCreate) SetDescription(sd schema.WorkspaceDescription) *WorkspaceCreate {
	wc.mutation.SetDescription(sd)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkspaceCreate) SetCreatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableCreatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkspaceCreate) SetUpdatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableUpdatedAt(t *time.Time) *WorkspaceCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkspaceCreate) SetID(u ulid.ID) *WorkspaceCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkspaceCreate) SetNillableID(u *ulid.ID) *WorkspaceCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetTeammateID sets the "teammate" edge to the Teammate entity by ID.
func (wc *WorkspaceCreate) SetTeammateID(id ulid.ID) *WorkspaceCreate {
	wc.mutation.SetTeammateID(id)
	return wc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (wc *WorkspaceCreate) SetTeammate(t *Teammate) *WorkspaceCreate {
	return wc.SetTeammateID(t.ID)
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wc *WorkspaceCreate) Mutation() *WorkspaceMutation {
	return wc.mutation
}

// Save creates the Workspace in the database.
func (wc *WorkspaceCreate) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkspaceCreate) SaveX(ctx context.Context) *Workspace {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkspaceCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkspaceCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkspaceCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workspace.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workspace.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workspace.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkspaceCreate) check() error {
	if _, ok := wc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "created_by"`)}
	}
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := workspace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := wc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	return nil
}

func (wc *WorkspaceCreate) sqlSave(ctx context.Context) (*Workspace, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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

func (wc *WorkspaceCreate) createSpec() (*Workspace, *sqlgraph.CreateSpec) {
	var (
		_node = &Workspace{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workspace.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: workspace.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
		_node.Name = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workspace.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workspace.TeammateTable,
			Columns: []string{workspace.TeammateColumn},
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
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkspaceCreateBulk is the builder for creating many Workspace entities in bulk.
type WorkspaceCreateBulk struct {
	config
	builders []*WorkspaceCreate
}

// Save creates the Workspace entities in the database.
func (wcb *WorkspaceCreateBulk) Save(ctx context.Context) ([]*Workspace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workspace, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) SaveX(ctx context.Context) []*Workspace {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkspaceCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
