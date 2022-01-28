// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/ent/teammatetasksection"
	"project-management-demo-backend/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeammateTaskSectionCreate is the builder for creating a TeammateTaskSection entity.
type TeammateTaskSectionCreate struct {
	config
	mutation *TeammateTaskSectionMutation
	hooks    []Hook
}

// SetTeammateID sets the "teammate_id" field.
func (ttsc *TeammateTaskSectionCreate) SetTeammateID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetTeammateID(u)
	return ttsc
}

// SetWorkspaceID sets the "workspace_id" field.
func (ttsc *TeammateTaskSectionCreate) SetWorkspaceID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetWorkspaceID(u)
	return ttsc
}

// SetName sets the "name" field.
func (ttsc *TeammateTaskSectionCreate) SetName(s string) *TeammateTaskSectionCreate {
	ttsc.mutation.SetName(s)
	return ttsc
}

// SetAssigned sets the "assigned" field.
func (ttsc *TeammateTaskSectionCreate) SetAssigned(b bool) *TeammateTaskSectionCreate {
	ttsc.mutation.SetAssigned(b)
	return ttsc
}

// SetCreatedAt sets the "created_at" field.
func (ttsc *TeammateTaskSectionCreate) SetCreatedAt(t time.Time) *TeammateTaskSectionCreate {
	ttsc.mutation.SetCreatedAt(t)
	return ttsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableCreatedAt(t *time.Time) *TeammateTaskSectionCreate {
	if t != nil {
		ttsc.SetCreatedAt(*t)
	}
	return ttsc
}

// SetUpdatedAt sets the "updated_at" field.
func (ttsc *TeammateTaskSectionCreate) SetUpdatedAt(t time.Time) *TeammateTaskSectionCreate {
	ttsc.mutation.SetUpdatedAt(t)
	return ttsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableUpdatedAt(t *time.Time) *TeammateTaskSectionCreate {
	if t != nil {
		ttsc.SetUpdatedAt(*t)
	}
	return ttsc
}

// SetID sets the "id" field.
func (ttsc *TeammateTaskSectionCreate) SetID(u ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.SetID(u)
	return ttsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttsc *TeammateTaskSectionCreate) SetNillableID(u *ulid.ID) *TeammateTaskSectionCreate {
	if u != nil {
		ttsc.SetID(*u)
	}
	return ttsc
}

// SetTeammate sets the "teammate" edge to the Teammate entity.
func (ttsc *TeammateTaskSectionCreate) SetTeammate(t *Teammate) *TeammateTaskSectionCreate {
	return ttsc.SetTeammateID(t.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (ttsc *TeammateTaskSectionCreate) SetWorkspace(w *Workspace) *TeammateTaskSectionCreate {
	return ttsc.SetWorkspaceID(w.ID)
}

// AddTeammateTaskIDs adds the "teammate_tasks" edge to the TeammateTask entity by IDs.
func (ttsc *TeammateTaskSectionCreate) AddTeammateTaskIDs(ids ...ulid.ID) *TeammateTaskSectionCreate {
	ttsc.mutation.AddTeammateTaskIDs(ids...)
	return ttsc
}

// AddTeammateTasks adds the "teammate_tasks" edges to the TeammateTask entity.
func (ttsc *TeammateTaskSectionCreate) AddTeammateTasks(t ...*TeammateTask) *TeammateTaskSectionCreate {
	ids := make([]ulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttsc.AddTeammateTaskIDs(ids...)
}

// Mutation returns the TeammateTaskSectionMutation object of the builder.
func (ttsc *TeammateTaskSectionCreate) Mutation() *TeammateTaskSectionMutation {
	return ttsc.mutation
}

// Save creates the TeammateTaskSection in the database.
func (ttsc *TeammateTaskSectionCreate) Save(ctx context.Context) (*TeammateTaskSection, error) {
	var (
		err  error
		node *TeammateTaskSection
	)
	ttsc.defaults()
	if len(ttsc.hooks) == 0 {
		if err = ttsc.check(); err != nil {
			return nil, err
		}
		node, err = ttsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeammateTaskSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttsc.check(); err != nil {
				return nil, err
			}
			ttsc.mutation = mutation
			if node, err = ttsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttsc.hooks) - 1; i >= 0; i-- {
			if ttsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttsc *TeammateTaskSectionCreate) SaveX(ctx context.Context) *TeammateTaskSection {
	v, err := ttsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttsc *TeammateTaskSectionCreate) Exec(ctx context.Context) error {
	_, err := ttsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttsc *TeammateTaskSectionCreate) ExecX(ctx context.Context) {
	if err := ttsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttsc *TeammateTaskSectionCreate) defaults() {
	if _, ok := ttsc.mutation.CreatedAt(); !ok {
		v := teammatetasksection.DefaultCreatedAt()
		ttsc.mutation.SetCreatedAt(v)
	}
	if _, ok := ttsc.mutation.UpdatedAt(); !ok {
		v := teammatetasksection.DefaultUpdatedAt()
		ttsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttsc.mutation.ID(); !ok {
		v := teammatetasksection.DefaultID()
		ttsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttsc *TeammateTaskSectionCreate) check() error {
	if _, ok := ttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate_id", err: errors.New(`ent: missing required field "teammate_id"`)}
	}
	if _, ok := ttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "workspace_id"`)}
	}
	if _, ok := ttsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := ttsc.mutation.Name(); ok {
		if err := teammatetasksection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := ttsc.mutation.Assigned(); !ok {
		return &ValidationError{Name: "assigned", err: errors.New(`ent: missing required field "assigned"`)}
	}
	if _, ok := ttsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ttsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ttsc.mutation.TeammateID(); !ok {
		return &ValidationError{Name: "teammate", err: errors.New("ent: missing required edge \"teammate\"")}
	}
	if _, ok := ttsc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace", err: errors.New("ent: missing required edge \"workspace\"")}
	}
	return nil
}

func (ttsc *TeammateTaskSectionCreate) sqlSave(ctx context.Context) (*TeammateTaskSection, error) {
	_node, _spec := ttsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttsc.driver, _spec); err != nil {
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

func (ttsc *TeammateTaskSectionCreate) createSpec() (*TeammateTaskSection, *sqlgraph.CreateSpec) {
	var (
		_node = &TeammateTaskSection{config: ttsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teammatetasksection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teammatetasksection.FieldID,
			},
		}
	)
	if id, ok := ttsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ttsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teammatetasksection.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ttsc.mutation.Assigned(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teammatetasksection.FieldAssigned,
		})
		_node.Assigned = value
	}
	if value, ok := ttsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasksection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ttsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teammatetasksection.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ttsc.mutation.TeammateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.TeammateTable,
			Columns: []string{teammatetasksection.TeammateColumn},
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
	if nodes := ttsc.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teammatetasksection.WorkspaceTable,
			Columns: []string{teammatetasksection.WorkspaceColumn},
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
	if nodes := ttsc.mutation.TeammateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teammatetasksection.TeammateTasksTable,
			Columns: []string{teammatetasksection.TeammateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teammatetask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TeammateTaskSectionCreateBulk is the builder for creating many TeammateTaskSection entities in bulk.
type TeammateTaskSectionCreateBulk struct {
	config
	builders []*TeammateTaskSectionCreate
}

// Save creates the TeammateTaskSection entities in the database.
func (ttscb *TeammateTaskSectionCreateBulk) Save(ctx context.Context) ([]*TeammateTaskSection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttscb.builders))
	nodes := make([]*TeammateTaskSection, len(ttscb.builders))
	mutators := make([]Mutator, len(ttscb.builders))
	for i := range ttscb.builders {
		func(i int, root context.Context) {
			builder := ttscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeammateTaskSectionMutation)
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
					_, err = mutators[i+1].Mutate(root, ttscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ttscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttscb *TeammateTaskSectionCreateBulk) SaveX(ctx context.Context) []*TeammateTaskSection {
	v, err := ttscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttscb *TeammateTaskSectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ttscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttscb *TeammateTaskSectionCreateBulk) ExecX(ctx context.Context) {
	if err := ttscb.Exec(ctx); err != nil {
		panic(err)
	}
}