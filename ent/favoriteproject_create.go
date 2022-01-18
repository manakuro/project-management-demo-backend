// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteProjectCreate is the builder for creating a FavoriteProject entity.
type FavoriteProjectCreate struct {
	config
	mutation *FavoriteProjectMutation
	hooks    []Hook
}

// SetProjectID sets the "project_id" field.
func (fpc *FavoriteProjectCreate) SetProjectID(u ulid.ID) *FavoriteProjectCreate {
	fpc.mutation.SetProjectID(u)
	return fpc
}

// SetTeammateID sets the "teammate_id" field.
func (fpc *FavoriteProjectCreate) SetTeammateID(u ulid.ID) *FavoriteProjectCreate {
	fpc.mutation.SetTeammateID(u)
	return fpc
}

// SetCreatedAt sets the "created_at" field.
func (fpc *FavoriteProjectCreate) SetCreatedAt(t time.Time) *FavoriteProjectCreate {
	fpc.mutation.SetCreatedAt(t)
	return fpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fpc *FavoriteProjectCreate) SetNillableCreatedAt(t *time.Time) *FavoriteProjectCreate {
	if t != nil {
		fpc.SetCreatedAt(*t)
	}
	return fpc
}

// SetUpdatedAt sets the "updated_at" field.
func (fpc *FavoriteProjectCreate) SetUpdatedAt(t time.Time) *FavoriteProjectCreate {
	fpc.mutation.SetUpdatedAt(t)
	return fpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fpc *FavoriteProjectCreate) SetNillableUpdatedAt(t *time.Time) *FavoriteProjectCreate {
	if t != nil {
		fpc.SetUpdatedAt(*t)
	}
	return fpc
}

// SetID sets the "id" field.
func (fpc *FavoriteProjectCreate) SetID(u ulid.ID) *FavoriteProjectCreate {
	fpc.mutation.SetID(u)
	return fpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fpc *FavoriteProjectCreate) SetNillableID(u *ulid.ID) *FavoriteProjectCreate {
	if u != nil {
		fpc.SetID(*u)
	}
	return fpc
}

// SetProject sets the "project" edge to the Project entity.
func (fpc *FavoriteProjectCreate) SetProject(p *Project) *FavoriteProjectCreate {
	return fpc.SetProjectID(p.ID)
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (fpc *FavoriteProjectCreate) SetTeammate(t *Teammate) *FavoriteProjectCreate {
	return fpc.SetTeammateID(t.ID)
}

// Mutation returns the FavoriteProjectMutation object of the builder.
func (fpc *FavoriteProjectCreate) Mutation() *FavoriteProjectMutation {
	return fpc.mutation
}

// Save creates the FavoriteProject in the database.
func (fpc *FavoriteProjectCreate) Save(ctx context.Context) (*FavoriteProject, error) {
	var (
		err  error
		node *FavoriteProject
	)
	fpc.defaults()
	if len(fpc.hooks) == 0 {
		if err = fpc.check(); err != nil {
			return nil, err
		}
		node, err = fpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fpc.check(); err != nil {
				return nil, err
			}
			fpc.mutation = mutation
			if node, err = fpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fpc.hooks) - 1; i >= 0; i-- {
			if fpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fpc *FavoriteProjectCreate) SaveX(ctx context.Context) *FavoriteProject {
	v, err := fpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fpc *FavoriteProjectCreate) Exec(ctx context.Context) error {
	_, err := fpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpc *FavoriteProjectCreate) ExecX(ctx context.Context) {
	if err := fpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fpc *FavoriteProjectCreate) defaults() {
	if _, ok := fpc.mutation.CreatedAt(); !ok {
		v := favoriteproject.DefaultCreatedAt()
		fpc.mutation.SetCreatedAt(v)
	}
	if _, ok := fpc.mutation.UpdatedAt(); !ok {
		v := favoriteproject.DefaultUpdatedAt()
		fpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fpc.mutation.ID(); !ok {
		v := favoriteproject.DefaultID()
		fpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fpc *FavoriteProjectCreate) check() error {
	if _, ok := fpc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "project_id"`)}
	}
	if _, ok := fpc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "teammate_id"`)}
	}
	if _, ok := fpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := fpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := fpc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New("ent: missing required edge \"project\"")}
	}
	if _, ok := fpc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	return nil
}

func (fpc *FavoriteProjectCreate) sqlSave(ctx context.Context) (*FavoriteProject, error) {
	_node, _spec := fpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fpc.driver, _spec); err != nil {
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

func (fpc *FavoriteProjectCreate) createSpec() (*FavoriteProject, *sqlgraph.CreateSpec) {
	var (
		_node = &FavoriteProject{config: fpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: favoriteproject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: favoriteproject.FieldID,
			},
		}
	)
	if id, ok := fpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: favoriteproject.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: favoriteproject.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := fpc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteproject.ProjectTable,
			Columns: []string{favoriteproject.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fpc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favoriteproject.TeammateTable,
			Columns: []string{favoriteproject.TeammateColumn},
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

// FavoriteProjectCreateBulk is the builder for creating many FavoriteProject entities in bulk.
type FavoriteProjectCreateBulk struct {
	config
	builders []*FavoriteProjectCreate
}

// Save creates the FavoriteProject entities in the database.
func (fpcb *FavoriteProjectCreateBulk) Save(ctx context.Context) ([]*FavoriteProject, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fpcb.builders))
	nodes := make([]*FavoriteProject, len(fpcb.builders))
	mutators := make([]Mutator, len(fpcb.builders))
	for i := range fpcb.builders {
		func(i int, root context.Context) {
			builder := fpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FavoriteProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, fpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fpcb *FavoriteProjectCreateBulk) SaveX(ctx context.Context) []*FavoriteProject {
	v, err := fpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fpcb *FavoriteProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := fpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpcb *FavoriteProjectCreateBulk) ExecX(ctx context.Context) {
	if err := fpcb.Exec(ctx); err != nil {
		panic(err)
	}
}