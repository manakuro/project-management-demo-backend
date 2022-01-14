// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecticon"
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectIconCreate is the builder for creating a ProjectIcon entity.
type ProjectIconCreate struct {
	config
	mutation *ProjectIconMutation
	hooks    []Hook
}

// SetIconID sets the "icon_id" field.
func (pic *ProjectIconCreate) SetIconID(u ulid.ID) *ProjectIconCreate {
	pic.mutation.SetIconID(u)
	return pic
}

// SetCreatedAt sets the "created_at" field.
func (pic *ProjectIconCreate) SetCreatedAt(t time.Time) *ProjectIconCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *ProjectIconCreate) SetNillableCreatedAt(t *time.Time) *ProjectIconCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *ProjectIconCreate) SetUpdatedAt(t time.Time) *ProjectIconCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *ProjectIconCreate) SetNillableUpdatedAt(t *time.Time) *ProjectIconCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *ProjectIconCreate) SetID(u ulid.ID) *ProjectIconCreate {
	pic.mutation.SetID(u)
	return pic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pic *ProjectIconCreate) SetNillableID(u *ulid.ID) *ProjectIconCreate {
	if u != nil {
		pic.SetID(*u)
	}
	return pic
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (pic *ProjectIconCreate) AddProjectIDs(ids ...ulid.ID) *ProjectIconCreate {
	pic.mutation.AddProjectIDs(ids...)
	return pic
}

// AddProjects adds the "projects" edges to the Project entity.
func (pic *ProjectIconCreate) AddProjects(p ...*Project) *ProjectIconCreate {
	ids := make([]ulid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pic.AddProjectIDs(ids...)
}

// SetIcon sets the "icon" edge to the Icon entity.
func (pic *ProjectIconCreate) SetIcon(i *Icon) *ProjectIconCreate {
	return pic.SetIconID(i.ID)
}

// Mutation returns the ProjectIconMutation object of the builder.
func (pic *ProjectIconCreate) Mutation() *ProjectIconMutation {
	return pic.mutation
}

// Save creates the ProjectIcon in the database.
func (pic *ProjectIconCreate) Save(ctx context.Context) (*ProjectIcon, error) {
	var (
		err  error
		node *ProjectIcon
	)
	pic.defaults()
	if len(pic.hooks) == 0 {
		if err = pic.check(); err != nil {
			return nil, err
		}
		node, err = pic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectIconMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pic.check(); err != nil {
				return nil, err
			}
			pic.mutation = mutation
			if node, err = pic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pic.hooks) - 1; i >= 0; i-- {
			if pic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pic *ProjectIconCreate) SaveX(ctx context.Context) *ProjectIcon {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *ProjectIconCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *ProjectIconCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *ProjectIconCreate) defaults() {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		v := projecticon.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		v := projecticon.DefaultUpdatedAt()
		pic.mutation.SetUpdatedAt(v)
	}
	if _, ok := pic.mutation.ID(); !ok {
		v := projecticon.DefaultID()
		pic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *ProjectIconCreate) check() error {
	if _, ok := pic.mutation.IconID(); !ok {
		return &ValidationError{Name: "icon_id", err: errors.New(`ent: missing required field "icon_id"`)}
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := pic.mutation.IconID(); !ok {
		return &ValidationError{Name: "icon", err: errors.New("ent: missing required edge \"icon\"")}
	}
	return nil
}

func (pic *ProjectIconCreate) sqlSave(ctx context.Context) (*ProjectIcon, error) {
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
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

func (pic *ProjectIconCreate) createSpec() (*ProjectIcon, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectIcon{config: pic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projecticon.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: projecticon.FieldID,
			},
		}
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecticon.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecticon.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pic.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecticon.ProjectsTable,
			Columns: []string{projecticon.ProjectsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.IconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecticon.IconTable,
			Columns: []string{projecticon.IconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: icon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.IconID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectIconCreateBulk is the builder for creating many ProjectIcon entities in bulk.
type ProjectIconCreateBulk struct {
	config
	builders []*ProjectIconCreate
}

// Save creates the ProjectIcon entities in the database.
func (picb *ProjectIconCreateBulk) Save(ctx context.Context) ([]*ProjectIcon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*ProjectIcon, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectIconMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *ProjectIconCreateBulk) SaveX(ctx context.Context) []*ProjectIcon {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *ProjectIconCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *ProjectIconCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
